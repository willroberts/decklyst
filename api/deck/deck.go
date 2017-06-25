package deck

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"github.com/willroberts/decklyst/api/card"
)

const (
	manaDenominator float64 = 39.0
)

type Deck struct {
	Faction         string
	General         string
	SpiritCost      int
	AverageManaCost float64
	Cards           []CardRepr
}

type CardRepr struct {
	ID    int
	Name  string
	Count int
}

// DecodeDeck assumes the name has not been included in the deck (e.g. [FOO]).
func DecodeDeck(d string) Deck {
	deck := Deck{}
	deck.Cards = make([]CardRepr, 0)

	csv, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return deck
	}

	spiritCost := 0
	totalManaCost := 0

	fields := strings.Split(string(csv), ",")
	for _, c := range fields {
		parts := strings.Split(c, ":")
		cardQty := ToInt(parts[0])
		cardID := ToInt(parts[1])
		card := card.GetByID(cardID)

		spiritCost += card.SpiritCost
		totalManaCost += card.Mana

		if card.IsGeneral {
			deck.General = card.Name
			deck.Faction = card.Faction
		} else {
			r := CardRepr{
				ID:    cardID,
				Name:  card.Name,
				Count: cardQty,
			}
			deck.Cards = append(deck.Cards, r)
		}
	}

	deck.SpiritCost = spiritCost
	avgManaCost := float64(totalManaCost) / manaDenominator
	deck.AverageManaCost = math.Trunc(10*avgManaCost) / 10 // Retain one digit.
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
