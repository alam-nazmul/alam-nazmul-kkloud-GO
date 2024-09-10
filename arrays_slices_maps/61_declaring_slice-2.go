package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 6, 9, 7, 54, 87}

	slice := arr[1:7] // Position-1 to position-6
	fmt.Println(slice)

	slice_1 := arr[0:3] // Position-0 to position-2
	fmt.Println(slice_1)

	slice_2 := arr[:4]
	fmt.Println(slice_2) // Position-0 to position-3

	slice_3 := arr[:]
	fmt.Println(slice_3) // All
}
