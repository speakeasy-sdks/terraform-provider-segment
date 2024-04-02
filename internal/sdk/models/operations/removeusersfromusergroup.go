// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
	"net/http"
)

type RemoveUsersFromUserGroupRequest struct {
	// The list of emails to remove from the user group.
	//
	// This parameter exists in v1.
	Emails      []string `queryParam:"style=deepObject,explode=true,name=emails"`
	UserGroupID string   `pathParam:"style=simple,explode=false,name=userGroupId"`
}

func (o *RemoveUsersFromUserGroupRequest) GetEmails() []string {
	if o == nil {
		return []string{}
	}
	return o.Emails
}

func (o *RemoveUsersFromUserGroupRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// RemoveUsersFromUserGroupIAMGroupsResponse200ResponseBody - OK
type RemoveUsersFromUserGroupIAMGroupsResponse200ResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *RemoveUsersFromUserGroupIAMGroupsResponse200ResponseBody) GetData() *shared.RemoveUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveUsersFromUserGroupIAMGroupsResponseResponseBody - OK
type RemoveUsersFromUserGroupIAMGroupsResponseResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *RemoveUsersFromUserGroupIAMGroupsResponseResponseBody) GetData() *shared.RemoveUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveUsersFromUserGroupIAMGroupsResponseBody - OK
type RemoveUsersFromUserGroupIAMGroupsResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *RemoveUsersFromUserGroupIAMGroupsResponseBody) GetData() *shared.RemoveUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveUsersFromUserGroupResponseBody - OK
type RemoveUsersFromUserGroupResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.RemoveUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *RemoveUsersFromUserGroupResponseBody) GetData() *shared.RemoveUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveUsersFromUserGroupResponse struct {
	// OK
	TwoHundredApplicationJSONObject *RemoveUsersFromUserGroupResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *RemoveUsersFromUserGroupIAMGroupsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *RemoveUsersFromUserGroupIAMGroupsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *RemoveUsersFromUserGroupIAMGroupsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveUsersFromUserGroupResponse) GetTwoHundredApplicationJSONObject() *RemoveUsersFromUserGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *RemoveUsersFromUserGroupResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *RemoveUsersFromUserGroupIAMGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *RemoveUsersFromUserGroupResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *RemoveUsersFromUserGroupIAMGroupsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *RemoveUsersFromUserGroupResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *RemoveUsersFromUserGroupIAMGroupsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *RemoveUsersFromUserGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveUsersFromUserGroupResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveUsersFromUserGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveUsersFromUserGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}