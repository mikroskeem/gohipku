package gohipku

import (
	"strings"
)

func toUpperFirst(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}
