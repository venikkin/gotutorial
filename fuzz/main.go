package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)

	fmt.Printf("original: %q\nreversed: %q\nreversed agin: %q\n",
		input, rev, doubleRev)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8 string")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
