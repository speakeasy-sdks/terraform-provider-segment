// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationMetadataV1LogosBeta - Represents a logo.
type DestinationMetadataV1LogosBeta struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *DestinationMetadataV1LogosBeta) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *DestinationMetadataV1LogosBeta) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *DestinationMetadataV1LogosBeta) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// DestinationMetadataV1Status - Support status of the Destination.
type DestinationMetadataV1Status string

const (
	DestinationMetadataV1StatusDeprecated       DestinationMetadataV1Status = "DEPRECATED"
	DestinationMetadataV1StatusPrivateBeta      DestinationMetadataV1Status = "PRIVATE_BETA"
	DestinationMetadataV1StatusPrivateBuilding  DestinationMetadataV1Status = "PRIVATE_BUILDING"
	DestinationMetadataV1StatusPrivateSubmitted DestinationMetadataV1Status = "PRIVATE_SUBMITTED"
	DestinationMetadataV1StatusPublic           DestinationMetadataV1Status = "PUBLIC"
	DestinationMetadataV1StatusPublicBeta       DestinationMetadataV1Status = "PUBLIC_BETA"
	DestinationMetadataV1StatusUnavailable      DestinationMetadataV1Status = "UNAVAILABLE"
)

func (e DestinationMetadataV1Status) ToPointer() *DestinationMetadataV1Status {
	return &e
}

func (e *DestinationMetadataV1Status) UnmarshalJSON(data []byte) error {
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
		*e = DestinationMetadataV1Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataV1Status: %v", v)
	}
}

// DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances string

const (
	DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesZero     DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "0"
	DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesOne      DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "1"
	DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesMultiple DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "MULTIPLE"
	DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesNone     DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "NONE"
	DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesSingle   DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "SINGLE"
)

func (e DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) ToPointer() *DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances {
	return &e
}

func (e *DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances: %v", v)
	}
}

// DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances string

const (
	DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesZero   DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "0"
	DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesOne    DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "1"
	DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesNone   DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "NONE"
	DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesSingle DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "SINGLE"
)

func (e DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) ToPointer() *DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances {
	return &e
}

func (e *DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances: %v", v)
	}
}

// DestinationMetadataV1DestinationMetadataFeaturesV1 - Represents features that a given Destination supports.
type DestinationMetadataV1DestinationMetadataFeaturesV1 struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

func (o *DestinationMetadataV1DestinationMetadataFeaturesV1) GetBrowserUnbundling() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundling
}

func (o *DestinationMetadataV1DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublic() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundlingPublic
}

func (o *DestinationMetadataV1DestinationMetadataFeaturesV1) GetCloudModeInstances() *DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances {
	if o == nil {
		return nil
	}
	return o.CloudModeInstances
}

func (o *DestinationMetadataV1DestinationMetadataFeaturesV1) GetDeviceModeInstances() *DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances {
	if o == nil {
		return nil
	}
	return o.DeviceModeInstances
}

func (o *DestinationMetadataV1DestinationMetadataFeaturesV1) GetReplay() *bool {
	if o == nil {
		return nil
	}
	return o.Replay
}

// DestinationMetadataV1DestinationMetadataMethodsV1 - Represents methods that a given Destination supports.
type DestinationMetadataV1DestinationMetadataMethodsV1 struct {
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

func (o *DestinationMetadataV1DestinationMetadataMethodsV1) GetAlias() *bool {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *DestinationMetadataV1DestinationMetadataMethodsV1) GetGroup() *bool {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *DestinationMetadataV1DestinationMetadataMethodsV1) GetIdentify() *bool {
	if o == nil {
		return nil
	}
	return o.Identify
}

func (o *DestinationMetadataV1DestinationMetadataMethodsV1) GetPageview() *bool {
	if o == nil {
		return nil
	}
	return o.Pageview
}

func (o *DestinationMetadataV1DestinationMetadataMethodsV1) GetTrack() *bool {
	if o == nil {
		return nil
	}
	return o.Track
}

// DestinationMetadataV1DestinationMetadataPlatformsV1 - Represents platforms that a given Destination supports.
type DestinationMetadataV1DestinationMetadataPlatformsV1 struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
}

func (o *DestinationMetadataV1DestinationMetadataPlatformsV1) GetBrowser() *bool {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *DestinationMetadataV1DestinationMetadataPlatformsV1) GetMobile() *bool {
	if o == nil {
		return nil
	}
	return o.Mobile
}

func (o *DestinationMetadataV1DestinationMetadataPlatformsV1) GetServer() *bool {
	if o == nil {
		return nil
	}
	return o.Server
}

// DestinationMetadataV1 - Represents a Destination within Segment.
//
// A Destination is a target for Segment to forward data to, and represents a tool or storage Destination.
type DestinationMetadataV1 struct {
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
	Logos DestinationMetadataV1LogosBeta `json:"logos"`
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
	Status DestinationMetadataV1Status `json:"status"`
	// Features that this Destination supports.
	//
	// Config API note: holds `browserUnbundling` fields.
	SupportedFeatures DestinationMetadataV1DestinationMetadataFeaturesV1 `json:"supportedFeatures"`
	// Methods that this Destination supports.
	//
	// Config API note: equal to `methods`.
	SupportedMethods DestinationMetadataV1DestinationMetadataMethodsV1 `json:"supportedMethods"`
	// Platforms from which the Destination receives events.
	//
	// Config API note: equal to `platforms`.
	SupportedPlatforms DestinationMetadataV1DestinationMetadataPlatformsV1 `json:"supportedPlatforms"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// A website URL for this Destination.
	Website string `json:"website"`
}

func (o *DestinationMetadataV1) GetActions() []DestinationMetadataActionV1 {
	if o == nil {
		return []DestinationMetadataActionV1{}
	}
	return o.Actions
}

func (o *DestinationMetadataV1) GetCategories() []string {
	if o == nil {
		return []string{}
	}
	return o.Categories
}

func (o *DestinationMetadataV1) GetComponents() []DestinationMetadataComponentV1 {
	if o == nil {
		return []DestinationMetadataComponentV1{}
	}
	return o.Components
}

func (o *DestinationMetadataV1) GetContacts() []Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *DestinationMetadataV1) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *DestinationMetadataV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DestinationMetadataV1) GetLogos() DestinationMetadataV1LogosBeta {
	if o == nil {
		return DestinationMetadataV1LogosBeta{}
	}
	return o.Logos
}

func (o *DestinationMetadataV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationMetadataV1) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *DestinationMetadataV1) GetPartnerOwned() *bool {
	if o == nil {
		return nil
	}
	return o.PartnerOwned
}

func (o *DestinationMetadataV1) GetPresets() []DestinationMetadataSubscriptionPresetV1 {
	if o == nil {
		return []DestinationMetadataSubscriptionPresetV1{}
	}
	return o.Presets
}

func (o *DestinationMetadataV1) GetPreviousNames() []string {
	if o == nil {
		return []string{}
	}
	return o.PreviousNames
}

func (o *DestinationMetadataV1) GetRegionEndpoints() []string {
	if o == nil {
		return nil
	}
	return o.RegionEndpoints
}

func (o *DestinationMetadataV1) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *DestinationMetadataV1) GetStatus() DestinationMetadataV1Status {
	if o == nil {
		return DestinationMetadataV1Status("")
	}
	return o.Status
}

func (o *DestinationMetadataV1) GetSupportedFeatures() DestinationMetadataV1DestinationMetadataFeaturesV1 {
	if o == nil {
		return DestinationMetadataV1DestinationMetadataFeaturesV1{}
	}
	return o.SupportedFeatures
}

func (o *DestinationMetadataV1) GetSupportedMethods() DestinationMetadataV1DestinationMetadataMethodsV1 {
	if o == nil {
		return DestinationMetadataV1DestinationMetadataMethodsV1{}
	}
	return o.SupportedMethods
}

func (o *DestinationMetadataV1) GetSupportedPlatforms() DestinationMetadataV1DestinationMetadataPlatformsV1 {
	if o == nil {
		return DestinationMetadataV1DestinationMetadataPlatformsV1{}
	}
	return o.SupportedPlatforms
}

func (o *DestinationMetadataV1) GetSupportedRegions() []string {
	if o == nil {
		return nil
	}
	return o.SupportedRegions
}

func (o *DestinationMetadataV1) GetWebsite() string {
	if o == nil {
		return ""
	}
	return o.Website
}
