package models

import (
	"fmt"
	"testing"
	"time"
)

func testingUserService() (*UserService, error) {
	const (
		host     = "localhost"
		port     = 5435
		user     = "juanluisrubalcava"
		password = "my-password"
		dbname   = "lenslocked_dev"
		sslmode  = "disable"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s"+" sslmode=disable", host, port, user, dbname)
	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}
	us.db.LogMode(false)
	// Clear the users table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "Dulce Gonzalez",
		Email: "dulcegonzalez@hola.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Errorf("Expected ID > 0. received %d", user.ID)
	}
	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Received %s", user.CreatedAt)
	}
	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Received %s", user.UpdatedAt)
	}
}
