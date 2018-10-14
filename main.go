/*
Main entry point for Boldly Go RESTful Application.

	Simple REST API to expose a few endpoints to get a feel for GoLang.

	Endpoints:
		- /api/v1/ping // HEALTH CHECK PING ENDPOINT
		- /api/v1/user/{owningUserId}/bank/{bankId} // GET THE USERS BANK RECORD
		- /api/v1/bank // SAVE A BANK RECORD
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const appPortKey = ":5002"

var awsSvc AwsConfig = &awsConf{}

func main() {
	// instantiate aws configuration
	awsSvc.Init()
	// instantiate mux router
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS").Schemes("http")
	// give me that api path-prefix
	api := router.PathPrefix("/api").Subrouter()
	// endpoint registry
	api.HandleFunc("/v1/ping", PingHandler).Methods("GET")
	api.HandleFunc("/v1/user/{owningUserId}/bank/{bankId}", GetBankHandler).Methods("GET")
	api.HandleFunc("/v1/bank", SaveBankHandler).Methods("POST")
	// add CORS acceptance to all requests
	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
	)(api)
	// start app
	fmt.Println(fmt.Sprintf("App Running on Port %s", appPortKey))
	log.Fatal(http.ListenAndServe(appPortKey, handlers.LoggingHandler(os.Stdout, handler)))
}
