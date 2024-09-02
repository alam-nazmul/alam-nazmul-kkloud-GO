package main

import "fmt"

func main() {
	var name string
	var is_mad bool
	fmt.Printf("Ener your name and are you mad: ")
	fmt.Scanf("%s %t", &name, &is_mad)
	fmt.Println(name, is_mad)
}
