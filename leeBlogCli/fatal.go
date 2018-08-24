package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	// CLRW 输出
	CLRW = ""
	// CLRR 输出
	CLRR = "\x1b[31;1m"
	// CLRG 输出
	CLRG = "\x1b[32;1m"
	// CLRB 输出
	CLRB = "\x1b[34;1m"
	// CLRY 输出
	CLRY = "\x1b[33;1m"
)

// exitCode 退出码
var exitCode int

// Fatal 报错并结束程序
func Fatal(info interface{}) {
	Error(info)
	os.Exit(1)
}

// Error 打印日志
func Error(info interface{}) {
	if runtime.GOOS == "windows" {
		fmt.Printf("ERR: %s\n", info)
	} else {
		fmt.Printf("%s%s\n%s", CLRR, info, "\x1b[0m")
	}
	exitCode = 1
}
