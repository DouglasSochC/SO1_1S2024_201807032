package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

var ctx = context.Background()
var mongoClient *mongo.Client

func main() {
	// Cargar variables de entorno
	godotenv.Load()

	// Obtener variables de entorno
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	redisAddr := os.Getenv("REDIS_ADDR")
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	mongoCollectionName := os.Getenv("MONGO_COLLECTION_NAME")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"))

	// Cliente Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Cliente MongoDB
	mongoClientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	mongoClient, err = mongo.Connect(ctx, mongoClientOptions)
	if err != nil {
		log.Fatalf("Error al conectar con MongoDB: %v", err)
		return
	}
	defer mongoClient.Disconnect(ctx)

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
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		var data Data
		if err := json.Unmarshal(m.Value, &data); err != nil {
			log.Fatal("Error deserializando JSON:", err)
		}

		insertMongoDB(mongoDBName, mongoCollectionName, data)
		insertRedis(rdb, data)
	}
}

func insertMongoDB(mongoDBName string, mongoCollectionName string, data Data) {
	collection := mongoClient.Database(mongoDBName).Collection(mongoCollectionName)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal("Error al insertar en MongoDB:", err)
	}
	log.Printf("Inserción en MongoDB: %v", data)
}

func insertRedis(rdb *redis.Client, data Data) {
	// Crear llave hash
	key := fmt.Sprintf("%s:%s:%s", data.Name, data.Album, data.Year)
	hash := sha256.Sum256([]byte(key))
	hashKey := hex.EncodeToString(hash[:])

	// Incrementar contador en Redis
	count, err := rdb.Incr(ctx, hashKey).Result()
	if err != nil {
		log.Fatal("Error al incrementar contador en Redis:", err)
	}

	log.Printf("Inserción en Redis, Llave: %s, Contador: %d", hashKey, count)
}
