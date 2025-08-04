package set1

import (
	"encoding/hex"
	"log"
)

func decodeString(input string) []byte {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}
