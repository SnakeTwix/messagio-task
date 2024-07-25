package ports

import (
	"context"
	"messagio/core/entity"
)

type ServiceMessage interface {
	CreateMessage(ctx context.Context, message *entity.CreateMessage) error
	GetMessage(ctx context.Context, messageId uint64) (*entity.Message, error)
	GetMessages(ctx context.Context, paginateOptions entity.PaginateRequest) (*entity.PaginateResponse[entity.Message], error)
}

type RepoMessage interface {
	CreateMessage(ctx context.Context, message *entity.Message) error
	GetMessage(ctx context.Context, messageId uint64) (*entity.Message, error)
	GetMessages(ctx context.Context, paginateOptions entity.Paginate) (*entity.PaginateResponse[entity.Message], error)
}
