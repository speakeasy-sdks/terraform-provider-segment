// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// EchoV1OutputMethod - The HTTP method used for this round-trip.
//
// Currently, this endpoint supports only `get` and `post` methods.
type EchoV1OutputMethod string

const (
	EchoV1OutputMethodGet  EchoV1OutputMethod = "get"
	EchoV1OutputMethodPost EchoV1OutputMethod = "post"
)

func (e EchoV1OutputMethod) ToPointer() *EchoV1OutputMethod {
	return &e
}

func (e *EchoV1OutputMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "get":
		fallthrough
	case "post":
		*e = EchoV1OutputMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EchoV1OutputMethod: %v", v)
	}
}

// EchoV1Output - Echo response.
type EchoV1Output struct {
	// The request's HTTP headers.
	Headers map[string]interface{} `json:"headers"`
	// The string passed in the `message` input field.
	Message string `json:"message"`
	// The HTTP method used for this round-trip.
	//
	// Currently, this endpoint supports only `get` and `post` methods.
	Method EchoV1OutputMethod `json:"method"`
}
