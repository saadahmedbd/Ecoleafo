package util

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPassword(HashedPass, plainPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPass), []byte(plainPass))
}
func ValidRole(role string) bool {
	r := strings.ToLower(role)
	return r == "seller" || r == "buyer" || r == "admin"
}
