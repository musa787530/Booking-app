package main

import "fmt"

// i am conference book application
// i am a ticket booking application

func main() {

	// name of the confernce is GO Conference
	var conferenceName = "GO Conference"

	// to number seat available is 50
	const conferenceTickets = 50 // it cant change
	// every time user book a ticket, the number of tickets will be reduced by
	var remainingTickets = 50 // this can change
	fmt.Println("Welcome to our", conferenceName, "booking application!")

	fmt.Print("we have total ", conferenceTickets, " tickets and ", remainingTickets, " tickets available for booking")
	fmt.Println("Get your ticket now!")

}
