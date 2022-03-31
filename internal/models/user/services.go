package user

func (s UserService) Add(u User) (int, error) {
	id, err := s.u.AddUser(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s UserService) Auth(email, password string) (int, string,error) {
	id,pass, err := s.u.Authenticate(email, password)
	if err != nil {
		return 0, "", err
	}
	return id, pass, nil
}


