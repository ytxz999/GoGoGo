package main

import "fmt"

func main() {
	var arr [10]int
	var standard int
	var sum int
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&standard)
	for i := range arr {
		if arr[i] <= standard+30 {
			sum++
		}
	}
	println(sum)
}
