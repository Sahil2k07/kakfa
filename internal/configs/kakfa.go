package configs

import (
	"log"
	"os"
	"strconv"
)

type kafkaConfig struct {
	Address   string `toml:"kafka_address"`
	Topic     string `toml:"kafka_topic"`
	Partition int    `toml:"kafka_partition"`
	GroupID   string `toml:"kafka_group_id"`
}

func loadKafkaConfig() kafkaConfig {
	partition, err := strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
	if err != nil {
		log.Fatal("invalid kakfa partition value")
	}

	return kafkaConfig{
		Address:   os.Getenv("KAFKA_ADDRESS"),
		Topic:     os.Getenv("KAFKA_TOPIC"),
		Partition: partition,
		GroupID:   os.Getenv("KAFKA_GROUP_ID"),
	}
}

func GetKafkaConfig() kafkaConfig {
	return globalConfig.Kafka
}
