package main

import "fmt"

func main() {
	codes := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(codes)
	delete(codes, "a")
	fmt.Println(codes)
}
