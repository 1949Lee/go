package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 整数转二进制，for循环圣略初始条件
func convertToBin(n int) string {
	bin := ""
	if n == 0 {
		bin = "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		bin = strconv.Itoa(lsb) + bin
	}
	return bin
}

func printlnByFile(name string) {
	if file, err := os.Open(name); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(0),  // 0
	)

	printlnByFile("note.md")

	// 死循环，并发编程时有大用
	//for {
	//	fmt.Println("looping")
	//}

}
