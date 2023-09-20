// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetEventsVolumeFromWorkspaceV1OutputPaginationOutput - Pagination metadata for a list response.
//
// Responses return this object alongside a list of resources, which provides the necessary metadata for manipulating a
// paginated collection. In operations that return lists, it's always present, though some of its fields might not be.
type GetEventsVolumeFromWorkspaceV1OutputPaginationOutput struct {
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

// GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity - Granularity corresponds to the requested bucket granularity.
type GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity string

const (
	GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularityDay    GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity = "DAY"
	GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularityHour   GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity = "HOUR"
	GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularityMinute GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity = "MINUTE"
)

func (e GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity) ToPointer() *GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity {
	return &e
}

func (e *GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAY":
		fallthrough
	case "HOUR":
		fallthrough
	case "MINUTE":
		*e = GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity: %v", v)
	}
}

// GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1Query - GetEventVolumeOutputQuery represents the input query sent to output.
type GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1Query struct {
	// AppVersion is a list of strings which allow you to restrict the result to just
	// the given application versions.
	AppVersion []string `json:"appVersion,omitempty"`
	// EndTime is the ISO8601 formatted timestamp corresponding to the
	// end of the requested time frame, noninclusive.
	EndTime string `json:"endTime"`
	// EventName is a list of strings which allow you to restrict the result to just
	// the given event names.
	EventName []string `json:"eventName,omitempty"`
	// EventType is a list of strings which allow you to restrict the result to just
	// the given event types.
	EventType []string `json:"eventType,omitempty"`
	// Granularity corresponds to the requested bucket granularity.
	Granularity GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1QueryGranularity `json:"granularity"`
	// GroupBy is a comma-delimited list of strings representing the dimensions
	// to group the result by. The current options are:
	// `eventName` or `eventType`.
	GroupBy []string `json:"groupBy,omitempty"`
	// Limit is the total number of items in the result.
	Limit *float64 `json:"limit,omitempty"`
	// List of strings which allow you to restrict the result to just
	// the given Sources.
	SourceID []string `json:"sourceId,omitempty"`
	// StartTime is the ISO8601 formatted timestamp corresponding to the
	// beginning of the requested time frame, inclusive.
	StartTime string `json:"startTime"`
	// Workspace being requested.
	WorkspaceID string `json:"workspaceId"`
}

// GetEventsVolumeFromWorkspaceV1Output - GetEventsVolumeFromWorkspaceV1Output represents the results given the input query.
type GetEventsVolumeFromWorkspaceV1Output struct {
	// Information about the pagination of this response.
	Pagination *GetEventsVolumeFromWorkspaceV1OutputPaginationOutput `json:"pagination,omitempty"`
	// Observability event volume path.
	Path string `json:"path"`
	// Input query returned.
	Query GetEventsVolumeFromWorkspaceV1OutputGetEventsVolumeFromWorkspaceV1Query `json:"query"`
	// The resultant list of series broken down by the
	// dimensions requested over the time frame requested and ordered by the total count of events in all series.
	// Note: The limit of entries returned is 5000.
	Result []SourceEventVolumeV1 `json:"result"`
}
