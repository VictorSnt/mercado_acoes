package main

import (
	"fmt"
	"mercado/acoes/configs"
	"mercado/acoes/database"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"

	"encoding/json"
)

func main() {
	fmt.Println(configs.GetDbUri())
	conn := database.GetConnection(configs.GetDbUri())
	// trasactionRepository := repositories.TransactionsRepository{Db: conn}
	// acoesRepository := repositories.EquitiesRepository{Db: conn}
	userRepository := repositories.UsersRepository{Db: conn}

	err := userRepository.Create(DTO.CreateUser{Name: "João", Balance: 1000.0})

	if err != nil {
		panic(err)
	}

	usuarios, seachUserErr := userRepository.FindAll()

	if seachUserErr != nil {
		panic(seachUserErr)
	}
	for _, usuario := range usuarios {
		jsonUser, err := json.MarshalIndent(usuario, "", "  ")

		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonUser))
	}

}
