// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SourceV1SourceMetadataV1LogosBeta - Represents a logo.
type SourceV1SourceMetadataV1LogosBeta struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *SourceV1SourceMetadataV1LogosBeta) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *SourceV1SourceMetadataV1LogosBeta) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *SourceV1SourceMetadataV1LogosBeta) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// SourceV1SourceMetadataV1 - A website, server library, mobile SDK, or cloud application which can send data into Segment.
type SourceV1SourceMetadataV1 struct {
	// A list of categories this Source belongs to.
	Categories []string `json:"categories"`
	// The description of this Source.
	Description string `json:"description"`
	// The id for this Source metadata in the Segment catalog.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// True if this is a Cloud Event Source.
	IsCloudEventSource bool `json:"isCloudEventSource"`
	// The logos for this Source.
	Logos SourceV1SourceMetadataV1LogosBeta `json:"logos"`
	// The user-friendly name of this Source.
	//
	// Config API note: equal to `displayName`.
	Name string `json:"name"`
	// Options for this Source.
	Options []IntegrationOptionBeta `json:"options"`
	// The slug that identifies this Source in the Segment app.
	//
	// Config API note: equal to `name`.
	Slug string `json:"slug"`
}

func (o *SourceV1SourceMetadataV1) GetCategories() []string {
	if o == nil {
		return []string{}
	}
	return o.Categories
}

func (o *SourceV1SourceMetadataV1) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *SourceV1SourceMetadataV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SourceV1SourceMetadataV1) GetIsCloudEventSource() bool {
	if o == nil {
		return false
	}
	return o.IsCloudEventSource
}

func (o *SourceV1SourceMetadataV1) GetLogos() SourceV1SourceMetadataV1LogosBeta {
	if o == nil {
		return SourceV1SourceMetadataV1LogosBeta{}
	}
	return o.Logos
}

func (o *SourceV1SourceMetadataV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceV1SourceMetadataV1) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *SourceV1SourceMetadataV1) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// SourceV1 - Defines a data Source for Segment data.
type SourceV1 struct {
	// Enable to receive data from the Source.
	Enabled bool `json:"enabled"`
	// The id of the Source.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// A list of labels applied to the Source.
	Labels []LabelV1 `json:"labels"`
	// The metadata for the Source.
	//
	// Config API note: includes `catalogName` and `catalogId`.
	Metadata SourceV1SourceMetadataV1 `json:"metadata"`
	// The name of the Source.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// The settings associated with the Source.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// The slug used to identify the Source in the Segment app.
	//
	// Config API note: equal to `name`.
	Slug string `json:"slug"`
	// The id of the Workspace that owns the Source.
	//
	// Config API note: equal to `parent`.
	WorkspaceID string `json:"workspaceId"`
	// The write keys used to send data from the Source. This field is left empty when the current token does not have the
	// 'source admin' permission.
	WriteKeys []string `json:"writeKeys"`
}

func (o *SourceV1) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *SourceV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SourceV1) GetLabels() []LabelV1 {
	if o == nil {
		return []LabelV1{}
	}
	return o.Labels
}

func (o *SourceV1) GetMetadata() SourceV1SourceMetadataV1 {
	if o == nil {
		return SourceV1SourceMetadataV1{}
	}
	return o.Metadata
}

func (o *SourceV1) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SourceV1) GetSettings() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *SourceV1) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *SourceV1) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceV1) GetWriteKeys() []string {
	if o == nil {
		return []string{}
	}
	return o.WriteKeys
}
