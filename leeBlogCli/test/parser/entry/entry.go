package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 0
	strings.TrimLeftFunc("   +", func(r rune) bool {
		if r == ' ' {
			count++
		}
		return r == ' '
	})
	fmt.Println(count)
}
