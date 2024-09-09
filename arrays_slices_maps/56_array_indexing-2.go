package main

import "fmt"

func main() {
	grades := [...]int{3, 5, 7, 1, 6, 8}
	fmt.Println(grades)

	grades[2] = 100
	fmt.Println(grades)
}
