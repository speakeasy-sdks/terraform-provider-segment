// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/operations"
	"github.com/scentregroup/terraform-provider-segment/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CreateFilterForDestinationV1InputResource{}
var _ resource.ResourceWithImportState = &CreateFilterForDestinationV1InputResource{}

func NewCreateFilterForDestinationV1InputResource() resource.Resource {
	return &CreateFilterForDestinationV1InputResource{}
}

// CreateFilterForDestinationV1InputResource defines the resource implementation.
type CreateFilterForDestinationV1InputResource struct {
	client *sdk.Segment
}

// CreateFilterForDestinationV1InputResourceModel describes the resource data model.
type CreateFilterForDestinationV1InputResourceModel struct {
	Actions       []DestinationFilterActionV1         `tfsdk:"actions"`
	Data          *CreateFilterForDestinationV1Output `tfsdk:"data"`
	Description   types.String                        `tfsdk:"description"`
	DestinationID types.String                        `tfsdk:"destination_id"`
	Enabled       types.Bool                          `tfsdk:"enabled"`
	If            types.String                        `tfsdk:"if"`
	SourceID      types.String                        `tfsdk:"source_id"`
	Title         types.String                        `tfsdk:"title"`
}

func (r *CreateFilterForDestinationV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_create_filter_for_destination_v1_input"
}

func (r *CreateFilterForDestinationV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CreateFilterForDestinationV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"actions": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fields": schema.MapAttribute{
							PlanModifiers: []planmodifier.Map{
								mapplanmodifier.RequiresReplaceIfConfigured(),
							},
							Optional:    true,
							ElementType: types.StringType,
							MarkdownDescription: `A dictionary of paths to object keys that this filter applies to.` + "\n" +
								`  The literal string '' represents the top level of the object.` + "\n" +
								`Requires replacement if changed. `,
							Validators: []validator.Map{
								mapvalidator.ValueStringsAre(validators.IsValidJSON()),
							},
						},
						"path": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
							},
							Optional: true,
							MarkdownDescription: `The JSON path to a property within a payload object from which Segment generates a deterministic` + "\n" +
								`sampling rate.` + "\n" +
								`Requires replacement if changed. `,
						},
						"percent": schema.NumberAttribute{
							PlanModifiers: []planmodifier.Number{
								numberplanmodifier.RequiresReplaceIfConfigured(),
							},
							Optional: true,
							MarkdownDescription: `A decimal between 0 and 1 used for 'sample' type events and` + "\n" +
								`influences the likelihood of sampling to occur.` + "\n" +
								`Requires replacement if changed. `,
						},
						"type": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `The kind of Transformation to apply to any matched properties. Requires replacement if changed. ; must be one of ["ALLOW_PROPERTIES", "DROP", "DROP_PROPERTIES", "SAMPLE"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ALLOW_PROPERTIES",
									"DROP",
									"DROP_PROPERTIES",
									"SAMPLE",
								),
							},
						},
					},
				},
				Description: `Actions for the Destination filter. Requires replacement if changed. `,
			},
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"filter": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"actions": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"fields": schema.MapAttribute{
											Computed:    true,
											ElementType: types.StringType,
											MarkdownDescription: `A dictionary of paths to object keys that this filter applies to.` + "\n" +
												`  The literal string '' represents the top level of the object.`,
											Validators: []validator.Map{
												mapvalidator.ValueStringsAre(validators.IsValidJSON()),
											},
										},
										"path": schema.StringAttribute{
											Computed: true,
											MarkdownDescription: `The JSON path to a property within a payload object from which Segment generates a deterministic` + "\n" +
												`sampling rate.`,
										},
										"percent": schema.NumberAttribute{
											Computed: true,
											MarkdownDescription: `A decimal between 0 and 1 used for 'sample' type events and` + "\n" +
												`influences the likelihood of sampling to occur.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `The kind of Transformation to apply to any matched properties. must be one of ["ALLOW_PROPERTIES", "DROP", "DROP_PROPERTIES", "SAMPLE"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"ALLOW_PROPERTIES",
													"DROP",
													"DROP_PROPERTIES",
													"SAMPLE",
												),
											},
										},
									},
								},
								Description: `A list of actions this filter performs.`,
							},
							"created_at": schema.StringAttribute{
								Computed:    true,
								Description: `The timestamp of this filter's creation.`,
							},
							"description": schema.StringAttribute{
								Computed:    true,
								Description: `A description for this filter.`,
							},
							"destination_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the Destination associated with this filter.`,
							},
							"enabled": schema.BoolAttribute{
								Computed:    true,
								Description: `When set to true, this filter is active.`,
							},
							"id": schema.StringAttribute{
								Computed:    true,
								Description: `The unique id of this filter.`,
							},
							"if": schema.StringAttribute{
								Computed:    true,
								Description: `A condition that defines whether to apply this filter to a payload.`,
							},
							"source_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the Source associated with this filter.`,
							},
							"title": schema.StringAttribute{
								Computed:    true,
								Description: `A title for this filter.`,
							},
							"updated_at": schema.StringAttribute{
								Computed:    true,
								Description: `The timestamp of this filter's last change.`,
							},
						},
						Description: `The newly created Destination filter.`,
					},
				},
				Description: `Output for CreateDestinationFiltersV1.`,
			},
			"description": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The description of the filter. Requires replacement if changed. `,
			},
			"destination_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"enabled": schema.BoolAttribute{
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `When set to true, the Destination filter is active. Requires replacement if changed. `,
			},
			"if": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The filter's condition. Requires replacement if changed. `,
			},
			"source_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The id of the Source associated with this filter. Requires replacement if changed. `,
			},
			"title": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The title of the filter. Requires replacement if changed. `,
			},
		},
	}
}

func (r *CreateFilterForDestinationV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CreateFilterForDestinationV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CreateFilterForDestinationV1InputResourceModel
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

	createFilterForDestinationV1Input := *data.ToSharedCreateFilterForDestinationV1Input()
	destinationID := data.DestinationID.ValueString()
	request := operations.CreateFilterForDestinationRequest{
		CreateFilterForDestinationV1Input: createFilterForDestinationV1Input,
		DestinationID:                     destinationID,
	}
	res, err := r.client.DestinationFilters.CreateFilterForDestination(ctx, request)
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
	data.RefreshFromOperationsCreateFilterForDestinationResponseBody(res.TwoHundredApplicationJSONObject)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CreateFilterForDestinationV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CreateFilterForDestinationV1InputResourceModel
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

func (r *CreateFilterForDestinationV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CreateFilterForDestinationV1InputResourceModel
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

func (r *CreateFilterForDestinationV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CreateFilterForDestinationV1InputResourceModel
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

func (r *CreateFilterForDestinationV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource create_filter_for_destination_v1_input.")
}
