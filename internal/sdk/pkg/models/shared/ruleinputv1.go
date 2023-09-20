// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RuleInputV1Type - The type for this Tracking Plan rule.
type RuleInputV1Type string

const (
	RuleInputV1TypeCommon   RuleInputV1Type = "COMMON"
	RuleInputV1TypeGroup    RuleInputV1Type = "GROUP"
	RuleInputV1TypeIdentify RuleInputV1Type = "IDENTIFY"
	RuleInputV1TypePage     RuleInputV1Type = "PAGE"
	RuleInputV1TypeScreen   RuleInputV1Type = "SCREEN"
	RuleInputV1TypeTrack    RuleInputV1Type = "TRACK"
)

func (e RuleInputV1Type) ToPointer() *RuleInputV1Type {
	return &e
}

func (e *RuleInputV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMMON":
		fallthrough
	case "GROUP":
		fallthrough
	case "IDENTIFY":
		fallthrough
	case "PAGE":
		fallthrough
	case "SCREEN":
		fallthrough
	case "TRACK":
		*e = RuleInputV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RuleInputV1Type: %v", v)
	}
}

// RuleInputV1 - Represents a rule to add to a Tracking Plan.
type RuleInputV1 struct {
	// JSON Schema of this rule.
	JSONSchema interface{} `json:"jsonSchema"`
	// Key to this rule (free-form string like 'Button clicked').
	Key *string `json:"key,omitempty"`
	// The type for this Tracking Plan rule.
	Type RuleInputV1Type `json:"type"`
	// Version of this rule.
	Version float64 `json:"version"`
}
