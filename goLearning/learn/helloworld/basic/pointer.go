package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 3
	b := 4
	swap(&a, &b)
	fmt.Println(a, b)

	var c string = "c"
	var pc *string = &c
	*pc = "d"
	fmt.Println(*pc)

	var s = "abc"
	fmt.Println(s[0] == 'a')
}
