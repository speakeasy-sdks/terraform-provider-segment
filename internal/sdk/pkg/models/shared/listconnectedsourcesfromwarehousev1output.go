// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListConnectedSourcesFromWarehouseV1OutputPagination - Information about the pagination of this response.
type ListConnectedSourcesFromWarehouseV1OutputPagination struct {
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

func (o *ListConnectedSourcesFromWarehouseV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListConnectedSourcesFromWarehouseV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListConnectedSourcesFromWarehouseV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListConnectedSourcesFromWarehouseV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListConnectedSourcesFromWarehouseV1Output - Returns a list of Sources connected to a Warehouse.
type ListConnectedSourcesFromWarehouseV1Output struct {
	// Information about the pagination of this response.
	Pagination ListConnectedSourcesFromWarehouseV1OutputPagination `json:"pagination"`
	// A list that contains the Sources connected to the requested Warehouse.
	Sources []SourceV1 `json:"sources"`
}

func (o *ListConnectedSourcesFromWarehouseV1Output) GetPagination() ListConnectedSourcesFromWarehouseV1OutputPagination {
	if o == nil {
		return ListConnectedSourcesFromWarehouseV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListConnectedSourcesFromWarehouseV1Output) GetSources() []SourceV1 {
	if o == nil {
		return []SourceV1{}
	}
	return o.Sources
}
