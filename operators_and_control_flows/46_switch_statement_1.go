package main

import "fmt"

func main() {
	var i int = 101
	switch i {
	case 10:

		fmt.Println("i is 10")
	case 20, 30:
		fmt.Println("i is 20 or 30")
	case 100, 200:
		fmt.Println("i is 100 or 200")
	default:
		fmt.Println("no matches")
	}
}
