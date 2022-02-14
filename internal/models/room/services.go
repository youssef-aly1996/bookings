package room

func (s Service) GetById(id int) (string, error) {
	roomName, err := s.r.GetRoomById(id)
	if err != nil {
		return "", err
	}
	return roomName, nil
}
