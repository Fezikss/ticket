package ticket

import (
	"fmt"
	"github.com/google/uuid"
)

type Ticket struct {
	Id   uuid.UUID
	From string
	To   string
	Date string
}

func (t Ticket) NewTicket() Ticket {
	var (
		id   string
		from string
		to   string
		date string
	)
	fmt.Print("id : ")
	fmt.Scan(&id)
	fmt.Print("from : ")
	fmt.Scan(&from)
	fmt.Print("to : ")
	fmt.Scan(&to)
	fmt.Print("date : ")
	fmt.Scan(&date)

	return Ticket{
		Id:   uuid.MustParse(id),
		From: from,
		To:   to,
		Date: date,
	}

}
