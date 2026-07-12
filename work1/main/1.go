package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
	//Go 的整数类型有：
	//int — 最常用，根据平台是 32 位或 64 位
	//int8 / int16 / int32 / int64
	//uint / uint8 / uint16 / uint32 / uint64
	//Go 语言中没有 long 类型
}
