# \VariableOperationsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSecureVariable**](VariableOperationsApi.md#DeleteSecureVariable) | **Delete** /var/{pathToVariable} | Delete a secure variable bag
[**GetSecureVariable**](VariableOperationsApi.md#GetSecureVariable) | **Get** /var/{pathToVariable} | Fetch a secure variables bag
[**ListVars**](VariableOperationsApi.md#ListVars) | **Get** /vars | List the variable bags
[**UpsertSecureVariable**](VariableOperationsApi.md#UpsertSecureVariable) | **Put** /var/{pathToVariable} | Store a secure variable bag



## DeleteSecureVariable

> DeleteSecureVariable(ctx, pathToVariable).Region(region).Namespace(namespace).Execute()

Delete a secure variable bag

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pathToVariable := "path/to/var" // string | Slash-delimited path to the read, store, or delete a variable
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableOperationsApi.DeleteSecureVariable(context.Background(), pathToVariable).Region(region).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableOperationsApi.DeleteSecureVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathToVariable** | **string** | Slash-delimited path to the read, store, or delete a variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecureVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecureVariable

> SecureVariable GetSecureVariable(ctx, pathToVariable).Execute()

Fetch a secure variables bag

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pathToVariable := "path/to/var" // string | Slash-delimited path to the read, store, or delete a variable

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableOperationsApi.GetSecureVariable(context.Background(), pathToVariable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableOperationsApi.GetSecureVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecureVariable`: SecureVariable
    fmt.Fprintf(os.Stdout, "Response from `VariableOperationsApi.GetSecureVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathToVariable** | **string** | Slash-delimited path to the read, store, or delete a variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecureVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecureVariable**](SecureVariable.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVars

> []SecureVariableStub ListVars(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).PerPage(perPage).NextToken(nextToken).Execute()

List the variable bags

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    index := int32(56) // int32 | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait := "wait_example" // string | Provided with IndexParam to wait for change. (optional)
    stale := "stale_example" // string | If present, results will include stale reads. (optional)
    prefix := "prefix_example" // string | Constrains results to jobs that start with the defined prefix (optional)
    perPage := int32(56) // int32 | Maximum number of results to return. (optional)
    nextToken := "nextToken_example" // string | Indicates where to start paging for queries that support pagination. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableOperationsApi.ListVars(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableOperationsApi.ListVars``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVars`: []SecureVariableStub
    fmt.Fprintf(os.Stdout, "Response from `VariableOperationsApi.ListVars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **index** | **int32** | If set, wait until query exceeds given index. Must be provided with WaitParam. | 
 **wait** | **string** | Provided with IndexParam to wait for change. | 
 **stale** | **string** | If present, results will include stale reads. | 
 **prefix** | **string** | Constrains results to jobs that start with the defined prefix | 
 **perPage** | **int32** | Maximum number of results to return. | 
 **nextToken** | **string** | Indicates where to start paging for queries that support pagination. | 

### Return type

[**[]SecureVariableStub**](SecureVariableStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertSecureVariable

> UpsertSecureVariable(ctx, pathToVariable).SecureVariable(secureVariable).Region(region).Namespace(namespace).Execute()

Store a secure variable bag

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pathToVariable := "path/to/var" // string | Slash-delimited path to the read, store, or delete a variable
    secureVariable := *openapiclient.NewSecureVariable() // SecureVariable | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableOperationsApi.UpsertSecureVariable(context.Background(), pathToVariable).SecureVariable(secureVariable).Region(region).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableOperationsApi.UpsertSecureVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathToVariable** | **string** | Slash-delimited path to the read, store, or delete a variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertSecureVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secureVariable** | [**SecureVariable**](SecureVariable.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

