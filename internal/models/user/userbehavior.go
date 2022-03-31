package user


type UserStore interface {
	AddUser(u User) (int, error)
	DeleteUserById(id int) (User, error)
	Authenticate(email, password string) (int, string,error)
}



type UserService struct {
	u UserStore
}

func NewServiceStore(s UserStore) *UserService {
	return &UserService{u:s}
}
