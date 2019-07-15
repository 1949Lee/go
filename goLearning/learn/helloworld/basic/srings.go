package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Hi我叫李佳轩!"

	//打印出是16，这个16是字符串s的字节数量。所以len(sting)会返回字节数量。
	fmt.Println(len(s))
	//字符串在go中默认采用utf-8的方式编码方式。中文三个字节，英文一个字节。所以上面这句话是3 * 5 + 1 = 16个字节

	// string的底层：下面的循环可以看出字符串在go中的utf-8编码16个字节的每个字节
	for _, b := range []byte(s) {
		fmt.Printf("%X, ", b)
	}
	fmt.Println()

	//直接遍历字符串类型，将的到每个字符对应的rune,你会发现i还是按照每个字符的字节开始分的，所以i会不连续
	for i, ch := range s {
		fmt.Printf("(%d, %c, %X) ", i, ch, ch)
	}
	fmt.Println()

	//手动按照字符遍历字符串
	str := s
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("(%c,%d)", r, size)

		str = str[size:]
	}
	fmt.Println()

	//字符串的处理：获取字符串的字符数量
	fmt.Println("s的字符数量：", utf8.RuneCountInString(s))

	//字符串的处理：按照字符处理字符串，i表示就是字符串中第i+1个字符，i是从0开始并且连续的整数。
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c, %x) ", i, ch, ch)
	}
	fmt.Println()

	//所有字符串的官方库都在strings包中。
	//以传入的字符（第二个）分割字符串为数组
	fmt.Println(len(strings.Split("1,2,3,4,   , ,", ",")))
	//以空格或者连续空格分割字符串为数组
	fmt.Println(len(strings.Fields("1,2,3,4,   , ,")))
}
