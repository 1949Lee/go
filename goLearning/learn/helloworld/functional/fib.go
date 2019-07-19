package main

import (
	"bufio"
	"fmt"
	"goLearning/learn/helloworld/functional/fib"
	"io"
)

func printLine(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	printLine(f)
	//fmt.Printf("f1 = %d\n", f())
	//fmt.Printf("f2 = %d\n", f())
	//fmt.Printf("f3 = %d\n", f())
	//fmt.Printf("f4 = %d\n", f())
	//fmt.Printf("f5 = %d\n", f())
	//fmt.Printf("f6 = %d\n", f())
}
