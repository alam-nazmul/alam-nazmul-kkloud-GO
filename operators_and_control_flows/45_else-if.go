package main

import "fmt"

func main() {
	fruit := "grapes"
	if fruit == "apples" {
		fmt.Println("I love apples")
	} else if fruit == "oranges" {
		fmt.Println("I love oranges ")
	} else {
		fmt.Println("no appetite")
	}
}
