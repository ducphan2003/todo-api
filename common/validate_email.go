package common

import (
	"net/mail"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
