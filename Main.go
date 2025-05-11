package main

import "fmt"

// This is a ticket booking application for a conference

func main() {

	// Name of the conference is GO Conference
	conferenceName := "GO Conference"

	// Total number of seats available
	const conferenceTickets = 50

	// Remaining tickets (changes with bookings)
	var remainingTickets = 50

	// %T is a typeholder of a variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to our", conferenceName, "booking application!")
	fmt.Printf("We have a total of %v tickets and %v are available for booking.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now!")

	// Declare variables to store user info
	var firstName string
	var lastName string
	var email string
	var userticket int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userticket)

	// Update remaining tickets
	remainingTickets -= userticket

	// Confirmation message
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, userticket, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
