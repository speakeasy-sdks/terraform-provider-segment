// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Pagination - Information about the pagination of this response.
type Pagination struct {
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

func (o *Pagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *Pagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *Pagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *Pagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// BatchQueryMessagingSubscriptionsForSpaceAlphaOutput - Batch get response.
type BatchQueryMessagingSubscriptionsForSpaceAlphaOutput struct {
	// General errors when making the request such as invalid payload or wrong http method errors.
	Errors []MessageSubscriptionResponseError `json:"errors"`
	// Validation errors due to invalid types or email/phone numbers.
	Failures []GetMessagingSubscriptionFailureResponse `json:"failures"`
	// Information about the pagination of this response.
	Pagination *Pagination `json:"pagination,omitempty"`
	// Array of successful subscription status.
	Successes []GetMessagingSubscriptionSuccessResponse `json:"successes"`
}

func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetErrors() []MessageSubscriptionResponseError {
	if o == nil {
		return []MessageSubscriptionResponseError{}
	}
	return o.Errors
}

func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetFailures() []GetMessagingSubscriptionFailureResponse {
	if o == nil {
		return []GetMessagingSubscriptionFailureResponse{}
	}
	return o.Failures
}

func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

func (o *BatchQueryMessagingSubscriptionsForSpaceAlphaOutput) GetSuccesses() []GetMessagingSubscriptionSuccessResponse {
	if o == nil {
		return []GetMessagingSubscriptionSuccessResponse{}
	}
	return o.Successes
}
