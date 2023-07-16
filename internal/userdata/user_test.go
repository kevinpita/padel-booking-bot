package userdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveLoginDataFromJson(t *testing.T) {
	login := loginFromFile("test.json")
	loginData, _ := login.get(1)

	assert.Equal(t, "email", loginData.Email, "Usernames should be equal")
	assert.Equal(t, "password", loginData.Password, "Passwords should be equal")
}
