package main

import "fmt"

func main() {
	const name = "Nazmul"
	name := "Tusher" // You can not modify the value of this variable
	fmt.Printf("%v: %T \n", name, name)
}

/*
You can't declare constant
Not initialize it a value
You can't assign it a value later on
CONST do not accept short hand variable assignment operator
*/
