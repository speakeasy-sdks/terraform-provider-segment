// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/scentregroup/terraform-provider-segment/internal/provider/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CreateFunctionV1InputResource{}
var _ resource.ResourceWithImportState = &CreateFunctionV1InputResource{}

func NewCreateFunctionV1InputResource() resource.Resource {
	return &CreateFunctionV1InputResource{}
}

// CreateFunctionV1InputResource defines the resource implementation.
type CreateFunctionV1InputResource struct {
	client *sdk.Segment
}

// CreateFunctionV1InputResourceModel describes the resource data model.
type CreateFunctionV1InputResourceModel struct {
	Code         types.String                    `tfsdk:"code"`
	Data         *tfTypes.CreateFunctionV1Output `tfsdk:"data"`
	Description  types.String                    `tfsdk:"description"`
	DisplayName  types.String                    `tfsdk:"display_name"`
	LogoURL      types.String                    `tfsdk:"logo_url"`
	ResourceType types.String                    `tfsdk:"resource_type"`
	Settings     []tfTypes.FunctionSettingV1     `tfsdk:"settings"`
}

func (r *CreateFunctionV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_create_function_v1_input"
}

func (r *CreateFunctionV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CreateFunctionV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"code": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The Function code. Requires replacement if changed. `,
			},
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"function": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"batch_max_count": schema.NumberAttribute{
								Computed:    true,
								Description: `The max count of the batch for this Function.`,
							},
							"catalog_id": schema.StringAttribute{
								Computed:    true,
								Description: `The catalog id of this Function.`,
							},
							"code": schema.StringAttribute{
								Computed:    true,
								Description: `The Function code.`,
							},
							"created_at": schema.StringAttribute{
								Computed:    true,
								Description: `The time this Function was created.`,
							},
							"created_by": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the user who created this Function.`,
							},
							"deployed_at": schema.StringAttribute{
								Computed:    true,
								Description: `The time of this Function's last deployment.`,
							},
							"description": schema.StringAttribute{
								Computed:    true,
								Description: `A description for this Function.`,
							},
							"display_name": schema.StringAttribute{
								Computed:    true,
								Description: `A display name for this Function.`,
							},
							"id": schema.StringAttribute{
								Computed:    true,
								Description: `An identifier for this Function.`,
							},
							"is_latest_version": schema.BoolAttribute{
								Computed:    true,
								Description: `Whether the deployment of this Function is the latest version.`,
							},
							"logo_url": schema.StringAttribute{
								Computed:    true,
								Description: `The URL of the logo for this Function.`,
							},
							"preview_webhook_url": schema.StringAttribute{
								Computed:    true,
								Description: `The preview webhook URL for this Function.`,
							},
							"resource_type": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `The Function type.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `type` + "`" + `.` + "\n" +
									`must be one of ["DESTINATION", "INSERT_DESTINATION", "SOURCE"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"DESTINATION",
										"INSERT_DESTINATION",
										"SOURCE",
									),
								},
							},
							"settings": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed:    true,
											Description: `A description of this Function Setting.`,
										},
										"label": schema.StringAttribute{
											Computed:    true,
											Description: `The label for this Function Setting.`,
										},
										"name": schema.StringAttribute{
											Computed:    true,
											Description: `The name of this Function Setting.`,
										},
										"required": schema.BoolAttribute{
											Computed:    true,
											Description: `Whether this Function Setting is required.`,
										},
										"sensitive": schema.BoolAttribute{
											Computed:    true,
											Description: `Whether this Function Setting contains sensitive information.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `The type of this Function Setting. must be one of ["ARRAY", "BOOLEAN", "STRING", "TEXT_MAP"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"ARRAY",
													"BOOLEAN",
													"STRING",
													"TEXT_MAP",
												),
											},
										},
									},
								},
								Description: `The list of settings for this Function.`,
							},
						},
						Description: `A Function object.`,
					},
				},
				Description: `Create a Function.`,
			},
			"description": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `A description for this Function. Requires replacement if changed. `,
			},
			"display_name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				MarkdownDescription: `A display name for this Function.` + "\n" +
					`` + "\n" +
					`Note that Destination Functions append the Workspace to the display name.` + "\n" +
					`Requires replacement if changed. `,
			},
			"logo_url": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The URL of the logo for this Function. Requires replacement if changed. `,
			},
			"resource_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				MarkdownDescription: `The Function type.` + "\n" +
					`` + "\n" +
					`Config API note: equal to ` + "`" + `type` + "`" + `.` + "\n" +
					`Requires replacement if changed. ; must be one of ["DESTINATION", "INSERT_DESTINATION", "SOURCE"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"DESTINATION",
						"INSERT_DESTINATION",
						"SOURCE",
					),
				},
			},
			"settings": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `A description of this Function Setting. Requires replacement if changed. `,
						},
						"label": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `The label for this Function Setting. Requires replacement if changed. `,
						},
						"name": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `The name of this Function Setting. Requires replacement if changed. `,
						},
						"required": schema.BoolAttribute{
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `Whether this Function Setting is required. Requires replacement if changed. `,
						},
						"sensitive": schema.BoolAttribute{
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `Whether this Function Setting contains sensitive information. Requires replacement if changed. `,
						},
						"type": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
							},
							Required:    true,
							Description: `The type of this Function Setting. Requires replacement if changed. ; must be one of ["ARRAY", "BOOLEAN", "STRING", "TEXT_MAP"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ARRAY",
									"BOOLEAN",
									"STRING",
									"TEXT_MAP",
								),
							},
						},
					},
				},
				Description: `The list of settings for this Function. Requires replacement if changed. `,
			},
		},
	}
}

func (r *CreateFunctionV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CreateFunctionV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CreateFunctionV1InputResourceModel
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

	request := *data.ToSharedCreateFunctionV1Input()
	res, err := r.client.Functions.CreateFunction(ctx, request)
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
	data.RefreshFromOperationsCreateFunctionResponseBody(res.TwoHundredApplicationJSONObject)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CreateFunctionV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CreateFunctionV1InputResourceModel
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

func (r *CreateFunctionV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CreateFunctionV1InputResourceModel
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

func (r *CreateFunctionV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CreateFunctionV1InputResourceModel
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

func (r *CreateFunctionV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource create_function_v1_input.")
}
