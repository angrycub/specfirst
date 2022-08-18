/*
 * Nomad Secure Variables
 *
 * Partial API specification for Nomad's secure variables feature
 *
 * API version: 1.4.0
 * Contact: support@hashicorp.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// VariableOperationsApiController binds http requests to an api service and writes the service results to the http response
type VariableOperationsApiController struct {
	service VariableOperationsApiServicer
	errorHandler ErrorHandler
}

// VariableOperationsApiOption for how the controller is set up.
type VariableOperationsApiOption func(*VariableOperationsApiController)

// WithVariableOperationsApiErrorHandler inject ErrorHandler into controller
func WithVariableOperationsApiErrorHandler(h ErrorHandler) VariableOperationsApiOption {
	return func(c *VariableOperationsApiController) {
		c.errorHandler = h
	}
}

// NewVariableOperationsApiController creates a default api controller
func NewVariableOperationsApiController(s VariableOperationsApiServicer, opts ...VariableOperationsApiOption) Router {
	controller := &VariableOperationsApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the VariableOperationsApiController
func (c *VariableOperationsApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteSecureVariable",
			strings.ToUpper("Delete"),
			"/v1/var/{pathToVariable}",
			c.DeleteSecureVariable,
		},
		{
			"GetSecureVariable",
			strings.ToUpper("Get"),
			"/v1/var/{pathToVariable}",
			c.GetSecureVariable,
		},
		{
			"ListVars",
			strings.ToUpper("Get"),
			"/v1/vars",
			c.ListVars,
		},
		{
			"UpsertSecureVariable",
			strings.ToUpper("Put"),
			"/v1/var/{pathToVariable}",
			c.UpsertSecureVariable,
		},
	}
}

// DeleteSecureVariable - Delete a secure variable bag
func (c *VariableOperationsApiController) DeleteSecureVariable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	pathToVariableParam := params["pathToVariable"]
	
	regionParam := query.Get("region")
	namespaceParam := query.Get("namespace")
	result, err := c.service.DeleteSecureVariable(r.Context(), pathToVariableParam, regionParam, namespaceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSecureVariable - Fetch a secure variables bag
func (c *VariableOperationsApiController) GetSecureVariable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pathToVariableParam := params["pathToVariable"]
	
	result, err := c.service.GetSecureVariable(r.Context(), pathToVariableParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListVars - List the variable bags
func (c *VariableOperationsApiController) ListVars(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	regionParam := query.Get("region")
	namespaceParam := query.Get("namespace")
	indexParam, err := parseInt32Parameter(query.Get("index"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	waitParam := query.Get("wait")
	staleParam := query.Get("stale")
	prefixParam := query.Get("prefix")
	perPageParam, err := parseInt32Parameter(query.Get("per_page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	nextTokenParam := query.Get("next_token")
	result, err := c.service.ListVars(r.Context(), regionParam, namespaceParam, indexParam, waitParam, staleParam, prefixParam, perPageParam, nextTokenParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpsertSecureVariable - Store a secure variable bag
func (c *VariableOperationsApiController) UpsertSecureVariable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	pathToVariableParam := params["pathToVariable"]
	
	secureVariableParam := SecureVariable{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&secureVariableParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSecureVariableRequired(secureVariableParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	regionParam := query.Get("region")
	namespaceParam := query.Get("namespace")
	result, err := c.service.UpsertSecureVariable(r.Context(), pathToVariableParam, secureVariableParam, regionParam, namespaceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}