package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // This CONTINUE  is override the IF condition
		}
		fmt.Println(i)
	}
}
