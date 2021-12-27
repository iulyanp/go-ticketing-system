package main

import (
	"ticketing-system/helper"
	"fmt"
	"time"
)

var conferanceName = "Go conferance"
const tickets int = 50
var remainingTickets uint = 50
// var bookings = []string{} // slices declaration
// var bookings [50]string // slices declaration

// var bookings = make([]map[string]string, 0) // this is a map, maps could only have elements of a single data type

var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {
	// var conferanceName = "Go conferance"
	// // conferanceName := "Go conf"
	// const tickets int = 50
	// var remainingTickets uint = 50
	// var bookings = []string{}
	// // var bookings [50]string
	// // bookings := []string{}

	var routine string
	
	greetUser()

	fmt.Println("Do you want to use goroutine?")
	fmt.Scanln(&routine)

	if (routine == "yes") {
		bookingWithGoRoutine()
		return
	}

	fmt.Println("Not using goroutines")

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(firstName, lastName, userTickets, email)
			go sendTicket(firstName, lastName, userTickets, email)

			firstNames := getFirstNameBooking()
			fmt.Printf("Bookings until now: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name is to short")
			}
			if !isValidEmail {
				fmt.Println("Email is incorect")
			}
			if !isValidUserTickets {
				fmt.Printf("The number of tickets is invalid. Only %v tickets are available.\n", remainingTickets)
			}

			fmt.Printf("There are only %v tickets available. We can't sell %v tickets.\n", remainingTickets, userTickets)
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to conferance %v.\nWe have %v tickets and %v are still available.\n", conferanceName, tickets, remainingTickets)
}

func getFirstNameBooking() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets = remainingTickets - userTickets

	// var userData = make(map[string]string) // make a map here with keys as string and values as strings
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an confirmation email.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferanceName)
}

func sendTicket(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("########")
	fmt.Printf("Sending %v, to email %v \n", ticket, email)
	fmt.Println("########")
}