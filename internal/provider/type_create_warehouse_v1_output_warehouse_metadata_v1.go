// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateWarehouseV1OutputWarehouseMetadataV1 struct {
	Description types.String                       `tfsdk:"description"`
	ID          types.String                       `tfsdk:"id"`
	Logos       CreateDestinationV1OutputLogosBeta `tfsdk:"logos"`
	Name        types.String                       `tfsdk:"name"`
	Options     []IntegrationOptionBeta            `tfsdk:"options"`
	Slug        types.String                       `tfsdk:"slug"`
}
