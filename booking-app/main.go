package main

import (
	"fmt"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("First Names of all the bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Booking for this year closed, Please comeback next here")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The First Name or Last Name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is InValid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainging for conference\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket\n %v\n to email address %v\n", ticket, email)
	fmt.Println("########################")
}
