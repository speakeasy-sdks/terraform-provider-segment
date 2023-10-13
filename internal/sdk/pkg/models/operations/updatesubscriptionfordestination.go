// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type UpdateSubscriptionForDestinationRequest struct {
	UpdateSubscriptionForDestinationAlphaInput shared.UpdateSubscriptionForDestinationAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	DestinationID                              string                                            `pathParam:"style=simple,explode=false,name=destinationId"`
	ID                                         string                                            `pathParam:"style=simple,explode=false,name=id"`
}

// UpdateSubscriptionForDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateSubscriptionForDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the updated Destination subscription.
	Data *shared.UpdateSubscriptionForDestinationAlphaOutput `json:"data,omitempty"`
}

type UpdateSubscriptionForDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	UpdateSubscriptionForDestination200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateSubscriptionForDestination200ApplicationVndSegmentV1alphaPlusJSON
}
