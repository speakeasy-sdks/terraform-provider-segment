// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/shared"
)

func (r *AddConnectionFromSourceToWarehouseV1OutputResourceModel) RefreshFromSharedAddConnectionFromSourceToWarehouseV1Output(resp *shared.AddConnectionFromSourceToWarehouseV1Output) {
	if resp != nil {
		r.Status = types.StringValue(string(resp.Status))
	}
}
