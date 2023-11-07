// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"segment/internal/sdk/pkg/models/operations"
	"segment/internal/sdk/pkg/models/shared"
)

func (r *CreateWarehouseV1InputResourceModel) ToCreateSDKType() *shared.CreateWarehouseV1Input {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	metadataID := r.MetadataID.ValueString()
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	settings := make(map[string]interface{})
	for settingsKey, settingsValue := range r.Settings {
		var settingsInst interface{}
		_ = json.Unmarshal([]byte(settingsValue.ValueString()), &settingsInst)
		settings[settingsKey] = settingsInst
	}
	out := shared.CreateWarehouseV1Input{
		Enabled:    enabled,
		MetadataID: metadataID,
		Name:       name,
		Settings:   settings,
	}
	return &out
}

func (r *CreateWarehouseV1InputResourceModel) RefreshFromCreateResponse(resp *operations.CreateWarehouseResponseBody) {
	if resp.Data == nil {
		r.Data = nil
	} else {
		r.Data = &CreateWarehouseV1Output{}
		r.Data.Warehouse.Enabled = types.BoolValue(resp.Data.Warehouse.Enabled)
		r.Data.Warehouse.ID = types.StringValue(resp.Data.Warehouse.ID)
		r.Data.Warehouse.Metadata.Description = types.StringValue(resp.Data.Warehouse.Metadata.Description)
		r.Data.Warehouse.Metadata.ID = types.StringValue(resp.Data.Warehouse.Metadata.ID)
		if resp.Data.Warehouse.Metadata.Logos.Alt != nil {
			r.Data.Warehouse.Metadata.Logos.Alt = types.StringValue(*resp.Data.Warehouse.Metadata.Logos.Alt)
		} else {
			r.Data.Warehouse.Metadata.Logos.Alt = types.StringNull()
		}
		r.Data.Warehouse.Metadata.Logos.Default = types.StringValue(resp.Data.Warehouse.Metadata.Logos.Default)
		if resp.Data.Warehouse.Metadata.Logos.Mark != nil {
			r.Data.Warehouse.Metadata.Logos.Mark = types.StringValue(*resp.Data.Warehouse.Metadata.Logos.Mark)
		} else {
			r.Data.Warehouse.Metadata.Logos.Mark = types.StringNull()
		}
		r.Data.Warehouse.Metadata.Name = types.StringValue(resp.Data.Warehouse.Metadata.Name)
		r.Data.Warehouse.Metadata.Options = nil
		for _, optionsItem := range resp.Data.Warehouse.Metadata.Options {
			var options1 IntegrationOptionBeta
			if optionsItem.DefaultValue == nil {
				options1.DefaultValue = types.StringNull()
			} else {
				defaultValueResult, _ := json.Marshal(optionsItem.DefaultValue)
				options1.DefaultValue = types.StringValue(string(defaultValueResult))
			}
			if optionsItem.Description != nil {
				options1.Description = types.StringValue(*optionsItem.Description)
			} else {
				options1.Description = types.StringNull()
			}
			if optionsItem.Label != nil {
				options1.Label = types.StringValue(*optionsItem.Label)
			} else {
				options1.Label = types.StringNull()
			}
			options1.Name = types.StringValue(optionsItem.Name)
			options1.Required = types.BoolValue(optionsItem.Required)
			options1.Type = types.StringValue(optionsItem.Type)
			r.Data.Warehouse.Metadata.Options = append(r.Data.Warehouse.Metadata.Options, options1)
		}
		r.Data.Warehouse.Metadata.Slug = types.StringValue(resp.Data.Warehouse.Metadata.Slug)
		if r.Data.Warehouse.Settings == nil && len(resp.Data.Warehouse.Settings) > 0 {
			r.Data.Warehouse.Settings = make(map[string]types.String)
			for key, value := range resp.Data.Warehouse.Settings {
				result, _ := json.Marshal(value)
				r.Data.Warehouse.Settings[key] = types.StringValue(string(result))
			}
		}
		r.Data.Warehouse.WorkspaceID = types.StringValue(resp.Data.Warehouse.WorkspaceID)
	}
}
