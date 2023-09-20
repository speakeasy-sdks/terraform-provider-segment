// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetEventsVolumeFromWorkspaceV1QueryGranularity - Granularity corresponds to the requested bucket granularity.
type GetEventsVolumeFromWorkspaceV1QueryGranularity string

const (
	GetEventsVolumeFromWorkspaceV1QueryGranularityDay    GetEventsVolumeFromWorkspaceV1QueryGranularity = "DAY"
	GetEventsVolumeFromWorkspaceV1QueryGranularityHour   GetEventsVolumeFromWorkspaceV1QueryGranularity = "HOUR"
	GetEventsVolumeFromWorkspaceV1QueryGranularityMinute GetEventsVolumeFromWorkspaceV1QueryGranularity = "MINUTE"
)

func (e GetEventsVolumeFromWorkspaceV1QueryGranularity) ToPointer() *GetEventsVolumeFromWorkspaceV1QueryGranularity {
	return &e
}

func (e *GetEventsVolumeFromWorkspaceV1QueryGranularity) UnmarshalJSON(data []byte) error {
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
		*e = GetEventsVolumeFromWorkspaceV1QueryGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetEventsVolumeFromWorkspaceV1QueryGranularity: %v", v)
	}
}

// GetEventsVolumeFromWorkspaceV1Query - GetEventVolumeOutputQuery represents the input query sent to output.
type GetEventsVolumeFromWorkspaceV1Query struct {
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
	Granularity GetEventsVolumeFromWorkspaceV1QueryGranularity `json:"granularity"`
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