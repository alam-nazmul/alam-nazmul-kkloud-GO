package main

import "fmt"

func printdetails(student string, subjects ...string) {
	fmt.Println("Hey, ", student, "here are your subjects...")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

func main() {
	printdetails("Jeo", "Math", "Physics")
}
