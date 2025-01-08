package main

import "fmt"

func sumnumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println(sumnumbers())
	fmt.Println(sumnumbers(10))
	fmt.Println(sumnumbers(10, 20))
	fmt.Println(sumnumbers(10, 20, 30))
}
