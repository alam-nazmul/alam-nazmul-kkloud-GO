package main

import "fmt"

func main() {
	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println(codes["en"])
	fmt.Println(codes["Fr"]) // This value will not print cause no key found as "Fr" in MAP
	fmt.Println(codes["hi"])
}
