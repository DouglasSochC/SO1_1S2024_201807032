package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func initRoutes() {

	http.HandleFunc("/", index)

	http.HandleFunc("/monitoreo-tiempo-real", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			getMonitoreoTiempoReal(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/monitoreo-historico", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			getObtenerHistoricos(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/procesos-actuales", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			getObtenerProcesosPadres(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/arbol-proceso/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:

			// Obtener la parte final de la URL (después de "/arbol-proceso/")
			idStr := strings.TrimPrefix(r.URL.Path, "/arbol-proceso/")

			// Verificar si se proporcionó el parámetro ID
			if idStr == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "ID is required")
				return
			}

			// Convertir la cadena ID a un entero
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid ID format")
				return
			}

			getObtenerArbolDeProceso(w, r, id)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/ver-procesos-segun-modulo", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			getObtenerProcesosSegunModulo(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/ver-procesos-generales", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			getObtenerProcesosGenerales(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})
}
