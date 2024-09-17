package main

import "fmt"

func main() {
	codes := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range codes {
		fmt.Println(key, "==", value)
	}
}
