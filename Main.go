package main

import (
	"fmt"
	"strings"
)

// pakage level variables
var conferenceName = "GO Conference"

const conferenceTickets = 50

var remainingTickets = 50
var bookings []string

func main() {

	greetUser()
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {
		firstName, lastName, email, userticket := getuseriput()
		// check if user input is valid

		//isValidcity:= city == "London" || city == "New York" || city == "Berlin"
		validateUserInput(firstName, lastName, email, userticket)
		isvalidName, isvalidemails, isvalidticketNumber := validateUserInput(firstName, lastName, email, userticket)
		if isvalidName && isvalidemails && isvalidticketNumber {

			booktickets(firstName, lastName, email, userticket)

			fmt.Printf("Current bookings: %v\n", bookings)
			fmt.Printf("Number of bookings: %v\n", len(bookings))
			fmt.Printf("Slice type: %T\n", bookings)
			// print first names of all bookings
			// call function to print first names of all bookings

			firstNames := getFirstNames()
			fmt.Printf("First names of all bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out! Come back next year.")
				break
			}
		} else if userticket == remainingTickets {
			fmt.Println("You just booked the last available tickets!")
			break
		} else {
			// we are gonna user what exactly they did wrong
			if !isvalidName {
				fmt.Printf("Invalid name. Please check your first and last name.\n")
			}
			if !isvalidemails {
				fmt.Printf("Invalid email. Please check your email address.\n")
			}
			if !isvalidticketNumber {
				fmt.Printf("Invalid ticket number. Please check the number of tickets you entered.\n")
			}

			//fmt.Println("Invalid input. Please check your details.\n")
		}

	}
}

func greetUser() {
	fmt.Println("========= GO Conference Booking =========")
	fmt.Println("Welcome to the %v booking app!", conferenceName)
	fmt.Println("Book your seats before they run out!")
	fmt.Println("=========================================")
	fmt.Printf("We have a total of %v tickets and %v are available for booking.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		if len(names) > 0 {
			firstNames = append(firstNames, names[0])
		}
	}
	//fmt.Printf("First names of all bookings: %v\n", firstNames)
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userticket int) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidemails := strings.Contains(email, "@")
	isvalidticketNumber := userticket > 0 && userticket <= remainingTickets

	return isvalidName, isvalidemails, isvalidticketNumber
}

func getuseriput() (string, string, string, int) {
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

	return firstName, lastName, email, userticket
}

func booktickets(firstName string, lastName string, email string, userticket int) {
	// book tickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userticket, email)
	fmt.Printf("You have booked %v tickets.\n", userticket)
}
