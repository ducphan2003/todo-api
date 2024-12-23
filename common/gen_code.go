package common

import (
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func GenerateCode(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	rand.Seed(time.Now().UnixNano())
	var password strings.Builder

	// set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	// set numberic character
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	// set uppercase character
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	// Shuffle
	shuffleString := []rune(password.String())
	rand.Shuffle(len(shuffleString), func(i, j int) {
		shuffleString[i], shuffleString[j] = shuffleString[j], shuffleString[i]
	})
	return string(shuffleString)
}
