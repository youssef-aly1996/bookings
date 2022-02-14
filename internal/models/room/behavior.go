package room

type Roomer interface {
	GetRoomById(id int) (string, error)
}

type Service struct {
	r Roomer
}

func New(r Roomer) Service {
	return Service{
		r: r,
	}
}
