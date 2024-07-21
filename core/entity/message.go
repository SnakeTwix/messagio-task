package entity

import "time"

type Message struct {
	ID        uint64        `gorm:"primaryKey" json:"id" validate:"required"`
	Content   string        `json:"content"`
	ReadTimes []MessageRead `json:"-"`
	CreatedAt time.Time     `json:"createdAt"`
}

type CreateMessage struct {
	Content string `json:"content" validate:"required"`
}
