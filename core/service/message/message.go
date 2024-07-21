package message

import (
	"context"
	"messagio/core/entity"
	"messagio/core/ports"
)

type Service struct {
	repoMessage ports.RepoMessage
}

var service *Service

func GetService(repoMessage ports.RepoMessage) *Service {
	if service != nil {
		return service
	}

	service = &Service{
		repoMessage: repoMessage,
	}

	return service
}

func (s *Service) CreateMessage(ctx context.Context, createMessage *entity.CreateMessage) error {
	message := entity.Message{
		Content: createMessage.Content,
	}

	return s.repoMessage.CreateMessage(ctx, &message)
}
