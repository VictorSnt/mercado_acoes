package main

import (
	"fmt"
	"mercado/acoes/database"
	"mercado/acoes/database/repositories"
	"os"

	"encoding/json"
)

func main() {
	dbUri := os.Getenv("DATABASE_URI")
	conn := database.GetConnection(dbUri)
	// trasactionRepository := repositories.TransactionsRepository{Db: conn}
	// acoesRepository := repositories.EquitieRepository{Db: conn}
	userRepository := repositories.UsersRepository{Db: conn}

	// err := userRepository.Create(DTO.CreateUser{Name: "João", Balance: 1000.0})

	// if err != nil {
	// 	panic(err)
	// }

	usuario, seachUserErr := userRepository.FindById(1)

	if seachUserErr != nil {
		panic(seachUserErr)
	}

	jsonUser, err := json.MarshalIndent(usuario, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonUser))
	// TODO: Implementar teste amanhã
}
