package set1

import (
	"encoding/hex"
	"fmt"
)

const errorLength = "byte slices must have equal length"

func ProduceXOR(input1 string, input2 string) (string, error) {
	// decode the strings to bytes
	d1 := decodeString(input1)
	d2 := decodeString(input2)

	// validate length
	if len(d1) != len(d2) {
		return "", fmt.Errorf(errorLength)
	}

	// create a result byte array
	result := make([]byte, len(d1))

	// loop through each element
	for i := range d1 {
		// bitwise XOR operation
		// A ^ B = C and C ^ B = A (you can reverse the operation)
		// specifically B can represent a key
		result[i] = d1[i] ^ d2[i]
	}

	// encode it back to a string to return
	encString := hex.EncodeToString(result)
	return encString, nil
}
