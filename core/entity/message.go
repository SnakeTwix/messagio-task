package entity

import "time"

type Message struct {
	ID        uint64        `gorm:"primaryKey" json:"id"`
	Content   string        `json:"content"`
	ReadTimes []MessageRead `json:"readTimes"`
	CreatedAt time.Time     `json:"createdAt"`
}

type CreateMessage struct {
	Content string `json:"content"`
}
