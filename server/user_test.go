package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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
func TestPingRoute(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"msg": "world",
	}
	// Grab our router
	router := setRouter()
	// Perform a GET request with that handler.
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
func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("email", "u1@gmail.com")
	params.Add("password", "p1dasdada")
	params.Add("username","hellfoh")
	return params.Encode()
}
func TestRegisterUnauthenticated(t *testing.T) {

	w := httptest.NewRecorder()

	r := setRouter()

	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	_,exists := response["jwt"]
	assert.Nil(t, err)
	assert.True(t, exists)
}

