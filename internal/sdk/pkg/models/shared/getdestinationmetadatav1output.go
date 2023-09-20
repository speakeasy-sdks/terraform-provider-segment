// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetDestinationMetadataV1OutputDestinationMetadataV1LogosBeta - Represents a logo.
type GetDestinationMetadataV1OutputDestinationMetadataV1LogosBeta struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

// GetDestinationMetadataV1OutputDestinationMetadataV1Status - Support status of the Destination.
type GetDestinationMetadataV1OutputDestinationMetadataV1Status string

const (
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusDeprecated       GetDestinationMetadataV1OutputDestinationMetadataV1Status = "DEPRECATED"
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusPrivateBeta      GetDestinationMetadataV1OutputDestinationMetadataV1Status = "PRIVATE_BETA"
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusPrivateBuilding  GetDestinationMetadataV1OutputDestinationMetadataV1Status = "PRIVATE_BUILDING"
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusPrivateSubmitted GetDestinationMetadataV1OutputDestinationMetadataV1Status = "PRIVATE_SUBMITTED"
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusPublic           GetDestinationMetadataV1OutputDestinationMetadataV1Status = "PUBLIC"
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusPublicBeta       GetDestinationMetadataV1OutputDestinationMetadataV1Status = "PUBLIC_BETA"
	GetDestinationMetadataV1OutputDestinationMetadataV1StatusUnavailable      GetDestinationMetadataV1OutputDestinationMetadataV1Status = "UNAVAILABLE"
)

func (e GetDestinationMetadataV1OutputDestinationMetadataV1Status) ToPointer() *GetDestinationMetadataV1OutputDestinationMetadataV1Status {
	return &e
}

func (e *GetDestinationMetadataV1OutputDestinationMetadataV1Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DEPRECATED":
		fallthrough
	case "PRIVATE_BETA":
		fallthrough
	case "PRIVATE_BUILDING":
		fallthrough
	case "PRIVATE_SUBMITTED":
		fallthrough
	case "PUBLIC":
		fallthrough
	case "PUBLIC_BETA":
		fallthrough
	case "UNAVAILABLE":
		*e = GetDestinationMetadataV1OutputDestinationMetadataV1Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationMetadataV1OutputDestinationMetadataV1Status: %v", v)
	}
}

// GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances string

const (
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesZero     GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "0"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesOne      GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "1"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesMultiple GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "MULTIPLE"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesNone     GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "NONE"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesSingle   GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "SINGLE"
)

func (e GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) ToPointer() *GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances {
	return &e
}

func (e *GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances: %v", v)
	}
}

// GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances string

const (
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesZero   GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "0"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesOne    GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "1"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesNone   GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "NONE"
	GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesSingle GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "SINGLE"
)

func (e GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) ToPointer() *GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances {
	return &e
}

func (e *GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances: %v", v)
	}
}

// GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1 - Represents features that a given Destination supports.
type GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1 struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

// GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataMethodsV1 - Represents methods that a given Destination supports.
type GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataMethodsV1 struct {
	// Identifies if the Destination supports the `alias` method.
	Alias *bool `json:"alias,omitempty"`
	// Identifies if the Destination supports the `group` method.
	Group *bool `json:"group,omitempty"`
	// Identifies if the Destination supports the `identify` method.
	Identify *bool `json:"identify,omitempty"`
	// Identifies if the Destination supports the `pageview` method.
	Pageview *bool `json:"pageview,omitempty"`
	// Identifies if the Destination supports the `track` method.
	Track *bool `json:"track,omitempty"`
}

// GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataPlatformsV1 - Represents platforms that a given Destination supports.
type GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataPlatformsV1 struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
}

// GetDestinationMetadataV1OutputDestinationMetadataV1 - Represents a Destination within Segment.
//
// A Destination is a target for Segment to forward data to, and represents a tool or storage Destination.
type GetDestinationMetadataV1OutputDestinationMetadataV1 struct {
	// Actions available for the Destination.
	Actions []DestinationMetadataActionV1 `json:"actions"`
	// A list of categories with which the Destination is associated.
	Categories []string `json:"categories"`
	// A list of components this Destination provides.
	Components []DestinationMetadataComponentV1 `json:"components"`
	// Contact info for Integration Owners.
	Contacts []Contact `json:"contacts,omitempty"`
	// The description of the Destination.
	Description string `json:"description"`
	// The id of the Destination metadata.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The Destination's logos.
	Logos GetDestinationMetadataV1OutputDestinationMetadataV1LogosBeta `json:"logos"`
	// The user-friendly name of the Destination.
	//
	// Config API note: equal to `displayName`.
	Name string `json:"name"`
	// Options configured for the Destination.
	Options []IntegrationOptionBeta `json:"options"`
	// Partner Owned flag.
	PartnerOwned *bool `json:"partnerOwned,omitempty"`
	// Predefined Destination subscriptions that can optionally be applied when connecting a new instance of the Destination.
	Presets []DestinationMetadataSubscriptionPresetV1 `json:"presets"`
	// A list of names previously used by the Destination.
	PreviousNames []string `json:"previousNames"`
	// The list of regional endpoints for this Destination.
	RegionEndpoints []string `json:"regionEndpoints,omitempty"`
	// The slug used to identify the Destination in the Segment app.
	Slug string `json:"slug"`
	// Support status of the Destination.
	Status GetDestinationMetadataV1OutputDestinationMetadataV1Status `json:"status"`
	// Features that this Destination supports.
	//
	// Config API note: holds `browserUnbundling` fields.
	SupportedFeatures GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataFeaturesV1 `json:"supportedFeatures"`
	// Methods that this Destination supports.
	//
	// Config API note: equal to `methods`.
	SupportedMethods GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataMethodsV1 `json:"supportedMethods"`
	// Platforms from which the Destination receives events.
	//
	// Config API note: equal to `platforms`.
	SupportedPlatforms GetDestinationMetadataV1OutputDestinationMetadataV1DestinationMetadataPlatformsV1 `json:"supportedPlatforms"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// A website URL for this Destination.
	Website string `json:"website"`
}

// GetDestinationMetadataV1Output - Returns a Destination catalog item.
type GetDestinationMetadataV1Output struct {
	// The catalog item matched by id.
	DestinationMetadata GetDestinationMetadataV1OutputDestinationMetadataV1 `json:"destinationMetadata"`
}
