# SecureVariableStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Nomad namespace containing the variable. | [optional] [default to "default"]
**Path** | Pointer to **string** | The path at which the variable is stored. | [optional] 
**CreateIndex** | Pointer to **int64** |  | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifyIndex** | Pointer to **int64** |  | [optional] [readonly] 
**ModifyTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Meta** | Pointer to **map[string]string** | Collection of user provided key/value pairs. These are stored unencrypted and can be seen by any user possessing the &#x60;list&#x60; permission. | [optional] 
**Version** | Pointer to **int32** | **RESERVED FOR FUTURE USE**   The version index for this secure variable. Currently, all objects will have a version of 1 | [optional] 

## Methods

### NewSecureVariableStub

`func NewSecureVariableStub() *SecureVariableStub`

NewSecureVariableStub instantiates a new SecureVariableStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureVariableStubWithDefaults

`func NewSecureVariableStubWithDefaults() *SecureVariableStub`

NewSecureVariableStubWithDefaults instantiates a new SecureVariableStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *SecureVariableStub) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SecureVariableStub) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SecureVariableStub) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SecureVariableStub) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPath

`func (o *SecureVariableStub) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecureVariableStub) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecureVariableStub) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecureVariableStub) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCreateIndex

`func (o *SecureVariableStub) GetCreateIndex() int64`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *SecureVariableStub) GetCreateIndexOk() (*int64, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *SecureVariableStub) SetCreateIndex(v int64)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *SecureVariableStub) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetCreateTime

`func (o *SecureVariableStub) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *SecureVariableStub) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *SecureVariableStub) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *SecureVariableStub) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetModifyIndex

`func (o *SecureVariableStub) GetModifyIndex() int64`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *SecureVariableStub) GetModifyIndexOk() (*int64, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *SecureVariableStub) SetModifyIndex(v int64)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *SecureVariableStub) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetModifyTime

`func (o *SecureVariableStub) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *SecureVariableStub) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *SecureVariableStub) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *SecureVariableStub) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetMeta

`func (o *SecureVariableStub) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecureVariableStub) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecureVariableStub) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecureVariableStub) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetVersion

`func (o *SecureVariableStub) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecureVariableStub) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecureVariableStub) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecureVariableStub) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


