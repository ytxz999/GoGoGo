package main

import "fmt"

func main() {
	var slice = make([]int, 50)
	var result []int

	for i := range slice {
		slice[i] = i + 1
	}
	for i := range slice {
		if slice[i]%3 != 0 {
			result = append(result, slice[i])
		}
	}
	//原本采用slice = append(slice[:i], slice[i+1:]...)的方式删除元素。
	//问题在于遍历切片的同时删除元素。
	//range slice 在开始时就已经确定了迭代范围（0~49），
	//但每次 append(slice[:i], slice[i+1:]...) 删除元素后，切片变短了，索引就会越界。
	//解决方法：用一个新的切片来保存结果，而不是在原切片上删除：
	result = append(result, 114514)
	fmt.Print(result)
}
