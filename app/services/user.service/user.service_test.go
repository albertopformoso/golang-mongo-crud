package user_service_test

import (
	"testing"
	"time"

	m "golang-mongo/models"
	userService "golang-mongo/services/user.service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	UserId = oid.Hex()

	user := m.User{
		ID:        oid,
		Name:      "John",
		Email:     "john@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("The data persistency for the user failed")
		t.Fail()
	} else {
		t.Log("The test ran succesfully!")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()

	if err != nil {
		t.Error("An error occured when reading users")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("The query didn't return data")
		t.Fail()
	} else {
		t.Log("The test ran succesfully!")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Johnny",
		Email: "johnny@test.com",
	}

	err := userService.Update(user, UserId)

	if err != nil {
		t.Error("An error occured while updating the user")
		t.Fail()
	} else {
		t.Log("The test ran succesfully!")
	}
}

func TestDelete(t *testing.T) {

	err := userService.Delete(UserId)

	if err != nil {
		t.Error("An error occured while deleteing the user")
		t.Fail()
	} else {
		t.Log("The test ran succesfully!")
	}
}
