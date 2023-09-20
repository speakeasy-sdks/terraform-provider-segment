// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListTransformationsBetaInputPaginationInput - Pagination parameters.
//
// Every resource that returns a list of items in its `Output` object may contain a `PaginationInput` in its `Input`
// object. Required, though some of its fields are optional.
type ListTransformationsBetaInputPaginationInput struct {
	// The number of items to retrieve in a page, between 1 and 200.
	Count float64 `json:"count"`
	// The page to request.
	//
	// Acceptable values to use here are in PaginationOutput objects, in the `current`, `next`, and `previous` keys.
	//
	// Consumers of the API must treat this value as opaque.
	Cursor *string `json:"cursor,omitempty"`
}

// ListTransformationsBetaInput - Lists the Transformations associated with the current Workspace.
type ListTransformationsBetaInput struct {
	// Pagination options.
	Pagination ListTransformationsBetaInputPaginationInput `json:"pagination"`
}