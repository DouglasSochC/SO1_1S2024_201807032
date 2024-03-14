package main

import (
	"backend/database"
	"backend/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	ServerDoneChan := make(chan os.Signal, 1)
	signal.Notify(ServerDoneChan, os.Interrupt, syscall.SIGTERM)

	// Cargar variables de entorno
	godotenv.Load()

	// Obtener variables de entorno
	serverPort := os.Getenv("SERVER_PORT")
	srv := server.New(":" + serverPort)

	// Routine para hacer los registros a la base de datos
	go func() {
		db, err := database.SetupDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		for {
			database.RegistrarHistoricoRAM(db)
			database.RegistrarHistoricoCPU(db)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Routine para el servidor
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Servidor Iniciado...")

	<-ServerDoneChan

	srv.Shutdown(ctx)
}
