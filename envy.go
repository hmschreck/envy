package envy

import (
	"os"
	"strconv"
)

func GetEnv(env_key, fallback string) (env string) {
	if value, ok := os.LookupEnv(env_key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(env_key, fallback string) (env int, err error) {
	env, err = strconv.Atoi(GetEnv(env_key, fallback))
	if err != nil {
		return
	}
	return
}
