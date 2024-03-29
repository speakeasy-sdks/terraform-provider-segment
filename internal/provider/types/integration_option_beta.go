// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type IntegrationOptionBeta struct {
	DefaultValue types.String `tfsdk:"default_value"`
	Description  types.String `tfsdk:"description"`
	Label        types.String `tfsdk:"label"`
	Name         types.String `tfsdk:"name"`
	Required     types.Bool   `tfsdk:"required"`
	Type         types.String `tfsdk:"type"`
}
