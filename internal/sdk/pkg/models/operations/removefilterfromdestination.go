// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RemoveFilterFromDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
	FilterID      string `pathParam:"style=simple,explode=false,name=filterId"`
}

func (o *RemoveFilterFromDestinationRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *RemoveFilterFromDestinationRequest) GetFilterID() string {
	if o == nil {
		return ""
	}
	return o.FilterID
}

// RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON - OK
type RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON - OK
type RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveFilterFromDestination200ApplicationJSON - OK
type RemoveFilterFromDestination200ApplicationJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestination200ApplicationJSON) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveFilterFromDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	RemoveFilterFromDestination200ApplicationJSONObject *RemoveFilterFromDestination200ApplicationJSON
	// OK
	RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSONObject *RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON
	// OK
	RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSONObject *RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON
}

func (o *RemoveFilterFromDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveFilterFromDestinationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveFilterFromDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveFilterFromDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveFilterFromDestinationResponse) GetRemoveFilterFromDestination200ApplicationJSONObject() *RemoveFilterFromDestination200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.RemoveFilterFromDestination200ApplicationJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetRemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSONObject() *RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetRemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject() *RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetRemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSONObject() *RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSONObject
}
