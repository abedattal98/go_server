package server

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"rgb/domain"
// 	"rgb/models"
// 	"rgb/repositories"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_GetUsers_EmptyResult(t *testing.T) {
// 	var UserRepository = repositories.ProvideUserRepository()
// 	req, w := setGetUserRouter(UserRepository)
// 	a := assert.New(t)
// 	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
// 	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
// 	actual := models.User{}
// 	// Convert the JSON response to a map
// 	var response = []models.User{}
// 	err := json.Unmarshal([]byte(w.Body, &response))
// 	// Grab the value & whether or not it exists
// 	// log.SetOutput((w.Body)
// 	assert.Nil(t, err)
// 	assert.True(t, exists)

// 	log.Println( actual)
// 	var UserStore = []models.User{
// 		models.User{
// 			ID:       1,
// 			Email:    "admin@admin.com",
// 			Password: "admin",
// 			Username: "admin",
// 		},
// 	}
// 	a.Equal(UserStore, value)
// }
// func setGetUserRouter(repo domain.IUserRepository) (*http.Request, *httptest.ResponseRecorder) {
// 	r := gin.Default()
// 	userAPI := initUserAPI(repo)
// 	r.GET("/api/users", userAPI.FindAll)

// 	// r.GET("/api/users", userAPI.FindAll)
// 	req, err := http.NewRequest(http.MethodGet, "/api/users", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	return req, w
// }

// import (
// 	"net/http"
// 	"rgb/models"

// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPizzasHandler(t *testing.T) {
// 	type Users []models.User

// 	tt := []struct {
// 		name       string
// 		method     string
// 		input      *Users
// 		want       string
// 		statusCode int
// 	}{
// 		{
// 			name:       "without pizzas",
// 			method:     http.MethodGet,
// 			input:      &Users{},
// 			want:       "Error: No pizzas found",
// 			statusCode: http.StatusNotFound,
// 		},
// 		{
// 			name:   "with pizzas",
// 			method: http.MethodGet,
// 			input: &Users{
// 				{
// 					ID:       1,
// 					Email:    "hello@gmail.com",
// 					Password: "hello",
// 					Username: "hello",
// 			},
// 			want:       `[{"id":1,"name":"Foo","price":10}]`,
// 			statusCode: http.StatusOK,
// 		},
// 		{
// 			name:       "with bad method",
// 			method:     http.MethodPost,
// 			input:      &Users{},
// 			want:       "Method not allowed",
// 			statusCode: http.StatusMethodNotAllowed,
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			request := httptest.NewRequest(tc.method, "/orders", nil)
// 			responseRecorder := httptest.NewRecorder()

// 			pizzasHandler{tc.input}.ServeHTTP(responseRecorder, request)

// 			if responseRecorder.Code != tc.statusCode {
// 				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
// 			}

// 			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
// 				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
// 			}
// 		})
// 	}
// }
