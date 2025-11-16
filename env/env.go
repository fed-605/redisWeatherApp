package env

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func GetEnvString(key string, defVal string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defVal
}

func Loadenv() error {
	var loadOnce sync.Once
	var loadErr error
	loadOnce.Do(func() {
		loadErr = godotenv.Load()
	})
	return loadErr
}
