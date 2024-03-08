// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/operations"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"math/big"
)

func (r *CreateFunctionV1InputResourceModel) ToSharedCreateFunctionV1Input() *shared.CreateFunctionV1Input {
	code := r.Code.ValueString()
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	displayName := r.DisplayName.ValueString()
	logoURL := new(string)
	if !r.LogoURL.IsUnknown() && !r.LogoURL.IsNull() {
		*logoURL = r.LogoURL.ValueString()
	} else {
		logoURL = nil
	}
	resourceType := shared.ResourceType(r.ResourceType.ValueString())
	var settings []shared.FunctionSettingV1 = nil
	for _, settingsItem := range r.Settings {
		description1 := settingsItem.Description.ValueString()
		label := settingsItem.Label.ValueString()
		name := settingsItem.Name.ValueString()
		required := settingsItem.Required.ValueBool()
		sensitive := settingsItem.Sensitive.ValueBool()
		typeVar := shared.FunctionSettingV1Type(settingsItem.Type.ValueString())
		settings = append(settings, shared.FunctionSettingV1{
			Description: description1,
			Label:       label,
			Name:        name,
			Required:    required,
			Sensitive:   sensitive,
			Type:        typeVar,
		})
	}
	out := shared.CreateFunctionV1Input{
		Code:         code,
		Description:  description,
		DisplayName:  displayName,
		LogoURL:      logoURL,
		ResourceType: resourceType,
		Settings:     settings,
	}
	return &out
}

func (r *CreateFunctionV1InputResourceModel) RefreshFromOperationsCreateFunctionResponseBody(resp *operations.CreateFunctionResponseBody) {
	if resp != nil {
		if resp.Data == nil {
			r.Data = nil
		} else {
			r.Data = &CreateFunctionV1Output{}
			if resp.Data.Function.BatchMaxCount != nil {
				r.Data.Function.BatchMaxCount = types.NumberValue(big.NewFloat(float64(*resp.Data.Function.BatchMaxCount)))
			} else {
				r.Data.Function.BatchMaxCount = types.NumberNull()
			}
			r.Data.Function.CatalogID = types.StringPointerValue(resp.Data.Function.CatalogID)
			r.Data.Function.Code = types.StringPointerValue(resp.Data.Function.Code)
			r.Data.Function.CreatedAt = types.StringPointerValue(resp.Data.Function.CreatedAt)
			r.Data.Function.CreatedBy = types.StringPointerValue(resp.Data.Function.CreatedBy)
			r.Data.Function.DeployedAt = types.StringPointerValue(resp.Data.Function.DeployedAt)
			r.Data.Function.Description = types.StringPointerValue(resp.Data.Function.Description)
			r.Data.Function.DisplayName = types.StringPointerValue(resp.Data.Function.DisplayName)
			r.Data.Function.ID = types.StringPointerValue(resp.Data.Function.ID)
			r.Data.Function.IsLatestVersion = types.BoolPointerValue(resp.Data.Function.IsLatestVersion)
			r.Data.Function.LogoURL = types.StringPointerValue(resp.Data.Function.LogoURL)
			r.Data.Function.PreviewWebhookURL = types.StringPointerValue(resp.Data.Function.PreviewWebhookURL)
			if resp.Data.Function.ResourceType != nil {
				r.Data.Function.ResourceType = types.StringValue(string(*resp.Data.Function.ResourceType))
			} else {
				r.Data.Function.ResourceType = types.StringNull()
			}
			if len(r.Data.Function.Settings) > len(resp.Data.Function.Settings) {
				r.Data.Function.Settings = r.Data.Function.Settings[:len(resp.Data.Function.Settings)]
			}
			for settingsCount, settingsItem := range resp.Data.Function.Settings {
				var settings1 FunctionSettingV1
				settings1.Description = types.StringValue(settingsItem.Description)
				settings1.Label = types.StringValue(settingsItem.Label)
				settings1.Name = types.StringValue(settingsItem.Name)
				settings1.Required = types.BoolValue(settingsItem.Required)
				settings1.Sensitive = types.BoolValue(settingsItem.Sensitive)
				settings1.Type = types.StringValue(string(settingsItem.Type))
				if settingsCount+1 > len(r.Data.Function.Settings) {
					r.Data.Function.Settings = append(r.Data.Function.Settings, settings1)
				} else {
					r.Data.Function.Settings[settingsCount].Description = settings1.Description
					r.Data.Function.Settings[settingsCount].Label = settings1.Label
					r.Data.Function.Settings[settingsCount].Name = settings1.Name
					r.Data.Function.Settings[settingsCount].Required = settings1.Required
					r.Data.Function.Settings[settingsCount].Sensitive = settings1.Sensitive
					r.Data.Function.Settings[settingsCount].Type = settings1.Type
				}
			}
		}
	}
}
