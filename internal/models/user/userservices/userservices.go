package userservices

import (
	"github.com/youssef-aly1996/bookings/internal/models/user"
	"github.com/youssef-aly1996/bookings/internal/repository/postgresdriver"
)

var (
	pgrepo = postgresdriver.NewPostgresRepo()
	us     = user.NewServiceStore(pgrepo)
)

func GetAllUsers() bool {
	return us.AllUsers()
}
