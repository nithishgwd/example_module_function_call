package randompsswd

import (
	"math/rand"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTVUWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*_+(){}[]:'><?,./'"
)

// function naming in Cap's can be used in other packages
func GenerateRandomPassword(length int) string {
	// seed the random number genarator with current time nanoseconds
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// combining all char
	charset := uppercaseLetters + lowercaseLetters + digits + specialChars

	// createing byte slice to store password
	password := make([]byte, length)

	for i := range password {
		// Generating random index
		randomIndex := rand.Intn(len(charset))

		// select char at the random index
		password[i] = charset[randomIndex]
	}

	// converting byte to string
	return string(password)
}
