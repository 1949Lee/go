package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok { // 如果真的是个错误
			fmt.Println("Error occurred: ", err)
		} else { // 如果不是错误。则继续panic
			//panic(r)
		}
	}()

	//panic(errors.New("This is an error!"))
	b := 0
	a := 1 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
