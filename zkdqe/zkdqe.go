// http://en.wikipedia.org/wiki/Autokey_cipher
package main

import (
	"fmt"
	"strings"
)

func Decode(msg, key string) string {
	msg = strings.Replace(msg, " ", "", -1)
	k := len(key)
	var r []byte
	for i, char := range msg {
		r = append(r, decode(byte(char), key[i%k]))
	}
	return string(r)
}

func decode(m, key byte) byte {
	m = m - 'a'
	key = key - 'a'
	dec := key + m + 'a'
	if dec > 'z' {
		dec -= 'z' - 'a'
	}
	return dec
}

func main() {
	m := "Z KDQE YC SS ZWEM MFI DBO O VFDH AJ ICIPWG HNZC VGZA RS KC ES DYFFBJSC"
	key := "rodolfo"
	fmt.Println(Decode(m, key))
}
