package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference "
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	// Welcome message to User
	greetUsers()

	for {
		// get user information
		firstName, lastName, email, userTickets := getUserInput()

		//validate user information
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Book the conference if validation passed.
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// Print First Names Of Bookings
			firstNames := getFirstNames()
			fmt.Printf("The First Names of  the bookings are : %v \n", firstNames)

			// quit the program if tickets finished successfully.
			if remainingTickets == 0 {

				//end  Program with the message.
				fmt.Println("Our conference is booked out , come back next year!!")
				break
			}
		} else {
			// message if firstname is not valid.
			if !isValidName {
				println("First Name or Last Name you entered is too short")
			}
			// message if the last name is not valid.
			if !isValidEmail {
				println("Your email address is not valid")
			}
			// message if the ticket number is not valid.
			if !isValidTicketNumber {
				println("Check your entered tickets number")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Println("We have total of ", conferenceTickets, " tickets and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//fmt.Printf("The First Names of  the bookings are : %v \n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask data from User.
	fmt.Println("Please Enter your  First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter yor lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your Email:")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you  %v %v for booking %v Ticket you will receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("There are %v remainingTickets for %v .\n", remainingTickets, conferenceName)
}
