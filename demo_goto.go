package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx
			}
			fmt.Printf("%v-%c \n", i, j)
		}
	}

	fmt.Println("xx After")
xx:{
	fmt.Println("this xx")
}
	return 
}
