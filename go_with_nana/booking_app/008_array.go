package main

import "fmt"

func main() {
	conferenceName := "Go Programming"
	totalTickets := 140
	remainingTickets := 100
	var bookings [50]string

	fmt.Printf("Welcome to %v language!\n", conferenceName)
	fmt.Printf("We have total of %v conference tickets and\n", totalTickets)
	fmt.Printf("we have remaining of %v conference tickets\n", remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var firstName string
	var lastName string
	var userTicketsNo int

	//ask the user information
	fmt.Println("What is your firstname:")
	fmt.Scan(&firstName)
	fmt.Println("What is your lastname:")
	fmt.Scan(&lastName)
	fmt.Println("How many tickets are you buying:")
	fmt.Scan(&userTicketsNo)

	fmt.Printf("Thank you %v %v, for booking your ticket. The number of your tickets is: %v \n", lastName, firstName, userTicketsNo)

	bookings[0] = lastName + " " + firstName
	remainingTickets = remainingTickets - userTicketsNo
	fmt.Printf("The number of remaining tickets are %v\n", remainingTickets)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value of array: %v\n", bookings[0])
	fmt.Printf("The array type: %T\n", bookings)
	fmt.Printf("The leangth of the array: %v\n", len(bookings))

}
