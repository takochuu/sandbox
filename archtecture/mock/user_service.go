package mock

import "github.com/takochuu/sandbox/archtecture"

type UserService struct {
	UserFn      func(id int) (*archtecture.User, error)
	UserInvoked bool
}

func (s *UserService) User(id int) (*archtecture.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}

func (s *UserService) Users() ([]*archtecture.User, error) {

}

func (s *UserService) Create(u *archtecture.User) error {

}

func (s *UserService) Delete(id int) error {

}
