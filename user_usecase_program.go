package ddd_tutorial

import (
	"errors"
	"fmt"
)

type Program struct {
	userRepository IUserRepository
}

func NewProgram(u IUserRepository) Program {
	return Program{userRepository: u}
}

func (p *Program) CreateUser(userName string) error {
	name, err := NewUserName(userName)
	if err != nil {
		return err
	}
	user := NewUser("", name)
	var userService = NewUserService(p.userRepository)
	if userService.Exists(user) {
		return errors.New(fmt.Sprintf("%s already exists"))
	}
	p.userRepository.Save(*user)
	return nil
}
