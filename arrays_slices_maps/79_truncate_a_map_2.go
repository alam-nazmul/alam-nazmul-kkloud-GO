package main

import "fmt"

func main() {
	codes := map[string]string{"a": "1", "b": "2", "c": "3"}
	codes = make(map[string]string)
	fmt.Println(codes)
}
