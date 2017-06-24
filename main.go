package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willroberts/decklyst/api/card"
	"github.com/willroberts/decklyst/api/deck"
)

var (
	httpPort int
)

func init() {
	flag.IntVar(&httpPort, "port", 8000, "bind to this port")
	flag.Parse()
}

func main() {
	if err := card.LoadCards(); err != nil {
		log.Fatal("error: failed to load cards:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/card/{id}", CardHandler)
	r.HandleFunc("/deck/{deck}", DeckHandler)

	log.Println("Serving HTTP on port", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r))
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	cardID := deck.ToInt(mux.Vars(r)["id"])
	resp := card.GetByID(cardID).Bytes()

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func DeckHandler(w http.ResponseWriter, r *http.Request) {
	encodedDeck := mux.Vars(r)["deck"]
	resp := deck.DecodeDeck(encodedDeck).Bytes()

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
