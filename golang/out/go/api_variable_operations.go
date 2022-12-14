/*
Nomad Secure Variables

Partial API specification for Nomad's secure variables feature

API version: 1.4.0
Contact: support@hashicorp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// VariableOperationsApiService VariableOperationsApi service
type VariableOperationsApiService service

type ApiDeleteSecureVariableRequest struct {
	ctx context.Context
	ApiService *VariableOperationsApiService
	pathToVariable string
	region *string
	namespace *string
}

// Filters results based on the specified region.
func (r ApiDeleteSecureVariableRequest) Region(region string) ApiDeleteSecureVariableRequest {
	r.region = &region
	return r
}

// Filters results based on the specified namespace.
func (r ApiDeleteSecureVariableRequest) Namespace(namespace string) ApiDeleteSecureVariableRequest {
	r.namespace = &namespace
	return r
}

func (r ApiDeleteSecureVariableRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSecureVariableExecute(r)
}

/*
DeleteSecureVariable Delete a secure variable bag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pathToVariable Slash-delimited path to the read, store, or delete a variable
 @return ApiDeleteSecureVariableRequest
*/
func (a *VariableOperationsApiService) DeleteSecureVariable(ctx context.Context, pathToVariable string) ApiDeleteSecureVariableRequest {
	return ApiDeleteSecureVariableRequest{
		ApiService: a,
		ctx: ctx,
		pathToVariable: pathToVariable,
	}
}

// Execute executes the request
func (a *VariableOperationsApiService) DeleteSecureVariableExecute(r ApiDeleteSecureVariableRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariableOperationsApiService.DeleteSecureVariable")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/var/{pathToVariable}"
	localVarPath = strings.Replace(localVarPath, "{"+"pathToVariable"+"}", url.PathEscape(parameterToString(r.pathToVariable, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSecureVariableRequest struct {
	ctx context.Context
	ApiService *VariableOperationsApiService
	pathToVariable string
}

func (r ApiGetSecureVariableRequest) Execute() (*SecureVariable, *http.Response, error) {
	return r.ApiService.GetSecureVariableExecute(r)
}

/*
GetSecureVariable Fetch a secure variables bag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pathToVariable Slash-delimited path to the read, store, or delete a variable
 @return ApiGetSecureVariableRequest
*/
func (a *VariableOperationsApiService) GetSecureVariable(ctx context.Context, pathToVariable string) ApiGetSecureVariableRequest {
	return ApiGetSecureVariableRequest{
		ApiService: a,
		ctx: ctx,
		pathToVariable: pathToVariable,
	}
}

// Execute executes the request
//  @return SecureVariable
func (a *VariableOperationsApiService) GetSecureVariableExecute(r ApiGetSecureVariableRequest) (*SecureVariable, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecureVariable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariableOperationsApiService.GetSecureVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/var/{pathToVariable}"
	localVarPath = strings.Replace(localVarPath, "{"+"pathToVariable"+"}", url.PathEscape(parameterToString(r.pathToVariable, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListVarsRequest struct {
	ctx context.Context
	ApiService *VariableOperationsApiService
	region *string
	namespace *string
	index *int32
	wait *string
	stale *string
	prefix *string
	perPage *int32
	nextToken *string
}

// Filters results based on the specified region.
func (r ApiListVarsRequest) Region(region string) ApiListVarsRequest {
	r.region = &region
	return r
}

// Filters results based on the specified namespace.
func (r ApiListVarsRequest) Namespace(namespace string) ApiListVarsRequest {
	r.namespace = &namespace
	return r
}

// If set, wait until query exceeds given index. Must be provided with WaitParam.
func (r ApiListVarsRequest) Index(index int32) ApiListVarsRequest {
	r.index = &index
	return r
}

// Provided with IndexParam to wait for change.
func (r ApiListVarsRequest) Wait(wait string) ApiListVarsRequest {
	r.wait = &wait
	return r
}

// If present, results will include stale reads.
func (r ApiListVarsRequest) Stale(stale string) ApiListVarsRequest {
	r.stale = &stale
	return r
}

// Constrains results to jobs that start with the defined prefix
func (r ApiListVarsRequest) Prefix(prefix string) ApiListVarsRequest {
	r.prefix = &prefix
	return r
}

// Maximum number of results to return.
func (r ApiListVarsRequest) PerPage(perPage int32) ApiListVarsRequest {
	r.perPage = &perPage
	return r
}

// Indicates where to start paging for queries that support pagination.
func (r ApiListVarsRequest) NextToken(nextToken string) ApiListVarsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListVarsRequest) Execute() ([]SecureVariableStub, *http.Response, error) {
	return r.ApiService.ListVarsExecute(r)
}

/*
ListVars List the variable bags

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListVarsRequest
*/
func (a *VariableOperationsApiService) ListVars(ctx context.Context) ApiListVarsRequest {
	return ApiListVarsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SecureVariableStub
func (a *VariableOperationsApiService) ListVarsExecute(r ApiListVarsRequest) ([]SecureVariableStub, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SecureVariableStub
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariableOperationsApiService.ListVars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.index != nil {
		localVarQueryParams.Add("index", parameterToString(*r.index, ""))
	}
	if r.wait != nil {
		localVarQueryParams.Add("wait", parameterToString(*r.wait, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpsertSecureVariableRequest struct {
	ctx context.Context
	ApiService *VariableOperationsApiService
	pathToVariable string
	secureVariable *SecureVariable
	region *string
	namespace *string
}

func (r ApiUpsertSecureVariableRequest) SecureVariable(secureVariable SecureVariable) ApiUpsertSecureVariableRequest {
	r.secureVariable = &secureVariable
	return r
}

// Filters results based on the specified region.
func (r ApiUpsertSecureVariableRequest) Region(region string) ApiUpsertSecureVariableRequest {
	r.region = &region
	return r
}

// Filters results based on the specified namespace.
func (r ApiUpsertSecureVariableRequest) Namespace(namespace string) ApiUpsertSecureVariableRequest {
	r.namespace = &namespace
	return r
}

func (r ApiUpsertSecureVariableRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpsertSecureVariableExecute(r)
}

/*
UpsertSecureVariable Store a secure variable bag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pathToVariable Slash-delimited path to the read, store, or delete a variable
 @return ApiUpsertSecureVariableRequest
*/
func (a *VariableOperationsApiService) UpsertSecureVariable(ctx context.Context, pathToVariable string) ApiUpsertSecureVariableRequest {
	return ApiUpsertSecureVariableRequest{
		ApiService: a,
		ctx: ctx,
		pathToVariable: pathToVariable,
	}
}

// Execute executes the request
func (a *VariableOperationsApiService) UpsertSecureVariableExecute(r ApiUpsertSecureVariableRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariableOperationsApiService.UpsertSecureVariable")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/var/{pathToVariable}"
	localVarPath = strings.Replace(localVarPath, "{"+"pathToVariable"+"}", url.PathEscape(parameterToString(r.pathToVariable, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.secureVariable == nil {
		return nil, reportError("secureVariable is required and must be specified")
	}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.secureVariable
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Nomad-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
