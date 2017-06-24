package deck

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/willroberts/decklyst/api/card"
)

type Deck struct {
	Faction string
	General string
	Cards   map[string]int
}

// DecodeDeck assumes the name has not been included in the deck (e.g. [FOO]).
func DecodeDeck(d string) Deck {
	deck := Deck{}
	deck.Cards = make(map[string]int)

	csv, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return deck
	}

	fields := strings.Split(string(csv), ",")
	for _, c := range fields {
		parts := strings.Split(c, ":")
		cardQty := ToInt(parts[0])
		cardID := ToInt(parts[1])

		card := card.GetByID(cardID)
		if card.IsGeneral {
			deck.General = card.Name
			deck.Faction = card.Faction
		} else {
			deck.Cards[card.Name] = cardQty
		}
	}

	return deck
}

func (d Deck) Bytes() []byte {
	b, err := json.Marshal(d)
	if err != nil {
		return []byte{}
	}
	return b
}

func ToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
