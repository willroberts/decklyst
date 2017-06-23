package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/willroberts/decklyst/card"
	"github.com/willroberts/decklyst/deck"
)

func main() {
	if err := card.LoadCards(); err != nil {
		log.Fatal("error: failed to load cards:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/card/{id}", CardHandler)
	r.HandleFunc("/deck/{deck}", DeckHandler)

	log.Println("Serving HTTP on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("error serving http:", err)
	}
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	cardID := toInt(mux.Vars(r)["id"])
	data := card.GetByID(cardID).Bytes()
	w.Write(data)
}

func DeckHandler(w http.ResponseWriter, r *http.Request) {
	encodedDeck := mux.Vars(r)["deck"]
	deckOut := deck.DecodeDeck(encodedDeck)
	w.Write([]byte(deckOut))
}

func toInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
