// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListUserGroupsV1OutputPagination - Information about the pagination of this response.
type ListUserGroupsV1OutputPagination struct {
	// The current cursor within a collection.
	//
	// Consumers of the API must treat this value as opaque.
	Current string `json:"current"`
	// A pointer to the next page.
	//
	// This does not return when you retrieve the last page.
	//
	// Consumers of the API must treat this value as opaque.
	Next *string `json:"next,omitempty"`
	// A pointer to the previous page.
	//
	// This does not return when you retrieve the first page.
	//
	// Consumers of the API must treat this value as opaque.
	Previous *string `json:"previous,omitempty"`
	// The total number of entries available in the collection.
	//
	// If calculating it impacts performance, the response may omit this field.
	TotalEntries *float64 `json:"totalEntries,omitempty"`
}

func (o *ListUserGroupsV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListUserGroupsV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListUserGroupsV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListUserGroupsV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListUserGroupsV1Output - Returns a list of user groups with the given Workspace id.
type ListUserGroupsV1Output struct {
	// Information about the pagination of this response.
	Pagination ListUserGroupsV1OutputPagination `json:"pagination"`
	// The user group returned from the query.
	UserGroups []UserGroupV1 `json:"userGroups"`
}

func (o *ListUserGroupsV1Output) GetPagination() ListUserGroupsV1OutputPagination {
	if o == nil {
		return ListUserGroupsV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListUserGroupsV1Output) GetUserGroups() []UserGroupV1 {
	if o == nil {
		return []UserGroupV1{}
	}
	return o.UserGroups
}