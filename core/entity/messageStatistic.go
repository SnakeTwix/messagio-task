package entity

import "time"

type MessageStatistic struct {
	ID             uint64        `json:"id"`
	ReadTimes      []MessageRead `json:"readTimes"`
	KafkaProcessed bool          `json:"kafkaProcessed"`
	CreatedAt      time.Time     `json:"createdAt"`
}
