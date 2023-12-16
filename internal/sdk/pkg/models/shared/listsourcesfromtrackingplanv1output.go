// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListSourcesFromTrackingPlanV1OutputPagination - Information about the pagination of this response.
type ListSourcesFromTrackingPlanV1OutputPagination struct {
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

func (o *ListSourcesFromTrackingPlanV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListSourcesFromTrackingPlanV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListSourcesFromTrackingPlanV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListSourcesFromTrackingPlanV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListSourcesFromTrackingPlanV1Output - Lists all Sources associated with a Tracking Plan.
type ListSourcesFromTrackingPlanV1Output struct {
	// Information about the pagination of this response.
	Pagination ListSourcesFromTrackingPlanV1OutputPagination `json:"pagination"`
	// A paginated list of Sources associated with the Tracking Plan.
	Sources []SourceV1 `json:"sources"`
}

func (o *ListSourcesFromTrackingPlanV1Output) GetPagination() ListSourcesFromTrackingPlanV1OutputPagination {
	if o == nil {
		return ListSourcesFromTrackingPlanV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListSourcesFromTrackingPlanV1Output) GetSources() []SourceV1 {
	if o == nil {
		return []SourceV1{}
	}
	return o.Sources
}
