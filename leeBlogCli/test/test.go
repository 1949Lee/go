package main

import (
	"fmt"
	uuid2 "github.com/google/uuid"
)

func main() {
	uuid := uuid2.New()
	fmt.Println(uuid.ID())
}
