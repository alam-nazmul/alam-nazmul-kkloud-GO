package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name) // Here "&" is a pointer
	fmt.Println("Hey there, ", name)
}
