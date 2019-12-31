package main 

import (
	"fmt"
)

func main() {
	age := 16

	if age > 18 {
		fmt.Println("Hello World!")
	} else {
		fmt.Println("~.~")
	}

	if age > 35 {
		fmt.Println("old ...")
	} else if age > 18 {
		fmt.Println("young")
	} else {
		fmt.Println("child hahha")
	}


	// 变量 a 只在 if 条件判断语句中生效
	if a := 10; a > 10 {
		fmt.Println("＞ 10")
	} else {
		fmt.Println("< 10")
	}

	// fmt.Println(a) // 变量a 未定义
}
