package main

import "fmt"

func test() (int, int) {
	return 80, 85
}

func main() {
	_, v := test() // Ignore the 1st value of the function
	fmt.Println(v)
}
