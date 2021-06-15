package user

import (
	"errors"
	"time"
)

// UserService
type UserService interface {
	Validate(u User) error
	Create(u *User) error
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) Validate(u User) error {
	if len(u.Name) <= 2 {
		return errors.New("Invalid user")
	}

	return nil
}

func (us *userService) Create(u *User) error {
	if err := us.Validate(*u); err != nil {
		return err
	}

	u.ID = time.Now().UnixNano()

	return nil
}
