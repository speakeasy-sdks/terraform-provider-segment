// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CreateInvitesV1InputResource{}
var _ resource.ResourceWithImportState = &CreateInvitesV1InputResource{}

func NewCreateInvitesV1InputResource() resource.Resource {
	return &CreateInvitesV1InputResource{}
}

// CreateInvitesV1InputResource defines the resource implementation.
type CreateInvitesV1InputResource struct {
	client *sdk.Segment
}

// CreateInvitesV1InputResourceModel describes the resource data model.
type CreateInvitesV1InputResourceModel struct {
	Data    *CreateInvitesV1Output `tfsdk:"data"`
	Invites []InviteV1             `tfsdk:"invites"`
}

func (r *CreateInvitesV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_create_invites_v1_input"
}

func (r *CreateInvitesV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CreateInvitesV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"emails": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `The list of emails invited to the Workspace.`,
					},
				},
				Description: `Returns the emails of the invited users.`,
			},
			"invites": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"email": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `The invited user's email to attach the permissions to.`,
						},
						"permissions": schema.ListNestedAttribute{
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplace(),
							},
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
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
										Description: `The labels that determine which resources to grant users access to.`,
									},
									"resources": schema.ListNestedAttribute{
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplace(),
										},
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Required:    true,
													Description: `The id of this resource.`,
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
														`The kind of resource this permission applies to.`,
												},
											},
										},
										Description: `The resources to grant the invited users access to.`,
									},
									"role_id": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required:    true,
										Description: `The id of the role.`,
									},
								},
							},
							Description: `The permissions to attach to the invited user.`,
						},
					},
				},
				Description: `The list of invites.`,
			},
		},
	}
}

func (r *CreateInvitesV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CreateInvitesV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CreateInvitesV1InputResourceModel
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
	res, err := r.client.IAMUsers.CreateInvites(ctx, request)
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
	data.RefreshFromCreateResponse(res.TwoHundredApplicationJSONObject)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CreateInvitesV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CreateInvitesV1InputResourceModel
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

func (r *CreateInvitesV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CreateInvitesV1InputResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CreateInvitesV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CreateInvitesV1InputResourceModel
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

func (r *CreateInvitesV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource create_invites_v1_input.")
}
