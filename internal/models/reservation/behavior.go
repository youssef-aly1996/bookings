package reservation

type Reservationer interface {
	InsertReservation(res Reservation) (int, error)
}

type Service struct {
	r Reservationer
}

func New(r Reservationer) Service {
	return Service{r: r}
}
