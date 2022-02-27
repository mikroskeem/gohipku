package gohipku

import (
	"net"
	"strings"

	dict "github.com/mikroskeem/gohipku/pkg/internal/dictionary"
	"inet.af/netaddr"
)

// Encodes an IPv4 or IPv6 address as hipku.
//
// scoped addressing zones on ipv6 is lost on encode (::0%ens18 ← ens18 is lost)
func Encode(stdip net.IP) (string, bool) {
	ip, ok := netaddr.FromStdIP(stdip)
	if !ok {
		return "", ok // invalid IP
	}

	if ip.Is4() {
		return encodeIPv4(ip), ok
	} else if ip.Is6() {
		return encodeIPv6(ip), ok
	}

	return "", false // invalid IP‽
}

// assumes IPv4 is already validated
func encodeIPv4(ip netaddr.IP) string {
	ipb := ip.As4()
	var factored [8]byte
	for n, b := range ipb {
		i := n * 2
		factored[i+1] = b % 16
		factored[i] = (b - factored[i+1]) / 16
	}

	var words [8]string
	for i, b := range factored {
		words[i] = dict.Hipku4[i][b]
	}

	return "The " + strings.Join(words[0:3], " ") + "\n" +
		words[3] + " in the " + strings.Join(words[4:6], " ") + ".\n" +
		capitalize(words[6]) + " " + strings.Join(words[7:8], " ") + ".\n"
}

// assumes IPv6 is already validated
func encodeIPv6(ip netaddr.IP) string {
	ipb := ip.As16()
	var words [16]string
	for i, b := range ipb {
		words[i] = dict.Hipku6[i][b]
	}

	return capitalize(words[0]) + " " + words[1] + " and " + strings.Join(words[2:4], " ") + "\n" +
		strings.Join(words[4:11], " ") + ".\n" +
		capitalize(words[11]) + " " + strings.Join(words[12:16], " ") + ".\n"
}
