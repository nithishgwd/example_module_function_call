package randompsswd

import (
	"errors"
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
// can return two variables
func GenerateRandomPassword(length int) (string, error) {

	// If no invalid was given, return an error with a message.
	if length <= 0 {
		return "", errors.New("Invalid password length")
	}

	if length > 15 {
		return "", errors.New("password length exceeds maximum")
	}

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
	return string(password), nil
}
