// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type GetUserRequest struct {
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *GetUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// GetUserIAMUsersResponse200ResponseBody - OK
type GetUserIAMUsersResponse200ResponseBody struct {
	// Returns the user.
	Data *shared.GetUserV1Output `json:"data,omitempty"`
}

func (o *GetUserIAMUsersResponse200ResponseBody) GetData() *shared.GetUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetUserIAMUsersResponseResponseBody - OK
type GetUserIAMUsersResponseResponseBody struct {
	// Returns the user.
	Data *shared.GetUserV1Output `json:"data,omitempty"`
}

func (o *GetUserIAMUsersResponseResponseBody) GetData() *shared.GetUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetUserIAMUsersResponseBody - OK
type GetUserIAMUsersResponseBody struct {
	// Returns the user.
	Data *shared.GetUserV1Output `json:"data,omitempty"`
}

func (o *GetUserIAMUsersResponseBody) GetData() *shared.GetUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetUserResponseBody - OK
type GetUserResponseBody struct {
	// Returns the user.
	Data *shared.GetUserV1Output `json:"data,omitempty"`
}

func (o *GetUserResponseBody) GetData() *shared.GetUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetUserResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetUserResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetUserIAMUsersResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetUserIAMUsersResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetUserIAMUsersResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetUserResponse) GetTwoHundredApplicationJSONObject() *GetUserResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetUserResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetUserIAMUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetUserResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetUserIAMUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetUserResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetUserIAMUsersResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *GetUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
