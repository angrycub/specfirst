/*
 * Nomad Secure Variables
 *
 * Partial API specification for Nomad's secure variables feature
 *
 * API version: 1.4.0
 * Contact: support@hashicorp.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	VariableOperationsApiService := openapi.NewVariableOperationsApiService()
	VariableOperationsApiController := openapi.NewVariableOperationsApiController(VariableOperationsApiService)

	router := openapi.NewRouter(VariableOperationsApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
