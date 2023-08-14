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
	// calling handleError function
	err := handleError(length)
	if err != nil {
		return "", err
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

// encpsulated function hidden from imported package
func handleError(length int) error {
	// If no invalid was given, return an error with a message.
	if length <= 0 {
		return errors.New("Invalid password length")
	}

	if length > 15 {
		return errors.New("password length exceeds maximum")
	}

	return nil
}

// testing, change package name to main before testing
/*
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("passwordGenarator: ")
	log.SetFlags(0)

	// Request a GenerateRandomPassword.
	// If need to edit imported package file will be present in C:\Users\name\go\pkg\mod\github.com\nithishgwd\folder_name\file.go
	password, err := GenerateRandomPassword(19)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Printf("Password Generated -->>   %s", password)
}
*/
