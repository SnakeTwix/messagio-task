package entity

import "time"

type MessageRead struct {
	MessageID uint64    `json:"-"`
	CreatedAt time.Time `json:"readAt"`
}
