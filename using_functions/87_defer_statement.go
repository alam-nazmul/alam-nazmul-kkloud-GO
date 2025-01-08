package main

import "fmt"

func printName(name string) {
	fmt.Println(name)
}

func printRoll(roll int) {
	fmt.Println(roll)
}

func printaddress(address string) {
	fmt.Println(address)
}

func main() {
	printName("Nazmul")
	defer printRoll(10) // DEFER will execute after return surrounding all functions results
	printaddress("BD")
}
