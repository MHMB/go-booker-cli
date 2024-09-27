package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, "are still available!")
	fmt.Println("Get your tickets here to attend")	

	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter you first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
