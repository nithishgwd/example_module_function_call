// refernce https://go.dev/doc/tutorial/add-a-test
// create a file called name_test.go
// run test (cmd)
package randompasswd

import (
	"testing"
)

// Test TestRandomPasswd calls UserNamePasswords with a name, checking
// for error and password not found
func TestRandomPasswd(t *testing.T) {
	name := []string{"Gladys", "Samantha", "Darrin"}
	length := 14

	namePasswd, err := UserNamePasswords(name, length)
	if err != nil {
		t.Fatalf("Error generating user name password: %v", err)
	}

	for _, name := range name {
		password, found := namePasswd[name]
		if !found {
			t.Errorf("password not found for name: %s", name)
		} else if len(password) != length {
			t.Errorf("Expected password length %d for name %s, but got %d", length, name, len(password))
		}
	}
}

// 2nd test TestPasswdLength check for length of password is correct
// go test -v
func TestPasswdLength(t *testing.T) {
	name := []string{"Gladys", "Samantha", "Darrin"}
	length := 18

	namePasswd, _ := UserNamePasswords(name, length)
	if len(namePasswd) != len(name) {
		t.Errorf("Expected %d passwords, but got %d", len(name), len(namePasswd))
	}
}
