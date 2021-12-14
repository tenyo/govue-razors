package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// PingHandler responds to ping requests
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// RazorsListHandler returns a list of all razors
func RazorsListHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}

	list, err := loadRazors()
	if err != nil {
		handleError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to load razors from file"))
		return
	}

	j, err := json.Marshal(list)
	if err != nil {
		handleError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to marshal json response"))
		return
	}

	w.Header().Set("X-Items", strconv.Itoa(len(list)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// RazorsShowHandler returns a specific razor by id
func RazorsShowHandler(w http.ResponseWriter, r *http.Request) {
	w = LogWriter{w}
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	entry, err := getRazor(id)
	if err != nil {
		handleError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to get razor"))
		return
	}
	if entry == nil {
		handleError(w, http.StatusNotFound, errors.New("razor not found"))
		return
	}

	j, err := json.Marshal(entry)
	if err != nil {
		handleError(w, http.StatusInternalServerError, errors.Wrap(err, "failed to marshal json response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func handleError(w http.ResponseWriter, errCode int, err error) {
	log.Print(err.Error())
	w.WriteHeader(errCode)
	w.Write([]byte(err.Error()))
}
