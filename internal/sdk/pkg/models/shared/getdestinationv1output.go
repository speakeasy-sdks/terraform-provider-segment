// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta - Represents a logo.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1Status - Support status of the Destination.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1Status string

const (
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusDeprecated       GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "DEPRECATED"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusPrivateBeta      GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PRIVATE_BETA"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusPrivateBuilding  GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PRIVATE_BUILDING"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusPrivateSubmitted GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PRIVATE_SUBMITTED"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusPublic           GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PUBLIC"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusPublicBeta       GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PUBLIC_BETA"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1StatusUnavailable      GetDestinationV1OutputDestinationV1DestinationMetadataV1Status = "UNAVAILABLE"
)

func (e GetDestinationV1OutputDestinationV1DestinationMetadataV1Status) ToPointer() *GetDestinationV1OutputDestinationV1DestinationMetadataV1Status {
	return &e
}

func (e *GetDestinationV1OutputDestinationV1DestinationMetadataV1Status) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationV1OutputDestinationV1DestinationMetadataV1Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationV1OutputDestinationV1DestinationMetadataV1Status: %v", v)
	}
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances string

const (
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesZero     GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "0"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesOne      GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "1"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesMultiple GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "MULTIPLE"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesNone     GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "NONE"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesSingle   GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "SINGLE"
)

func (e GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) ToPointer() *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances {
	return &e
}

func (e *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances: %v", v)
	}
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances string

const (
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesZero   GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "0"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesOne    GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "1"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesNone   GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "NONE"
	GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesSingle GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "SINGLE"
)

func (e GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) ToPointer() *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances {
	return &e
}

func (e *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances: %v", v)
	}
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 - Represents features that a given Destination supports.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1) GetBrowserUnbundling() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundling
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublic() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundlingPublic
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1) GetCloudModeInstances() *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances {
	if o == nil {
		return nil
	}
	return o.CloudModeInstances
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1) GetDeviceModeInstances() *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances {
	if o == nil {
		return nil
	}
	return o.DeviceModeInstances
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1) GetReplay() *bool {
	if o == nil {
		return nil
	}
	return o.Replay
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 - Represents methods that a given Destination supports.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 struct {
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

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1) GetAlias() *bool {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1) GetGroup() *bool {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1) GetIdentify() *bool {
	if o == nil {
		return nil
	}
	return o.Identify
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1) GetPageview() *bool {
	if o == nil {
		return nil
	}
	return o.Pageview
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1) GetTrack() *bool {
	if o == nil {
		return nil
	}
	return o.Track
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 - Represents platforms that a given Destination supports.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1) GetBrowser() *bool {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1) GetMobile() *bool {
	if o == nil {
		return nil
	}
	return o.Mobile
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1) GetServer() *bool {
	if o == nil {
		return nil
	}
	return o.Server
}

// GetDestinationV1OutputDestinationV1DestinationMetadataV1 - Represents a Destination within Segment.
//
// A Destination is a target for Segment to forward data to, and represents a tool or storage Destination.
type GetDestinationV1OutputDestinationV1DestinationMetadataV1 struct {
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
	Logos GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta `json:"logos"`
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
	Status GetDestinationV1OutputDestinationV1DestinationMetadataV1Status `json:"status"`
	// Features that this Destination supports.
	//
	// Config API note: holds `browserUnbundling` fields.
	SupportedFeatures GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 `json:"supportedFeatures"`
	// Methods that this Destination supports.
	//
	// Config API note: equal to `methods`.
	SupportedMethods GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 `json:"supportedMethods"`
	// Platforms from which the Destination receives events.
	//
	// Config API note: equal to `platforms`.
	SupportedPlatforms GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 `json:"supportedPlatforms"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// A website URL for this Destination.
	Website string `json:"website"`
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetActions() []DestinationMetadataActionV1 {
	if o == nil {
		return []DestinationMetadataActionV1{}
	}
	return o.Actions
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetCategories() []string {
	if o == nil {
		return []string{}
	}
	return o.Categories
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetComponents() []DestinationMetadataComponentV1 {
	if o == nil {
		return []DestinationMetadataComponentV1{}
	}
	return o.Components
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetContacts() []Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetLogos() GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta {
	if o == nil {
		return GetDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta{}
	}
	return o.Logos
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetPartnerOwned() *bool {
	if o == nil {
		return nil
	}
	return o.PartnerOwned
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetPresets() []DestinationMetadataSubscriptionPresetV1 {
	if o == nil {
		return []DestinationMetadataSubscriptionPresetV1{}
	}
	return o.Presets
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetPreviousNames() []string {
	if o == nil {
		return []string{}
	}
	return o.PreviousNames
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetRegionEndpoints() []string {
	if o == nil {
		return nil
	}
	return o.RegionEndpoints
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetStatus() GetDestinationV1OutputDestinationV1DestinationMetadataV1Status {
	if o == nil {
		return GetDestinationV1OutputDestinationV1DestinationMetadataV1Status("")
	}
	return o.Status
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetSupportedFeatures() GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 {
	if o == nil {
		return GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1{}
	}
	return o.SupportedFeatures
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetSupportedMethods() GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 {
	if o == nil {
		return GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1{}
	}
	return o.SupportedMethods
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetSupportedPlatforms() GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 {
	if o == nil {
		return GetDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1{}
	}
	return o.SupportedPlatforms
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetSupportedRegions() []string {
	if o == nil {
		return nil
	}
	return o.SupportedRegions
}

func (o *GetDestinationV1OutputDestinationV1DestinationMetadataV1) GetWebsite() string {
	if o == nil {
		return ""
	}
	return o.Website
}

// GetDestinationV1OutputDestinationV1 - Business tools or apps that you can connect to the data flowing through Segment.
//
// This is equal to the Destination object in Config API, with the following fields omitted:
// - catalogId
// - createTime
// - updateTime
// - connectionMode.
type GetDestinationV1OutputDestinationV1 struct {
	// Whether this instance of a Destination receives data.
	Enabled bool `json:"enabled"`
	// The unique identifier of this instance of a Destination.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The metadata of the Destination of which this Destination is an instance of. For example, Google Analytics or Amplitude.
	Metadata GetDestinationV1OutputDestinationV1DestinationMetadataV1 `json:"metadata"`
	// The name of this instance of a Destination.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// The collection of settings associated with a Destination.
	//
	// Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings"`
	// The id of a Source connected to this instance of a Destination.
	//
	// Config API note: analogous to `parent`.
	SourceID string `json:"sourceId"`
}

func (o *GetDestinationV1OutputDestinationV1) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *GetDestinationV1OutputDestinationV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetDestinationV1OutputDestinationV1) GetMetadata() GetDestinationV1OutputDestinationV1DestinationMetadataV1 {
	if o == nil {
		return GetDestinationV1OutputDestinationV1DestinationMetadataV1{}
	}
	return o.Metadata
}

func (o *GetDestinationV1OutputDestinationV1) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetDestinationV1OutputDestinationV1) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}

func (o *GetDestinationV1OutputDestinationV1) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// GetDestinationV1Output - Returns a single Destination by its id.
type GetDestinationV1Output struct {
	// The Destination looked up.
	Destination GetDestinationV1OutputDestinationV1 `json:"destination"`
}

func (o *GetDestinationV1Output) GetDestination() GetDestinationV1OutputDestinationV1 {
	if o == nil {
		return GetDestinationV1OutputDestinationV1{}
	}
	return o.Destination
}
