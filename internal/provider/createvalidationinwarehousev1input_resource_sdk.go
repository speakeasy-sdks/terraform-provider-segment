// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"segment/internal/sdk/pkg/models/shared"
)

func (r *CreateValidationInWarehouseV1InputResourceModel) ToCreateSDKType() *shared.CreateValidationInWarehouseV1Input {
	metadataID := r.MetadataID.ValueString()
	settings := make(map[string]interface{})
	for settingsKey, settingsValue := range r.Settings {
		var settingsInst interface{}
		_ = json.Unmarshal([]byte(settingsValue.ValueString()), &settingsInst)
		settings[settingsKey] = settingsInst
	}
	out := shared.CreateValidationInWarehouseV1Input{
		MetadataID: metadataID,
		Settings:   settings,
	}
	return &out
}

func (r *CreateValidationInWarehouseV1InputResourceModel) RefreshFromCreateResponse(resp *shared.RequestErrorEnvelope) {
	r.Errors = nil
	for _, errorsItem := range resp.Errors {
		var errors1 RequestError
		if errorsItem.Data == nil {
			errors1.Data = types.StringNull()
		} else {
			dataResult, _ := json.Marshal(errorsItem.Data)
			errors1.Data = types.StringValue(string(dataResult))
		}
		if errorsItem.Field != nil {
			errors1.Field = types.StringValue(*errorsItem.Field)
		} else {
			errors1.Field = types.StringNull()
		}
		if errorsItem.Message != nil {
			errors1.Message = types.StringValue(*errorsItem.Message)
		} else {
			errors1.Message = types.StringNull()
		}
		if errorsItem.Status != nil {
			errors1.Status = types.NumberValue(big.NewFloat(float64(*errorsItem.Status)))
		} else {
			errors1.Status = types.NumberNull()
		}
		errors1.Type = types.StringValue(errorsItem.Type)
		r.Errors = append(r.Errors, errors1)
	}
}
