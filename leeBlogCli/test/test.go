package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := ">>> a dsada"
	re := regexp.MustCompile(`>+\s*(.+)`)
	fmt.Println(re.FindAllStringSubmatch(a, -1)[0][1])
}
