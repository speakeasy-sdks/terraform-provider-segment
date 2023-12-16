// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateDestinationV1OutputMetadata struct {
	Actions            []DestinationMetadataActionV1               `tfsdk:"actions"`
	Categories         []types.String                              `tfsdk:"categories"`
	Components         []DestinationMetadataComponentV1            `tfsdk:"components"`
	Contacts           []Contact                                   `tfsdk:"contacts"`
	Description        types.String                                `tfsdk:"description"`
	ID                 types.String                                `tfsdk:"id"`
	Logos              CreateDestinationV1OutputLogos              `tfsdk:"logos"`
	Name               types.String                                `tfsdk:"name"`
	Options            []IntegrationOptionBeta                     `tfsdk:"options"`
	PartnerOwned       types.Bool                                  `tfsdk:"partner_owned"`
	Presets            []DestinationMetadataSubscriptionPresetV1   `tfsdk:"presets"`
	PreviousNames      []types.String                              `tfsdk:"previous_names"`
	RegionEndpoints    []types.String                              `tfsdk:"region_endpoints"`
	Slug               types.String                                `tfsdk:"slug"`
	Status             types.String                                `tfsdk:"status"`
	SupportedFeatures  CreateDestinationV1OutputSupportedFeatures  `tfsdk:"supported_features"`
	SupportedMethods   CreateDestinationV1OutputSupportedMethods   `tfsdk:"supported_methods"`
	SupportedPlatforms CreateDestinationV1OutputSupportedPlatforms `tfsdk:"supported_platforms"`
	SupportedRegions   []types.String                              `tfsdk:"supported_regions"`
	Website            types.String                                `tfsdk:"website"`
}
