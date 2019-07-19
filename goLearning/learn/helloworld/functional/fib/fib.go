package fib

import (
	"fmt"
	"io"
	"strings"
)

type fibGen func() int

//为斐波那契数列的生成器实现一个reader接口。
func (f fibGen) Read(p []byte) (n int, err error) {
	next := f()

	//某一个数大于100万时视为读到了流的最后
	if next > 1000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func Fibonacci() fibGen {
	i := 0
	j := 1
	return func() int {
		i, j = j, i+j
		return i
	}
}
