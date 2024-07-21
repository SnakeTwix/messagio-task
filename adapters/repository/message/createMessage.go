package message

import (
	"context"
	"messagio/core/entity"
)

func (r *Repo) CreateMessage(ctx context.Context, message *entity.Message) error {
	result := r.db.WithContext(ctx).Select("Content").Create(message)

	return result.Error
}
