package main

import (
	"fmt"
)

func main() {
	s1 := "人生苦短"
	s2 := []rune(s1)  // 把字符串强制转换为rune切片
	s2[0] = '一'

	fmt.Println(string(s2))  // 将rune切片强制转换为字符串

	c1 := "磊"	// string
	c2 := '磊'	// int32
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"	// string
	c4 := 'H'	// int32
	c5 := byte('H') // byte(uint8)
	fmt.Printf("c3:%T c4:%T c5:%T\n", c3, c4, c5)
	fmt.Printf("%d\n", c4)
}
