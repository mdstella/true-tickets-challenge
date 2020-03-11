// Package errors will have the errors definitions and the different way to create different status code errors.
package errors

import (
	"fmt"
	"net/http"
	"strconv"
)

// CoreError represent all the business errors will be from this struct type
//{
//   "message": "Wrong metric key",
//   "type": "BAD_PARAM_ERROR",
//   "code": 400
//}
type CoreError struct {
	// TODO consider having an array of errors instead of just one
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (err CoreError) Error() string {
	return fmt.Sprintf("Message:'%s', Type:%s, Code:%d", err.Message, err.Type, err.Code)
}

// StatusCode is the implementation of StatusCoder interface to be able to override the status code on the output by go-kit
func (err CoreError) StatusCode() int {
	return err.Code
}

// MarshalJSON is the implementation of Marshaler interface to be able to marshal the JSON
func (err CoreError) MarshalJSON() ([]byte, error) {
	return []byte(`{"message": "` + err.Message + `", "type": "` + err.Type + `", "code": ` + strconv.Itoa(err.Code) + `}`), nil
}

// NewBadParamError creates a new bad param error on the core
func NewBadParamError(message string) *CoreError {
	return &CoreError{Message: message, Code: http.StatusBadRequest, Type: "BAD_PARAM_ERROR"}
}

// GetCoreError retrieve the pointer to be able to use when we do the error http encoding in main file
func GetCoreError(err error) *CoreError {
	if err == nil {
		return nil
	}
	if nerr2, ok := err.(*CoreError); ok {
		return nerr2
	}
	if nerr2, ok := err.(CoreError); ok {
		return &nerr2
	}
	return nil
}
