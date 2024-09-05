package main

import "fmt"

func main() {
	var i int = 10
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough // Though this case is matches. But "FALLTHROUGH" force to execution the next case
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("default")
	}
}
