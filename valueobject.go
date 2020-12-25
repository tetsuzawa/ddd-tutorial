package ddd_tutorial

import "errors"

// Value Object
type FullName struct {
	firstName FirstName `readonly`
	lastName  LastName  `readonly`
}

func NewFullName(firstName FirstName, lastName LastName) FullName {
	return FullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (f FullName) FirstName() FirstName {
	return f.firstName
}

func (f FullName) LastName() LastName {
	return f.lastName
}

type FirstName string

func NewFirstName(v string) (FirstName, error) {
	if len(v) < 3 {
		return "", errors.New("first name length must be longer than 3")
	}
	return FirstName(v), nil
}

type LastName string

func NewLastName(v string) (LastName, error) {
	if len(v) < 3 {
		return "", errors.New("last name length must be longer than 3")
	}
	return LastName(v), nil
}
