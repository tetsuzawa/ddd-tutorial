package ddd_tutorial

import (
	"testing"
)

func TestUseCase(t *testing.T) {
	const wantName = "tetsuzawa"

	userRepository := NewInMemoryUserRepository()
	program := NewProgram(userRepository)
	if err := program.CreateUser(wantName); err != nil {
		t.Error(err)
	}
	var head User
	for _, head = range userRepository.Store {
		break
	}
	if head.name != wantName {
		t.Errorf("head value of repository is differrent. want %v, got: %v", wantName, head.name)
	}
}
