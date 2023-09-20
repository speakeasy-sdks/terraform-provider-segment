// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// IdentifySourceSettingsV1CommonEventOnViolations - The common identify event on violations.
//
// Config API note: equal to `commonIdentifyEventOnViolations`.
type IdentifySourceSettingsV1CommonEventOnViolations string

const (
	IdentifySourceSettingsV1CommonEventOnViolationsAllow      IdentifySourceSettingsV1CommonEventOnViolations = "ALLOW"
	IdentifySourceSettingsV1CommonEventOnViolationsBlock      IdentifySourceSettingsV1CommonEventOnViolations = "BLOCK"
	IdentifySourceSettingsV1CommonEventOnViolationsOmitTraits IdentifySourceSettingsV1CommonEventOnViolations = "OMIT_TRAITS"
)

func (e IdentifySourceSettingsV1CommonEventOnViolations) ToPointer() *IdentifySourceSettingsV1CommonEventOnViolations {
	return &e
}

func (e *IdentifySourceSettingsV1CommonEventOnViolations) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALLOW":
		fallthrough
	case "BLOCK":
		fallthrough
	case "OMIT_TRAITS":
		*e = IdentifySourceSettingsV1CommonEventOnViolations(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IdentifySourceSettingsV1CommonEventOnViolations: %v", v)
	}
}

type IdentifySourceSettingsV1 struct {
	// Enable to allow identify traits on violations.
	//
	// Config API note: equal to `allowIdentifyTraitsOnViolations`.
	AllowTraitsOnViolations *bool `json:"allowTraitsOnViolations,omitempty"`
	// Enable to allow unplanned identify traits.
	//
	// Config API note: equal to `allowUnplannedIdentifyTraits`.
	AllowUnplannedTraits *bool `json:"allowUnplannedTraits,omitempty"`
	// The common identify event on violations.
	//
	// Config API note: equal to `commonIdentifyEventOnViolations`.
	CommonEventOnViolations *IdentifySourceSettingsV1CommonEventOnViolations `json:"commonEventOnViolations,omitempty"`
}