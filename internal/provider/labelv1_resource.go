// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"segment/internal/sdk"
	"segment/internal/sdk/pkg/models/operations"
	"segment/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"segment/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LabelV1Resource{}
var _ resource.ResourceWithImportState = &LabelV1Resource{}

func NewLabelV1Resource() resource.Resource {
	return &LabelV1Resource{}
}

// LabelV1Resource defines the resource implementation.
type LabelV1Resource struct {
	client *sdk.Segment
}

// LabelV1ResourceModel describes the resource data model.
type LabelV1ResourceModel struct {
	Description types.String   `tfsdk:"description"`
	Errors      []RequestError `tfsdk:"errors"`
	Key         types.String   `tfsdk:"key"`
	Labels      []LabelV1      `tfsdk:"labels"`
	SourceID    types.String   `tfsdk:"source_id"`
	Value       types.String   `tfsdk:"value"`
}

func (r *LabelV1Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_label_v1"
}

func (r *LabelV1Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "LabelV1 Resource",

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Optional:    true,
				Description: `An optional description of the purpose of this label.`,
			},
			"errors": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsValidJSON(),
							},
							MarkdownDescription: `Parsed as JSON.` + "\n" +
								`Any extra data associated with this error.`,
						},
						"field": schema.StringAttribute{
							Computed:    true,
							Description: `The name of an input field from the request that triggered this error.`,
						},
						"message": schema.StringAttribute{
							Computed:    true,
							Description: `An error message attached to this error.`,
						},
						"status": schema.NumberAttribute{
							Computed:    true,
							Description: `Http status code.`,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `The type for this error (validation, server, unknown, etc).`,
						},
					},
				},
			},
			"key": schema.StringAttribute{
				Required:    true,
				Description: `The key that represents the name of this label.`,
			},
			"labels": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Optional:    true,
							Description: `An optional description of the purpose of this label.`,
						},
						"key": schema.StringAttribute{
							Required:    true,
							Description: `The key that represents the name of this label.`,
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: `The value associated with the key of this label.`,
						},
					},
				},
				Description: `The labels to associate with a Source.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: `The value associated with the key of this label.`,
			},
		},
	}
}

func (r *LabelV1Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Segment)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Segment, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *LabelV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *LabelV1ResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var labels []shared.LabelV1 = nil
	for _, labelsItem := range data.Labels {
		description := new(string)
		if !labelsItem.Description.IsUnknown() && !labelsItem.Description.IsNull() {
			*description = labelsItem.Description.ValueString()
		} else {
			description = nil
		}
		key := labelsItem.Key.ValueString()
		value := labelsItem.Value.ValueString()
		labels = append(labels, shared.LabelV1{
			Description: description,
			Key:         key,
			Value:       value,
		})
	}
	addLabelsToSourceV1Input := shared.AddLabelsToSourceV1Input{
		Labels: labels,
	}
	sourceID := data.SourceID.ValueString()
	request := operations.AddLabelsToSourceRequest{
		AddLabelsToSourceV1Input: addLabelsToSourceV1Input,
		SourceID:                 sourceID,
	}
	res, err := r.client.Sources.AddLabelsToSource(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RequestErrorEnvelope == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.RequestErrorEnvelope)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LabelV1Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *LabelV1ResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LabelV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *LabelV1ResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var labels []shared.LabelV1 = nil
	for _, labelsItem := range data.Labels {
		description := new(string)
		if !labelsItem.Description.IsUnknown() && !labelsItem.Description.IsNull() {
			*description = labelsItem.Description.ValueString()
		} else {
			description = nil
		}
		key := labelsItem.Key.ValueString()
		value := labelsItem.Value.ValueString()
		labels = append(labels, shared.LabelV1{
			Description: description,
			Key:         key,
			Value:       value,
		})
	}
	replaceLabelsInSourceV1Input := shared.ReplaceLabelsInSourceV1Input{
		Labels: labels,
	}
	sourceID := data.SourceID.ValueString()
	request := operations.ReplaceLabelsInSourceRequest{
		ReplaceLabelsInSourceV1Input: replaceLabelsInSourceV1Input,
		SourceID:                     sourceID,
	}
	res, err := r.client.Sources.ReplaceLabelsInSource(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RequestErrorEnvelope == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.RequestErrorEnvelope)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LabelV1Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *LabelV1ResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *LabelV1Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource label_v1.")
}
