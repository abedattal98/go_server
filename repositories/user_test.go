package repositories

import (
	"reflect"
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

func TestUserRepository_Save(t *testing.T) {
	type fields struct {
		db *MemoryStorage
	}
	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		{ // Test case 1
			name: "Test case 1",
			fields: fields{
				db: NewMemoryStorage(),
			},
			args: args{
				user: models.User{
					Username: "rand.String(10)",
					Password: "secret123",
				},
			},
			want: models.User{
				Username: "rand.String(10)",
				Password: "secret123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &UserRepository{
				db: tt.fields.db,
			}
			got, err := p.Save(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
