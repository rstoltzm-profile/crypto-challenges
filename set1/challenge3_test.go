package set1

import (
	"fmt"
	"testing"
)

// The hex encoded string:
// ... has been XOR'd against a single character. Find the key, decrypt the message.
func TestFindTheKey(t *testing.T) {
	// just validate we don't return an error
	input1 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	got, err := DecryptMessage(input1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	want := "Cooking MC's like a pound of bacon"
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
	//fmt.Println(got)
}
