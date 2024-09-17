package main

import "fmt"

func main() {
	codes := map[string]string{"a": "1", "b": "2", "c": "3"}
	for key, value := range codes {
		delete(codes, key)
		delete(codes, value)
	}
	fmt.Println(codes)
}
