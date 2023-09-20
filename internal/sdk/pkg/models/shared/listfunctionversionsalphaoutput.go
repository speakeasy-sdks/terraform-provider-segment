// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListFunctionVersionsAlphaOutputPaginationOutput - Pagination metadata for a list response.
//
// Responses return this object alongside a list of resources, which provides the necessary metadata for manipulating a
// paginated collection. In operations that return lists, it's always present, though some of its fields might not be.
type ListFunctionVersionsAlphaOutputPaginationOutput struct {
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

// ListFunctionVersionsAlphaOutput - Lists Versions of a Function.
type ListFunctionVersionsAlphaOutput struct {
	// Information about the pagination of this response.
	Pagination ListFunctionVersionsAlphaOutputPaginationOutput `json:"pagination"`
	// An array of Functions.
	Versions []Version `json:"versions"`
}
