package gohipku

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHipku4(t *testing.T) {
	ip, err := Decode("The silent black ape\n" +
		"eats in the ancient grasslands.\n" +
		"Autumn colors grow.\n")
	assert.Nil(t, err)
	assert.Equal(t, net.ParseIP("193.4.5.6"), ip)
}

func TestHipku6(t *testing.T) {
	ip, err := Decode("Last sprouts and plain boys\n" +
		"pelt rushed suave torn pale stiff drakes.\n" +
		"Ripe heads drip good grooms.\n")
	assert.Nil(t, err)
	assert.Equal(t, net.ParseIP("67d1:8f17:b5a9:dbee:89d4:42a6:7154:4f69"), ip)
}

func TestErrCutShort(t *testing.T) {
	expected := "ran out of words to decode"

	_, err := Decode("The silent black ape\n")
	// https://github.com/stretchr/testify/issues/497
	assert.ErrorIs(t, err, ErrInvalidAddr)
	assert.Contains(t, err.Error(), expected)

	_, err = Decode("The silent black ape\n" +
		"eats in the ancient grasslands.\n" +
		"Autumn") // cut short; "autumn colors" is regarded as a single word
	assert.ErrorIs(t, err, ErrInvalidAddr)
	assert.Contains(t, err.Error(), expected)

	_, err = Decode("The silent black ape\n" +
		"eats in the ancient grasslands.\n" +
		"Autumn grow tall.\n") // "autumn colors" is regarded as a single word
	assert.ErrorIs(t, err, ErrInvalidAddr)
	assert.Contains(t, err.Error(), expected) // even though grow and tall are valid,
	// the last line is taken as a single PlantNouns word (with maxLen 4); error nevertheless
}

func TestErrUnknownWord(t *testing.T) {
	_, err := Decode("The silent black ape\n" +
		"eats in the turtle grasslands.\n" + // turtle
		"Autumn colors grow.\n")
	assert.ErrorIs(t, err, ErrInvalidAddr)
	assert.Contains(t, err.Error(), "is not of")

	_, err = Decode("The silent black ape\n" +
		"eats in the ancient grasslands.\n" +
		"Autumn foo bar baz grow tall.\n") // "autumn colors" is regarded as a single word
	assert.ErrorIs(t, err, ErrInvalidAddr)
	assert.Contains(t, err.Error(), "is not of")
}

func TestErrExtraWords(t *testing.T) {
	_, err := Decode("The silent black ape\n" +
		"eats in the ancient grasslands.\n" +
		"Autumn colors grow tall.\n") // tall
	assert.ErrorIs(t, err, ErrInvalidAddr)
	assert.Contains(t, err.Error(), "decoded all words, ended with")
}
