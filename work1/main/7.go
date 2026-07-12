package main

//切片和数组
//数组的长度更加是固定的，切片的长度更加灵活
//切片实质是指向数组的结构体，包含指向数组的指针，长度和容量
//切片的容量可以动态扩充,使用频率更高
//创建切片
//1. var numbers = make([]int,3,5) 创建一个长度为3，容量为5的切片
//2. var numbers = []int{1,2,3,4,5} 创建一个包含5个元素的切片（与数组的区别是【】没有写长度）
//创建map
//1.MyMap := make(map[string]string) 创建一个key和value都是string类型的map
//MyMap["你好"]=“再见”
//2. MyMap := map[string]string{
//    "你好": “再见”
//}
