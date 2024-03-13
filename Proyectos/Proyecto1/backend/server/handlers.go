package server

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
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
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Children []ProcesoInfo `json:"children"`
}

var process *exec.Cmd

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

func getObtenerProcesosPadres(w http.ResponseWriter, r *http.Request) {

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
		nuevoPadre := ProcesoInfo{ID: process.PID, Name: strconv.Itoa(process.PID) + " - " + process.Name}
		if !existeID(arbolGenealogico, nuevoPadre.ID) {
			agregarNuevoPadreNivelSuperior(&arbolGenealogico, nuevoPadre)
		}

		// Iterar sobre los procesos hijos
		for _, child := range process.Children {
			agregarNuevoHijo(&arbolGenealogico, child.PidPadre, ProcesoInfo{ID: child.PID, Name: strconv.Itoa(child.PID) + " - " + child.Name})
		}
	}

	// Obtener nombres e IDs
	resultado := obtenerNombresIDs(arbolGenealogico)

	// Convertir a JSON e imprimir
	jsonData, err := json.Marshal(resultado)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}

func getObtenerProcesosGenerales(w http.ResponseWriter, r *http.Request) {

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

func getObtenerProcesosSegunModulo(w http.ResponseWriter, r *http.Request) {

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

	// Convertir a JSON
	jsonData, err := json.Marshal(cpuInfo)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}

func getObtenerArbolDeProceso(w http.ResponseWriter, r *http.Request, identificador int) {

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
		nuevoPadre := ProcesoInfo{ID: process.PID, Name: strconv.Itoa(process.PID) + " - " + process.Name}
		if !existeID(arbolGenealogico, nuevoPadre.ID) {
			agregarNuevoPadreNivelSuperior(&arbolGenealogico, nuevoPadre)
		}

		// Iterar sobre los procesos hijos
		for _, child := range process.Children {
			agregarNuevoHijo(&arbolGenealogico, child.PidPadre, ProcesoInfo{ID: child.PID, Name: strconv.Itoa(child.PID) + " - " + child.Name})
		}
	}

	// Buscar por ID
	resultado := buscarPorID(arbolGenealogico, identificador)

	// Convertir a JSON e imprimir
	jsonData, err := json.Marshal(resultado)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Fprintf(w, string(jsonData))
}

func manejadorInicioProceso(w http.ResponseWriter, r *http.Request) {

	// Crear un nuevo proceso con un comando de espera
	cmd := exec.Command("bash", "-c", "while true; do echo 'Ejecutando proceso...'; done")
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error al iniciar el proceso", http.StatusInternalServerError)
		return
	}

	// Obtener el comando con PID
	process = cmd

	// Respuesta
	fmt.Fprintf(w, strconv.Itoa(process.Process.Pid))
}

func manejadorPararProceso(w http.ResponseWriter, r *http.Request, pidStr string) {

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGSTOP al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGSTOP", pidStr)
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al detener el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d detenido", pid)
}

func manejadorIniciarProceso(w http.ResponseWriter, r *http.Request, pidStr string) {

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al reanudar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d reanudado", pid)
}

func manejadorMatarProceso(w http.ResponseWriter, r *http.Request, pidStr string) {

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-9", pidStr)
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d ha terminado", pid)
}

// Función para agregar un nuevo padre al mismo nivel que el Abuelo
func agregarNuevoPadreNivelSuperior(arbol *[]ProcesoInfo, nuevoPadre ProcesoInfo) {
	*arbol = append(*arbol, nuevoPadre)
}

// Función para agregar un nuevo hijo a la proceso con el ID especificado
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
	for _, proceso := range arbol {
		if proceso.ID == id {
			return true
		}
		if len(proceso.Children) > 0 && existeID(proceso.Children, id) {
			return true
		}
	}
	return false
}

// Función para obtener nombres e IDs
func obtenerNombresIDs(arbol []ProcesoInfo) []map[string]interface{} {
	var resultado []map[string]interface{}

	for _, proceso := range arbol {
		// Crear un mapa para almacenar el par key-value
		info := make(map[string]interface{})
		info["value"] = proceso.ID
		info["label"] = proceso.Name

		// Agregar el mapa al resultado
		resultado = append(resultado, info)

		// Si hay hijos, realizar la llamada recursiva
		if len(proceso.Children) > 0 {
			hijos := obtenerNombresIDs(proceso.Children)
			resultado = append(resultado, hijos...)
		}
	}

	return resultado
}

// Función para buscar por ID y retornar toda la información de la proceso
func buscarPorID(arbol []ProcesoInfo, id int) *ProcesoInfo {
	for _, proceso := range arbol {
		if proceso.ID == id {
			return &proceso
		}

		// Si hay hijos, realizar la llamada recursiva
		if len(proceso.Children) > 0 {
			if resultado := buscarPorID(proceso.Children, id); resultado != nil {
				return resultado
			}
		}
	}

	// Si no se encuentra, retornar nil
	return nil
}
