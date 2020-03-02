package main

import "fmt"

// Array

// 存放元素的容器
// 必须指定存放元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T a2:%T", a1, a2)
	//fmt.Println(a1 == a2)
	var a3 [4]bool
	fmt.Println(a2 == a3)

	// 数组的初始化
	// 如果不初始化：默认元素是零值(bool:false, string:"")
	a1[0] = true
	fmt.Println(a1)

	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 根据初始值自动判断长度
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)

	//
	a4 := [5]int{1, 2}
	fmt.Println(a4)

	a5 := [5]int{0: 1, 4: 3}
	fmt.Println(a5)

	// 数组的遍历
	citys := [...]string{"BeiJing", "ShangHai", "NanJing", "Xi'An"}
	fmt.Println(citys)
	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// 2. for range
	for i, v := range(citys) {
		fmt.Println(i, v)
	}


}
