package kubenurse

import (
	"os"
	"strconv"
)

func getIntEnv(envStr string, defaultValue int) int {
	if value, ok := os.LookupEnv(envStr); ok {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}

		return intValue
	}

	return defaultValue
}
