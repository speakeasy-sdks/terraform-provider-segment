// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMetadataSubscriptionPresetV1 struct {
	ActionID types.String            `tfsdk:"action_id"`
	Fields   map[string]types.String `tfsdk:"fields"`
	Name     types.String            `tfsdk:"name"`
	Trigger  types.String            `tfsdk:"trigger"`
}
