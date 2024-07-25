package repository

import (
	"context"
	"encoding/binary"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"messagio/core/entity"
	"time"
)

type Message struct {
	db                 *gorm.DB
	kafkaMessageReader *kafka.Reader
	kafkaMessageWriter *kafka.Writer
}

var repo *Message

func GetMessage(db *gorm.DB, kafkaMessageReader *kafka.Reader, kafkaMessageWrite *kafka.Writer) *Message {
	if repo != nil {
		return repo
	}

	repo = &Message{
		db:                 db,
		kafkaMessageWriter: kafkaMessageWrite,
		kafkaMessageReader: kafkaMessageReader,
	}

	return repo
}

func (r *Message) CreateMessage(ctx context.Context, message *entity.Message) error {
	result := r.db.WithContext(ctx).Select("Content").Create(message)

	if result.Error != nil {
		return result.Error
	}

	byteId := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteId, message.ID)

	err := r.kafkaMessageWriter.WriteMessages(ctx,
		kafka.Message{
			Value: byteId,
		},
	)

	return err
}

func (r *Message) GetMessage(ctx context.Context, messageId uint64) (*entity.Message, error) {
	var message entity.Message

	result := r.db.WithContext(ctx).First(&message, messageId)

	return &message, result.Error
}

func (r *Message) GetMessages(ctx context.Context, paginateOptions entity.Paginate) (*entity.PaginateResponse[entity.Message], error) {
	var messages []entity.Message
	var totalCount int64

	result := r.db.WithContext(ctx).Order("id").Limit(paginateOptions.Limit).Offset(paginateOptions.Offset).Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.WithContext(ctx).Model(&entity.Message{}).Count(&totalCount)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity.PaginateResponse[entity.Message]{
		Values: messages,
		Total:  int(totalCount),
	}, nil
}

func (r *Message) GetNewMessages(ctx context.Context) ([]entity.Message, error) {
	messages := []entity.Message{}
	messageIds := []uint64{}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	var lastMessage *kafka.Message

	// Максимум обрабатывает 100 за раз
	for i := 0; i < 100; i++ {
		message, err := r.kafkaMessageReader.FetchMessage(timeoutCtx)
		if err != nil {
			break
		}

		lastMessage = &message
		parsedId, _ := binary.Uvarint(message.Value)
		messageIds = append(messageIds, parsedId)
	}

	if lastMessage != nil {
		err := r.kafkaMessageReader.CommitMessages(ctx, *lastMessage)
		if err != nil {
			return nil, err
		}
	}

	for _, messageId := range messageIds {

		// Could batch get and update this instead of sending many sql queries to the database
		// but ehh... There is no native functionality for this
		message, err := r.GetMessage(ctx, messageId)
		if err != nil {
			continue
		}

		message.KafkaProcessed = true

		err = r.UpdateMessage(ctx, message)
		if err != nil {
			continue
		}

		// This is 3 queries PER message. Quite a lot, would be ideal to batch, yeah...
		err = r.ReadMessage(ctx, messageId)
		if err != nil {
			continue
		}

		messages = append(messages, *message)
	}

	return messages, nil
}

func (r *Message) UpdateMessage(ctx context.Context, message *entity.Message) error {
	result := r.db.WithContext(ctx).Save(message)
	return result.Error
}

func (r *Message) ReadMessage(ctx context.Context, messageId uint64) error {
	messageRead := entity.MessageRead{MessageID: messageId}

	result := r.db.WithContext(ctx).Create(&messageRead)
	return result.Error
}
