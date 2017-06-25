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
	Faction         string      `json:"faction"`
	General         string      `json:"general"`
	SpiritCost      int         `json:"spiritCost"`
	AverageManaCost float64     `json:"averageManaCost"`
	ManaCurve       map[int]int `json:"manaCurve"`
	Cards           []CardRepr  `json:"cards"`
}

type CardRepr struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
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
	curve := map[int]int{}

	fields := strings.Split(string(csv), ",")
	for _, c := range fields {
		parts := strings.Split(c, ":")
		cardQty := ToInt(parts[0])
		cardID := ToInt(parts[1])
		card := card.GetByID(cardID)

		if card.IsGeneral {
			deck.General = card.Name
			deck.Faction = card.Faction
			continue
		}

		r := CardRepr{
			ID:    cardID,
			Name:  card.Name,
			Count: cardQty,
		}
		deck.Cards = append(deck.Cards, r)

		// Accumulate spirit cost, mana cost, and mana curve for deck.
		spiritCost += card.SpiritCost
		totalManaCost += card.Mana * cardQty
		found := false
		for k, v := range curve {
			if k == card.Mana {
				found = true
				curve[k] = v + cardQty
			}
		}
		if !found {
			curve[card.Mana] = cardQty
		}
	}

	deck.SpiritCost = spiritCost
	avgManaCost := float64(totalManaCost) / manaDenominator
	deck.AverageManaCost = math.Trunc(10*avgManaCost) / 10 // Retain one digit.
	deck.ManaCurve = curve

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
