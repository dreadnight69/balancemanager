package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getStringFromEnv(key string) (res string, err error) {
	res = os.Getenv(key)
	if res == "" {
		return "", fmt.Errorf("env variable with name %s not found", key)
	}
	return
}

func getSecondsDurationFromEnv(key string) (res time.Duration, err error) {
	resString, err := getStringFromEnv(key)
	if err != nil {
		return
	}
	resInt, err := strToInt(resString)
	if err != nil {
		return
	}
	return time.Duration(resInt) * time.Second, nil
}

func strToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
