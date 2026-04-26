package main

import "fmt"

func main()	{
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)

	fmt.Println("The first value is:", x + y)
	fmt.Println("The first value is:", x - y)
}