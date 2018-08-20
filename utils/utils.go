package utils

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "secret/app.rsa"
	pubKeyPath  = "secret/app.rsa.pub"
)

var (
	// VerifyKey is the corresponding public key for the given private key
	VerifyKey *rsa.PublicKey
	// SignKey is the private key used for authentication
	SignKey *rsa.PrivateKey
)

// InitKeys initializes the public and private keys
func InitKeys() {
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
}

// JSONResponse marshals the received response in json format
func JSONResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
