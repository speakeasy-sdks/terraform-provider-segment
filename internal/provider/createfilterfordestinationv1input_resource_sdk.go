// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/operations"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
	"math/big"
)

func (r *CreateFilterForDestinationV1InputResourceModel) ToCreateSDKType() *shared.CreateFilterForDestinationV1Input {
	var actions []shared.DestinationFilterActionV1 = nil
	for _, actionsItem := range r.Actions {
		fields := make(map[string]interface{})
		for fieldsKey, fieldsValue := range actionsItem.Fields {
			var fieldsInst interface{}
			_ = json.Unmarshal([]byte(fieldsValue.ValueString()), &fieldsInst)
			fields[fieldsKey] = fieldsInst
		}
		path := new(string)
		if !actionsItem.Path.IsUnknown() && !actionsItem.Path.IsNull() {
			*path = actionsItem.Path.ValueString()
		} else {
			path = nil
		}
		percent := new(float64)
		if !actionsItem.Percent.IsUnknown() && !actionsItem.Percent.IsNull() {
			*percent, _ = actionsItem.Percent.ValueBigFloat().Float64()
		} else {
			percent = nil
		}
		typeVar := shared.DestinationFilterActionV1Type(actionsItem.Type.ValueString())
		actions = append(actions, shared.DestinationFilterActionV1{
			Fields:  fields,
			Path:    path,
			Percent: percent,
			Type:    typeVar,
		})
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	enabled := r.Enabled.ValueBool()
	ifVar := r.If.ValueString()
	sourceID := r.SourceID.ValueString()
	title := r.Title.ValueString()
	out := shared.CreateFilterForDestinationV1Input{
		Actions:     actions,
		Description: description,
		Enabled:     enabled,
		If:          ifVar,
		SourceID:    sourceID,
		Title:       title,
	}
	return &out
}

func (r *CreateFilterForDestinationV1InputResourceModel) RefreshFromCreateResponse(resp *operations.CreateFilterForDestinationResponseBody) {
	if resp.Data == nil {
		r.Data = nil
	} else {
		r.Data = &CreateFilterForDestinationV1Output{}
		r.Data.Filter.Actions = nil
		for _, actionsItem := range resp.Data.Filter.Actions {
			var actions1 DestinationFilterActionV1
			if actions1.Fields == nil && len(actionsItem.Fields) > 0 {
				actions1.Fields = make(map[string]types.String)
				for key, value := range actionsItem.Fields {
					result, _ := json.Marshal(value)
					actions1.Fields[key] = types.StringValue(string(result))
				}
			}
			if actionsItem.Path != nil {
				actions1.Path = types.StringValue(*actionsItem.Path)
			} else {
				actions1.Path = types.StringNull()
			}
			if actionsItem.Percent != nil {
				actions1.Percent = types.NumberValue(big.NewFloat(float64(*actionsItem.Percent)))
			} else {
				actions1.Percent = types.NumberNull()
			}
			actions1.Type = types.StringValue(string(actionsItem.Type))
			r.Data.Filter.Actions = append(r.Data.Filter.Actions, actions1)
		}
		r.Data.Filter.CreatedAt = types.StringValue(resp.Data.Filter.CreatedAt)
		if resp.Data.Filter.Description != nil {
			r.Data.Filter.Description = types.StringValue(*resp.Data.Filter.Description)
		} else {
			r.Data.Filter.Description = types.StringNull()
		}
		r.Data.Filter.DestinationID = types.StringValue(resp.Data.Filter.DestinationID)
		r.Data.Filter.Enabled = types.BoolValue(resp.Data.Filter.Enabled)
		r.Data.Filter.ID = types.StringValue(resp.Data.Filter.ID)
		r.Data.Filter.If = types.StringValue(resp.Data.Filter.If)
		r.Data.Filter.SourceID = types.StringValue(resp.Data.Filter.SourceID)
		r.Data.Filter.Title = types.StringValue(resp.Data.Filter.Title)
		r.Data.Filter.UpdatedAt = types.StringValue(resp.Data.Filter.UpdatedAt)
	}
}
