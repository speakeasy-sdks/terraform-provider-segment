// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateDestinationV1OutputDestinationMetadataMethodsV1 struct {
	Alias    types.Bool `tfsdk:"alias"`
	Group    types.Bool `tfsdk:"group"`
	Identify types.Bool `tfsdk:"identify"`
	Pageview types.Bool `tfsdk:"pageview"`
	Track    types.Bool `tfsdk:"track"`
}
