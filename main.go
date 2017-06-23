package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	cards map[int]Card
)

func main() {
	var err error
	cards, err = loadCards()
	if err != nil {
		log.Fatal("error: failed to load cards:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/card/{id}", CardHandler)
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("error serving http:", err)
	}
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := cards[toInt(vars["id"])].Bytes()

	w.Write(data)
	w.Write([]byte("\n"))
}

func toInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
