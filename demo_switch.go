package main

import "fmt"

func main() {
	// var n = 2

	switch n := 4; n {
	case 1:
		fmt.Println("i'm 1")
	case 2:
		fmt.Println("i'm 2")
	case 3:
		fmt.Println("i'm 3")
	case 4:
		fmt.Println("i'm 4")
		// break
		fmt.Println("break ...")
		fallthrough
	case 5:
		fmt.Println("i'm 5")
		fallthrough
	default:
		fmt.Println("i'm other")
	}
}
