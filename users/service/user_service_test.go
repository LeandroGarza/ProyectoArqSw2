package services

import (
	"errors"
	"testing"
	"users/dto"
	"users/model"
	e "users/utils/errors"

	"github.com/stretchr/testify/assert"
)

// creacion de un cliente de usuario mock
type MockUserClient struct{}

func (m *MockUserClient) GetUserById(id int) (model.User, error) {
	if id == 1 {
		return model.User{
			Id:       id,
			Username: "test_user_1",
			Email:    "test_user_1@test.com",
			Password: "testpassword123",
		}, nil
	} else {
		return model.User{}, errors.New("error")
	}
}

func (m *MockUserClient) DeleteUser(user model.User) error {
	if user.Id == 2 {
		return nil
	} else {
		return errors.New("error")
	}
}

func (m *MockUserClient) GetUserByUsername(username string) (model.User, error) {
	return model.User{
		Id:       1,
		Username: username,
		Email:    "test_user_1@test.com",
		Password: "testpassword123",
	}, nil
}

func (m *MockUserClient) InsertUser(user model.User) model.User {
	if user.Username == "testinsertuser" {
		return model.User{
			Id:       2,
			Username: "testinsertuser",
			Email:    "testinsertuser@test.com",
			Password: "testpassword123",
		}
	} else {
		return model.User{
			Id:       0,
			Username: "testinsertuser",
			Email:    "testinsertuser@test.com",
			Password: "testpassword123",
		}
	}
}

// creacion de un cliente de rabbit mock
type MockQueueClient struct{}

func (mq *MockQueueClient) SendMessage(userid int, action string, message string) e.ApiError {
	return nil
}

var mockuserclient = &MockUserClient{}
var mockqueueclient = &MockQueueClient{}

var TestServiceImpl = NewUserServiceImpl(mockuserclient, mockqueueclient)

func TestGetUserById(t *testing.T) {

	// caso de prueba usuario existente

	id := 1
	expecteduserid := 1

	user, err := TestServiceImpl.GetUserById(id)

	assert.NoError(t, err)
	assert.Equal(t, user.Id, expecteduserid)

	// caso de prueba para usuario inexistente

	id = 666
	user, er := TestServiceImpl.GetUserById(id)

	assert.Error(t, er)
	assert.Equal(t, user.Id, 0)
}

func TestDeleteUser(t *testing.T) {
	// caso de prueba para id valido

	id := 2
	err := TestServiceImpl.DeleteUser(id)

	assert.NoError(t, err)

	// caso de prueba para id no valido

	id = 999
	er := TestServiceImpl.DeleteUser(id)

	assert.Error(t, er)
}

func TestInsertUser(t *testing.T) {
	// caso de prueba para usuario valido
	validuserdto := dto.UserDto{
		Username: "testinsertuser",
		Email:    "testinsertuser@test.com",
		Password: "testpassword123",
	}

	expecteduserid := 2

	userdto, err := TestServiceImpl.InsertUser(validuserdto)

	assert.NoError(t, err)
	assert.Equal(t, userdto.Id, expecteduserid)

	// caso de prueba para usuario no valido
	invaliduser := dto.UserDto{
		Username: "testinsertuserinvalid",
		Email:    "testinsertuser@test.com",
		Password: "testpassword123",
	}

	userdto, _ = TestServiceImpl.InsertUser(invaliduser)

	assert.Equal(t, userdto.Id, 0)
}
