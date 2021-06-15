package user

import (
	"fmt"
)

// User
type User struct {
	ID     int64
	Name   string
	Active bool `json:"enable"`
}

func (u User) Metodo1() {
	fmt.Println("Metodo 1")
}

func (u *User) Metodo2() {
	fmt.Println("Metodo 2")
}
