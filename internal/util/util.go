package util

import (
	"os"
	"strconv"
)

func GetenvDefault(key string, default_value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_value
	} else {
		return val
	}
}

func VarByEnv(env, dev_value string, prod_value string) string {
	if env == "dev" {
		return dev_value
	}
	return prod_value
}

func GetenvIntDefault(key string, default_value int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_value
	} else {
		v, err := strconv.Atoi(val)
		if err != nil {
			return default_value
		}
		return v
	}
}
