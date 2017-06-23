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
	Faction     string `json:"faction"`
	SetName     string `json:"setName"`
	Rarity      string `json:"rarity"`
	Mana        int    `json:"mana"`
	Attack      int    `json:"attack"`
	HP          int    `json:"hp"`
	Type        string `json:"type"`           // General, Minion, Spell, etc.
	Race        string `json:"race,omitempty"` // E.g. Arcanyst
	Description string `json:"description"`

	// Debatable utility.
	//FactionID         int    `json:"factionId"`
	//RarityID          int    `json:"rarityId"`
	//Category          string `json:"category"` // unit, spell, etc.
	//SearchableContent string `json:"searchableContent"`
	//IsGeneral         bool   `json:"isGeneral"`
	//IsHidden          bool   `json:"isHidden"`
	//Frame             string `json:"frame"`
	//PList             string `json:"plist"`
	//Sprite            string `json:"sprite"`
	//Slug              string `json:"slug"`
	//FactionSlug       string `json:"factionSlug"`
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

	decoder := json.NewDecoder(buf)
	data := CardData{}
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}

	allCards = data.Cards
	return nil
}

func GetByID(id int) Card {
	return allCards[id]
}

func IDToInt() {}
