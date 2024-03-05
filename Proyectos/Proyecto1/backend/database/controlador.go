package database

import (
	"database/sql"
	"encoding/json"
	"log"
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

func RegistrarHistoricoRAM(db *sql.DB) {

	// Utiliza un modulo para obtener la informacion
	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Procesar el output y construir la estructura RAMInfo
	var memInfo RAMInfo
	err = json.Unmarshal(out, &memInfo)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Realizar registro en la base de datos
	err = insertHistoricoRAM(db, memInfo.MemoriaTotal, memInfo.MemoriaLibre, memInfo.MemoriaUso, memInfo.MemoriaPorcentajeUso)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func RegistrarHistoricoCPU(db *sql.DB) {

	// Utiliza un modulo para obtener la informacion
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Procesar el output y construir la estructura CPUInfo
	var cpuInfo CPUInfo
	err = json.Unmarshal(out, &cpuInfo)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Realizar registro en la base de datos
	err = insertHistoricoCPU(db, cpuInfo.Total, cpuInfo.EnUSo)
	if err != nil {
		log.Fatal(err)
		return
	}
}
