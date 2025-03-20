// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Music info
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/donskova1ex/effective_mobile/openapi"
)

func main() {
	log.Printf("Server started")

	DefaultAPIService := openapi.NewDefaultAPIService()
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
