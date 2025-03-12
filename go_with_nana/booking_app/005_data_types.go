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

	var userName string
	var userTicketsNo int

	userName = "Nazmul Alam"
	userTicketsNo = 111
	fmt.Printf("User name: %v \nTicket no: %v\n", userName, userTicketsNo)

	fmt.Printf("conference name: %T\ntotal tickets: %T\nremaining ticket: %T\nusername: %T\nuserticket no: %T\n", conferenceName, totalTickets, remainingTickets, userName, userTicketsNo)
}
