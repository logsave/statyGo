package main

import "fmt"

func main() {
	var s string = "Hello World"

	fmt.Println(s)

	fmt.Printf("%c, %T\n", s[1], s[1])
	fmt.Printf("%c, %T\n", s[2], s[2])

}
