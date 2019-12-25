package main

import "fmt"

const (
    a1 = iota
    a2
    _
    a3 = iota
)

const (
    b1 = 1
    b2 = 2
    b3 = 3
    b4 = iota
)

const (
    c1, c2 = iota + 1, iota + 2
    c3, c4 = iota + 1, iota + 2
)

const (
    _ = iota
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
    TB = 1 << (10 * iota)
    PB = 1 << (10 * iota)
)

func main() {
    fmt.Println(a1)
    fmt.Println(a2)
    fmt.Println(a3)
    fmt.Println(b4)

    fmt.Println(c1)
    fmt.Println(c2)
    fmt.Println(c3)
    fmt.Println(c4)

}
