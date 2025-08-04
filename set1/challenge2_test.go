package set1

import (
	"fmt"
	"testing"
)

// Write a function that takes two equal-length buffers and produces their XOR combination.
func TestProduceXOR(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got, err := ProduceXOR(input1, input2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestProduceXORBadSize(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c2773206579"
	want := errorLength
	_, got := ProduceXOR(input1, input2)
	if got == nil {
		fmt.Println("Expected error")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
