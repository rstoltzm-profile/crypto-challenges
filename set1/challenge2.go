package set1

import "fmt"

func ProduceXOR(input1 string, input2 string) (string, error) {
	d1 := decodeString(input1)
	d2 := decodeString(input1)
	if len(d1) != len(d2) {
		return "", fmt.Errorf("byte slices must have equal length")
	}

	return "", nil
}
