package main

import "fmt"

const PI float64 = 3.14 //Global constant

func main() {
	var redius float64 = 5.0
	var area float64
	area = PI * redius * redius
	fmt.Printf("Area of circle is: %v and type of value is %T \n", area, area)
}
