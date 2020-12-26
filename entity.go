package ddd_tutorial

import (
	"errors"
	"github.com/google/uuid"
)

// Entity
type User struct {
	userId UserId `readonly`
	name   UserName
}

func NewUser(name UserName) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &User{userId: NewUserId(id.String()), name: name}, nil
}

func (u *User) Equal(other *User) bool {
	if u == nil {
		return false
	}
	if other == nil {
		return false
	}
	return u.userId == other.userId
}

type UserId string

func NewUserId(v string) UserId {
	return UserId(v)
}

type UserName string

func NewUserName(v string) (UserName, error) {
	var u UserName
	if err := u.ChangeName(v); err != nil {
		return "", err
	}
	return u, nil
}

func (u *UserName) ChangeName(name string) error {
	if len(name) < 3 {
		return errors.New("first name length must be longer than 3")
	}
	*u = UserName(name)
	return nil
}
