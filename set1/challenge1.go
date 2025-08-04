package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func ConvertHexToBase64(hexString string) string {
	// input hexString
	// convert to array of byte
	hexByte := []byte(hexString)
	dst := make([]byte, hex.DecodedLen(len(hexString)))
	_, err := hex.Decode(dst, hexByte)
	if err != nil {
		log.Fatal(err)
	}
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(dst)))
	base64.StdEncoding.Encode(eb, dst)
	// dst := make([]byte, hex.EncodedLen(len(decoded)))
	return base64.StdEncoding.EncodeToString(dst)
}
