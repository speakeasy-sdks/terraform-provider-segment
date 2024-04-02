// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type ListUserGroupsRequest struct {
	// Pagination for user groups.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListUserGroupsRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListUserGroupsIAMGroupsResponse200ResponseBody - OK
type ListUserGroupsIAMGroupsResponse200ResponseBody struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

func (o *ListUserGroupsIAMGroupsResponse200ResponseBody) GetData() *shared.ListUserGroupsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListUserGroupsIAMGroupsResponseResponseBody - OK
type ListUserGroupsIAMGroupsResponseResponseBody struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

func (o *ListUserGroupsIAMGroupsResponseResponseBody) GetData() *shared.ListUserGroupsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListUserGroupsIAMGroupsResponseBody - OK
type ListUserGroupsIAMGroupsResponseBody struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

func (o *ListUserGroupsIAMGroupsResponseBody) GetData() *shared.ListUserGroupsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListUserGroupsResponseBody - OK
type ListUserGroupsResponseBody struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

func (o *ListUserGroupsResponseBody) GetData() *shared.ListUserGroupsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListUserGroupsResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListUserGroupsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListUserGroupsIAMGroupsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListUserGroupsIAMGroupsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListUserGroupsIAMGroupsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUserGroupsResponse) GetTwoHundredApplicationJSONObject() *ListUserGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListUserGroupsResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListUserGroupsIAMGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListUserGroupsResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListUserGroupsIAMGroupsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListUserGroupsResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListUserGroupsIAMGroupsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListUserGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserGroupsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListUserGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}