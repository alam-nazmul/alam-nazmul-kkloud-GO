package main

import (
	"fmt"
	"reflect"
)

func main() {
	grade := 42
	message := "Hello World"

	fmt.Printf("Variable grades = %v is of type %v \n", grade, reflect.TypeOf(grade))
	fmt.Printf("Variable message = %v is of type %v \n", message, reflect.TypeOf(message))
}
