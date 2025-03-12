package main

import "fmt"

func main() {
	conferenceName := "Go Programming"
	totalTickets := 140
	remainingTickets := 10

	fmt.Printf("Welcome to %v language!\n", conferenceName)
	fmt.Printf("We have total of %v conference tickets and\n", totalTickets)
	fmt.Printf("we have remaining of %v conference tickets\n", remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}
