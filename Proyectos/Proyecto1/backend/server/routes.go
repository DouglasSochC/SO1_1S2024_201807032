package server

import (
	"fmt"
	"net/http"
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
			getObtenerProcesosPadre(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})
}
