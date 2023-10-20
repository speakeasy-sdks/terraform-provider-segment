// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationMetadataFeaturesV1CloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type DestinationMetadataFeaturesV1CloudModeInstances string

const (
	DestinationMetadataFeaturesV1CloudModeInstancesZero     DestinationMetadataFeaturesV1CloudModeInstances = "0"
	DestinationMetadataFeaturesV1CloudModeInstancesOne      DestinationMetadataFeaturesV1CloudModeInstances = "1"
	DestinationMetadataFeaturesV1CloudModeInstancesMultiple DestinationMetadataFeaturesV1CloudModeInstances = "MULTIPLE"
	DestinationMetadataFeaturesV1CloudModeInstancesNone     DestinationMetadataFeaturesV1CloudModeInstances = "NONE"
	DestinationMetadataFeaturesV1CloudModeInstancesSingle   DestinationMetadataFeaturesV1CloudModeInstances = "SINGLE"
)

func (e DestinationMetadataFeaturesV1CloudModeInstances) ToPointer() *DestinationMetadataFeaturesV1CloudModeInstances {
	return &e
}

func (e *DestinationMetadataFeaturesV1CloudModeInstances) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "MULTIPLE":
		fallthrough
	case "NONE":
		fallthrough
	case "SINGLE":
		*e = DestinationMetadataFeaturesV1CloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataFeaturesV1CloudModeInstances: %v", v)
	}
}

// DestinationMetadataFeaturesV1DeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type DestinationMetadataFeaturesV1DeviceModeInstances string

const (
	DestinationMetadataFeaturesV1DeviceModeInstancesZero   DestinationMetadataFeaturesV1DeviceModeInstances = "0"
	DestinationMetadataFeaturesV1DeviceModeInstancesOne    DestinationMetadataFeaturesV1DeviceModeInstances = "1"
	DestinationMetadataFeaturesV1DeviceModeInstancesNone   DestinationMetadataFeaturesV1DeviceModeInstances = "NONE"
	DestinationMetadataFeaturesV1DeviceModeInstancesSingle DestinationMetadataFeaturesV1DeviceModeInstances = "SINGLE"
)

func (e DestinationMetadataFeaturesV1DeviceModeInstances) ToPointer() *DestinationMetadataFeaturesV1DeviceModeInstances {
	return &e
}

func (e *DestinationMetadataFeaturesV1DeviceModeInstances) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "NONE":
		fallthrough
	case "SINGLE":
		*e = DestinationMetadataFeaturesV1DeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataFeaturesV1DeviceModeInstances: %v", v)
	}
}

// DestinationMetadataFeaturesV1 - Represents features that a given Destination supports.
type DestinationMetadataFeaturesV1 struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *DestinationMetadataFeaturesV1CloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *DestinationMetadataFeaturesV1DeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundling() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundling
}

func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublic() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundlingPublic
}

func (o *DestinationMetadataFeaturesV1) GetCloudModeInstances() *DestinationMetadataFeaturesV1CloudModeInstances {
	if o == nil {
		return nil
	}
	return o.CloudModeInstances
}

func (o *DestinationMetadataFeaturesV1) GetDeviceModeInstances() *DestinationMetadataFeaturesV1DeviceModeInstances {
	if o == nil {
		return nil
	}
	return o.DeviceModeInstances
}

func (o *DestinationMetadataFeaturesV1) GetReplay() *bool {
	if o == nil {
		return nil
	}
	return o.Replay
}
