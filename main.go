package main

import (
	"fmt"

	"github.com/dobau/golearn/user"
)

func main() {
	user1 := user.User{}

	user2 := user.User{
		Name: "Rafael",
	}

	user3 := user.User{123, "Rafael", true}

	fmt.Printf("user1 = %v\n", user1)
	fmt.Printf("user2 = %v\n", user2)
	fmt.Printf("user3 = %+v\n", user3)

	user.User{}.Metodo1()
	//user.User{}.Metodo2()
	user4 := &user.User{}
	user4.Metodo1()
	user4.Metodo2()

	service := user.NewUserService()

	err := service.Validate(user1)
	if err != nil {
		fmt.Println("Erro ao validar user1")
	}

	err = service.Validate(user2)
	if err != nil {
		fmt.Println("Erro ao validar user2")
	}

	if createErr := service.Create(&user2); createErr != nil {
		fmt.Println("Erro ao criar user2")
	}
	fmt.Printf("user2 criado = %v\n", user2)

	users := []user.User{
		{
			ID:   1,
			Name: "Marcos",
		},
		{
			ID:   2,
			Name: "Carlos",
		},
		{
			ID:   3,
			Name: "Pedro",
		},
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
