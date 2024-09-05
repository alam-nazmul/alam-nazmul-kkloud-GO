package main

import "fmt"

func main() {
	for x := 1; x <= 5; x++ {
		if x == 3 {
			break // when x get 3, then the loop will be break and output will be the previous state
		}
		fmt.Println(x)
	}
}
