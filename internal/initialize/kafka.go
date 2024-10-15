package initialize

import (
	"log"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/segmentio/kafka-go"
)

// Init kafka Producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost: 19092"),
		Topic:    "otp-auth-topic", // top name
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer: %v", err)
	}
}
