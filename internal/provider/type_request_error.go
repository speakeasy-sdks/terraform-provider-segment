// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type RequestError struct {
	Data    types.String `tfsdk:"data"`
	Field   types.String `tfsdk:"field"`
	Message types.String `tfsdk:"message"`
	Status  types.Number `tfsdk:"status"`
	Type    types.String `tfsdk:"type"`
}
