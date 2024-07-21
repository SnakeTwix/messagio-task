package entity

import "time"

type MessageRead struct {
	MessageID uint64
	CreatedAt time.Time
}
