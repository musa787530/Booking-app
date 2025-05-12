package helper

import "strings"

// Validate user input (exported function name)
func ValidateUserInput(firstName string, lastName string, email string, userticket int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userticket > 0 && userticket <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
