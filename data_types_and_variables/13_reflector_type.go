package main

import "fmt"
import "reflect"

func main() {
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Nazmul"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(58.9))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}
