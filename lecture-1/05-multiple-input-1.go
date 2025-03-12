package main

import "fmt"

func main() {
	var name string
	var is_moody bool

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s, %t", &name, &is_moody)
	fmt.Printf("%s, %t", name, is_moody)
}
