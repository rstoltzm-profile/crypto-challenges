package set1

import (
	"fmt"
)

func ConvertHexToBase64(hex string) string {
	hexByte := []byte(hex)
	hexDecoded, err := decodeHex(hexByte)
	if err != nil {
		fmt.Printf("Error decoding Hex")
	}
	base64Encoded, err := encodeHex(hexDecoded)
	if err != nil {
		fmt.Printf("error encoding to base64")
	}
	return toStr(base64Encoded)
}

func decodeHex(input []byte) ([]byte, error) {
	return input, nil
}

func encodeHex(input []byte) ([]byte, error) {
	return input, nil
}
