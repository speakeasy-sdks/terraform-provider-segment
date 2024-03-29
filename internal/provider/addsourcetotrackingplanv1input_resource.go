// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/scentregroup/terraform-provider-segment/internal/provider/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AddSourceToTrackingPlanV1InputResource{}
var _ resource.ResourceWithImportState = &AddSourceToTrackingPlanV1InputResource{}

func NewAddSourceToTrackingPlanV1InputResource() resource.Resource {
	return &AddSourceToTrackingPlanV1InputResource{}
}

// AddSourceToTrackingPlanV1InputResource defines the resource implementation.
type AddSourceToTrackingPlanV1InputResource struct {
	client *sdk.Segment
}

// AddSourceToTrackingPlanV1InputResourceModel describes the resource data model.
type AddSourceToTrackingPlanV1InputResourceModel struct {
	Data           *tfTypes.AddSourceToTrackingPlanV1Output `tfsdk:"data"`
	SourceID       types.String                             `tfsdk:"source_id"`
	TrackingPlanID types.String                             `tfsdk:"tracking_plan_id"`
}

func (r *AddSourceToTrackingPlanV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_add_source_to_tracking_plan_v1_input"
}

func (r *AddSourceToTrackingPlanV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AddSourceToTrackingPlanV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:    true,
						Description: `The operation status. must be one of ["SUCCESS"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"SUCCESS",
							),
						},
					},
				},
				Description: `Connects a Source to a Tracking Plan.`,
			},
			"source_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				MarkdownDescription: `The id of the Source associated with the Tracking Plan.` + "\n" +
					`` + "\n" +
					`Config API note: analogous to ` + "`" + `sourceName` + "`" + `.` + "\n" +
					`Requires replacement if changed. `,
			},
			"tracking_plan_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
		},
	}
}

func (r *AddSourceToTrackingPlanV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AddSourceToTrackingPlanV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AddSourceToTrackingPlanV1InputResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	addSourceToTrackingPlanV1Input := *data.ToSharedAddSourceToTrackingPlanV1Input()
	trackingPlanID := data.TrackingPlanID.ValueString()
	request := operations.AddSourceToTrackingPlanRequest{
		AddSourceToTrackingPlanV1Input: addSourceToTrackingPlanV1Input,
		TrackingPlanID:                 trackingPlanID,
	}
	res, err := r.client.TrackingPlans.AddSourceToTrackingPlan(ctx, request)
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
	if res.TwoHundredApplicationJSONObject == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromOperationsAddSourceToTrackingPlanResponseBody(res.TwoHundredApplicationJSONObject)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AddSourceToTrackingPlanV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AddSourceToTrackingPlanV1InputResourceModel
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

func (r *AddSourceToTrackingPlanV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AddSourceToTrackingPlanV1InputResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AddSourceToTrackingPlanV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AddSourceToTrackingPlanV1InputResourceModel
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

func (r *AddSourceToTrackingPlanV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource add_source_to_tracking_plan_v1_input.")
}
