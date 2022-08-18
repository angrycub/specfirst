/*
Nomad Secure Variables

Partial API specification for Nomad's secure variables feature

API version: 1.4.0
Contact: support@hashicorp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// SecureVariableStub A `SecureVariableStub` object contains: - the `UserMeta` collection - version information - create and modify index and timestamp  
type SecureVariableStub struct {
	// Nomad namespace containing the variable.
	Namespace *string `json:"Namespace,omitempty"`
	// The path at which the variable is stored.
	Path *string `json:"Path,omitempty"`
	CreateIndex *int64 `json:"CreateIndex,omitempty"`
	CreateTime *time.Time `json:"CreateTime,omitempty"`
	ModifyIndex *int64 `json:"ModifyIndex,omitempty"`
	ModifyTime *time.Time `json:"ModifyTime,omitempty"`
	// Collection of user provided key/value pairs. These are stored unencrypted and can be seen by any user possessing the `list` permission.
	Meta *map[string]string `json:"Meta,omitempty"`
	// **RESERVED FOR FUTURE USE**   The version index for this secure variable. Currently, all objects will have a version of 1
	Version *int32 `json:"Version,omitempty"`
}

// NewSecureVariableStub instantiates a new SecureVariableStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecureVariableStub() *SecureVariableStub {
	this := SecureVariableStub{}
	var namespace string = "default"
	this.Namespace = &namespace
	return &this
}

// NewSecureVariableStubWithDefaults instantiates a new SecureVariableStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecureVariableStubWithDefaults() *SecureVariableStub {
	this := SecureVariableStub{}
	var namespace string = "default"
	this.Namespace = &namespace
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SecureVariableStub) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SecureVariableStub) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SecureVariableStub) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SecureVariableStub) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SecureVariableStub) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SecureVariableStub) SetPath(v string) {
	o.Path = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *SecureVariableStub) GetCreateIndex() int64 {
	if o == nil || o.CreateIndex == nil {
		var ret int64
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetCreateIndexOk() (*int64, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *SecureVariableStub) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int64 and assigns it to the CreateIndex field.
func (o *SecureVariableStub) SetCreateIndex(v int64) {
	o.CreateIndex = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *SecureVariableStub) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *SecureVariableStub) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *SecureVariableStub) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *SecureVariableStub) GetModifyIndex() int64 {
	if o == nil || o.ModifyIndex == nil {
		var ret int64
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetModifyIndexOk() (*int64, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *SecureVariableStub) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int64 and assigns it to the ModifyIndex field.
func (o *SecureVariableStub) SetModifyIndex(v int64) {
	o.ModifyIndex = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *SecureVariableStub) GetModifyTime() time.Time {
	if o == nil || o.ModifyTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetModifyTimeOk() (*time.Time, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *SecureVariableStub) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given time.Time and assigns it to the ModifyTime field.
func (o *SecureVariableStub) SetModifyTime(v time.Time) {
	o.ModifyTime = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SecureVariableStub) GetMeta() map[string]string {
	if o == nil || o.Meta == nil {
		var ret map[string]string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetMetaOk() (*map[string]string, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SecureVariableStub) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *SecureVariableStub) SetMeta(v map[string]string) {
	o.Meta = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecureVariableStub) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureVariableStub) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecureVariableStub) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecureVariableStub) SetVersion(v int32) {
	o.Version = &v
}

func (o SecureVariableStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.CreateTime != nil {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.ModifyTime != nil {
		toSerialize["ModifyTime"] = o.ModifyTime
	}
	if o.Meta != nil {
		toSerialize["Meta"] = o.Meta
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSecureVariableStub struct {
	value *SecureVariableStub
	isSet bool
}

func (v NullableSecureVariableStub) Get() *SecureVariableStub {
	return v.value
}

func (v *NullableSecureVariableStub) Set(val *SecureVariableStub) {
	v.value = val
	v.isSet = true
}

func (v NullableSecureVariableStub) IsSet() bool {
	return v.isSet
}

func (v *NullableSecureVariableStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecureVariableStub(val *SecureVariableStub) *NullableSecureVariableStub {
	return &NullableSecureVariableStub{value: val, isSet: true}
}

func (v NullableSecureVariableStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecureVariableStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

