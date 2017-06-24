package card

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

var (
	allCards map[int]Card
)

type CardData struct {
	Cards map[int]Card `json:"cards"`
}

type Card struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FactionID   int    `json:"factionId"`
	Faction     string `json:"faction"`
	SetName     string `json:"setName"`
	RarityID    int    `json:"rarityId"`
	Rarity      string `json:"rarity"`
	Mana        int    `json:"mana"`
	Attack      int    `json:"attack"`
	HP          int    `json:"hp"`
	Category    string `json:"category"` // unit, spell, etc.
	Type        string `json:"type"`     // General, Minion, Golem, Spell, etc.
	IsGeneral   bool   `json:"isGeneral"`
	Race        string `json:"race,omitempty"` // E.g. Arcanyst
	Description string `json:"description"`

	// Useful for deckbuilding sites.
	SearchableContent string `json:"searchableContent"`
	Frame             string `json:"frame"`
	PList             string `json:"plist"`
	Sprite            string `json:"sprite"`

	// Debatable utility.
	//IsHidden          bool   `json:"isHidden"`
	//FactionSlug string `json:"factionSlug"`
	//Slug              string `json:"slug"`
}

func (c Card) Bytes() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return b
}

func LoadCards() error {
	b, err := ioutil.ReadFile("assets/cards/v1.86.0.json")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)

	data := CardData{}
	decoder := json.NewDecoder(buf)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	allCards = data.Cards
	return nil
}

func GetByID(id int) Card {
	return allCards[id]
}
