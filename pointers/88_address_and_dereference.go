package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%T %v \n", &x, &x)       // & is the address of operators
	fmt.Printf("%T %v \n", *(&x), *(&x)) // * is the deference of operators
}
