package main

import "fmt"

const (
	A float64 = 2 * iota
	B
	C int = 10
	D
)


func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

}
