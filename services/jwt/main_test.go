package jwt

import (
	"github.com/gin-gonic/gin"
	"log"
	"rgb/models"
	"rgb/services"
	"rgb/store"
)

func testSetup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	store.Users = nil
	JwtSetup("secret")
	return gin.Default()
}

func addTestUser() *models.UserEntity {
	user := &models.User{
		Username: "batman",
		Password: "secret123",
	}
	UserEntity, err := services.AddUser(user)
	if err != nil {
		log.Println("Error adding test user", err)
	}
	return UserEntity
}


	// "bytes"
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"
	// "strings"
	// "rgb/internal/conf"

// func addTestUser2() *models.User {
// 	user := &models.User{
// 		Username: "superman",
// 		Password: "secret123",
// 	}
// 	user,err := services.AddUser(user)
// 	if err != nil {
// 		log.Println("Error adding test user", err)
// 	}
// 	return user
// }

// func addTestPost(user *models.User) *models.Post {
// 	post := &models.Post{
// 		Title:   "Gotham cronicles",
// 		Content: "Joker is planning a big hit tonight.",
// 	}
// 	err := services.AddPost(user, post)
// 	if err != nil {
// 		log.Println("Error adding test post user", err)
// 	}
// 	return post
// }

// func addTestPost2(user *models.User) *models.Post {
// 	post := &models.Post{
// 		Title:   "Justice league meeting",
// 		Content: "Darkseid is plotting again.",
// 	}
// 	err := services.AddPost(user, post)
// 	if err != nil {
// 		log.Println("Error adding test post user", err)
// 	}
// 	return post
// }

// func userJSON(user models.User) string {
// 	body, err := json.Marshal(map[string]interface{}{
// 		"Username": user.Username,
// 		"Password": user.Password,
// 	})
// 	if err != nil {
// 		log.Println("Error marshalling JSON body", err)
// 	}
// 	return string(body)
// }

// func postJSON(post models.Post) string {
// 	body, err := json.Marshal(map[string]interface{}{
// 		"ID":      post.ID,
// 		"Title":   post.Title,
// 		"Content": post.Content,
// 	})
// 	if err != nil {
// 		log.Println("Error marshalling JSON body", err)
// 	}
// 	return string(body)
// }

// func jsonRes(body *bytes.Buffer) map[string]interface{} {
// 	jsonRes := &map[string]interface{}{}
// 	err := json.Unmarshal(body.Bytes(), jsonRes)
// 	if err != nil {
// 		log.Println("Error unmarshalling JSON body", err)
// 	}
// 	return *jsonRes
// }

// func jsonDataSlice(body *bytes.Buffer) []map[string]interface{} {
// 	jsonRes := jsonRes(body)
// 	_jsonDataSlice, ok := jsonRes["data"].([]interface{})
// 	if !ok {
// 		log.Println("JSON response data is not a slice.",jsonRes)
// 	}
// 	jsonSliceMaps := make([]map[string]interface{}, 0)
// 	for _, _jsonSliceMap := range _jsonDataSlice {
// 		jsonSliceMap, ok := _jsonSliceMap.(map[string]interface{})
// 		if !ok {
// 			log.Println("JSON object in slice is not a map.",_jsonSliceMap)
// 		}
// 		jsonSliceMaps = append(jsonSliceMaps, jsonSliceMap)
// 	}
// 	return jsonSliceMaps
// }

// func jsonFieldError(jsonRes map[string]interface{}, field string) interface{} {
// 	jsonError, ok := jsonRes["error"].(map[string]interface{})
// 	if !ok {
// 		log.Println("JSON response error is not a map.",jsonRes["error"])
// 	}
// 	return jsonError[field]
// }

// func jsonFieldData(jsonRes map[string]interface{}, field string) interface{} {
// 	jsonData, ok := jsonRes["data"].(map[string]interface{})
// 	if !ok {
// 		log.Println("JSON response error is not a map.",jsonRes["data"])
// 	}
// 	return jsonData[field]
// }

// func NewRequest(router *gin.Engine, method, path, body string) *http.Request {
// 	req, err := http.NewRequest(method, path, strings.NewReader(body))
// 	if err != nil {
// 		log.Println("Error creating new request")
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	return req
// }

// func performRequest(router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
// 	req := NewRequest(router, method, path, body)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)
// 	return rec
// }

// func PerformAuthorizedRequest(router *gin.Engine, token, method, path, body string) *httptest.ResponseRecorder {
// 	req := NewRequest(router, method, path, body)
// 	rec := httptest.NewRecorder()
// 	req.Header.Add("Authorization", "Bearer "+token)
// 	router.ServeHTTP(rec, req)
// 	return rec
// }
