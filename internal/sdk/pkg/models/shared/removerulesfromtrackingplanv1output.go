// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RemoveRulesFromTrackingPlanV1OutputStatus - The status of the operation.
type RemoveRulesFromTrackingPlanV1OutputStatus string

const (
	RemoveRulesFromTrackingPlanV1OutputStatusSuccess RemoveRulesFromTrackingPlanV1OutputStatus = "SUCCESS"
)

func (e RemoveRulesFromTrackingPlanV1OutputStatus) ToPointer() *RemoveRulesFromTrackingPlanV1OutputStatus {
	return &e
}

func (e *RemoveRulesFromTrackingPlanV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = RemoveRulesFromTrackingPlanV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RemoveRulesFromTrackingPlanV1OutputStatus: %v", v)
	}
}

// RemoveRulesFromTrackingPlanV1Output - Remove specified rules from a Tracking Plan.
type RemoveRulesFromTrackingPlanV1Output struct {
	// The status of the operation.
	Status RemoveRulesFromTrackingPlanV1OutputStatus `json:"status"`
}