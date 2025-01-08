package main

import (
	"fmt"
)

func operation(a int, b int) (sum int, diff int) { // This concept has build when a functions has to return multiple values
	sum = a + b
	diff = a - b
	return
}

func main() {
	sum, difference := operation(80, 20)
	fmt.Println(sum)
	fmt.Println(difference)
}
