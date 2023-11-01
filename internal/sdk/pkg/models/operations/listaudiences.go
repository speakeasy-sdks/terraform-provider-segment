// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ListAudiencesRequest struct {
	// Information about the pagination of this response.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID    string                 `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *ListAudiencesRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListAudiencesRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// ListAudiences200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListAudiences200ApplicationVndSegmentV1alphaPlusJSON struct {
	// List audiences endpoint output.
	Data *shared.ListAudiencesAlphaOutput `json:"data,omitempty"`
}

func (o *ListAudiences200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.ListAudiencesAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListAudiencesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	ListAudiences200ApplicationVndSegmentV1alphaPlusJSONObject *ListAudiences200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *ListAudiencesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAudiencesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAudiencesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAudiencesResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListAudiencesResponse) GetListAudiences200ApplicationVndSegmentV1alphaPlusJSONObject() *ListAudiences200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ListAudiences200ApplicationVndSegmentV1alphaPlusJSONObject
}
