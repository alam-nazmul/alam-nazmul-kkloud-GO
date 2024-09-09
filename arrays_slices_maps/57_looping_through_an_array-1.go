package main

import "fmt"

func main() {
	grades := [...]int{80, 68, 87, 45, 87}

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
}
