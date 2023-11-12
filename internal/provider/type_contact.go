// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Contact struct {
	Email     types.String `tfsdk:"email"`
	IsPrimary types.Bool   `tfsdk:"is_primary"`
	Name      types.String `tfsdk:"name"`
	Role      types.String `tfsdk:"role"`
}