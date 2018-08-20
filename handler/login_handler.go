package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hemanik/go-oauth2-jwt/models"
	"github.com/hemanik/go-oauth2-jwt/utils"
)

const (
	// DefaultUser for authentication
	DefaultUser = "developer"
	// DefaultPassword used for authentication
	DefaultPassword = "developer"
	// DefaultIssuer for JWT Token is OSIO
	DefaultIssuer = "OSIO"
)

// LoginHandler validate user credentials and generates access token for authentication
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user models.OauthUser

	//decode request into OauthUser struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	fmt.Println(user.Username, user.Password)

	//validate user credentials
	if strings.ToLower(user.Username) != DefaultUser {
		if user.Password != DefaultPassword {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}

	tokenString, err := GenerateToken()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	//create a token instance using the token string
	response := models.OauthAccessToken{Token: tokenString}
	utils.JSONResponse(response, w)
}

// GenerateToken generates a jwt token using RS256 algorithm
func GenerateToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Issuer:    DefaultIssuer,
	}

	//create a rsa 256 signer
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(utils.SignKey)
}
