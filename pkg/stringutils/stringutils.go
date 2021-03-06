// Package stringutils provides helper functions for dealing with strings.
package stringutils

import (
	"bytes"
	"math/rand"
	"strings"
)

// GenerateRandomASCIIString generates an ASCII random string with length n.
func GenerateRandomASCIIString(n int) string {
	chars := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:` "
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = chars[rand.Intn(len(chars))]
	}
	return string(res)
}

// Ellipsis truncates a string to fit within maxlen, and appends ellipsis (...).
// For maxlen of 3 and lower, no ellipsis is appended.
func Ellipsis(s string, maxlen int) string {
	r := []rune(s)
	if len(r) <= maxlen {
		return s
	}
	if maxlen <= 3 {
		return string(r[:maxlen])
	}
	return string(r[:maxlen-3]) + "..."
}

// Truncate truncates a string to maxlen.
func Truncate(s string, maxlen int) string {
	r := []rune(s)
	if len(r) <= maxlen {
		return s
	}
	return string(r[:maxlen])
}

func quote(word string, buf *bytes.Buffer) {
	// Bail out early for "simple" strings
	if word != "" && !strings.ContainsAny(word, "\\'\"`${[|&;<>()~*?! \t\n") {
		buf.WriteString(word)
		return
	}

	buf.WriteString("'")

	for i := 0; i < len(word); i++ {
		b := word[i]
		if b == '\'' {
			// Replace literal ' with a close ', a \', and an open '
			buf.WriteString("'\\''")
		} else {
			buf.WriteByte(b)
		}
	}

	buf.WriteString("'")
}

// ShellQuoteArguments takes a list of strings and escapes them so they will be
// handled right when passed as arguments to a program via a shell
func ShellQuoteArguments(args []string) string {
	var buf bytes.Buffer
	for i, arg := range args {
		if i != 0 {
			buf.WriteByte(' ')
		}
		quote(arg, &buf)
	}
	return buf.String()
}
