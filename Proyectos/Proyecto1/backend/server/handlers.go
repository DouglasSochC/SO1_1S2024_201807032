package server

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"log"
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

type ProcesoHijoAuxiliar struct {
	PID      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	PidPadre int    `json:"pidPadre"`
}

type ProcesoAuxiliar struct {
	PID      int                   `json:"pid"`
	Name     string                `json:"name"`
	State    int                   `json:"state"`
	RAM      int                   `json:"ram"`
	Children []ProcesoHijoAuxiliar `json:"child"`
}

type CPUInfo struct {
	Total    int               `json:"cpu_total"`
	EnUso    int               `json:"cpu_porcentaje"`
	Procesos []ProcesoAuxiliar `json:"processes"`
}

type ProcesoInfo struct {
	ID       int
	Name     string
	Children []ProcesoInfo
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

func getObtenerHistoricos(w http.ResponseWriter, r *http.Request) {

	// Configurar el encabezado para indicar que el contenido es JSON
	w.Header().Set("Content-Type", "application/json")
	// Permitir solicitudes desde cualquier origen
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := database.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	historicoRAM, err := database.ObtenerHistoricoRAM(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	historicoCPU, err := database.ObtenerHistoricoCPU(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Combinar datos de RAM y CPU en una respuesta JSON
	responseData := fmt.Sprintf(`{"ram": %s, "cpu": %s}`, historicoRAM, historicoCPU)

	// Escribir el JSON en la respuesta HTTP
	fmt.Fprintf(w, responseData)
}

func getObtenerProcesosPadre(w http.ResponseWriter, r *http.Request) {

	// Configurar el encabezado para indicar que el contenido es JSON
	w.Header().Set("Content-Type", "application/json")
	// Permitir solicitudes desde cualquier origen
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Obtener información de la CPU
	cpuCmd := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	cpuOut, cpuErr := cpuCmd.CombinedOutput()
	if cpuErr != nil {
		fmt.Println(cpuErr)
		http.Error(w, "Error al obtener la información de la CPU", http.StatusInternalServerError)
		return
	}

	// Convertir la cadena JSON en la estructura Go
	var cpuInfo CPUInfo
	err := json.Unmarshal([]byte(cpuOut), &cpuInfo)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return
	}

	// Crear una estructura para almacenar toda la información
	var arbolGenealogico []ProcesoInfo

	// Iterar sobre los procesos y sus hijos
	for _, process := range cpuInfo.Procesos {

		// Agregar un nuevo padre
		nuevoPadre := ProcesoInfo{ID: process.PID, Name: process.Name}
		if !existeID(arbolGenealogico, nuevoPadre.ID) {
			agregarNuevoPadreNivelSuperior(&arbolGenealogico, nuevoPadre)
		}

		// Iterar sobre los procesos hijos
		for _, child := range process.Children {
			agregarNuevoHijo(&arbolGenealogico, child.PidPadre, ProcesoInfo{ID: child.PID, Name: child.Name})
		}
	}

	// Convertir a JSON
	jsonData, err := json.Marshal(arbolGenealogico)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}

// Función para agregar un nuevo padre al mismo nivel que el Abuelo
func agregarNuevoPadreNivelSuperior(arbol *[]ProcesoInfo, nuevoPadre ProcesoInfo) {
	*arbol = append(*arbol, nuevoPadre)
}

// Función para agregar un nuevo hijo a la persona con el ID especificado
func agregarNuevoHijo(arbol *[]ProcesoInfo, idPadre int, nuevoHijo ProcesoInfo) {
	for i := range *arbol {
		if (*arbol)[i].ID == idPadre {
			(*arbol)[i].Children = append((*arbol)[i].Children, nuevoHijo)
			return
		}
		if len((*arbol)[i].Children) > 0 {
			agregarNuevoHijo(&((*arbol)[i].Children), idPadre, nuevoHijo)
		}
	}
}

// Función para verificar si un ID existe en el árbol genealógico
func existeID(arbol []ProcesoInfo, id int) bool {
	for _, persona := range arbol {
		if persona.ID == id {
			return true
		}
		if len(persona.Children) > 0 && existeID(persona.Children, id) {
			return true
		}
	}
	return false
}
