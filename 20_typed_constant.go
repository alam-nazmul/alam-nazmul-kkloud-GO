package main

import "fmt"

func main() {
	const name string = "Nazmul"
	const age int64 = 32
	fmt.Printf("%v, %T \n", name, name)
	fmt.Printf("%v, %T \n", age, age)
}
