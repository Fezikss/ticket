package user

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Ticket_id uuid.UUID
}

func (u User) NewUser() User {
	var (
		id        string
		firstName string
		lastName  string
		email     string
		ticketId  string
	)

	fmt.Print("Id : ")
	fmt.Scan(&id)
	fmt.Print("First name : ")
	fmt.Scan(&firstName)
	fmt.Print("Last name : ")
	fmt.Scan(&lastName)
	fmt.Print("email : ")
	fmt.Scan(&email)
	fmt.Print("Ticket id : ")
	fmt.Scan(&ticketId)

	return User{
		Id:        uuid.MustParse(id),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Ticket_id: uuid.MustParse(ticketId),
	}
}
