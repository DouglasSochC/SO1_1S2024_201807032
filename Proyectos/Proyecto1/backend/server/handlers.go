package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

// Estructura para almacenar la informacion de la RAM
type RAMInfo struct {
	MemoriaTotal         int `json:"memoria_total"`
	MemoriaUso           int `json:"memoria_uso"`
	MemoriaPorcentajeUso int `json:"memoria_porcentaje_uso"`
	MemoriaLibre         int `json:"memoria_libre"`
}

// Estructura para almacenar la informacion de la CPU
type CPUInfo struct {
	Total    int             `json:"cpu_total"`
	EnUSo    int             `json:"cpu_porcentaje"`
	Procesos json.RawMessage `json:"processes"`
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
	var memInfo RAMInfo
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

	// Obtener información de la RAM
	ramCmd := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	ramOut, ramErr := ramCmd.CombinedOutput()
	if ramErr != nil {
		fmt.Println(ramErr)
		http.Error(w, "Error al obtener la información de la memoria", http.StatusInternalServerError)
		return
	}

	// Procesar el output y construir la estructura RAMInfo
	var ramInfo RAMInfo
	err := json.Unmarshal(ramOut, &ramInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir el output a JSON", http.StatusInternalServerError)
		return
	}

	// Obtener información de la CPU
	cpuCmd := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	cpuOut, cpuErr := cpuCmd.CombinedOutput()
	if cpuErr != nil {
		fmt.Println(cpuErr)
		http.Error(w, "Error al obtener la información de la CPU", http.StatusInternalServerError)
		return
	}

	// Procesar el output y construir la estructura CPUInfo
	var cpuInfo CPUInfo
	err = json.Unmarshal(cpuOut, &cpuInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir el output de CPU a JSON", http.StatusInternalServerError)
		return
	}

	// Configurar el encabezado para indicar que el contenido es JSON
	w.Header().Set("Content-Type", "application/json")

	// Permitir solicitudes desde cualquier origen
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Convertir la estructura RAMInfo a JSON y escribir en el ResponseWriter
	ramData, err := json.Marshal(ramInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir la estructura de RAM a JSON", http.StatusInternalServerError)
		return
	}

	// Convertir la estructura CPUInfo a JSON y añadir al ResponseWriter
	cpuData, err := json.Marshal(cpuInfo)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al convertir la estructura de CPU a JSON", http.StatusInternalServerError)
		return
	}

	// Combinar datos de RAM y CPU en una respuesta JSON
	responseData := fmt.Sprintf(`{"ram": %s, "cpu": %s}`, ramData, cpuData)

	fmt.Fprintf(w, responseData)
}
