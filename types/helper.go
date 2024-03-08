package types

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsPhoneValid(e string) bool {
	if len(e) < minPhoneLen {
		return false
	}
	return true
}
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
func IsValidPassword(encryptedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil
}
