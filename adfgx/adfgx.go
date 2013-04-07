// http://en.wikipedia.org/wiki/ADFGVX_cipher
package main

import (
	"fmt"
	"strings"
)

func Decode(code string) byte {
	table := "btalpdhozkqfvsngicuxmrewy"
	label := "ADFGX"
	row := strings.Index(label, code[0:1])
	col := strings.Index(label, code[1:2])
	return table[row*5+col]
}

func main() {
	m := "AD XF AF XA DF DF XA XF GG FD DF FD GD FX DD DF"
	list := strings.Split(m, " ")
	for _, value := range list { 
		fmt.Print(string(Decode(value)))
	}
	fmt.Println()
}
