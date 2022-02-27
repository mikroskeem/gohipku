package gohipku

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"inet.af/netaddr"
)

var ipv4Set = map[string]string{
	"0.0.0.0":         "The agile beige ape\naches in the ancient canyon.\nAutumn colors blow.\n",
	"127.0.0.1":       "The hungry white ape\naches in the ancient canyon.\nAutumn colors crunch.\n",
	"82.158.98.2":     "The fearful blue newt\nwakes in the foggy desert.\nAutumn colors dance.\n",
	"255.255.255.255": "The weary white wolf\nyawns in the wind-swept wetlands.\nYellowwood leaves twist.\n",
}
var ipv6Set = map[string]string{
	"0:0:0:0:0:0:0:0":                         "Ace ants and ace ants\naid ace ace ace ace ace ants.\nAce ants aid ace ants.\n",
	"2c8f:27aa:61fd:56ec:7ebe:d03a:1f50:475f": "Cursed mobs and crazed queens\nfeel wrong gruff tired moist slow sprats.\nFaint bulls dread fond fruits.\n",
	"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff": "Young yaks and young yaks\ntend young young young young young yaks.\nYoung yaks tend young yaks.\n",
}

func TestIPv4(t *testing.T) {
	ip, err := netaddr.ParseIP("193.4.5.6") // lo: 0xffff00000000..0xffffffffffff
	assert.NotNil(t, ip)
	assert.Nil(t, err)

	assert.Equal(t, "The silent black ape\n"+
		"eats in the ancient grasslands.\n"+
		"Autumn colors grow.\n",
		encodeIPv4(ip))
}

func TestIPv6(t *testing.T) {
	ip, err := netaddr.ParseIP("67d1:8f17:b5a9:dbee:89d4:42a6:7154:4f69")
	assert.NotNil(t, ip)
	assert.Nil(t, err)

	assert.Equal(t, "Last sprouts and plain boys\n"+
		"pelt rushed suave torn pale stiff drakes.\n"+
		"Ripe heads drip good grooms.\n",
		encodeIPv6(ip))
}
