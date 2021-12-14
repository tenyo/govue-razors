package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	// context context.Context
}

func (s *Server) SetRoutes() {
	apiv1 := s.Router.PathPrefix("/api/v1").Subrouter()
	apiv1.HandleFunc("/ping", PingHandler).Methods(http.MethodGet)
	apiv1.HandleFunc("/razors", RazorsListHandler).Methods(http.MethodGet)
	apiv1.HandleFunc("/razors/{id}", RazorsShowHandler).Methods(http.MethodGet)
}

// LogWriter is an http.ResponseWriter
type LogWriter struct {
	http.ResponseWriter
}

// Write log message if http response writer returns an error
func (w LogWriter) Write(p []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(p)
	if err != nil {
		log.Printf("write failed: %v", err)
	}
	return
}
