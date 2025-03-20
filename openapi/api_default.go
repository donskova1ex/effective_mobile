// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Music info
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service DefaultAPIServicer
	errorHandler ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) *DefaultAPIController {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
		"InfoGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/info",
			c.InfoGet,
		},
		"InfoPut": Route{
			strings.ToUpper("Put"),
			"/api/v1/info",
			c.InfoPut,
		},
		"InfoPost": Route{
			strings.ToUpper("Post"),
			"/api/v1/info",
			c.InfoPost,
		},
		"InfoDelete": Route{
			strings.ToUpper("Delete"),
			"/api/v1/info",
			c.InfoDelete,
		},
	}
}

// InfoGet - Get song details
func (c *DefaultAPIController) InfoGet(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var groupParam string
	if query.Has("group") {
		param := query.Get("group")

		groupParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "group"}, nil)
		return
	}
	var songParam string
	if query.Has("song") {
		param := query.Get("song")

		songParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "song"}, nil)
		return
	}
	result, err := c.service.InfoGet(r.Context(), groupParam, songParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// InfoPut - Update an existing song
func (c *DefaultAPIController) InfoPut(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var groupParam string
	if query.Has("group") {
		param := query.Get("group")

		groupParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "group"}, nil)
		return
	}
	var songParam string
	if query.Has("song") {
		param := query.Get("song")

		songParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "song"}, nil)
		return
	}
	var songDetailParam SongDetail
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&songDetailParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSongDetailRequired(songDetailParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSongDetailConstraints(songDetailParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.InfoPut(r.Context(), groupParam, songParam, songDetailParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// InfoPost - Create a new song
func (c *DefaultAPIController) InfoPost(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var groupParam string
	if query.Has("group") {
		param := query.Get("group")

		groupParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "group"}, nil)
		return
	}
	var songParam string
	if query.Has("song") {
		param := query.Get("song")

		songParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "song"}, nil)
		return
	}
	var songDetailParam SongDetail
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&songDetailParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSongDetailRequired(songDetailParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSongDetailConstraints(songDetailParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.InfoPost(r.Context(), groupParam, songParam, songDetailParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// InfoDelete - Delete an existing song
func (c *DefaultAPIController) InfoDelete(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var groupParam string
	if query.Has("group") {
		param := query.Get("group")

		groupParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "group"}, nil)
		return
	}
	var songParam string
	if query.Has("song") {
		param := query.Get("song")

		songParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "song"}, nil)
		return
	}
	result, err := c.service.InfoDelete(r.Context(), groupParam, songParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
