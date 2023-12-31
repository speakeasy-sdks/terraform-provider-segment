// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpsertRuleV1Type - The type for this Tracking Plan rule.
type UpsertRuleV1Type string

const (
	UpsertRuleV1TypeCommon   UpsertRuleV1Type = "COMMON"
	UpsertRuleV1TypeGroup    UpsertRuleV1Type = "GROUP"
	UpsertRuleV1TypeIdentify UpsertRuleV1Type = "IDENTIFY"
	UpsertRuleV1TypePage     UpsertRuleV1Type = "PAGE"
	UpsertRuleV1TypeScreen   UpsertRuleV1Type = "SCREEN"
	UpsertRuleV1TypeTrack    UpsertRuleV1Type = "TRACK"
)

func (e UpsertRuleV1Type) ToPointer() *UpsertRuleV1Type {
	return &e
}

func (e *UpsertRuleV1Type) UnmarshalJSON(data []byte) error {
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
		*e = UpsertRuleV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpsertRuleV1Type: %v", v)
	}
}

type UpsertRuleV1 struct {
	// JSON Schema of this rule.
	JSONSchema interface{} `json:"jsonSchema"`
	// Key to this rule (free-form string like 'Button clicked').
	Key *string `json:"key,omitempty"`
	// This rule's new intended key.
	NewKey *string `json:"newKey,omitempty"`
	// The type for this Tracking Plan rule.
	Type UpsertRuleV1Type `json:"type"`
	// Version of this rule.
	Version float64 `json:"version"`
}

func (o *UpsertRuleV1) GetJSONSchema() interface{} {
	if o == nil {
		return nil
	}
	return o.JSONSchema
}

func (o *UpsertRuleV1) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *UpsertRuleV1) GetNewKey() *string {
	if o == nil {
		return nil
	}
	return o.NewKey
}

func (o *UpsertRuleV1) GetType() UpsertRuleV1Type {
	if o == nil {
		return UpsertRuleV1Type("")
	}
	return o.Type
}

func (o *UpsertRuleV1) GetVersion() float64 {
	if o == nil {
		return 0.0
	}
	return o.Version
}
