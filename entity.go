package ddd_tutorial

import "errors"

// Entity
type User struct {
	userId UserId `readonly`
	name   string
}

func NewUser(id UserId, name string) (User, error) {
	u := User{userId: id}
	if err := u.ChangeName(name); err != nil {
		return User{}, err
	}
	return u, nil
}

func (u *User) ChangeName(name string) error {
	if len(name) < 3 {
		return errors.New("first name length must be longer than 3")
	}
	return nil
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
	return NewUserId(v)
}
