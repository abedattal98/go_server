package services

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

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



