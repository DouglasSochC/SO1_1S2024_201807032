package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	pb "servidor/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

var ctx = context.Background()
var db *sql.DB

type server struct {
	pb.UnimplementedGetInfoServer
}

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func mysqlConnect() {

	// Obtener variables de entorno para la conexion a MySQL
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Configurando conexion de MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Conexión a MySQL exitosa")
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetRank())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	fmt.Println(data)
	insertMySQL(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func insertMySQL(voto Data) {
	// Prepara la consulta SQL para la inserción en MySQL
	query := "INSERT INTO votos (name_v, album_v, year_v, rank_v) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, voto.Name, voto.Album, voto.Year, voto.Rank)
	if err != nil {
		log.Println("Error al insertar en MySQL:", err)
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

	mysqlConnect()

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
