package main

import "fmt"

func main() {
	src_slice := []int{12, 654, 78, 36, 14}
	dest_slice := make([]int, 3)

	fmt.Println(dest_slice)

	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Printf("Number of elements copied: %v", num)

}
