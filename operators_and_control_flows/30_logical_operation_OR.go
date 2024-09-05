package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println((x > 6) || (x < 10))
	fmt.Println((x > 10) || (x > 20))
}
