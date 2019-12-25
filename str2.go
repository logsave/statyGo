package main

import (
	"fmt"
)

func main() {
	s1 := "人生苦短"
	s2 := []rune(s1)  // 把字符串强制转换为rune切片
	s2[0] = '一'

	fmt.Println(string(s2))  // 将rune切片强制转换为字符串
}
