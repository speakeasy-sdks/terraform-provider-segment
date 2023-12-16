// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListUsersFromUserGroupV1OutputPagination - Information about the pagination of this response.
type ListUsersFromUserGroupV1OutputPagination struct {
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

func (o *ListUsersFromUserGroupV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListUsersFromUserGroupV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListUsersFromUserGroupV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListUsersFromUserGroupV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListUsersFromUserGroupV1Output - Returns the members of a user group with the given group id.
type ListUsersFromUserGroupV1Output struct {
	// Information about the pagination of this response.
	Pagination ListUsersFromUserGroupV1OutputPagination `json:"pagination"`
	// The users of the user group.
	Users []MinimalUserV1 `json:"users"`
}

func (o *ListUsersFromUserGroupV1Output) GetPagination() ListUsersFromUserGroupV1OutputPagination {
	if o == nil {
		return ListUsersFromUserGroupV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListUsersFromUserGroupV1Output) GetUsers() []MinimalUserV1 {
	if o == nil {
		return []MinimalUserV1{}
	}
	return o.Users
}
