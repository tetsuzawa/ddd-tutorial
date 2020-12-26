package ddd_tutorial

import "log"

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s UserService) Exists(user *User) bool {
	_, err := s.userRepository.Find(user.name)
	log.Println(err)
	if err != nil {
		return false
	}
	return true
}
