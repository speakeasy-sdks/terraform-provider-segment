// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListSubscriptionsFromDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
	// Pagination options.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListSubscriptionsFromDestinationRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *ListSubscriptionsFromDestinationRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for ListDestinationSubscriptionsAlpha.
	Data *shared.ListSubscriptionsFromDestinationAlphaOutput `json:"data,omitempty"`
}

func (o *ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.ListSubscriptionsFromDestinationAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListSubscriptionsFromDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject *ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *ListSubscriptionsFromDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSubscriptionsFromDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSubscriptionsFromDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSubscriptionsFromDestinationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListSubscriptionsFromDestinationResponse) GetListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject() *ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject
}
