package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9091"
	kafkaTopic = "user_topic_001"
)

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,           // List of Kafka brokers, e.g., []string{"localhost:9092"}
		GroupID:        groupID,           // Consumer group ID for Kafka
		Topic:          topic,             // Topic to consume messages from
		MinBytes:       10e3,              // Minimum batch size (10KB) to fetch in each request
		MaxBytes:       10e6,              // Maximum batch size (10MB) to fetch in each request
		CommitInterval: time.Second,       // Time interval between offset commits
		StartOffset:    kafka.FirstOffset, // Start consuming from the earliest offset (first available message in the topic)
	})
}

// buy or sell stock
func NewStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s
}

func actionStock(c *gin.Context) {
	s := NewStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	// create msg
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// write message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer waiting to buy ATC
func RegisterConsumerATC(id int64) {
	// group consumer
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) waiting ATC section::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v", id, err)
		}
		fmt.Printf("Consumer(%d), waiting topic: %v, partition: %v, offerset: %v, time: %d %s = %s\n", id, m.Topic, m.Partition,
			m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close() // release connection regardless of success or failure

	r.POST("action/stock", actionStock)

	// register two users(1,2) to buy stock in ATC
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4) // do not retrieve old messages

	r.Run(":8999")
}
