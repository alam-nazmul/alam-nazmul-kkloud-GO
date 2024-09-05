package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println((x < 100) && (x < 200))
	fmt.Println((x < 10) && (x < 1))
}
