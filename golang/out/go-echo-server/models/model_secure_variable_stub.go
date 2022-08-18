package models

import (
	"time"
)

// SecureVariableStub - A `SecureVariableStub` object contains: - the `UserMeta` collection - version information - create and modify index and timestamp  
type SecureVariableStub struct {

	// Nomad namespace containing the variable.
	Namespace string `json:"Namespace,omitempty"`

	// The path at which the variable is stored.
	Path string `json:"Path,omitempty"`

	CreateIndex int64 `json:"CreateIndex,omitempty"`

	CreateTime time.Time `json:"CreateTime,omitempty"`

	ModifyIndex int64 `json:"ModifyIndex,omitempty"`

	ModifyTime time.Time `json:"ModifyTime,omitempty"`

	// Collection of user provided key/value pairs. These are stored unencrypted and can be seen by any user possessing the `list` permission.
	Meta map[string]string `json:"Meta,omitempty"`

	// **RESERVED FOR FUTURE USE**   The version index for this secure variable. Currently, all objects will have a version of 1
	Version int32 `json:"Version,omitempty"`
}
