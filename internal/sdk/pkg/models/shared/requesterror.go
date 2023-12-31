// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestError - Represents any error that could have occurred while performing a request.
type RequestError struct {
	// Any extra data associated with this error.
	Data interface{} `json:"data,omitempty"`
	// The name of an input field from the request that triggered this error.
	Field *string `json:"field,omitempty"`
	// An error message attached to this error.
	Message *string `json:"message,omitempty"`
	// Http status code.
	Status *float64 `json:"status,omitempty"`
	// The type for this error (validation, server, unknown, etc).
	Type string `json:"type"`
}

func (o *RequestError) GetData() interface{} {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *RequestError) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

func (o *RequestError) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RequestError) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *RequestError) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
