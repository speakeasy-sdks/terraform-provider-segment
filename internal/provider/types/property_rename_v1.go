// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PropertyRenameV1 struct {
	NewName types.String `tfsdk:"new_name"`
	OldName types.String `tfsdk:"old_name"`
}