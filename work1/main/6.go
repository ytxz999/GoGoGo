package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	file, err := os.Create("ninenine.txt")
	if err != nil {
		fmt.Println("打开失败:", err)
	}
	defer file.Close()
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			file.WriteString(fmt.Sprintf("%d*%d=%d\t", j, i, i*j))
		}
		file.WriteString("\n")
	}
}
