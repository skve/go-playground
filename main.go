package main

import "fmt"

const (
	a = iota
)

func main() {
	fmt.Printf("%v, %T\n", a, a)
}
