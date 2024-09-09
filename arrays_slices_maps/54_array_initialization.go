package main

import "fmt"

func main() {
	var fruits [2]string = [2]string{"apples", "oranges"} // The number of values should be always be less than or equal to the size of array. If they are grater, the compiler will throw us an error. If lesser the rest of the elements would be assigned the zero value.
	fmt.Println(fruits)

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	names := [...]string{"Nazmul", "Alam", "Tusher"} // Ellipses: We need to specify the size. Compiler will implicitly calculate the length of the array.
	fmt.Println(names)
}
