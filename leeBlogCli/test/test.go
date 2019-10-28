package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "左对 齐 dsad :+1+:"
	re := regexp.MustCompile(`:-(\d+)-:|:\+(\d)+\+:`)
	arr := re.FindAllStringSubmatch(a, -1)
	fmt.Println(arr)
}
