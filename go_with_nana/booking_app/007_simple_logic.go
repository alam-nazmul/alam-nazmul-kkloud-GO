package main

import "fmt"

func main() {
	conferenceName := "Go Programming"
	totalTickets := 140
	remainingTickets := 100

	fmt.Printf("Welcome to %v language!\n", conferenceName)
	fmt.Printf("We have total of %v conference tickets and\n", totalTickets)
	fmt.Printf("we have remaining of %v conference tickets\n", remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var userName string
	var userTicketsNo int

	//ask the user information
	fmt.Println("What is your username:")
	fmt.Scan(&userName)
	fmt.Println("How many tickets are you buying:")
	fmt.Scan(&userTicketsNo)

	fmt.Printf("Thank you %v, for booking your ticket. The number of your tickets is: %v \n", userName, userTicketsNo)

	remainingTickets = remainingTickets - userTicketsNo
	fmt.Printf("The number of remaining tickets are %v\n", remainingTickets)
}
