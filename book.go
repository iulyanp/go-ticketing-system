package main

import (
	"ticketing-system/helper"
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

func bookingWithGoRoutine() {
	firstName, lastName, email, userTickets := helper.GetUserInput()
	isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets {
		bookTicket(firstName, lastName, userTickets, email)

		wg.Add(1)
		go sendTicketWithGoroutine(firstName, lastName, userTickets, email)

		firstNames := getFirstNameBooking()
		fmt.Printf("Bookings until now: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out!")
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

	wg.Wait() // waits for all the threds to be finished before exit the main thread
}

func sendTicketWithGoroutine(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("########")
	fmt.Printf("Sending %v, to email %v \n", ticket, email)
	fmt.Println("########")

	wg.Done()
}