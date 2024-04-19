package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	pb "servidor/proto"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {

	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	enviarMensaje(data)
	log.Printf("Recibí la información: %v", data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibí la informacion que me enviaste"}, nil
}

func enviarMensaje(data Data) {

	// Se convierte el 'data' en un JSON
	jsonData, errjson := json.Marshal(data)
	if errjson != nil {
		log.Fatal("Error marshaling data:", errjson)
		return
	}

	// Obtener variables de entorno para la comunicacion con kafka
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	// Crear un escritor de Kafka especificando el tópico y la dirección del servidor
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBrokers},
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	// Enviar un mensaje al tópico
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: jsonData,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}

func main() {

	// Cargar variables de entorno
	godotenv.Load()

	// Obtener variable de entorno para el puerto del servidor
	port := os.Getenv("SERVER_PORT")

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	log.Println("Servidor escuchando en el puerto: " + port)

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
