package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/rayspock/go-answer/config"
	"github.com/rayspock/go-answer/models"
	"github.com/rayspock/go-answer/routes"
)

const (
	endpoint = "/api"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	config.LoadENV()

	log.Println("Open connection to database")
	dbConfig := config.BuildDBConfig()
	dbURL := config.DbURL(dbConfig)
	log.Printf("Connection detail: %v", dbURL)
	config.DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Println("Status:", err)
	}
	defer config.DB.Close()

	log.Println("Drop table")
	models.DropTableIfExists(config.DB)

	log.Println("Ensure table exists")
	models.Init(config.DB)

	log.Println("Setup Router")
	router = routes.SetupRouter()

	log.Println("Run test cases")
	code := m.Run()

	log.Println("Leave Test")
	os.Exit(code)
}

func TestPingRoute(t *testing.T) {
	w := performRequest(router, "/ping")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestCreateAnswerByKey(t *testing.T) {
	path := "/answer"
	payload := []byte(`{
		"key": "name",
		"value": "John"
	}`)
	// Everything work as expected
	w := performPost("POST", router, path, bytes.NewBuffer(payload))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "success")

	// Key already exists
	w = performPost("POST", router, path, bytes.NewBuffer(payload))
	var resp map[string]interface{}
	fromJSON(&resp, w.Body)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, resp["message"], "the answer already exists")
}

func TestGetAnswerByKey(t *testing.T) {
	path := "/answer/name"

	// Everything work as expected
	w := performRequest(router, path)
	var resp map[string]interface{}
	fromJSON(&resp, w.Body)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resp["value"], "John")

	// Key doest not exists
	path = "/answer/title"
	w = performRequest(router, path)
	fromJSON(&resp, w.Body)
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, resp["message"], "record not found")
}

func TestGetAllAnswer(t *testing.T) {
	path := "/answer"

	// Create 1 record
	payload := []byte(`{
		"key": "country",
		"value": "UK"
	}`)
	w := performPost("POST", router, path, bytes.NewBuffer(payload))

	// Everything work as expected
	w = performRequest(router, path)
	var resp []models.Answer
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		log.Fatalf("JSON Decode fail: %v", err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resp[0].Val, "John")
	assert.GreaterOrEqual(t, len(resp), 2)
}

func TestGetAnswerHistoryByKey(t *testing.T) {
	path := "/answer/name/history"

	// Everything work as expected
	w := performRequest(router, path)
	var resp []models.History
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		log.Fatalf("JSON Decode fail: %v", err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resp[0].Event, "create")
	
	for _, item := range resp {
		// Make sure there is no "get" events
		assert.NotEqual(t, item, "get")
	}
}

func TestUpdateAnswerByKey(t *testing.T) {
	path := "/answer/name"
	payload := []byte(`{
		"value": "Ray"
	}`)
	// Everything work as expected
	w := performPost("PUT", router, path, bytes.NewBuffer(payload))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "success")

	// Update the key does not exist
	w = performPost("DELETE", router, path, nil)	
	w = performPost("PUT", router, path, bytes.NewBuffer(payload))
	var resp map[string]interface{}
	fromJSON(&resp, w.Body)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, resp["message"], "the answer doesn't exist or has been deleted")
}

func TestDeleteAnswerByKey(t *testing.T) {
	path := "/answer"
	payload := []byte(`{
		"key": "name",
		"value": "John"
	}`)
	performPost("POST", router, path, bytes.NewBuffer(payload))

	path = "/answer/name"
	// Everything work as expected
	w := performPost("DELETE", router, path, nil)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "success")

	// Deleted key does not exist
	w = performPost("DELETE", router, path, nil)
	var resp map[string]interface{}
	fromJSON(&resp, w.Body)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, resp["message"], "the answer doesn't exist or has been deleted")
}

func fromJSON(resp *map[string]interface{}, body *bytes.Buffer) {
	err := json.NewDecoder(body).Decode(resp)
	if err != nil {
		log.Fatalf("JSON Decode fail: %v", err)
	}
}

func performPost(method string, r http.Handler, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, endpoint+url, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performRequest(r http.Handler, url string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", endpoint+url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
