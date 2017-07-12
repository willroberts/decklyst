package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/willroberts/decklyst/api/card"
	"github.com/willroberts/decklyst/api/deck"
)

const (
	version string = "v1.87.1"
)

var (
	// Defaults.
	defaultPath string = fmt.Sprintf("assets/cards/%s.json", version)

	// Variables.
	httpPort int
	dataFile string
	logFile  string

	// Metrics.
	cardHits uint = 0
	deckHits uint = 0
)

func init() {
	flag.IntVar(&httpPort, "port", 8000, "bind to this port")
	flag.StringVar(&dataFile, "data", defaultPath, "cards json file")
	flag.StringVar(&logFile, "log", "", "where to write log output")
	flag.Parse()
}

func main() {
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal("error: failed to open log file")
		}
		defer f.Close()
		log.SetOutput(f)
	}

	if err := card.LoadCards(dataFile); err != nil {
		log.Fatal("error: failed to load cards:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/card/{id}", CardHandler)
	r.HandleFunc("/deck/{deck}", DeckHandler)

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for _ = range ticker.C {
			log.Println("Total requests to /card/ since startup:", cardHits)
			log.Println("Total requests to /deck/ since startup:", deckHits)
		}
	}()

	log.Println("Serving HTTP on port", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r))
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	cardHits++
	cardID := deck.ToInt(mux.Vars(r)["id"])
	resp := card.GetByID(cardID).Bytes()

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func DeckHandler(w http.ResponseWriter, r *http.Request) {
	deckHits++
	encodedDeck := mux.Vars(r)["deck"]
	resp := deck.DecodeDeck(encodedDeck).Bytes()

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
