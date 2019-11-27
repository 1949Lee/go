package main

import "fmt"

func main() {
	//uuid := uuid2.New()
	//fmt.Println(uuid.ID())

	a := struct {
		name string
	}{
		name: "a",
	}

	var b *struct {
		name string
	}
	b = &a

	b.name = "b"

	fmt.Println(a)
}
