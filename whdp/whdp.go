package main

import (
	"fmt"
	"strings"
	"../caesar"
)

func main() {
	msg := "WHDPRRPHXIRILQKR"
	fmt.Println(strings.Map(caesar.Shift(-3), msg))
}
