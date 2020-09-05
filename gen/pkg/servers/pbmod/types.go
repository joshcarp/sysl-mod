// Code generated by sysl DO NOT EDIT.
package pbmod

import (
	"time"

	"github.com/anz-bank/sysl-go/validator"

	"github.com/rickb777/date"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// KeyValue ...
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// RetrieveResponse ...
type RetrieveResponse struct {
	Content []KeyValue `json:"content"`
}

// GetResourceRequest ...
type GetResourceRequest struct {
	Resource string
	Version  string
}

// *KeyValue validator
func (s *KeyValue) Validate() error {
	return validator.Validate(s)
}

// *RetrieveResponse validator
func (s *RetrieveResponse) Validate() error {
	return validator.Validate(s)
}
