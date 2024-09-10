package main

import "fmt"

func main() {
	arr := []int{67, 45, 23, 56, 34, 435, 12341}
	slice := arr[:3]
	fmt.Println("Before modification")
	fmt.Println(arr)
	fmt.Println(slice)

	slice[2] = 1000
	fmt.Println("After modification")
	fmt.Println(arr)
	fmt.Println(slice)
}
