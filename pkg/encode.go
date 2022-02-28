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

	words := encodeByKey(factored[:], dict.Hipku4)

	words[0] = capitalize(words[0])
	words[9] = capitalize(words[9])

	return strings.Join(words[0:4], " ") + "\n" +
		strings.Join(words[4:9], " ") + ".\n" +
		strings.Join(words[9:], " ") + ".\n"
}

// assumes IPv6 is already validated
func encodeIPv6(ip netaddr.IP) string {
	ipb := ip.As16()

	words := encodeByKey(ipb[:], dict.Hipku6)

	words[0] = capitalize(words[0])
	words[12] = capitalize(words[12])

	return strings.Join(words[0:5], " ") + "\n" +
		strings.Join(words[5:12], " ") + ".\n" +
		strings.Join(words[12:], " ") + ".\n"
}

func encodeByKey(ipb []byte, key []dict.DictObj) (words []string) {
	var ipbPos int
	for _, t := range key {
		if t.Dict == nil {
			words = append(words, t.MapName)
			continue
		}
		words = append(words, t.Dict[ipb[ipbPos]])
		ipbPos++
	}
	return
}
