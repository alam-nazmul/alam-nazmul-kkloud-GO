package main

import (
	"fmt"
	"strconv"
)

func main() {
	integer := 45
	string := strconv.Itoa(integer)
	fmt.Printf("%q \n", string)
}

/*
Itoa()
	converts int to string
	return one value - string formed with the given integer
*/
