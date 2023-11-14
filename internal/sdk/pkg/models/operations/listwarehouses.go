// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListWarehousesRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListWarehousesRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListWarehousesWarehousesResponse200ResponseBody - OK
type ListWarehousesWarehousesResponse200ResponseBody struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

func (o *ListWarehousesWarehousesResponse200ResponseBody) GetData() *shared.ListWarehousesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListWarehousesWarehousesResponseResponseBody - OK
type ListWarehousesWarehousesResponseResponseBody struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

func (o *ListWarehousesWarehousesResponseResponseBody) GetData() *shared.ListWarehousesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListWarehousesWarehousesResponseBody - OK
type ListWarehousesWarehousesResponseBody struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

func (o *ListWarehousesWarehousesResponseBody) GetData() *shared.ListWarehousesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListWarehousesResponseBody - OK
type ListWarehousesResponseBody struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

func (o *ListWarehousesResponseBody) GetData() *shared.ListWarehousesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListWarehousesResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListWarehousesWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListWarehousesWarehousesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListWarehousesWarehousesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListWarehousesResponse) GetTwoHundredApplicationJSONObject() *ListWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListWarehousesResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListWarehousesWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListWarehousesResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListWarehousesWarehousesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListWarehousesResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListWarehousesWarehousesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListWarehousesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWarehousesResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListWarehousesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWarehousesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
