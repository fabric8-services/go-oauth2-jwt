package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hemanik/go-oauth2-jwt/handler"
	"github.com/hemanik/go-oauth2-jwt/utils"
)

// Router ....
func Router() *mux.Router {
	router := mux.NewRouter()

	//PUBLIC ENDPOINTS
	router.HandleFunc("/api/authenticate", handler.LoginHandler).Methods("POST")

	//PROTECTED ENDPOINTS
	router.Handle("/api/resource", negroni.New(
		negroni.HandlerFunc(handler.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(handler.ProtectedHandler)),
	)).Methods("GET")

	return router
}

func main() {
	utils.InitKeys()
	log.Println("Now listening...")
	log.Fatal(http.ListenAndServe(":8000", Router()))
}
