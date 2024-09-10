package main

import "fmt"

func main() {
	arr := []int{45, 56, 654, 879, 3422}
	fmt.Println(arr)

	i := 2

	slice_1 := arr[:i]
	slice_2 := arr[i+1:]
	fmt.Println(slice_1)
	fmt.Println(slice_2)

	new_slice := append(slice_1, slice_2...)
	fmt.Println(new_slice)

}
