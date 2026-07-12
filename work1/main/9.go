package main

import (
	"fmt"
)

// 一个不断生成整数的协程
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in chan int, out chan int, prime int) {
	for {
		//从通道in中读数据
		num := <-in
		// 过滤掉能被prime整除的数
		if num%prime != 0 {
			// 将不能被prime整除的数写入通道out
			out <- num
		}
	}
}

func main() {
	// 创建一个无缓冲的通道ch
	ch := make(chan int)
	//创建协程
	go generate(ch)
	for i := 0; i < 6; i++ {
		//从通道ch中读数据
		prime := <-ch
		fmt.Printf("prime:%d\n", prime)
		//创建一个无缓冲的通道out
		out := make(chan int)
		//创建协程
		go filter(ch, out, prime)
		// 将ch覆盖为out
		ch = out
	}
}

//1.功能：生成素数
//2.利用了go协程和通道的特性，通过通信来共享数据
//3.由于存在协程之间的调度，所以性能会弱于普通写法
