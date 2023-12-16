// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Function struct {
	BatchMaxCount     types.Number        `tfsdk:"batch_max_count"`
	CatalogID         types.String        `tfsdk:"catalog_id"`
	Code              types.String        `tfsdk:"code"`
	CreatedAt         types.String        `tfsdk:"created_at"`
	CreatedBy         types.String        `tfsdk:"created_by"`
	DeployedAt        types.String        `tfsdk:"deployed_at"`
	Description       types.String        `tfsdk:"description"`
	DisplayName       types.String        `tfsdk:"display_name"`
	ID                types.String        `tfsdk:"id"`
	IsLatestVersion   types.Bool          `tfsdk:"is_latest_version"`
	LogoURL           types.String        `tfsdk:"logo_url"`
	PreviewWebhookURL types.String        `tfsdk:"preview_webhook_url"`
	ResourceType      types.String        `tfsdk:"resource_type"`
	Settings          []FunctionSettingV1 `tfsdk:"settings"`
}
