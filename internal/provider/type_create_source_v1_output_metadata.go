// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateSourceV1OutputMetadata struct {
	Categories         []types.String                 `tfsdk:"categories"`
	Description        types.String                   `tfsdk:"description"`
	ID                 types.String                   `tfsdk:"id"`
	IsCloudEventSource types.Bool                     `tfsdk:"is_cloud_event_source"`
	Logos              CreateDestinationV1OutputLogos `tfsdk:"logos"`
	Name               types.String                   `tfsdk:"name"`
	Options            []IntegrationOptionBeta        `tfsdk:"options"`
	Slug               types.String                   `tfsdk:"slug"`
}
