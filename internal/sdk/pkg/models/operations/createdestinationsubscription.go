// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type CreateDestinationSubscriptionRequest struct {
	CreateDestinationSubscriptionAlphaInput shared.CreateDestinationSubscriptionAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	DestinationID                           string                                         `pathParam:"style=simple,explode=false,name=destinationId"`
}

// CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a newly created Destination subscription.
	Data *shared.CreateDestinationSubscriptionAlphaOutput `json:"data,omitempty"`
}

type CreateDestinationSubscriptionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSONObject *CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSON
}