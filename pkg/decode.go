package gohipku

import (
	"fmt"
	"net"
	"strings"

	dict "github.com/mikroskeem/gohipku/pkg/internal/dictionary"
)

var (
	ErrInvalidAddr = fmt.Errorf("invalid address") // user error
)

var indexes = generateIndexes()

func generateIndexes() map[string]indexedStringSlice {
	ixs := make(map[string]indexedStringSlice)
	for _, d := range append(dict.Hipku4, dict.Hipku6...) {
		if _, indexed := ixs[d.MapName]; !indexed {
			ixs[d.MapName] = createStringSliceIndex(d.Dict)
		}
	}
	return ixs
}

// Decodes from hipku to an IPv4 or IPv6 address.
//
// error: nil, ErrInvalidAddr
func Decode(hipku string) (net.IP, error) {
	hipku = strings.ReplaceAll(hipku, ".", "")
	hipku = strings.ToLower(hipku)
	hraw := strings.Fields(hipku)
	var h []string
	for _, x := range hraw { // 2 whitespaces would create an empty 'word'
		if x != "" {
			h = append(h, x)
		}
	}

	factors, err := deocdeToFactors(h)
	if err != nil {
		return nil, err
	}

	switch len(factors) {
	case 8: // IPv4
		var i4 []byte
		for n := 0; n < 8; n = n + 2 {
			octet := factors[n]*16 + factors[n+1]
			i4 = append(i4, byte(octet))
		}
		return net.IPv4(i4[0], i4[1], i4[2], i4[3]), nil

	case 16: // IPv6
		return net.IP(factors[0:16]), nil
	}

	return nil, fmt.Errorf("this should be impossible‽")
}

func deocdeToFactors(words []string) (factors []byte, _ error) {
	var wordsIndex int // for lookahead of words; indicates nth in var words

	var typ []dict.DictObj
	if strings.EqualFold(words[0], "The") { // IPv4?
		typ = dict.Hipku4
	} else { // IPv6?
		typ = dict.Hipku6
	}

typ:
	for i, t := range typ { // have to duplicate the whole codeblock since no generics
		if t.Dict == nil {
			wordsIndex++
			continue // non-data word
		}

		var factor int    // see: after for {}
		var buffer string // for potential lookahead, where MaxLen > 1
		for n := 1; n <= t.MaxLen; n++ {
			nword, ok := safeSliceIndexGenericStr(&words, wordsIndex)
			if !ok {
				return nil, fmt.Errorf("ran out of words to decode for %d-%q:%d/%d: %w", i, t.MapName, n, t.MaxLen, ErrInvalidAddr)
			}
			if n != 1 {
				buffer += " "
			}
			buffer += nword
			wordsIndex++

			factor = indexes[t.MapName].getIndex(buffer)
			if factor >= 0 {
				factors = append(factors, byte(factor))
				continue typ
			}
		}

		// no match found
		return nil, fmt.Errorf("hipku word %q is not of %q: %w", buffer, t, ErrInvalidAddr)
	}

	if wordsIndex != len(words) {
		return nil, fmt.Errorf("decoded all words, ended with %d extra (%q): %w", len(words)-wordsIndex, words[wordsIndex:], ErrInvalidAddr)
	}

	return factors, nil
}
