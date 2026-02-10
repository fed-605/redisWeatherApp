package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvString(key string, defVal string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defVal
}

func Loadenv() error {
	return godotenv.Load()
}
