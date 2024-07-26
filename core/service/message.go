package service

import (
	"context"
	"messagio/core/entity"
	"messagio/core/ports"
)

type Message struct {
	repoMessage ports.RepoMessage
}

var s *Message

func GetMessage(repoMessage ports.RepoMessage) *Message {
	if s != nil {
		return s
	}

	s = &Message{
		repoMessage: repoMessage,
	}

	return s
}

func (s *Message) CreateMessage(ctx context.Context, createMessage *entity.CreateMessage) error {
	message := entity.Message{
		Content: createMessage.Content,
	}

	return s.repoMessage.CreateMessage(ctx, &message)
}

func (s *Message) GetMessage(ctx context.Context, messageId uint64) (*entity.Message, error) {
	message, err := s.repoMessage.GetMessage(ctx, messageId)
	if err != nil {
		return nil, err
	}

	err = s.repoMessage.ReadMessage(ctx, message.ID)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *Message) GetMessages(ctx context.Context, paginateOptions entity.PaginateRequest) (*entity.PaginateResponse[entity.Message], error) {
	computedOptions := entity.Paginate{
		Limit:  paginateOptions.Limit,
		Offset: (paginateOptions.Page - 1) * paginateOptions.Limit,
	}

	paginatedMessages, err := s.repoMessage.GetMessages(ctx, computedOptions)
	if err != nil {
		return nil, err
	}

	for index := range paginatedMessages.Values {
		message := &paginatedMessages.Values[index]

		err = s.repoMessage.ReadMessage(ctx, message.ID)
		if err != nil {
			continue
		}
	}

	return paginatedMessages, nil
}

func (s *Message) GetNewMessages(ctx context.Context) ([]entity.Message, error) {
	messages, err := s.repoMessage.GetNewMessages(ctx)
	if err != nil {
		return nil, err
	}

	for index := range messages {
		message := &messages[index]
		message.KafkaProcessed = true

		err = s.repoMessage.UpdateMessage(ctx, message)
		if err != nil {
			continue
		}

		// This is 3 queries PER message. Quite a lot, would be ideal to batch, yeah...
		err = s.repoMessage.ReadMessage(ctx, message.ID)
		if err != nil {
			continue
		}
	}

	return messages, nil
}

func (s *Message) GetFullMessage(ctx context.Context, messageId uint64) (*entity.Message, error) {
	return s.repoMessage.GetFullMessage(ctx, messageId)
}
