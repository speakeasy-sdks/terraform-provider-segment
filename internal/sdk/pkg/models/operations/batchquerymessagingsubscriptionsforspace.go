// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type BatchQueryMessagingSubscriptionsForSpaceRequest struct {
	BatchQueryMessagingSubscriptionsForSpaceAlphaInput shared.BatchQueryMessagingSubscriptionsForSpaceAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	SpaceID                                            string                                                    `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *BatchQueryMessagingSubscriptionsForSpaceRequest) GetBatchQueryMessagingSubscriptionsForSpaceAlphaInput() shared.BatchQueryMessagingSubscriptionsForSpaceAlphaInput {
	if o == nil {
		return shared.BatchQueryMessagingSubscriptionsForSpaceAlphaInput{}
	}
	return o.BatchQueryMessagingSubscriptionsForSpaceAlphaInput
}

func (o *BatchQueryMessagingSubscriptionsForSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// BatchQueryMessagingSubscriptionsForSpaceResponseBody - OK
type BatchQueryMessagingSubscriptionsForSpaceResponseBody struct {
	// Batch get response.
	Data *shared.BatchQueryMessagingSubscriptionsForSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *BatchQueryMessagingSubscriptionsForSpaceResponseBody) GetData() *shared.BatchQueryMessagingSubscriptionsForSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type BatchQueryMessagingSubscriptionsForSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *BatchQueryMessagingSubscriptionsForSpaceResponseBody
}

func (o *BatchQueryMessagingSubscriptionsForSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BatchQueryMessagingSubscriptionsForSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BatchQueryMessagingSubscriptionsForSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *BatchQueryMessagingSubscriptionsForSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *BatchQueryMessagingSubscriptionsForSpaceResponse) GetObject() *BatchQueryMessagingSubscriptionsForSpaceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
