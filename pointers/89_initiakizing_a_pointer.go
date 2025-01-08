package main

import "fmt"

func main() {
	x := "hello"
	a := &x
	fmt.Println(a)
	b := &x
	fmt.Println(b)
}
