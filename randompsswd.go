// reference --->>> https://go.dev/doc/tutorial/create-module
package main

import (
	"errors"
	"math/rand"
	"time"
)

/*
const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTVUWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*_+(){}[]:'><?,./'"
)*/

// function naming in Cap's can be used in other packages
// can return two variables
// After creating UserNamePasswords function, made this function hidden from other package
func generateRandomPassword(length int) (string, error) {
	// calling handleError function
	err := handleError(length)
	if err != nil {
		return "", err
	}

	// seed the random number genarator with current time nanoseconds
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// inizilized only whn this function is called
	charSets := map[string]string{
		"uppercaseLetters": "ABCDEFGHIJKLMNOPQRSTVUWXYZ",
		"lowercaseLetters": "abcdefghijklmnopqrstuvwxyz",
		"digits ":          "0123456789",
		"specialChars":     "!@#$%^&*_+(){}[]:'><?,./'",
	}

	// combining all strings
	charsets := ""
	// index variable is not in use, so left _
	for _, set := range charSets {
		charsets += set
	}

	// createing byte slice to store password
	password := make([]byte, length)

	for i := range password {
		// Generating random index
		randomIndex := rand.Intn(len(charsets))

		// select char at the random index
		password[i] = charsets[randomIndex]
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

// takes slice of names, length of psswd and associate each name with passwordGenartor
// Return map associated each of name with password
func UserNamePasswords(names []string, length int) (map[string]string, error) {
	// map to associate names with GenerateRandomPassword
	namePassw := make(map[string]string)

	// looking through the recived slice of names,calling
	// GenerateRandomPassword to get password to each name
	for _, name := range names {
		password, err := generateRandomPassword(length)
		if err != nil {
			return nil, err
		}
		// in the map, associated the random password to each
		// names
		namePassw[name] = password
	}
	return namePassw, nil
}

// testing, change package name to main before testing
/*
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("passwordGenarator: ")
	log.SetFlags(0)

	names := []string{"Om", "Nithish", "Kishan", "Suraj"}
	// Request a GenerateRandomPassword.
	// If need to edit imported package file will be present in C:\Users\name\go\pkg\mod\github.com\nithishgwd\folder_name\file.go

	// password, err := GenerateRandomPassword(19) -------> now we can make this function encapsulated

	// calling new map function maping names and password
	password, err := UserNamePasswords(names, 14)

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
