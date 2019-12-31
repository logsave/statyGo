package main

import "fmt"

func main() {
	var f1 = 1.234
	fmt.Printf("%T \n", f1)
	var f2 = float32(1.234)
	fmt.Printf("%T \n", f2)

	// 不同类型不能进行比较
}
