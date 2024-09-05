package main

import "fmt"

func main() {
	var a, b int = 10, 21
	switch {
	case a+b == 30:
		fmt.Println("Equal to 30")
	case a+b <= 40:
		fmt.Println("less than or equal to 30")
		fallthrough
	default:
		fmt.Println("Grater than 30")
	}
}
