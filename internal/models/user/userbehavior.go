package user

type UserStoreReader interface {
	AllUsers() bool
}

type UserStoreWriter interface {
	AllUsers() bool
}

type UserService struct {
	UserStoreReader
}

func NewServiceStore(s UserStoreReader) *UserService {
	return &UserService{s}
}
