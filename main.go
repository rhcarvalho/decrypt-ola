package main

import (
	"fmt"
	"strings"
	"./caesar"
)

func main() {
	msg := "GVZNLLNVFULURMSL"
	for i := range make([]interface{}, 26) {
		fmt.Println(strings.Map(caesar.Shift(i), msg))
	}
}
