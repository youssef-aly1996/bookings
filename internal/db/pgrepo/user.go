package pgrepo

import (
	"context"
	"errors"
	"time"

	"github.com/youssef-aly1996/bookings/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)

const (
	query = `select * from user where id = $1`
	insertUser = `insert into users (first_name, last_name, email, password, access_level) 
	values ($1,$2,$3,$4,$5) returning id`
	deleteUser = `delete from users where id = $1`
	authen = `select id, password from users where email = $1`
)
func (pgr *PgRepo) GetUserById(id int) (user.User, error) {
	var user user.User
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	row, err := pgr.DbPool.Query(ctx, query, id)
	if err != nil {
		return user, err
	}
	err = row.Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (pgr *PgRepo) AddUser(u user.User) (int, error) {
	var id int
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	err := pgr.DbPool.QueryRow(ctx, insertUser, 
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.AccessLevel,
	).Scan(&id)
	
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (pgr *PgRepo) DeleteUserById(id int) (user.User, error) {
	var user user.User
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	row, err := pgr.DbPool.Query(ctx, deleteUser, id)
	if err != nil {
		return user, err
	}
	err = row.Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (pgr *PgRepo) Authenticate(email, password string) (int, string,error) {
	var id int
	var hashedPass string
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	row := pgr.DbPool.QueryRow(ctx, authen, email)
	
	err := row.Scan(&id, &hashedPass)
	if err != nil {
		return 0, "email not found", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err !=nil {
		return 0, "", err
	}
	return id, hashedPass, nil
}