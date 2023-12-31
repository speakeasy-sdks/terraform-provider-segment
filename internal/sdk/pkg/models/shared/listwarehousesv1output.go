// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListWarehousesV1OutputPagination - Information about the pagination of this response.
type ListWarehousesV1OutputPagination struct {
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

func (o *ListWarehousesV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListWarehousesV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListWarehousesV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListWarehousesV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListWarehousesV1Output - Returns a list of Warehouses that belong to a Workspace.
type ListWarehousesV1Output struct {
	// Information about the pagination of this response.
	Pagination ListWarehousesV1OutputPagination `json:"pagination"`
	// A list of Warehouses that belong to the Workspace.
	Warehouses []WarehouseV1 `json:"warehouses"`
}

func (o *ListWarehousesV1Output) GetPagination() ListWarehousesV1OutputPagination {
	if o == nil {
		return ListWarehousesV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListWarehousesV1Output) GetWarehouses() []WarehouseV1 {
	if o == nil {
		return []WarehouseV1{}
	}
	return o.Warehouses
}
