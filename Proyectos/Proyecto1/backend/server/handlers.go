package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

// Estructura para almacenar la información
type MemInfo struct {
	CantidadMemoria      int `json:"cantidad_memoria"`
	MemoriaUso           int `json:"memoria_uso"`
	MemoriaPorcentajeUso int `json:"memoria_porcentaje_uso"`
	MemoriaLibre         int `json:"memoria_libre"`
}

func index(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al obtener la información de la memoria", http.StatusInternalServerError)
		return
	}

	// Procesar el output y construir la estructura MemInfo
	var memInfo MemInfo
	err = json.Unmarshal(out, &memInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir el output a JSON", http.StatusInternalServerError)
		return
	}

	// Configurar el encabezado para indicar que el contenido es JSON
	w.Header().Set("Content-Type", "application/json")

	// Convertir la estructura a JSON y escribir en el ResponseWriter
	jsonData, err := json.Marshal(memInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir la estructura a JSON", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}

func getMonitoreoTiempoReal(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al obtener la información de la memoria", http.StatusInternalServerError)
		return
	}

	// Procesar el output y construir la estructura MemInfo
	var memInfo MemInfo
	err = json.Unmarshal(out, &memInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir el output a JSON", http.StatusInternalServerError)
		return
	}

	// Configurar el encabezado para indicar que el contenido es JSON
	w.Header().Set("Content-Type", "application/json")

	// Convertir la estructura a JSON y escribir en el ResponseWriter
	jsonData, err := json.Marshal(memInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir la estructura a JSON", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}
