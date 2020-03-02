package main

import "fmt"

func main() {
	// 多维数组
	// [[1, 2], [3, 4], [5, 6]]
	var all [3][2]int

	all = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(all)

	for _, v := range(all) {
		for i := 0; len(v) > i; i++ {
			fmt.Println(v[i])
		}
		fmt.Println("===")
	}

	// 通过递归调用对不确定维度的数组进行遍历


	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1

	b2[0] = 100
	fmt.Println(b1, b2)

	// 数组支持 == , != 操作
	c1 := [3]int{1,2,3}
	c2 := [3]int{2,3,4}

	fmt.Println(c1,c2)
	fmt.Println(c1 == c2)

}
