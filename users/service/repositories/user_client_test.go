package repositories

import (
	"testing"
	"users/config"
	"users/model"

	"github.com/stretchr/testify/assert"
)

var TestUserClient = NewUserClient(config.DBTESTUSER, config.DBTESTPASS, config.DBTESTHOST, config.DBTESTPORT, config.DBTESTNAME)

func TestGetUserById(t *testing.T) {

	// caso de prueba: usuario existente

	expecteduserid := 1

	user, err := TestUserClient.GetUserById(1)

	assert.NoError(t, err)
	assert.Equal(t, user.Id, expecteduserid)

	// caso de prueba: usuario inexistente

	expecteduserid = 0
	user, er := TestUserClient.GetUserById(666)

	assert.Error(t, er)
	assert.Equal(t, user.Id, expecteduserid)
}

func TestGetUserByUsername(t *testing.T) {

	// caso de prueba: usuario existente

	username := "test_user_1"

	expecteduserid := 1

	user, err := TestUserClient.GetUserByUsername(username)

	assert.NoError(t, err)
	assert.Equal(t, user.Id, expecteduserid)

	// caso de prueba: usuario inexistente

	username = "inexistent_test_user"

	user, er := TestUserClient.GetUserByUsername(username)

	assert.Error(t, er)
	assert.Equal(t, user.Id, 0)
}

func TestInsertUser(t *testing.T) {

	// insertar usuario valido

	validuser := model.User{
		Username: "testinsertuser",
		Email:    "testinsertuser@test.com",
		Password: "testpassword123",
	}

	user := TestUserClient.InsertUser(validuser)

	assert.Equal(t, user.Id, 2)
}
