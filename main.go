package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidInput := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidInput {

			bookTickets(userTickets, firstName, lastName, email)

			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()

			fmt.Printf("These are all our bookings : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is bookes out!!!. Come back next year 😃")
				break
			} else {
				fmt.Printf("%v tickets remining for %v\n\n", remainingTickets, conferenceName)

			}

		} else {
			if !isValidName {

				fmt.Printf("Invalid first name or last name\n")
			}
			if !isValidEmail {
				fmt.Printf("Invalid email\n")
			}
			if !isValidInput {
				fmt.Printf("%v tickets left, but you entered %v\n", remainingTickets, userTickets)
			}

		}
	}

}

func greetUser() {

	fmt.Println("Welcome to", conferenceName, " BOOKING App")
	fmt.Printf("We have %v openings and %v tickets left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

}

func getFirstNames() []string {
	firstNames := []string{}
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
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation mail at %v\n", firstName, lastName, userTickets, email)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	fmt.Printf("Generating ticket")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	var ticket = fmt.Sprintf(" %v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n#########################")
	fmt.Printf("Sending ticket :\n%v \nto email address: %v\n", ticket, email)
	fmt.Printf("#########################\n")

}
