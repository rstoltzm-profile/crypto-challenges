package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBase64String(hexString string) string {
	// convert hexString to raw bytes
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}

	// Encode to String
	encString := base64.StdEncoding.EncodeToString(decoded)
	return encString
}

func HexToBase64Bytes(hexString string) []byte {
	// convert hexString to raw bytes
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}

	// calculate len of enc
	encLen := base64.StdEncoding.EncodedLen(len(decoded))

	// create a byte array of size encLen
	encByte := make([]byte, encLen)

	// encode to base64
	base64.StdEncoding.Encode(encByte, decoded)

	return encByte
}
