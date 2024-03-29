// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type ListSourcesRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListSourcesRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListSourcesSourcesResponse200ResponseBody - OK
type ListSourcesSourcesResponse200ResponseBody struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesV1Output `json:"data,omitempty"`
}

func (o *ListSourcesSourcesResponse200ResponseBody) GetData() *shared.ListSourcesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListSourcesSourcesResponseResponseBody - OK
type ListSourcesSourcesResponseResponseBody struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesAlphaOutput `json:"data,omitempty"`
}

func (o *ListSourcesSourcesResponseResponseBody) GetData() *shared.ListSourcesAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListSourcesSourcesResponseBody - OK
type ListSourcesSourcesResponseBody struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesV1Output `json:"data,omitempty"`
}

func (o *ListSourcesSourcesResponseBody) GetData() *shared.ListSourcesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListSourcesResponseBody - OK
type ListSourcesResponseBody struct {
	// Returns a list of Sources that belong to the current Workspace.
	Data *shared.ListSourcesV1Output `json:"data,omitempty"`
}

func (o *ListSourcesResponseBody) GetData() *shared.ListSourcesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListSourcesResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListSourcesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListSourcesSourcesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListSourcesSourcesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListSourcesSourcesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListSourcesResponse) GetTwoHundredApplicationJSONObject() *ListSourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListSourcesResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListSourcesSourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListSourcesResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListSourcesSourcesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListSourcesResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListSourcesSourcesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListSourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSourcesResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListSourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
