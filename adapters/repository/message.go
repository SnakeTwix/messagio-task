package repository

import (
	"context"
	"gorm.io/gorm"
	"messagio/core/entity"
)

type Message struct {
	db *gorm.DB
}

var repo *Message

func GetMessage(db *gorm.DB) *Message {
	if repo != nil {
		return repo
	}

	repo = &Message{
		db: db,
	}

	return repo
}

func (r *Message) CreateMessage(ctx context.Context, message *entity.Message) error {
	result := r.db.WithContext(ctx).Select("Content").Create(message)

	return result.Error
}

func (r *Message) GetMessage(ctx context.Context, messageId uint64) (*entity.Message, error) {
	var message entity.Message

	result := r.db.WithContext(ctx).First(&message, messageId)

	return &message, result.Error
}
