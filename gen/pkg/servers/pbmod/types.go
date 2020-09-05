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

// RetrieveResponse ...
type RetrieveResponse struct {
	Content []string `json:"content"`
}

// GetResourceRequest ...
type GetResourceRequest struct {
	Resource string
	Version  string
}

// *RetrieveResponse validator
func (s *RetrieveResponse) Validate() error {
	return validator.Validate(s)
}