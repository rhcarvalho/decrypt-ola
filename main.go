package main

import (
	"fmt"
	"strings"
)

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}

func rotN(n int) func(rune) rune {
	return func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+rune(n))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(n))%26
		}
		return r
	}
}

func main() {
	msg := "IESCAUED LYOHTHNE O OMTYA O A VUUHOVI"
	//msg := "GVZNL L NVFULURMSL"
	for i := range make([]interface{}, 26) {
		fmt.Println(strings.Map(rotN(i), msg))
	}
}
