// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// EchoAlphaOutputMethod - The HTTP method used for this round-trip.
//
// Currently, this endpoint supports only `get` and `post` methods.
type EchoAlphaOutputMethod string

const (
	EchoAlphaOutputMethodGet  EchoAlphaOutputMethod = "get"
	EchoAlphaOutputMethodPost EchoAlphaOutputMethod = "post"
)

func (e EchoAlphaOutputMethod) ToPointer() *EchoAlphaOutputMethod {
	return &e
}

func (e *EchoAlphaOutputMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "get":
		fallthrough
	case "post":
		*e = EchoAlphaOutputMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EchoAlphaOutputMethod: %v", v)
	}
}

// EchoAlphaOutput - Echo response.
type EchoAlphaOutput struct {
	// The request's HTTP headers.
	Headers map[string]interface{} `json:"headers"`
	// The string passed in the `message` input field.
	Message string `json:"message"`
	// The HTTP method used for this round-trip.
	//
	// Currently, this endpoint supports only `get` and `post` methods.
	Method EchoAlphaOutputMethod `json:"method"`
}