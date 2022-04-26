package repositories

// import (

// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/stretchr/testify/require"
// )


// func TestAddUser(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	createUser, err := AddTestUser(repo)
// 	if err != nil {
// 		t.Fatalf("expected: %v, got: %v", nil, err)
// 	}
// 	require.Equal(t, createUser, nil)
// }
// func TestAddUserWithExistingUsername(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	createUser, err := AddTestUser(repo)
// 	assert.NoError(t, err)
// 	assert.Greater(t, createUser.ID, 0)
// 	_, err = AddTestUser(repo)
// 	assert.Error(t, err)
// 	assert.Equal(t, "User already exists", err.Error())
// }
// func TestAuthenticateUserInvalidUsername(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)
// 	authUser, err := repo.Authenticate("invalid", user.Password)
// 	assert.Error(t, err)
// 	assert.Empty(t, authUser)
// }
// func TestAuthenticateUserInvalidPassword(t *testing.T) {
// 	repo := ProvideUserRepository()
// 	user, err := AddTestUser(repo)
// 	assert.NoError(t, err)
// 	authUser, err := repo.Authenticate(user.Username, "invalid")
// 	assert.Error(t, err)
// 	assert.Empty(t, authUser)
// }

