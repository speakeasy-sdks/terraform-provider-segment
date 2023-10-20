// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TrackSourceSettingsV1CommonEventOnViolations - The common track event on violations.
//
// Config API note: equal to `commonTrackEventOnViolations`.
type TrackSourceSettingsV1CommonEventOnViolations string

const (
	TrackSourceSettingsV1CommonEventOnViolationsAllow          TrackSourceSettingsV1CommonEventOnViolations = "ALLOW"
	TrackSourceSettingsV1CommonEventOnViolationsBlock          TrackSourceSettingsV1CommonEventOnViolations = "BLOCK"
	TrackSourceSettingsV1CommonEventOnViolationsOmitProperties TrackSourceSettingsV1CommonEventOnViolations = "OMIT_PROPERTIES"
)

func (e TrackSourceSettingsV1CommonEventOnViolations) ToPointer() *TrackSourceSettingsV1CommonEventOnViolations {
	return &e
}

func (e *TrackSourceSettingsV1CommonEventOnViolations) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALLOW":
		fallthrough
	case "BLOCK":
		fallthrough
	case "OMIT_PROPERTIES":
		*e = TrackSourceSettingsV1CommonEventOnViolations(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TrackSourceSettingsV1CommonEventOnViolations: %v", v)
	}
}

type TrackSourceSettingsV1 struct {
	// Allow track event on violations.
	//
	// Config API note: equal to `allowTrackEventOnViolations`.
	AllowEventOnViolations *bool `json:"allowEventOnViolations,omitempty"`
	// Enable to allow track properties on violations.
	//
	// Config API note: equal to `allowTrackEventPropertiesOnViolations`.
	AllowPropertiesOnViolations *bool `json:"allowPropertiesOnViolations,omitempty"`
	// Enable to allow unplanned track event properties.
	//
	// Config API note: equal to `allowUnplannedTrackEventProperties`.
	AllowUnplannedEventProperties *bool `json:"allowUnplannedEventProperties,omitempty"`
	// Enable to allow unplanned track events.
	//
	// Config API note: equal to `allowUnplannedTrackEvents`.
	AllowUnplannedEvents *bool `json:"allowUnplannedEvents,omitempty"`
	// The common track event on violations.
	//
	// Config API note: equal to `commonTrackEventOnViolations`.
	CommonEventOnViolations *TrackSourceSettingsV1CommonEventOnViolations `json:"commonEventOnViolations,omitempty"`
}

func (o *TrackSourceSettingsV1) GetAllowEventOnViolations() *bool {
	if o == nil {
		return nil
	}
	return o.AllowEventOnViolations
}

func (o *TrackSourceSettingsV1) GetAllowPropertiesOnViolations() *bool {
	if o == nil {
		return nil
	}
	return o.AllowPropertiesOnViolations
}

func (o *TrackSourceSettingsV1) GetAllowUnplannedEventProperties() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUnplannedEventProperties
}

func (o *TrackSourceSettingsV1) GetAllowUnplannedEvents() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUnplannedEvents
}

func (o *TrackSourceSettingsV1) GetCommonEventOnViolations() *TrackSourceSettingsV1CommonEventOnViolations {
	if o == nil {
		return nil
	}
	return o.CommonEventOnViolations
}
