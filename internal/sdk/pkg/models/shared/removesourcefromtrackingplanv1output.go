// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RemoveSourceFromTrackingPlanV1OutputStatus - The operation status.
type RemoveSourceFromTrackingPlanV1OutputStatus string

const (
	RemoveSourceFromTrackingPlanV1OutputStatusSuccess RemoveSourceFromTrackingPlanV1OutputStatus = "SUCCESS"
)

func (e RemoveSourceFromTrackingPlanV1OutputStatus) ToPointer() *RemoveSourceFromTrackingPlanV1OutputStatus {
	return &e
}

func (e *RemoveSourceFromTrackingPlanV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = RemoveSourceFromTrackingPlanV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RemoveSourceFromTrackingPlanV1OutputStatus: %v", v)
	}
}

// RemoveSourceFromTrackingPlanV1Output - Disconnects a Source from a Tracking Plan.
type RemoveSourceFromTrackingPlanV1Output struct {
	// The operation status.
	Status RemoveSourceFromTrackingPlanV1OutputStatus `json:"status"`
}
