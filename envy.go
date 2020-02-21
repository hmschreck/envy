package envy

import "os"

func GetEnv(env_key, fallback string) (env string) {
	if value, ok := os.LookupEnv(env_key); ok {
		return value
	}
	return fallback
}
