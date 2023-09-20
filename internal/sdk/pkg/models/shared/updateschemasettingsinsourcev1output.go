// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations - The common group event on violations.
//
// Config API note: equal to `commonGroupEventOnViolations`.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations string

const (
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolationsAllow      UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations = "ALLOW"
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolationsBlock      UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations = "BLOCK"
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolationsOmitTraits UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations = "OMIT_TRAITS"
)

func (e UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations) ToPointer() *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations {
	return &e
}

func (e *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations) UnmarshalJSON(data []byte) error {
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
		*e = UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations: %v", v)
	}
}

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1 - Group settings.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1 struct {
	// Enable to allow group traits on violations.
	//
	// Config API note: equal to `allowGroupTraitsOnViolations`.
	AllowTraitsOnViolations *bool `json:"allowTraitsOnViolations,omitempty"`
	// Enable to allow unplanned group traits.
	//
	// Config API note: equal to `allowUnplannedGroupTraits`.
	AllowUnplannedTraits *bool `json:"allowUnplannedTraits,omitempty"`
	// The common group event on violations.
	//
	// Config API note: equal to `commonGroupEventOnViolations`.
	CommonEventOnViolations *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1CommonEventOnViolations `json:"commonEventOnViolations,omitempty"`
}

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations - The common identify event on violations.
//
// Config API note: equal to `commonIdentifyEventOnViolations`.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations string

const (
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolationsAllow      UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations = "ALLOW"
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolationsBlock      UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations = "BLOCK"
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolationsOmitTraits UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations = "OMIT_TRAITS"
)

func (e UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations) ToPointer() *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations {
	return &e
}

func (e *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations) UnmarshalJSON(data []byte) error {
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
		*e = UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations: %v", v)
	}
}

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1 - Identify settings.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1 struct {
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
	CommonEventOnViolations *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1CommonEventOnViolations `json:"commonEventOnViolations,omitempty"`
}

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations - The common track event on violations.
//
// Config API note: equal to `commonTrackEventOnViolations`.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations string

const (
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolationsAllow          UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations = "ALLOW"
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolationsBlock          UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations = "BLOCK"
	UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolationsOmitProperties UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations = "OMIT_PROPERTIES"
)

func (e UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations) ToPointer() *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations {
	return &e
}

func (e *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations) UnmarshalJSON(data []byte) error {
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
		*e = UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations: %v", v)
	}
}

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1 - Track settings.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1 struct {
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
	CommonEventOnViolations *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1CommonEventOnViolations `json:"commonEventOnViolations,omitempty"`
}

// UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1 - The output of Source settings.
type UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1 struct {
	// SourceId to forward blocked events to.
	ForwardingBlockedEventsTo *string `json:"forwardingBlockedEventsTo,omitempty"`
	// SourceId to forward violations to.
	ForwardingViolationsTo *string `json:"forwardingViolationsTo,omitempty"`
	// Group settings.
	Group *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1GroupSourceSettingsV1 `json:"group,omitempty"`
	// Identify settings.
	Identify *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1IdentifySourceSettingsV1 `json:"identify,omitempty"`
	// Track settings.
	Track *UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1TrackSourceSettingsV1 `json:"track,omitempty"`
}

// UpdateSchemaSettingsInSourceV1Output - Output of the Source with updated settings.
type UpdateSchemaSettingsInSourceV1Output struct {
	// The output of Source settings.
	Settings UpdateSchemaSettingsInSourceV1OutputSourceSettingsOutputV1 `json:"settings"`
	// The id of the updated Source.
	//
	// Config API note: analogous to `parent` and `name`.
	SourceID string `json:"sourceId"`
}