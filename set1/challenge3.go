package set1

import (
	"fmt"
)

func DecryptMessage(input string) (string, error) {
	// hint: XOR'd from a single character
	// hint2: score character frequency for english plaintext
	// hint3: evaluate each output, choose one with best score

	// first let's convert string to raw bytes
	d := decodeString(input)

	// create some result variables,
	rFinal := make([]byte, len(d))
	var score int
	var key int

	// loop through each byte in decoded, check against a single character
	// 0 to 256 represents all possible single byte values
	// 0 to 9; A to Z; a to z; symbols
	// could do 32 to 126 for printable ASCII
	for i := 0; i < 256; i++ {
		result := make([]byte, len(d))
		var s int
		for j := range d {
			// check the character XOR'd against i
			c := d[j] ^ byte(i)
			// get the weight of the character
			s += GetCharWeight(c)
			result[j] = c
		}
		if s > score {
			key = i
			score = s
			rFinal = result
		}
		// fmt.Println(string(result), score, key)
		s = 0
	}
	fmt.Println(string(rFinal), score, string(byte(key)))
	rFinalString := string(rFinal)

	return rFinalString, nil
}

func GetCharWeight(char byte) int {
	// function to weight most common characters in english language and phrases
	// Most common english letters in order? ETAOIN SHRDLU
	wm := map[byte]int{
		byte('?'):  1,
		byte('\''): 1,
		byte('!'):  1,
		byte('.'):  1,
		byte('\n'): 1,
		byte('U'):  2,
		byte('u'):  2,
		byte('L'):  3,
		byte('l'):  3,
		byte('D'):  4,
		byte('d'):  4,
		byte('R'):  5,
		byte('r'):  5,
		byte('H'):  6,
		byte('h'):  6,
		byte('S'):  7,
		byte('s'):  7,
		byte(' '):  8,
		byte('N'):  9,
		byte('n'):  9,
		byte('I'):  10,
		byte('i'):  10,
		byte('O'):  11,
		byte('o'):  11,
		byte('A'):  12,
		byte('a'):  12,
		byte('T'):  13,
		byte('t'):  13,
		byte('E'):  14,
		byte('e'):  14,
	}
	return wm[char]
}
