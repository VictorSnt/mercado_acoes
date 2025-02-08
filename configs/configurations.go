package configs

import "os"

func GetDbUri() string {
	return os.Getenv("DATABASE_URI")
}
