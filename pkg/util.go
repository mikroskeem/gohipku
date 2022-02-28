package gohipku

import (
	"strings"
)

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

// 1.17 doesn't go generics; also this code wouldn't work on 1.18
// type indexedSlice[T comparable] map[T]int

// func indexSlice[T comparable](slice []T) (index indexedSlice) {
// 	for i, x := range slice {
// 		index[x] = i
// 	}
// 	return
// }

// returns -1 if not found
// func (index *indexedSlice) getIndex[T comparable](x T) int {
// 	if i, ok := index[x]; ok {
// 		return i
// 	}
// 	return -1
// }

type indexedStringSlice map[string]int

func createStringSliceIndex(slice []string) indexedStringSlice {
	index := make(indexedStringSlice)
	for i, s := range slice {
		index[s] = i
	}
	return index
}

// returns -1 if not found
func (index indexedStringSlice) getIndex(s string) int {
	if i, ok := index[s]; ok {
		return i
	}
	return -1
}

// bool: is safe?
func safeSliceIndexGenericStr(slice *[]string, i int) (string, bool) {
	if i < 0 || i >= len(*slice) {
		return "", false
	}
	return (*slice)[i], true
}
