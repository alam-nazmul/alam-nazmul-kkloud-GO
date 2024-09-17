package main

import "fmt"

func main() {
	codes := map[string]int{"a": 1, "b": 2, "c": 3}
	codes["c"] = 4
	fmt.Println(codes)
}
