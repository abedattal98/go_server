package services

import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
 )

// func TestAddUser(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	assert.Greater(t, user.ID, 0)
// }
// func TestAddUserWithExistingUsername(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	assert.Greater(t, user.ID, 0)
// 	user, err = addTestUser()
// 	assert.Error(t, err)
// 	assert.Equal(t, "User already exists", err.Error())
// }
// func TestAuthenticateUserInvalidUsername(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)
// 	authUser, err := Authenticate("invalid", user.Password)
// 	assert.Error(t, err)
// 	assert.Nil(t, authUser)
// }


// func TestAuthenticateUserInvalidPassword(t *testing.T) {
// 	testSetup()
// 	user, err := addTestUser()
// 	assert.NoError(t, err)

// 	authUser, err := Authenticate(user.Username, "invalid")
// 	assert.Error(t, err)
// 	assert.Nil(t, authUser)
// }



// // func TestFetchUser(t *testing.T) {
// // 	testSetup()
// // 	user, err := addTestUser()
// // 	assert.NoError(t, err)

// // 	fetchedUser, err := FetchUser(user.ID)
// // 	assert.NoError(t, err)
// // 	assert.Equal(t, user.ID, fetchedUser.ID)
// // 	assert.Equal(t, user.Username, fetchedUser.Username)
// // 	assert.Empty(t, fetchedUser.Password)
// // 	assert.Equal(t, user.Salt, fetchedUser.Salt)
// // 	assert.Equal(t, user.HashedPassword, fetchedUser.HashedPassword)
// // }

// // func TestFetchNotExistingUser(t *testing.T) {
// // 	testSetup()

// // 	fetchedUser, err := FetchUser(1)
// // 	assert.Error(t, err)
// // 	assert.Nil(t, fetchedUser)
// // 	assert.Equal(t, "Not found.", err.Error())
// // }
