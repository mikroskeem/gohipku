package gohipku

import (
	"net"
	"strings"

	dict "github.com/mikroskeem/gohipku/pkg/internal/dictionary"
)

// Encodes an IPv4 or IPv6 address as hipku.
//
// scoped addressing zones on ipv6 is lost on encode (::0%ens18 ← ens18 is lost)
func Encode(ip net.IP) (hipku string, ok bool) {
	if ip == nil {
		return "", false
	}

	if v4 := ip.To4(); v4 != nil {
		return encodeIPv4(v4), true
	}
	return encodeIPv6(ip), true
}

// assumes IPv4 is already validated
func encodeIPv4(ip net.IP) string {
	ipb := ip
	if len(ipb) == 16 {
		ipb = ip[12:16]
	}

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
func encodeIPv6(ip net.IP) string {
	words := encodeByKey(ip, dict.Hipku6)

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
