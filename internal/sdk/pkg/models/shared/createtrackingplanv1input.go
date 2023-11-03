// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateTrackingPlanV1InputType - The Tracking Plan's type.
type CreateTrackingPlanV1InputType string

const (
	CreateTrackingPlanV1InputTypeEngage          CreateTrackingPlanV1InputType = "ENGAGE"
	CreateTrackingPlanV1InputTypeLive            CreateTrackingPlanV1InputType = "LIVE"
	CreateTrackingPlanV1InputTypePropertyLibrary CreateTrackingPlanV1InputType = "PROPERTY_LIBRARY"
	CreateTrackingPlanV1InputTypeRuleLibrary     CreateTrackingPlanV1InputType = "RULE_LIBRARY"
	CreateTrackingPlanV1InputTypeTemplate        CreateTrackingPlanV1InputType = "TEMPLATE"
)

func (e CreateTrackingPlanV1InputType) ToPointer() *CreateTrackingPlanV1InputType {
	return &e
}

func (e *CreateTrackingPlanV1InputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ENGAGE":
		fallthrough
	case "LIVE":
		fallthrough
	case "PROPERTY_LIBRARY":
		fallthrough
	case "RULE_LIBRARY":
		fallthrough
	case "TEMPLATE":
		*e = CreateTrackingPlanV1InputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTrackingPlanV1InputType: %v", v)
	}
}

// CreateTrackingPlanV1Input - Creates a Tracking Plan in the Workspace.
type CreateTrackingPlanV1Input struct {
	// The Tracking Plan's description.
	Description *string `json:"description,omitempty"`
	// The Tracking Plan's name.
	//
	// Config API note: equal to `displayName`.
	Name string `json:"name"`
	// The Tracking Plan's type.
	Type CreateTrackingPlanV1InputType `json:"type"`
}

func (o *CreateTrackingPlanV1Input) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateTrackingPlanV1Input) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTrackingPlanV1Input) GetType() CreateTrackingPlanV1InputType {
	if o == nil {
		return CreateTrackingPlanV1InputType("")
	}
	return o.Type
}
