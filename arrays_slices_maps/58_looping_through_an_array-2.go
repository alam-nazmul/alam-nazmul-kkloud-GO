package main

import "fmt"

func main() {
	grades := [...]int{45, 21, 89, 23, 12}

	for index, elements := range grades {
		fmt.Println(index, elements)
	}
}
