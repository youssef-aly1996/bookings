package roomrestriction

func (rr Service) Insert(r RoomRestriction) error {
	err := rr.r.InsertRoomRestriction(r)
	if err != nil {
		return err
	}

	return nil
}
