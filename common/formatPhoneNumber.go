package common

import (
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func FormatPhoneNumber(phone string) (string, error) {
	num, err := phonenumbers.Parse(phone, "VN")
	return strings.ReplaceAll(phonenumbers.Format(num, phonenumbers.INTERNATIONAL), " ", ""), err
}
