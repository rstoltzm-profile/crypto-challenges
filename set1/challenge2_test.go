package set1

import "testing"

// Write a function that takes two equal-length buffers and produces their XOR combination.
func TestProduceXOR(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got := ProduceXOR(input1, input2)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
