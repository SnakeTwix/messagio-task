package repository

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"messagio/core/entity"
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
	fmt.Println(message.ID)

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
