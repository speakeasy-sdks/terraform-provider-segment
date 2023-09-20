// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"segment/internal/sdk"

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
	"segment/internal/validators"
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
	Code         types.String        `tfsdk:"code"`
	Description  types.String        `tfsdk:"description"`
	DisplayName  types.String        `tfsdk:"display_name"`
	Errors       []RequestError      `tfsdk:"errors"`
	LogoURL      types.String        `tfsdk:"logo_url"`
	ResourceType types.String        `tfsdk:"resource_type"`
	Settings     []FunctionSettingV1 `tfsdk:"settings"`
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
					stringplanmodifier.RequiresReplace(),
				},
				Required:    true,
				Description: `The Function code.`,
			},
			"description": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `A description for this Function.`,
			},
			"display_name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
				MarkdownDescription: `A display name for this Function.` + "\n" +
					`` + "\n" +
					`Note that Destination Functions append the Workspace to the display name.`,
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
			"logo_url": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The URL of the logo for this Function.`,
			},
			"resource_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"DESTINATION",
						"INSERT_DESTINATION",
						"SOURCE",
					),
				},
				MarkdownDescription: `must be one of ["DESTINATION", "INSERT_DESTINATION", "SOURCE"]` + "\n" +
					`The Function type.` + "\n" +
					`` + "\n" +
					`Config API note: equal to ` + "`" + `type` + "`" + `.`,
			},
			"settings": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `A description of this Function Setting.`,
						},
						"label": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `The label for this Function Setting.`,
						},
						"name": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `The name of this Function Setting.`,
						},
						"required": schema.BoolAttribute{
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `Whether this Function Setting is required.`,
						},
						"sensitive": schema.BoolAttribute{
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `Whether this Function Setting contains sensitive information.`,
						},
						"type": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ARRAY",
									"BOOLEAN",
									"STRING",
									"TEXT_MAP",
								),
							},
							MarkdownDescription: `must be one of ["ARRAY", "BOOLEAN", "STRING", "TEXT_MAP"]` + "\n" +
								`The type of this Function Setting.`,
						},
					},
				},
				Description: `The list of settings for this Function.`,
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

	request := *data.ToCreateSDKType()
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
	if res.RequestErrorEnvelope == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.RequestErrorEnvelope)

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
