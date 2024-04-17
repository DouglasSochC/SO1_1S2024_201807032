package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables de entorno
	godotenv.Load()

	// Obtener variables de entorno para la conexion a Kafka
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"group.id":          "test-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	err = c.SubscribeTopics([]string{kafkaTopic}, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to subscribe to topics: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Listening to topic: %s\n", kafkaTopic)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message on topic %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Fprintf(os.Stderr, "Error reading message: %s\n", err)
		}
	}
}
