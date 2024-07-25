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
	return s.repoMessage.GetMessage(ctx, messageId)
}

func (s *Message) GetMessages(ctx context.Context, paginateOptions entity.PaginateRequest) (*entity.PaginateResponse[entity.Message], error) {
	computedOptions := entity.Paginate{
		Limit:  paginateOptions.Limit,
		Offset: (paginateOptions.Page - 1) * paginateOptions.Limit,
	}

	return s.repoMessage.GetMessages(ctx, computedOptions)
}
