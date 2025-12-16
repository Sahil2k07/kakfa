package connections

import (
	"github.com/Sahil2k07/kakfa/internal/configs"
	"github.com/segmentio/kafka-go"
)

var (
	KafkaWriter *kafka.Writer
	KafkaReader *kafka.Reader
)

func ConnectKafkaReader() {
	config := configs.GetKafkaConfig()

	KafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.Address},
		Topic:   config.Topic,
		GroupID: config.GroupID,
	})
}

func ConnectKafkaWriter() {
	config := configs.GetKafkaConfig()

	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(config.Address),
		Topic:    config.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}
