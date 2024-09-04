package main

import (
	"fmt"
	"strconv"
)

func main() {
	string := "200abc"
	integer, error := strconv.Atoi(string)
	fmt.Printf("%v, %T \n", integer, integer)
	fmt.Printf("%v, %T", error, error)
}
