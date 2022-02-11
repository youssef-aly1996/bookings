package reservation

func (s Service) Insert(res Reservation) (int, error) {
	id, err := s.r.InsertReservation(res)
	if err != nil {
		return 0, err
	}
	return id, nil
}
