package repositories

import (
	"rgb/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddUser(t *testing.T) {

	repos := NewRepositoriesTest(Memory)

	user := models.User{
		Username: "rand.String(10)",
		Password: "secret123",
		ID:       int(time.Now().Unix()),
	}

	createUser, err := addTestUser(repos)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, createUser, user)
}

func TestAddUserWithExistingUsername(t *testing.T) {
	repos := NewRepositoriesTest(Memory)

	createUser, err := addTestUser(repos)
	assert.NoError(t, err)
	assert.Greater(t, createUser.ID, 0)
	_, err = addTestUser(repos)
	assert.Error(t, err)
	assert.Equal(t, "User already exists", err.Error())
}
func TestAuthenticateUserInvalidUsername(t *testing.T) {
	repos := NewRepositoriesTest(Memory)

	createUser, err := addTestUser(repos)
	assert.NoError(t, err)
	authUser, err := repos.Users.Authenticate("invalid", createUser.Password)
	assert.Error(t, err)
	assert.Empty(t, authUser)
}
func TestAuthenticateUserInvalidPassword(t *testing.T) {
	repos := NewRepositoriesTest(Memory)

	createUser, err := addTestUser(repos)

	assert.NoError(t, err)
	authUser, err := repos.Users.Authenticate(createUser.Username, "invalid")
	assert.Error(t, err)
	assert.Empty(t, authUser)
}
