// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type AddPermissionsToUserGroupRequest struct {
	AddPermissionsToUserGroupV1Input shared.AddPermissionsToUserGroupV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserGroupID                      string                                  `pathParam:"style=simple,explode=false,name=userGroupId"`
}

func (o *AddPermissionsToUserGroupRequest) GetAddPermissionsToUserGroupV1Input() shared.AddPermissionsToUserGroupV1Input {
	if o == nil {
		return shared.AddPermissionsToUserGroupV1Input{}
	}
	return o.AddPermissionsToUserGroupV1Input
}

func (o *AddPermissionsToUserGroupRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.AddPermissionsToUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.AddPermissionsToUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON) GetData() *shared.AddPermissionsToUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddPermissionsToUserGroup200ApplicationJSON - OK
type AddPermissionsToUserGroup200ApplicationJSON struct {
	// Returns the group's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserGroupV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserGroup200ApplicationJSON) GetData() *shared.AddPermissionsToUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type AddPermissionsToUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AddPermissionsToUserGroup200ApplicationJSONObject *AddPermissionsToUserGroup200ApplicationJSON
	// OK
	AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSONObject *AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON
}

func (o *AddPermissionsToUserGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddPermissionsToUserGroupResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *AddPermissionsToUserGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddPermissionsToUserGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddPermissionsToUserGroupResponse) GetAddPermissionsToUserGroup200ApplicationJSONObject() *AddPermissionsToUserGroup200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AddPermissionsToUserGroup200ApplicationJSONObject
}

func (o *AddPermissionsToUserGroupResponse) GetAddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSONObject() *AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.AddPermissionsToUserGroup200ApplicationVndSegmentV1PlusJSONObject
}

func (o *AddPermissionsToUserGroupResponse) GetAddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject() *AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.AddPermissionsToUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *AddPermissionsToUserGroupResponse) GetAddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSONObject() *AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.AddPermissionsToUserGroup200ApplicationVndSegmentV1betaPlusJSONObject
}
