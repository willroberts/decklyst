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
	Name        string `json:"name,omitempty"`
	FactionID   int    `json:"factionId,omitempty"`
	Faction     string `json:"faction,omitempty"`
	SetName     string `json:"setName,omitempty"`
	RarityID    int    `json:"rarityId,omitempty"`
	Rarity      string `json:"rarity,omitempty"`
	SpiritCost  int    `json:"spiritCost,omitempty"`
	Mana        int    `json:"mana,omitempty"`
	Attack      int    `json:"attack,omitempty"`
	HP          int    `json:"hp,omitempty"`
	Category    string `json:"category,omitempty"` // unit, spell, etc.
	Type        string `json:"type,omitempty"`     // General, Minion, Golem, Spell, etc.
	IsGeneral   bool   `json:"isGeneral,omitempty"`
	Race        string `json:"race,omitempty"` // E.g. Arcanyst
	Description string `json:"description,omitempty"`

	// Useful for deckbuilding sites.
	SearchableContent string `json:"searchableContent,omitempty"`
	Frame             string `json:"frame,omitempty"`
	PList             string `json:"plist,omitempty"`
	Sprite            string `json:"sprite,omitempty"`

	// Debatable utility.
	//IsHidden    bool   `json:"isHidden,omitempty"`
	//FactionSlug string `json:"factionSlug,omitempty"`
	//Slug        string `json:"slug,omitempty"`
}

func (c Card) Bytes() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return b
}

func LoadCards(dataFile string) error {
	b, err := ioutil.ReadFile(dataFile)
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

	// Populate SpiritCost.
	for id, c := range allCards {
		if c.RarityID == 1 {
			c.SpiritCost = 40
			allCards[id] = c
		} else if c.RarityID == 2 {
			c.SpiritCost = 100
			allCards[id] = c
		} else if c.RarityID == 3 {
			c.SpiritCost = 350
			allCards[id] = c
		} else if c.RarityID == 4 {
			c.SpiritCost = 900
			allCards[id] = c
		}
	}

	return nil
}

func GetByID(id int) Card {
	return allCards[id]
}
