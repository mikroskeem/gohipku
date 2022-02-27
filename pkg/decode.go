package gohipku

import (
	"fmt"
	"net"
	"strings"

	dict "github.com/mikroskeem/gohipku/pkg/internal/dictionary"
	"inet.af/netaddr"
)

var (
	ErrInvalidAddr = fmt.Errorf("invalid address")
)

// func (index indexesT) index(k []string) {
// 	s := reflect.ValueOf(k).Type().Name()
// 	if _, ok := index[s]; !ok {
// 		index[s] = createStringSliceIndex(k)
// 	}
// }

type indexesT map[string]indexedStringSlice

var indexes = make(indexesT)

func init() {
	// let's give up after an hour and hope it is not kept as such for 10 years
	indexes["Adjectives"] = createStringSliceIndex(dict.Adjectives)
	indexes["Nouns"] = createStringSliceIndex(dict.Nouns)
	indexes["Verbs"] = createStringSliceIndex(dict.Verbs)
	indexes["AnimalAdjectives"] = createStringSliceIndex(dict.AnimalAdjectives)
	indexes["AnimalColors"] = createStringSliceIndex(dict.AnimalColors)
	indexes["AnimalNouns"] = createStringSliceIndex(dict.AnimalNouns)
	indexes["AnimalVerbs"] = createStringSliceIndex(dict.AnimalVerbs)
	indexes["NatureAdjectives"] = createStringSliceIndex(dict.NatureAdjectives)
	indexes["NatureNouns"] = createStringSliceIndex(dict.NatureNouns)
	indexes["PlantNouns"] = createStringSliceIndex(dict.PlantNouns)
	indexes["PlantVerbs"] = createStringSliceIndex(dict.PlantVerbs)

	// gimmeGenerics := func(index indexesT, k []string) {
	// 	s := reflect.ValueOf(k).String()
	// 	if _, ok := index[s]; !ok {
	// 		index[s] = createStringSliceIndex(k)
	// 	}
	// }
	// for i, _ := range dict.Hipku4 {
	// 	s := reflect.TypeOf(dict.Hipku4).FieldByIndex([]int{i})
	// 	fmt.Print(s)
	// 	// if _, ok := indexes[s]; !ok {
	// 	// 	indexes[s] = createStringSliceIndex(k)
	// 	// }
	// 	// indexes.index(k)
	// 	// gimmeGenerics(indexes, k)
	// }
	// for _, k := range dict.Hipku6 {
	// 	indexes.index(k)
	// 	// gimmeGenerics(indexes, k)
	// }
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
	case 8:
		var i4 []byte
		for n := 0; n < 8; n = n + 2 {
			octet := factors[n]*16 + factors[n+1]
			i4 = append(i4, byte(octet))
		}
		return net.IPv4(i4[0], i4[1], i4[2], i4[3]), nil

	case 16: // IPv6
		var f16 [16]byte
		copy(f16[:], factors[0:16])
		return netaddr.IPFrom16(f16).IPAddr().IP, nil
	}
	return nil, fmt.Errorf("this should be impossible‽")
}

func deocdeToFactors(words []string) (factors []byte, _ error) {
	// everything on from here is just plain ugly
	var inputWordIndex = -1 // for lookahead of words, -1 for 1st loop
	// can't use len(h) since the dictionary uses multiple words inside a word :(
	var typ []string
	if strings.EqualFold(words[0], "The") {
		typ = dict.Hipku4String // IPv4?
	} else {
		typ = dict.Hipku6String
	}
	for i, t := range typ { // have to duplicate the whole codeblock since no generics
		inputWordIndex++
		switch t {
		default:
			input, ok := safeSliceIndexGenericStr(&words, inputWordIndex)
			if !ok {
				return nil, fmt.Errorf("ran out of words to decode for %d-%q: %w", i, t, ErrInvalidAddr)
			}
			pos := indexes[t].getIndex(input)
			if pos < 0 {
				return nil, fmt.Errorf("hipku word %q is not of %q: %w", input, t, ErrInvalidAddr)
			}
			factors = append(factors, byte(pos))

		case "": // non-data word

		case "PlantNouns":
			var pos int
			suckedInput := words[inputWordIndex]
			for l := 1; l < dict.MaxLen; l++ {
				pos = indexes[t].getIndex(suckedInput)
				if pos >= 0 {
					break // match found
				}

				next, ok := safeSliceIndexGenericStr(&words, inputWordIndex+l)
				if !ok {
					return nil, fmt.Errorf("ran out of words to decode %d-%q: %w", i, t, ErrInvalidAddr)
				}
				suckedInput += " " + next
				inputWordIndex++
			}
			if pos < 0 { // if exited before match
				return nil, fmt.Errorf("hipku word %q is not of %q: %w", suckedInput, t, ErrInvalidAddr)
			}
			factors = append(factors, byte(pos))
		}
	}
	if inputWordIndex+1 != len(words) {
		return nil, fmt.Errorf("decoded all words, but left with %d extra: %w", len(words)-inputWordIndex+1, ErrInvalidAddr)
	}

	return
}
