package main

import (
	"../caesar"
	"fmt"
	"strings"
)

func main() {
	msg := "GWSWGTQUGTEQOXQEGRCTCUGORTG"
	fmt.Println(strings.Map(caesar.Shift(-2), msg))
}
