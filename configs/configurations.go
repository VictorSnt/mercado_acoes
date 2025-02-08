package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDbUri() string {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Print("Erro ao carregar o arquivo .env")
		}
	}

	return os.Getenv("DATABASE_URI")
}
