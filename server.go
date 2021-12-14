package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tenyo/govue-razors/api"
)

//go:embed frontend/dist
var feFS embed.FS

type frontendHandler struct {
	devmode bool
	fs      embed.FS
	path    string
}

func (fe frontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var frontend fs.FS

	if fe.devmode {
		frontend = os.DirFS(fe.path)
	} else {
		var err error
		frontend, err = fs.Sub(fe.fs, fe.path)
		if err != nil {
			log.Fatalln(err)
		}
	}

	http.FileServer(http.FS(frontend)).ServeHTTP(w, r)
}

// NewServer creates a new HTTP server and starts it
func NewServer(port int, dev bool) error {
	r := mux.NewRouter()

	api := api.Server{Router: r}
	api.SetRoutes()

	r.PathPrefix("/").Handler(frontendHandler{
		devmode: dev,
		fs:      feFS,
		path:    "frontend/dist",
	}).Methods(http.MethodGet)

	handler := handlers.RecoveryHandler()(handlers.LoggingHandler(os.Stdout, r))
	httpServer := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Starting listener on %s", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
