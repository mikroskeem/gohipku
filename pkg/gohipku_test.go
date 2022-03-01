package gohipku

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPortabilityIPv4(t *testing.T) {
	testPortability(t, ipv4Set)
}
func TestPortabilityIPv6(t *testing.T) {
	testPortability(t, ipv6Set)
}

func testPortability(t *testing.T, hipkuIP map[string]string) {
	for ip, hipku := range hipkuIP {
		ip := net.ParseIP(ip)
		assert.NotNil(t, ip)

		h, ok := Encode(ip)
		assert.True(t, ok)
		assert.Equal(t, hipku, h)

		ip2, err := Decode(hipku)
		assert.Nil(t, err)
		assert.Equal(t, ip, ip2)
	}
}
