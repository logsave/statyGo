package main

import "fmt"

func main() {
	for i := 1;; {
		fmt.Println(i)
		i++
		if i > 10000 {
			break
		}
	}
}
