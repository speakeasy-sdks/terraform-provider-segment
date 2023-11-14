// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateCloudSourceRegulationRequest struct {
	CreateCloudSourceRegulationV1Input shared.CreateCloudSourceRegulationV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	SourceID                           string                                    `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *CreateCloudSourceRegulationRequest) GetCreateCloudSourceRegulationV1Input() shared.CreateCloudSourceRegulationV1Input {
	if o == nil {
		return shared.CreateCloudSourceRegulationV1Input{}
	}
	return o.CreateCloudSourceRegulationV1Input
}

func (o *CreateCloudSourceRegulationRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// CreateCloudSourceRegulationDeletionAndSuppressionResponse200ResponseBody - OK
type CreateCloudSourceRegulationDeletionAndSuppressionResponse200ResponseBody struct {
	// The output of a create Cloud Source regulation call.
	Data *shared.CreateCloudSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateCloudSourceRegulationDeletionAndSuppressionResponse200ResponseBody) GetData() *shared.CreateCloudSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateCloudSourceRegulationDeletionAndSuppressionResponseResponseBody - OK
type CreateCloudSourceRegulationDeletionAndSuppressionResponseResponseBody struct {
	// The output of a create Cloud Source regulation call.
	Data *shared.CreateCloudSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateCloudSourceRegulationDeletionAndSuppressionResponseResponseBody) GetData() *shared.CreateCloudSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateCloudSourceRegulationDeletionAndSuppressionResponseBody - OK
type CreateCloudSourceRegulationDeletionAndSuppressionResponseBody struct {
	// The output of a create Cloud Source regulation call.
	Data *shared.CreateCloudSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateCloudSourceRegulationDeletionAndSuppressionResponseBody) GetData() *shared.CreateCloudSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateCloudSourceRegulationResponseBody - OK
type CreateCloudSourceRegulationResponseBody struct {
	// The output of a create Cloud Source regulation call.
	Data *shared.CreateCloudSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateCloudSourceRegulationResponseBody) GetData() *shared.CreateCloudSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateCloudSourceRegulationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateCloudSourceRegulationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateCloudSourceRegulationDeletionAndSuppressionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateCloudSourceRegulationDeletionAndSuppressionResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateCloudSourceRegulationDeletionAndSuppressionResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCloudSourceRegulationResponse) GetTwoHundredApplicationJSONObject() *CreateCloudSourceRegulationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateCloudSourceRegulationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateCloudSourceRegulationDeletionAndSuppressionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateCloudSourceRegulationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateCloudSourceRegulationDeletionAndSuppressionResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateCloudSourceRegulationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateCloudSourceRegulationDeletionAndSuppressionResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateCloudSourceRegulationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCloudSourceRegulationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateCloudSourceRegulationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCloudSourceRegulationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
