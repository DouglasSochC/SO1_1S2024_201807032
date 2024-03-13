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

	http.HandleFunc("/crear-proceso", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			manejadorInicioProceso(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

	})

	http.HandleFunc("/parar-proceso", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			// Leer los datos del cuerpo de la solicitud
			err := r.ParseForm()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error parsing form data")
				return
			}

			// Acceder a los datos del formulario
			pid := r.Form.Get("pid")
			manejadorPararProceso(w, r, pid)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/iniciar-proceso", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			// Leer los datos del cuerpo de la solicitud
			err := r.ParseForm()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error parsing form data")
				return
			}

			// Acceder a los datos del formulario
			pid := r.Form.Get("pid")
			manejadorIniciarProceso(w, r, pid)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

	http.HandleFunc("/matar-proceso", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			// Leer los datos del cuerpo de la solicitud
			err := r.ParseForm()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error parsing form data")
				return
			}

			// Acceder a los datos del formulario
			pid := r.Form.Get("pid")
			manejadorMatarProceso(w, r, pid)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})

}
