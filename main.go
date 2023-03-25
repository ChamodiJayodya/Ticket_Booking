package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	userEmail string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, userEmail, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, userEmail)
			go sendTicket(userTickets, firstName, lastName, userEmail)

			//call getfirstname function
			firstName := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstName)

			if remainingTickets == 0 {
				//
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email you enterd in not in correct format.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your first name ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)

	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		userEmail: userEmail,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(" Thank you %v %v for booking %v tickets. The confirmation email will be sent to %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending tickets : %v \n to %v address", ticket, userEmail)
	fmt.Println("########")
}