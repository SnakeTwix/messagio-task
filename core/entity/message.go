package entity

import "time"

type Message struct {
	ID             uint64        `gorm:"primaryKey" json:"id" validate:"required"`
	Content        string        `json:"content"`
	ReadTimes      []MessageRead `json:"readTimes,omitempty"`
	KafkaProcessed bool          `json:"kafkaProcessed" gorm:"default:false"`
	CreatedAt      time.Time     `json:"createdAt"`
}

type CreateMessage struct {
	Content string `json:"content" validate:"required"`
}
