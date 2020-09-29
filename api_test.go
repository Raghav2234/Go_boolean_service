package main

import (
	"Go_boolean_service/auth"
	"Go_boolean_service/db"
	"Go_boolean_service/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
)

func performPostRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	var jsonStr = []byte(`{"key":"test_bool", "value": true}`)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestPostData(t *testing.T) {
	database, _ := db.CreateConnection()
	app := serverSetup(database)
	w := performPostRequest(app, "POST", "/")
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	if err == nil {
		t.Fatalf("Error should be nil")
	}
	_, IDExists := response["id"]
	if IDExists != true {
		t.Fatalf("Id should be there in response")
	}
	_, TokenExists := response["token"]
	if TokenExists != true {
		t.Fatalf("Authentication token not returned")
	}
}
func performGetRequest(r http.Handler, method, path string, token string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	var bearer = "Bearer " + token
	req.Header.Add("Authorization", bearer)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestGetRequest(t *testing.T) {
	database, _ := db.CreateConnection()
	app := serverSetup(database)
	ID := db.CreateUUID()
	token, _ := auth.GenerateAccessToken(ID)
	boolObj := models.Boolean{Id: ID, Key: "test_get", Value: true}
	db.CreateBoolean(database, boolObj)
	w := performGetRequest(app, "GET", "/"+ID, token)
	var response models.Boolean
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		t.Fatalf("Error should be nil")
	}

	assert.Equal(t, http.StatusFound, w.Code)
	if response.Id != ID {
		t.Errorf(response.Id + "should be same as" + ID)
	}
	if response.Value != boolObj.Value {
		t.Errorf("Value response should be" + strconv.FormatBool(boolObj.Value))
	}
}
func performDeleteRequest(r http.Handler, method, path string, token string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	var bearer = "Bearer " + token
	req.Header.Add("Authorization", bearer)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestDeleteRequest(t *testing.T) {
	database, _ := db.CreateConnection()
	app := serverSetup(database)
	ID := db.CreateUUID()
	token, _ := auth.GenerateAccessToken(ID)
	boolObj := models.Boolean{Id: ID, Key: "test_delete", Value: true}
	db.CreateBoolean(database, boolObj)
	w := performGetRequest(app, "DELETE", "/"+ID, token)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
func performUpdateRequest(r http.Handler, method, path string, token string) *httptest.ResponseRecorder {
	var jsonStr = []byte(`{"key":"updated_val", "value": false}`)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonStr))
	var bearer = "Bearer " + token
	req.Header.Add("Authorization", bearer)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestUpdateRequest(t *testing.T) {
	database, _ := db.CreateConnection()
	app := serverSetup(database)
	ID := db.CreateUUID()
	token, _ := auth.GenerateAccessToken(ID)
	boolObj := models.Boolean{Id: ID, Key: "test_update", Value: true}
	db.CreateBoolean(database, boolObj)
	w := performUpdateRequest(app, "PATCH", "/"+ID, token)
	assert.Equal(t, http.StatusOK, w.Code)
}
