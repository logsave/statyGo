
package main

import (
	"fmt"
)

func main () {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("=======")

	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("=======")

	for i < 20 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	// for {	}

	s := "Hello World你好!!"
	for i,v := range s{
		fmt.Printf("%d %c\n", i, v)
	}

}
