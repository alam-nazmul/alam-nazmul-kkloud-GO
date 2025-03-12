package main

import "fmt"

func main() {
	var a string
	var b int
	fmt.Printf("Enter the value of a and b :")
	count_correct_inputs, error_inputs := fmt.Scanf("%s %d", &a, &b)
	fmt.Printf("count_correct_inputs: %d \n", count_correct_inputs)
	fmt.Println("error_inputs", error_inputs)
	fmt.Printf("The vaule of 1st input is: %s \n", a)
	fmt.Println("The vaule of 2nd input is: ", b)

}
