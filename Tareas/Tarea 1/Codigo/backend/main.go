package main

import (
	"encoding/json"
	"net/http"
)

// Struct para representar la información que se convertirá a JSON
type Response struct {
	Carnet string `json:"carnet"`
	Nombre string `json:"nombre"`
}

func main() {
	// Manejador de la ruta "/json" que devuelve un JSON
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Crear una instancia de la estructura Response
		response := Response{
			Carnet: "201807032",
			Nombre: "Douglas Alexander Soch Catalán",
		}

		// Convertir la estructura a JSON
		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
			return
		}

		// Establecer el encabezado Content-Type como application/json
		w.Header().Set("Content-Type", "application/json")

		// Permitir solicitudes desde cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Escribir el JSON en la respuesta HTTP
		w.Write(jsonData)
	})

	// Iniciar el servidor en el puerto 8080
	http.ListenAndServe(":8080", nil)
}
