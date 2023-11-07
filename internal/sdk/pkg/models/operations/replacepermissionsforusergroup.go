// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type ReplacePermissionsForUserGroupRequest struct {
	ReplacePermissionsForUserGroupV1Input shared.ReplacePermissionsForUserGroupV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserGroupID                           string                                       `pathParam:"style=simple,explode=false,name=userGroupId"`
}

func (o *ReplacePermissionsForUserGroupRequest) GetReplacePermissionsForUserGroupV1Input() shared.ReplacePermissionsForUserGroupV1Input {
	if o == nil {
		return shared.ReplacePermissionsForUserGroupV1Input{}
	}
	return o.ReplacePermissionsForUserGroupV1Input
}

func (o *ReplacePermissionsForUserGroupRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// ReplacePermissionsForUserGroupIAMGroupsResponse200ResponseBody - OK
type ReplacePermissionsForUserGroupIAMGroupsResponse200ResponseBody struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserGroupIAMGroupsResponse200ResponseBody) GetData() *shared.ReplacePermissionsForUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplacePermissionsForUserGroupIAMGroupsResponseResponseBody - OK
type ReplacePermissionsForUserGroupIAMGroupsResponseResponseBody struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserGroupIAMGroupsResponseResponseBody) GetData() *shared.ReplacePermissionsForUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplacePermissionsForUserGroupIAMGroupsResponseBody - OK
type ReplacePermissionsForUserGroupIAMGroupsResponseBody struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserGroupIAMGroupsResponseBody) GetData() *shared.ReplacePermissionsForUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplacePermissionsForUserGroupResponseBody - OK
type ReplacePermissionsForUserGroupResponseBody struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserGroupResponseBody) GetData() *shared.ReplacePermissionsForUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ReplacePermissionsForUserGroupResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ReplacePermissionsForUserGroupResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ReplacePermissionsForUserGroupIAMGroupsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ReplacePermissionsForUserGroupIAMGroupsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ReplacePermissionsForUserGroupIAMGroupsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReplacePermissionsForUserGroupResponse) GetTwoHundredApplicationJSONObject() *ReplacePermissionsForUserGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ReplacePermissionsForUserGroupResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ReplacePermissionsForUserGroupIAMGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ReplacePermissionsForUserGroupResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ReplacePermissionsForUserGroupIAMGroupsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ReplacePermissionsForUserGroupResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ReplacePermissionsForUserGroupIAMGroupsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ReplacePermissionsForUserGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplacePermissionsForUserGroupResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ReplacePermissionsForUserGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplacePermissionsForUserGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
