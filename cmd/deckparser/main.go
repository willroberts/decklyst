package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	// I exported a generic Arcanyst Faie deck and copied the output here.
	exportOut := "[1.84 ARCANYST FAIE]MTo1MDEsMzo1MTcsMjo1MzgsMzo1NDAsMzoxMDMwMiwyOjEwMzAzLDM6MTAzMDUsMjoxMTA5NCwzOjExMDk3LDM6MjAxMzQsMzoyMDEzOSwyOjIwMTQ0LDI6MjAxNDcsMjoyMDIwNywzOjIwMjM3LDM6MjAyNjE="

	// After the [DECK TITLE], the rest of the string is base64-encoded.
	encodedStr := strings.Split(exportOut, "]")[1]

	// Decoding the base64 reveals CSV data.
	sDec, _ := base64.StdEncoding.DecodeString(encodedStr)
	csv := string(sDec)
	fields := strings.Split(csv, ",")

	// The CSV fields have a small number (1-3 in my deck mapped to a larger number like 20134).
	// Turn this into a native map.
	deckData := make(map[string]string, len(fields))
	for _, f := range fields {
		kv := strings.Split(f, ":")
		cardQty := kv[0]
		cardID := kv[1]
		deckData[cardID] = cardQty
	}

	// Print out the result.
	for k, v := range deckData {
		fmt.Println(k, "->", v)
	}

	// What might the first number represent?
	// What could the encoding of the second number be?
}
