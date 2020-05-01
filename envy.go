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

// returns a boolean value
// Any value that isn't YES, yes, TRUE, or true will return as false
func GetEnvBool(env_key, fallback string) (env bool) {
	env = false
	env_read := GetEnv(env_key, fallback)
	if env_read == "YES" || env_read == "yes" || env_read == "TRUE" || env_read == "true" {
		env = true
	}
	return
}

func GetEnvFloat(env_key, fallback string) (env float64, err error) {
	env, err = strconv.ParseFloat(GetEnv(env_key, fallback), 64)
	return
}
