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
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CreateTrackingPlanV1InputResource{}
var _ resource.ResourceWithImportState = &CreateTrackingPlanV1InputResource{}

func NewCreateTrackingPlanV1InputResource() resource.Resource {
	return &CreateTrackingPlanV1InputResource{}
}

// CreateTrackingPlanV1InputResource defines the resource implementation.
type CreateTrackingPlanV1InputResource struct {
	client *sdk.Segment
}

// CreateTrackingPlanV1InputResourceModel describes the resource data model.
type CreateTrackingPlanV1InputResourceModel struct {
	Data        *CreateTrackingPlanV1Output `tfsdk:"data"`
	Description types.String                `tfsdk:"description"`
	Name        types.String                `tfsdk:"name"`
	Type        types.String                `tfsdk:"type"`
}

func (r *CreateTrackingPlanV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_create_tracking_plan_v1_input"
}

func (r *CreateTrackingPlanV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CreateTrackingPlanV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"tracking_plan": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"created_at": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `The timestamp of this Tracking Plan's creation.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `createTime` + "`" + `.`,
							},
							"description": schema.StringAttribute{
								Computed:    true,
								Description: `The Tracking Plan's description.`,
							},
							"id": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `The Tracking Plan's identifier.` + "\n" +
									`` + "\n" +
									`Config API note: analogous to ` + "`" + `name` + "`" + `.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `The Tracking Plan's name.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `displayName` + "`" + `.`,
							},
							"slug": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `URL-friendly slug of this Tracking Plan.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `name` + "`" + `.`,
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Description: `The Tracking Plan's type. must be one of ["ENGAGE", "LIVE", "PROPERTY_LIBRARY", "RULE_LIBRARY", "TEMPLATE"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"ENGAGE",
										"LIVE",
										"PROPERTY_LIBRARY",
										"RULE_LIBRARY",
										"TEMPLATE",
									),
								},
							},
							"updated_at": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `The timestamp of the last change to the Tracking Plan.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `updateTime` + "`" + `.`,
							},
						},
						Description: `The created Tracking Plan.`,
					},
				},
				Description: `Result of a CreateTrackingPlan call.`,
			},
			"description": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The Tracking Plan's description. Requires replacement if changed. `,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				MarkdownDescription: `The Tracking Plan's name.` + "\n" +
					`` + "\n" +
					`Config API note: equal to ` + "`" + `displayName` + "`" + `.` + "\n" +
					`Requires replacement if changed. `,
			},
			"type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The Tracking Plan's type. Requires replacement if changed. ; must be one of ["ENGAGE", "LIVE", "PROPERTY_LIBRARY", "RULE_LIBRARY", "TEMPLATE"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ENGAGE",
						"LIVE",
						"PROPERTY_LIBRARY",
						"RULE_LIBRARY",
						"TEMPLATE",
					),
				},
			},
		},
	}
}

func (r *CreateTrackingPlanV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CreateTrackingPlanV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CreateTrackingPlanV1InputResourceModel
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

	request := *data.ToSharedCreateTrackingPlanV1Input()
	res, err := r.client.TrackingPlans.CreateTrackingPlan(ctx, request)
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
	data.RefreshFromOperationsCreateTrackingPlanResponseBody(res.TwoHundredApplicationJSONObject)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CreateTrackingPlanV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CreateTrackingPlanV1InputResourceModel
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

func (r *CreateTrackingPlanV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CreateTrackingPlanV1InputResourceModel
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

func (r *CreateTrackingPlanV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CreateTrackingPlanV1InputResourceModel
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

func (r *CreateTrackingPlanV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource create_tracking_plan_v1_input.")
}
