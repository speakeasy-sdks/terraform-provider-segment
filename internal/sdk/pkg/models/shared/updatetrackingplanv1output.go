// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateTrackingPlanV1OutputStatus - The operation status.
type UpdateTrackingPlanV1OutputStatus string

const (
	UpdateTrackingPlanV1OutputStatusSuccess UpdateTrackingPlanV1OutputStatus = "SUCCESS"
)

func (e UpdateTrackingPlanV1OutputStatus) ToPointer() *UpdateTrackingPlanV1OutputStatus {
	return &e
}

func (e *UpdateTrackingPlanV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = UpdateTrackingPlanV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateTrackingPlanV1OutputStatus: %v", v)
	}
}

// UpdateTrackingPlanV1Output - Result of an UpdateTrackingPlan call.
type UpdateTrackingPlanV1Output struct {
	// The operation status.
	Status UpdateTrackingPlanV1OutputStatus `json:"status"`
}