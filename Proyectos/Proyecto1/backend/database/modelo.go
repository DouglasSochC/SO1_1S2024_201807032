package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySQL
)

// Configuración de la conexión a la base de datos
func SetupDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root_password@tcp(localhost:3306)/so1_proyecto1")
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

// // Función para realizar una operación de SELECT
// func selectData(db *sql.DB) error {
// 	selectQuery := "SELECT campo1, campo2 FROM ejemplos"
// 	rows, err := db.Query(selectQuery)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("Resultados de la consulta SELECT:")
// 	for rows.Next() {
// 		var campo1, campo2 string
// 		err := rows.Scan(&campo1, &campo2)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Printf("Campo1: %s, Campo2: %s\n", campo1, campo2)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }
