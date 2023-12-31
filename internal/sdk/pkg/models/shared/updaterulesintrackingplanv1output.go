// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateRulesInTrackingPlanV1OutputStatus - The operation status.
type UpdateRulesInTrackingPlanV1OutputStatus string

const (
	UpdateRulesInTrackingPlanV1OutputStatusSuccess UpdateRulesInTrackingPlanV1OutputStatus = "SUCCESS"
)

func (e UpdateRulesInTrackingPlanV1OutputStatus) ToPointer() *UpdateRulesInTrackingPlanV1OutputStatus {
	return &e
}

func (e *UpdateRulesInTrackingPlanV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = UpdateRulesInTrackingPlanV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateRulesInTrackingPlanV1OutputStatus: %v", v)
	}
}

// UpdateRulesInTrackingPlanV1Output - Updates Tracking Plan rules. Non-existent rules are added.
type UpdateRulesInTrackingPlanV1Output struct {
	// The operation status.
	Status UpdateRulesInTrackingPlanV1OutputStatus `json:"status"`
}

func (o *UpdateRulesInTrackingPlanV1Output) GetStatus() UpdateRulesInTrackingPlanV1OutputStatus {
	if o == nil {
		return UpdateRulesInTrackingPlanV1OutputStatus("")
	}
	return o.Status
}
