package main

import (
	"fmt"
	"time"

	//"strconv"

	//"strings"

	// Import the helper package
	"LANG/helper" // Correct import path
)

var conferenceName = "GO Conference"

const conferenceTickets = 50

var remainingTickets = 50
var bookings = make([]Userdata, 0)

type Userdata struct {
	firstName string
	lastName  string
	email     string
	tickets   int
}

func main() {
	greetUser()

	firstName, lastName, email, userticket := getUserInput()

	// Correct function name here
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userticket, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(firstName, lastName, email, userticket)
		go sendTicket(userticket, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("First names of all bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out! Come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid name. Please enter at least 2 characters for first and last name.")
		}
		if !isValidEmail {
			fmt.Println("Invalid email. It should contain '@'.")
		}
		if !isValidTicketNumber {
			fmt.Printf("Invalid ticket number. Please enter a number between 1 and %v.\n", remainingTickets)
		}
	}

}

func greetUser() {
	fmt.Println("========= GO Conference Booking =========")
	fmt.Printf("Welcome to the %v booking app!\n", conferenceName)
	fmt.Println("Book your seats before they run out!")
	fmt.Println("=========================================")
	fmt.Printf("We have a total of %v tickets and %v are available for booking.\n", conferenceTickets, remainingTickets)
}

func getUserInput() (string, string, string, int) {
	var firstName, lastName, email string
	var userticket int

	fmt.Println("\nEnter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userticket)

	return firstName, lastName, email, userticket
}

func bookTickets(firstName string, lastName string, email string, userticket int) {
	remainingTickets -= userticket

	// create a map to store user data
	var userdata = Userdata{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userticket,
	}

	bookings = append(bookings, userdata)
	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userticket, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

// maps are used to store data in key-value pairs
// maps are unordered collections of data
// maps are reference types
// maps are created using the make function
// maps are not safe for concurrent use

// struct data type is used to create a custom data type
// struct is a collection of fields
// struct is a value type
// struct is created using the type keyword
// struct is used to create a custom data type

// Concurrency /Go routines
// Concurrency is the ability to run multiple tasks at the same time
// Go routines are lightweight threads managed by the Go runtime

// Go routines are created using the go keyword
// Go routines are used to run functions concurrently
// Go routines are used to run functions in the background

func sendTicket(userTicket int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Simulate a delay for sending the ticket
	ticket := fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n  %v\n  to email: %v\n...\n", ticket, email)
	fmt.Println("#########################")
}
