package main

import (
	"bytes"
	"fmt"
	"strings"
)

var MorseAlphabet = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
	"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
	"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
	"-.--", "--.."}

func decipher(msg string) string {
	var b bytes.Buffer
F:
	for _, w := range strings.Fields(msg) {
		for i, m := range MorseAlphabet {
			if w == m {
				b.WriteByte(byte('A' + i))
				continue F
			}
		}
		b.WriteString(w)
	}
	return b.String()
}

func main() {
	msg := "- . .- -- ---"
	fmt.Println(decipher(msg))
}
