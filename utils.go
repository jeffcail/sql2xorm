package main

import (
	"strings"
	"unicode"
)

func compactStr(str ...string) string {
	var builder strings.Builder
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}

func toUpperCamelCase(s string) string {
	var b strings.Builder

	prevUnderscore := false
	for _, r := range s {
		if r == '_' {
			prevUnderscore = true
			continue
		}

		if prevUnderscore {
			b.WriteRune(unicode.ToUpper(r))
			prevUnderscore = false
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}
