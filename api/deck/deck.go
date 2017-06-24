package deck

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/willroberts/decklyst/api/card"
)

func DecodeDeck(d string) string {
	//deckParts := strings.Split(d, "]")
	//deckName := deckParts[0]
	//deckEnc := deckParts[1]
	deckEnc := d
	deckName := "test"
	out := fmt.Sprintf("%s\n", deckName)

	deck, err := base64.StdEncoding.DecodeString(deckEnc)
	if err != nil {
		return deckName
	}

	for _, c := range strings.Split(string(deck), ",") {
		cardParts := strings.Split(c, ":")
		cardQty := toInt(cardParts[0])
		cardID := toInt(cardParts[1])
		cardName := card.GetByID(cardID).Name
		out = fmt.Sprintf("%s%dx %s\n", out, cardQty, cardName)
	}

	return out
}

func toInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
