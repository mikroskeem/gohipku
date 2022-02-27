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
