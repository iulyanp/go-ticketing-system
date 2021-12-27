package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickets
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scanln(&firstName)
	
	fmt.Println("Enter your last name")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email")
	fmt.Scanln(&email)

	fmt.Println("How many tickets?")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func switchExample() {
	city := "London"

	switch city {
		case "London":
			fmt.Println("You selected London!")
		case "Berlin", "New York":
			fmt.Println("You selected Berlin, New York")
		case "Paris":
			fmt.Println("You selected Paris!")
		default: 
			fmt.Println("No valid city selected!")
	}
}