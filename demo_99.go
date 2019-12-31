package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println("=================")

	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
