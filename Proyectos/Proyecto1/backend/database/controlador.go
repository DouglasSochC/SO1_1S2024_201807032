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

// Estructura para almacenar los historicos
type Historicos struct {
	Labels []string `json:"labels"`
	Data   []string `json:"data"`
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

func ObtenerHistoricoRAM(db *sql.DB) ([]byte, error) {

	rows, err := selectHistoricoRAM(db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados Historicos
	for rows.Next() {
		var campo1, campo2 string
		err := rows.Scan(&campo1, &campo2)
		if err != nil {
			return nil, err
		}

		resultados.Labels = append(resultados.Labels, campo1)
		resultados.Data = append(resultados.Data, campo2)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convierte los resultados a formato JSON
	jsonResult, err := json.Marshal(resultados)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil
}

func ObtenerHistoricoCPU(db *sql.DB) ([]byte, error) {

	rows, err := selectHistoricoCPU(db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados Historicos
	for rows.Next() {
		var campo1, campo2 string
		err := rows.Scan(&campo1, &campo2)
		if err != nil {
			return nil, err
		}

		resultados.Labels = append(resultados.Labels, campo1)
		resultados.Data = append(resultados.Data, campo2)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convierte los resultados a formato JSON
	jsonResult, err := json.Marshal(resultados)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil
}
