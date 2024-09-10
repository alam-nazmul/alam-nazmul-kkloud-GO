package main

import "fmt"

func main() {
	arr := []int{5, 45, 34, 78}
	slice := arr[1:3]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 900, -100, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
