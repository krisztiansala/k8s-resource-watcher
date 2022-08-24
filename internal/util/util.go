package util

import (
	"os"
	"strconv"
	"unicode"
)

func GetenvDefault(key string, default_value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_value
	} else {
		return val
	}
}

func GetenvIntDefault(key string, default_value int) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_value, nil
	} else {
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		return v, nil
	}
}

func GetPort(flagPort int, envPort int) int {
	if envPort != 0 {
		return envPort
	}
	if flagPort != 0 {
		return flagPort
	}
	return 8080
}

func ParseName(originalName string) string {
	name := ""
	first := true
	for _, char := range originalName {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) && !first {
				name += " "
			}
			name += string(char)
			if first {
				first = false
			}
		}
	}
	return name
}
