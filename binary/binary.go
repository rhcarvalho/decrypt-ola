package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	m := "BAABAAABAAAAAAAABABBABBAB"
	m = strings.Map(func(r rune) rune {
		if r == 'A' {
			return '0'
		}
		return '1'
	}, m)
	for i := 0; i < len(m); i += 5 {
		d, _ := strconv.ParseInt(m[i:i+5], 2, 0)
		fmt.Print(string('A' + d))
	}
}
