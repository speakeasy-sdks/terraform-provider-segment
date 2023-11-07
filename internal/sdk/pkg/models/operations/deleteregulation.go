// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type DeleteRegulationRequest struct {
	RegulateID string `pathParam:"style=simple,explode=false,name=regulateId"`
}

func (o *DeleteRegulationRequest) GetRegulateID() string {
	if o == nil {
		return ""
	}
	return o.RegulateID
}

// DeleteRegulationDeletionAndSuppressionResponse200ResponseBody - OK
type DeleteRegulationDeletionAndSuppressionResponse200ResponseBody struct {
	// The output of the delete regulation call.
	Data *shared.DeleteRegulationV1Output `json:"data,omitempty"`
}

func (o *DeleteRegulationDeletionAndSuppressionResponse200ResponseBody) GetData() *shared.DeleteRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteRegulationDeletionAndSuppressionResponseResponseBody - OK
type DeleteRegulationDeletionAndSuppressionResponseResponseBody struct {
	// The output of the delete regulation call.
	Data *shared.DeleteRegulationV1Output `json:"data,omitempty"`
}

func (o *DeleteRegulationDeletionAndSuppressionResponseResponseBody) GetData() *shared.DeleteRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteRegulationDeletionAndSuppressionResponseBody - OK
type DeleteRegulationDeletionAndSuppressionResponseBody struct {
	// The output of the delete regulation call.
	Data *shared.DeleteRegulationV1Output `json:"data,omitempty"`
}

func (o *DeleteRegulationDeletionAndSuppressionResponseBody) GetData() *shared.DeleteRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteRegulationResponseBody - OK
type DeleteRegulationResponseBody struct {
	// The output of the delete regulation call.
	Data *shared.DeleteRegulationV1Output `json:"data,omitempty"`
}

func (o *DeleteRegulationResponseBody) GetData() *shared.DeleteRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteRegulationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *DeleteRegulationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *DeleteRegulationDeletionAndSuppressionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *DeleteRegulationDeletionAndSuppressionResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *DeleteRegulationDeletionAndSuppressionResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteRegulationResponse) GetTwoHundredApplicationJSONObject() *DeleteRegulationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *DeleteRegulationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *DeleteRegulationDeletionAndSuppressionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteRegulationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *DeleteRegulationDeletionAndSuppressionResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteRegulationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *DeleteRegulationDeletionAndSuppressionResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *DeleteRegulationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteRegulationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteRegulationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteRegulationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
