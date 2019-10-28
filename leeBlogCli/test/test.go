package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "| 左对 齐 dsad | 右对齐 ||"
	re := regexp.MustCompile(`\s*([^|]*)\s*\|`)
	arr := re.FindAllStringSubmatch(a, -1)
	fmt.Println(arr)
}
