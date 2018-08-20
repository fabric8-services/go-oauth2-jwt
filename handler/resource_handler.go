package handler

import (
	"net/http"

	"github.com/hemanik/go-oauth2-jwt/models"
	"github.com/hemanik/go-oauth2-jwt/utils"
)

// ProtectedHandler grants access to the protected resource
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Message: "Gained access to protected resource"}
	utils.JSONResponse(response, w)
}
