package main

import (
	"fmt"
	"strconv"
)

func main() {
	string := "200"
	integer, error := strconv.Atoi(string)
	fmt.Printf("%v, %T \n", integer, integer)
	fmt.Printf("%v, %T", error, error)
}

/*
Atoi()
	converts string to integer
	returns two values - the corresponding integer, error (if any)
*/
