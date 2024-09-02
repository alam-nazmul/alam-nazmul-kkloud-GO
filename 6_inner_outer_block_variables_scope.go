package main

import "fmt"

func main() {
	city := "Dhaka"
	{
		country := "Bangladesh"
		fmt.Println(country)
		fmt.Println(city)
	}
	fmt.Println(city)
	//fmt.Println(country)		// Outer block cannot access variables declared within innner blocks
}
