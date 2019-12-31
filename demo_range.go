package main

import "fmt"

func main() {
	s := "helloWorld"
	fmt.Println(len(s))
	for _ , v := range s {
		// fmt.Printf("%d %c\n", i, v)
		fmt.Printf("%c\n", v)
	}
}
