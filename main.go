package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/card/%d", CardHandler)
	http.Handle("/", r)
}

func CardHandler(r *http.Request, w *http.ResponseWriter) {
	return
}