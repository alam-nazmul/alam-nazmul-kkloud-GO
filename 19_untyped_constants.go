package main

import "fmt"

func main() {
	const name = "Nazmul"
	const is_okay = false
	const age = 32

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v, %T \n", is_okay, is_okay)
	fmt.Printf("%v, %T \n", age, age)
}
