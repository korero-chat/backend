package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/korero-chat/backend/database"
	"github.com/korero-chat/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// before tests
	// using the "test" db by default
	users := database.ConnectToDB().Database("test").Collection("users")
	users.Drop(context.TODO())

	os.Exit(m.Run())

	// after tests
	users.Drop(context.TODO())
}

func TestRegisterUserEndpointWithValidData(t *testing.T) {
	payload, err := json.Marshal(models.User{
		Username: "kicia",
		Password: "miaumiau",
		Email: "miau@email.com",
	})
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", "/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()

	RegisterUserEndpoint(rec, req)

	assert.Equal(
		t,
		http.StatusCreated,
		rec.Result().StatusCode,
		"should return 201 Created upon creation",
	)

	var jsonData models.ResponseModel
	json.Unmarshal(rec.Body.Bytes(), &jsonData) 
	assert.Empty(
		t,
		jsonData.Error,
		"'errors' value should be falsy",
	)
}