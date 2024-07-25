package server

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"messagio/core/enums/env"
	"messagio/utils"
)

type KafkaConn struct {
	MessageReader *kafka.Reader
	MessageWriter *kafka.Writer
}

func initKafka() *KafkaConn {
	initKafkaTopics()

	kafkaMessageReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{utils.GetEnv(env.KafkaAddress)},
		Topic:     utils.GetEnv(env.KafkaMessageTopic),
		Partition: 0,
		MaxBytes:  10e6, // 10 mb
	})

	kafkaMessageWriter := &kafka.Writer{
		Addr:      kafka.TCP(utils.GetEnv(env.KafkaAddress)),
		Topic:     utils.GetEnv(env.KafkaMessageTopic),
		Balancer:  &kafka.LeastBytes{},
		BatchSize: 1,
	}

	return &KafkaConn{
		MessageReader: kafkaMessageReader,
		MessageWriter: kafkaMessageWriter,
	}
}

func initKafkaTopics() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", utils.GetEnv(env.KafkaAddress), utils.GetEnv(env.KafkaMessageTopic), 0)

	if err != nil {
		log.Fatal("couldn't connect to kafka leader: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("couldn't close kafka connection: ", err)
	}
}
