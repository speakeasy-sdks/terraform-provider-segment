// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type CreateSourceRegulationRequest struct {
	CreateSourceRegulationV1Input shared.CreateSourceRegulationV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	SourceID                      string                               `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *CreateSourceRegulationRequest) GetCreateSourceRegulationV1Input() shared.CreateSourceRegulationV1Input {
	if o == nil {
		return shared.CreateSourceRegulationV1Input{}
	}
	return o.CreateSourceRegulationV1Input
}

func (o *CreateSourceRegulationRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// CreateSourceRegulationDeletionAndSuppressionResponse200ResponseBody - OK
type CreateSourceRegulationDeletionAndSuppressionResponse200ResponseBody struct {
	// The output of a create Source regulation call.
	Data *shared.CreateSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateSourceRegulationDeletionAndSuppressionResponse200ResponseBody) GetData() *shared.CreateSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateSourceRegulationDeletionAndSuppressionResponseResponseBody - OK
type CreateSourceRegulationDeletionAndSuppressionResponseResponseBody struct {
	// The output of a create Source regulation call.
	Data *shared.CreateSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateSourceRegulationDeletionAndSuppressionResponseResponseBody) GetData() *shared.CreateSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateSourceRegulationDeletionAndSuppressionResponseBody - OK
type CreateSourceRegulationDeletionAndSuppressionResponseBody struct {
	// The output of a create Source regulation call.
	Data *shared.CreateSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateSourceRegulationDeletionAndSuppressionResponseBody) GetData() *shared.CreateSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateSourceRegulationResponseBody - OK
type CreateSourceRegulationResponseBody struct {
	// The output of a create Source regulation call.
	Data *shared.CreateSourceRegulationV1Output `json:"data,omitempty"`
}

func (o *CreateSourceRegulationResponseBody) GetData() *shared.CreateSourceRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateSourceRegulationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateSourceRegulationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateSourceRegulationDeletionAndSuppressionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateSourceRegulationDeletionAndSuppressionResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateSourceRegulationDeletionAndSuppressionResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSourceRegulationResponse) GetTwoHundredApplicationJSONObject() *CreateSourceRegulationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateSourceRegulationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateSourceRegulationDeletionAndSuppressionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateSourceRegulationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateSourceRegulationDeletionAndSuppressionResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateSourceRegulationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateSourceRegulationDeletionAndSuppressionResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateSourceRegulationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceRegulationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateSourceRegulationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceRegulationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
