package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySQL
)

// Configuración de la conexión a la base de datos
func SetupDB() (*sql.DB, error) {

	// Obtener variables de entorno
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	// Construir la cadena de conexión
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, "so1_proyecto1")

	// Intentar conectar a la base de datos
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Función para realizar el registro de un historico de la RAM
func insertHistoricoRAM(db *sql.DB, ram_total, ram_libre, ram_utilizada, porcentaje_utilizacion int) error {

	insertQuery := "INSERT INTO HISTORICO_RAM (RAM_TOTAL, RAM_LIBRE, RAM_UTILIZADA, PORCENTAJE_UTILIZACION) VALUES (?, ?, ?, ?)"
	insertStmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(ram_total, ram_libre, ram_utilizada, porcentaje_utilizacion)
	if err != nil {
		return err
	}

	return nil
}

// Función para realizar el registro de un historico de la CPU
func insertHistoricoCPU(db *sql.DB, cpu_total, porcentaje_utilizacion int) error {
	insertQuery := "INSERT INTO HISTORICO_CPU (CPU_TOTAL, PORCENTAJE_UTILIZACION) VALUES (?, ?)"
	insertStmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(cpu_total, porcentaje_utilizacion)
	if err != nil {
		return err
	}

	return nil
}

// Función para obtener los ultimos 20 registros del historico de la RAM
func selectHistoricoRAM(db *sql.DB) (*sql.Rows, error) {

	selectQuery := "SELECT DATE_FORMAT(hr.FECHA, '%d/%m/%Y %H:%i:%s') AS FECHA_FORMATEADA, hr.PORCENTAJE_UTILIZACION FROM HISTORICO_RAM hr ORDER BY hr.FECHA DESC LIMIT 20"
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Función para obtener los ultimos 20 registros del historico de la CPU
func selectHistoricoCPU(db *sql.DB) (*sql.Rows, error) {

	selectQuery := "SELECT DATE_FORMAT(hc.FECHA, '%d/%m/%Y %H:%i:%s') AS FECHA_FORMATEADA, hc.PORCENTAJE_UTILIZACION FROM HISTORICO_CPU hc ORDER BY hc.FECHA DESC LIMIT 20"
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
