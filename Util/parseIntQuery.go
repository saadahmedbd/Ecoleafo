package util

import (
	"net/http"
	"strconv"
)

func ParseIntQuery(r *http.Request, key string, defaultValue int) int {
	value, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil || value < 1 {
		return defaultValue
	}
	return value
}
func ParseUintQuery(r *http.Request, key string) uint {
	value, err := strconv.ParseUint(r.URL.Query().Get(key), 10, 64)
	if err != nil {
		return 0
	}
	return uint(value)
}
