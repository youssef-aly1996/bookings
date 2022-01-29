package models

//Reservation holds reservatin data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func NewReservation() *Reservation {
	return &Reservation{}
}
