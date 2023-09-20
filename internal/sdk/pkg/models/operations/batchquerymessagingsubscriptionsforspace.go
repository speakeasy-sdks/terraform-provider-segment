// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type BatchQueryMessagingSubscriptionsForSpaceRequest struct {
	BatchQueryMessagingSubscriptionsForSpaceAlphaInput shared.BatchQueryMessagingSubscriptionsForSpaceAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	SpaceID                                            string                                                    `pathParam:"style=simple,explode=false,name=spaceId"`
}

// BatchQueryMessagingSubscriptionsForSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type BatchQueryMessagingSubscriptionsForSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Batch get response.
	Data *shared.BatchQueryMessagingSubscriptionsForSpaceAlphaOutput `json:"data,omitempty"`
}

type BatchQueryMessagingSubscriptionsForSpaceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	BatchQueryMessagingSubscriptionsForSpace200ApplicationVndSegmentV1alphaPlusJSONObject *BatchQueryMessagingSubscriptionsForSpace200ApplicationVndSegmentV1alphaPlusJSON
}