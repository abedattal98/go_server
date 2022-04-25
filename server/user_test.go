package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"rgb/repositories"
	"rgb/services/jwt"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"msg": "worldd",
	}
	// Grab our router
	var UserRepository = repositories.ProvideUserRepository()
	router := setRouter(UserRepository) // Perform a GET request with that handler.
	w := performRequest(router, "GET", "/api/hello", nil)

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["msg"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["msg"], value)
}
func TestUserRoute(t *testing.T) {
	var UserRepository = repositories.ProvideUserRepository()
	router := setRouter(UserRepository) // Perform a GET request with that handler.
	w := performRequest(router, "GET", "/api/users", nil)

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"users":[{"id":"1","username":"admin","email":"admin@admin.com","CreatedAt":"0001-01-01T00:00:00Z","ModifiedAt":"0001-01-01T00:00:00Z"}]}`
	expectedResponse, _ := UnmarshalDataResponse([]byte(expected))
	// Make some assertions on the correctness of the response.
	assert.Equal(t, expected, w.Body.String())
	new, _ := UnmarshalDataResponse([]byte(w.Body.String()))
	assert.Equal(t, expectedResponse.Users, new.Users)

}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("email", "u1@gmail.com")
	params.Add("password", "p1dasdada")
	params.Add("username", "Hello12w3")
	return params.Encode()
}
func TestRegisterUnauthenticated(t *testing.T) {
	//set up the payload
	registrationPayload := `{"email":"Hellloo22worldd","password": "Hello12w3","username":"21312"}`
	//set up the server
	var UserRepository = repositories.ProvideUserRepository()
	jwt.JwtSetup("test")
	router := setRouter(UserRepository)

	// Perform a POST request with that handler.
	w := performRequest(router, "POST", "/api/signup", strings.NewReader(registrationPayload))

	if w.Code != http.StatusOK {
		t.Fail()
	}
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	_, exists := response["jwt"]
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestAddUser(t *testing.T) {
	//set up the payload
	registrationPayload := `{"email":"Hellloo22worldd","password": "Hello12w3","username":"21312"}`
	//set up the server
	var UserRepository = repositories.ProvideUserRepository()
	jwt.JwtSetup("test")
	router := setRouter(UserRepository)

	// Perform a POST request with that handler.
	w := performRequest(router, "POST", "/api/users", strings.NewReader(registrationPayload))
	//Check response code is 200
	if w.Code != http.StatusOK {
		t.Fail()
	}
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["msg"]
	assert.Equal(t, value, "user created successfully")
	assert.True(t, exists)
}

func UnmarshalDataResponse(data []byte) (Data, error) {
	var r Data
	err := json.Unmarshal(data, &r)
	return r, err
}
type Data struct {
	Users []User `json:"users"`
}
type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	CreatedAt  string `json:"CreatedAt"`
	ModifiedAt string `json:"ModifiedAt"`
}
