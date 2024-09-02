package main

import "fmt"

var last_name string = "Alam" // Global variable

func main() {
	first_name := "Nazmul" // Local variable
	fmt.Println(first_name)
	fmt.Println(last_name)
}

/* Global variable can call from local variable.
But local variable can't call from global variable.
*/
