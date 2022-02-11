package roomrestriction

type RoomRestrictioner interface {
	InsertRoomRestriction(r RoomRestriction) error
}

type Service struct {
	r RoomRestrictioner
}

func New(r RoomRestrictioner) Service {
	return Service{r: r}
}
