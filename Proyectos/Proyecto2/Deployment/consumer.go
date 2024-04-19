package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func main() {
	// Cargar variables de entorno
	godotenv.Load()

	// Obtener variables de entorno
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	// Crear un lector de Kafka especificando el tópico y la dirección del servidor
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaBrokers},
		Topic:     kafkaTopic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer r.Close()

	// Leer mensajes
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var data Data
		if err := json.Unmarshal(m.Value, &data); err != nil {
			log.Fatal("Error unmarshalling JSON:", err)
		}

		insertMongoDB(data)
		insertRedis(data)
	}
}

func insertMongoDB(data Data) {
	log.Printf("Inserción en MongoDB: %v", data)
}

func insertRedis(data Data) {
	log.Printf("Inserción en Redis: %v", data)
}
