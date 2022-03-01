package gohipku

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPv4(t *testing.T) {
	t.Parallel()
	ip := net.ParseIP("193.4.5.6") // lo: 0xffff00000000..0xffffffffffff
	assert.NotNil(t, ip)

	assert.Equal(t, "The silent black ape\n"+
		"eats in the ancient grasslands.\n"+
		"Autumn colors grow.\n",
		encodeIPv4(ip))
}

func TestIPv6(t *testing.T) {
	t.Parallel()
	ip := net.ParseIP("67d1:8f17:b5a9:dbee:89d4:42a6:7154:4f69")
	assert.NotNil(t, ip)

	assert.Equal(t, "Last sprouts and plain boys\n"+
		"pelt rushed suave torn pale stiff drakes.\n"+
		"Ripe heads drip good grooms.\n",
		encodeIPv6(ip))
}
