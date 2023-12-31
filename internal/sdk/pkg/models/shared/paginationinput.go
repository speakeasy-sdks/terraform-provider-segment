// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PaginationInput - Pagination parameters.
//
// Every resource that returns a list of items in its `Output` object may contain a `PaginationInput` in its `Input`
// object. Required, though some of its fields are optional.
type PaginationInput struct {
	// The number of items to retrieve in a page, between 1 and 200.
	Count float64 `queryParam:"name=count"`
	// The page to request.
	//
	// Acceptable values to use here are in PaginationOutput objects, in the `current`, `next`, and `previous` keys.
	//
	// Consumers of the API must treat this value as opaque.
	Cursor *string `queryParam:"name=cursor"`
}

func (o *PaginationInput) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *PaginationInput) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}
