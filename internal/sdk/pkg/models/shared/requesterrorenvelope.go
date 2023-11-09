// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestErrorEnvelope - Envelope type for RequestErrors.
type RequestErrorEnvelope struct {
	Errors []RequestError `json:"errors"`
}

func (o *RequestErrorEnvelope) GetErrors() []RequestError {
	if o == nil {
		return []RequestError{}
	}
	return o.Errors
}
