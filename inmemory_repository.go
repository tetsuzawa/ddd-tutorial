package ddd_tutorial

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type InMemoryUserRepository struct {
	Store map[UserId]User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{Store: make(map[UserId]User)}
}

func (r InMemoryUserRepository) Save(user User) (err error) {
	r.Store[user.userId] = user
	return nil
}

func (r InMemoryUserRepository) Find(userName UserName) (u *User, err error) {
	for _, v := range r.Store {
		if v.name == userName {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}
