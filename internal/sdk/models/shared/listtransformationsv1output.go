// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListTransformationsV1OutputPagination - Information about the pagination of this response.
type ListTransformationsV1OutputPagination struct {
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

func (o *ListTransformationsV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListTransformationsV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListTransformationsV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListTransformationsV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListTransformationsV1Output - Lists the Transformations associated with the current Workspace.
type ListTransformationsV1Output struct {
	// Information about the pagination of this response.
	Pagination ListTransformationsV1OutputPagination `json:"pagination"`
	// A paginated list of Transformations.
	Transformations []TransformationV1 `json:"transformations"`
}

func (o *ListTransformationsV1Output) GetPagination() ListTransformationsV1OutputPagination {
	if o == nil {
		return ListTransformationsV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListTransformationsV1Output) GetTransformations() []TransformationV1 {
	if o == nil {
		return []TransformationV1{}
	}
	return o.Transformations
}