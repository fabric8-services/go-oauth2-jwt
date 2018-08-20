package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hemanik/go-oauth2-jwt/handler"

	"github.com/hemanik/go-oauth2-jwt/models"
	"github.com/hemanik/go-oauth2-jwt/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	/* load keys as test data */
	utils.InitKeys()
}

func TestAuthenticateEndpoint(t *testing.T) {
	// Given
	user, _ := json.Marshal(&models.OauthUser{
		Username: handler.DefaultUser,
		Password: handler.DefaultPassword,
	})

	// When
	request, _ := http.NewRequest("POST", "/api/authenticate", bytes.NewBuffer(user))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	// Then
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.NotEmpty(t, response.Body)
}

func TestResourceEndpoint(t *testing.T) {
	// Given
	token, _ := handler.GenerateToken()

	// When
	request, _ := http.NewRequest("GET", "/api/resource", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	// Then
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, `{"message":"Gained access to protected resource"}`, response.Body.String())
}
