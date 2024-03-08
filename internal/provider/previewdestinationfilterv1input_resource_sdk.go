// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/operations"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
)

func (r *PreviewDestinationFilterV1InputResourceModel) ToSharedPreviewDestinationFilterV1Input() *shared.PreviewDestinationFilterV1Input {
	var actions []shared.DestinationFilterActionV1 = nil
	for _, actionsItem := range r.Filter.Actions {
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
	ifVar := r.Filter.If.ValueString()
	filter := shared.PreviewDestinationFilterV1InputFilter{
		Actions: actions,
		If:      ifVar,
	}
	payload := make(map[string]interface{})
	for payloadKey, payloadValue := range r.Payload {
		var payloadInst interface{}
		_ = json.Unmarshal([]byte(payloadValue.ValueString()), &payloadInst)
		payload[payloadKey] = payloadInst
	}
	out := shared.PreviewDestinationFilterV1Input{
		Filter:  filter,
		Payload: payload,
	}
	return &out
}

func (r *PreviewDestinationFilterV1InputResourceModel) RefreshFromOperationsPreviewDestinationFilterResponseBody(resp *operations.PreviewDestinationFilterResponseBody) {
	if resp != nil {
		if resp.Data == nil {
			r.Data = nil
		} else {
			r.Data = &PreviewDestinationFilterV1Output{}
			if len(resp.Data.InputPayload) > 0 {
				r.Data.InputPayload = make(map[string]types.String)
				for key, value := range resp.Data.InputPayload {
					result, _ := json.Marshal(value)
					r.Data.InputPayload[key] = types.StringValue(string(result))
				}
			}
			if resp.Data.Result == nil {
				r.Data.Result = nil
			} else {
				r.Data.Result = &Result{}
			}
		}
	}
}
