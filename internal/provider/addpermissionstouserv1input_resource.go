// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"segment/internal/sdk"
	"segment/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"segment/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AddPermissionsToUserV1InputResource{}
var _ resource.ResourceWithImportState = &AddPermissionsToUserV1InputResource{}

func NewAddPermissionsToUserV1InputResource() resource.Resource {
	return &AddPermissionsToUserV1InputResource{}
}

// AddPermissionsToUserV1InputResource defines the resource implementation.
type AddPermissionsToUserV1InputResource struct {
	client *sdk.Segment
}

// AddPermissionsToUserV1InputResourceModel describes the resource data model.
type AddPermissionsToUserV1InputResourceModel struct {
	Errors      []RequestError      `tfsdk:"errors"`
	Permissions []PermissionInputV1 `tfsdk:"permissions"`
	UserID      types.String        `tfsdk:"user_id"`
}

func (r *AddPermissionsToUserV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_add_permissions_to_user_v1_input"
}

func (r *AddPermissionsToUserV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AddPermissionsToUserV1Input Resource",

		Attributes: map[string]schema.Attribute{
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
			"permissions": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"resources": schema.ListNestedAttribute{
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplace(),
							},
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required:    true,
										Description: `The id of this resource.`,
									},
									"labels": schema.ListNestedAttribute{
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
													Optional:    true,
													Description: `A description of what this label represents.`,
												},
												"key": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Required:    true,
													Description: `The key identifier for this label.`,
												},
												"value": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Required:    true,
													Description: `The value of this label.`,
												},
											},
										},
										Description: `The labels that further refine access to this resource. Labels are exclusive to Workspace-level permissions.`,
									},
									"type": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"FUNCTION",
												"SOURCE",
												"SPACE",
												"WAREHOUSE",
												"WORKSPACE",
											),
										},
										MarkdownDescription: `must be one of ["FUNCTION", "SOURCE", "SPACE", "WAREHOUSE", "WORKSPACE"]` + "\n" +
											`The type for this resource.`,
									},
								},
							},
							Description: `The resources to link the selected role to.`,
						},
						"role_id": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `The role id of this permission.`,
						},
					},
				},
				Description: `The permissions to add.`,
			},
			"user_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
		},
	}
}

func (r *AddPermissionsToUserV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AddPermissionsToUserV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AddPermissionsToUserV1InputResourceModel
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

	addPermissionsToUserV1Input := *data.ToCreateSDKType()
	userID := data.UserID.ValueString()
	request := operations.AddPermissionsToUserRequest{
		AddPermissionsToUserV1Input: addPermissionsToUserV1Input,
		UserID:                      userID,
	}
	res, err := r.client.IAMUsers.AddPermissionsToUser(ctx, request)
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

func (r *AddPermissionsToUserV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AddPermissionsToUserV1InputResourceModel
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

func (r *AddPermissionsToUserV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AddPermissionsToUserV1InputResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AddPermissionsToUserV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AddPermissionsToUserV1InputResourceModel
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

func (r *AddPermissionsToUserV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource add_permissions_to_user_v1_input.")
}
