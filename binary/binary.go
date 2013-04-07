package main

import (
	"fmt"
	"strconv"
	"strings"
)

// http://en.wikipedia.org/wiki/Bacon's_cipher

func abCodeToChar(abCode string) string {
	binaryCode := strings.Map(func(r rune) rune {
		if r == 'A' {
			return '0'
		}
		return '1'
	}, abCode)
	d, _ := strconv.ParseInt(binaryCode, 2, 0)
	char := 'A' + d
	if char > 'I' {
		char += 1
	}
	if char > 'U' {
		char += 1
	}
	return string(char)
}

func main() {
	m := "BAABAAABAAAAAAAABABBABBAB"
	for i := 0; i < len(m); i += 5 {
		fmt.Print(abCodeToChar(m[i:i+5]))
	}
	fmt.Println()
}
