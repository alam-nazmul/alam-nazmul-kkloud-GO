package main

import "fmt"

func main() {
	arr_1 := []int{10, 25, 23, 17}
	slice_1 := arr_1[:2]

	arr_2 := []int{45, 65, 78, 12}
	slice_2 := arr_2[:2]

	new_slice := append(slice_1, slice_2...)
	fmt.Println(new_slice)
}
