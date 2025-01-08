package main

import "fmt"

func f() (int, int) {
	return 80, 85
}

func main() {
	a, b := f()
	fmt.Println(a, b)
}
