package main

import "fmt"

func main() {
	codes := map[string]int{"en": 1, "fr": 2, "hi": 3}
	codes["it"] = 4
	fmt.Println(codes)
}
