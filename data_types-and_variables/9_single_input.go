package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name) // Here "&" is a pointer as a scanner
	fmt.Println("Hey there, ", name)
}
