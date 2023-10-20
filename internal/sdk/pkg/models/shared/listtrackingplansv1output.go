// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListTrackingPlansV1OutputPaginationOutput - Pagination metadata for a list response.
//
// Responses return this object alongside a list of resources, which provides the necessary metadata for manipulating a
// paginated collection. In operations that return lists, it's always present, though some of its fields might not be.
type ListTrackingPlansV1OutputPaginationOutput struct {
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

func (o *ListTrackingPlansV1OutputPaginationOutput) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListTrackingPlansV1OutputPaginationOutput) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListTrackingPlansV1OutputPaginationOutput) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListTrackingPlansV1OutputPaginationOutput) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListTrackingPlansV1Output - Lists the Tracking Plans associated with the current Workspace.
type ListTrackingPlansV1Output struct {
	// Information about the pagination of this response.
	Pagination ListTrackingPlansV1OutputPaginationOutput `json:"pagination"`
	// A paginated list of Tracking Plans.
	TrackingPlans []TrackingPlanV1 `json:"trackingPlans"`
}

func (o *ListTrackingPlansV1Output) GetPagination() ListTrackingPlansV1OutputPaginationOutput {
	if o == nil {
		return ListTrackingPlansV1OutputPaginationOutput{}
	}
	return o.Pagination
}

func (o *ListTrackingPlansV1Output) GetTrackingPlans() []TrackingPlanV1 {
	if o == nil {
		return []TrackingPlanV1{}
	}
	return o.TrackingPlans
}
