// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_boolplanmodifier "github.com/scentregroup/terraform-provider-segment/internal/planmodifiers/boolplanmodifier"
	speakeasy_listplanmodifier "github.com/scentregroup/terraform-provider-segment/internal/planmodifiers/listplanmodifier"
	speakeasy_mapplanmodifier "github.com/scentregroup/terraform-provider-segment/internal/planmodifiers/mapplanmodifier"
	speakeasy_objectplanmodifier "github.com/scentregroup/terraform-provider-segment/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/scentregroup/terraform-provider-segment/internal/planmodifiers/stringplanmodifier"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"
	"github.com/scentregroup/terraform-provider-segment/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CreateSourceV1InputResource{}
var _ resource.ResourceWithImportState = &CreateSourceV1InputResource{}

func NewCreateSourceV1InputResource() resource.Resource {
	return &CreateSourceV1InputResource{}
}

// CreateSourceV1InputResource defines the resource implementation.
type CreateSourceV1InputResource struct {
	client *sdk.Segment
}

// CreateSourceV1InputResourceModel describes the resource data model.
type CreateSourceV1InputResourceModel struct {
	Data       *CreateSourceV1Output   `tfsdk:"data"`
	Enabled    types.Bool              `tfsdk:"enabled"`
	MetadataID types.String            `tfsdk:"metadata_id"`
	Settings   map[string]types.String `tfsdk:"settings"`
	Slug       types.String            `tfsdk:"slug"`
}

func (r *CreateSourceV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_create_source_v1_input"
}

func (r *CreateSourceV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CreateSourceV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.Standard),
				},
				Attributes: map[string]schema.Attribute{
					"source": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.Standard),
						},
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Bool{
									speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.Standard),
								},
								Description: `Enable to receive data from the Source.`,
							},
							"id": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
								},
								MarkdownDescription: `The id of the Source.` + "\n" +
									`` + "\n" +
									`Config API note: analogous to ` + "`" + `name` + "`" + `.`,
							},
							"labels": schema.ListNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.List{
									speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.Standard),
								},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
											},
											Description: `An optional description of the purpose of this label.`,
										},
										"key": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
											},
											Description: `The key that represents the name of this label.`,
										},
										"value": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
											},
											Description: `The value associated with the key of this label.`,
										},
									},
								},
								Description: `A list of labels applied to the Source.`,
							},
							"metadata": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.Standard),
								},
								Attributes: map[string]schema.Attribute{
									"categories": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.Standard),
										},
										ElementType: types.StringType,
										Description: `A list of categories this Source belongs to.`,
									},
									"description": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
										},
										Description: `The description of this Source.`,
									},
									"id": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
										},
										MarkdownDescription: `The id for this Source metadata in the Segment catalog.` + "\n" +
											`` + "\n" +
											`Config API note: analogous to ` + "`" + `name` + "`" + `.`,
									},
									"is_cloud_event_source": schema.BoolAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.Bool{
											speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.Standard),
										},
										Description: `True if this is a Cloud Event Source.`,
									},
									"logos": schema.SingleNestedAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.Object{
											speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.Standard),
										},
										Attributes: map[string]schema.Attribute{
											"alt": schema.StringAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.String{
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
												},
												Description: `The alternative text for this logo.`,
											},
											"default": schema.StringAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.String{
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
												},
												Description: `The default URL for this logo.`,
											},
											"mark": schema.StringAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.String{
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
												},
												Description: `The logo mark.`,
											},
										},
										Description: `The logos for this Source.`,
									},
									"name": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
										},
										MarkdownDescription: `The user-friendly name of this Source.` + "\n" +
											`` + "\n" +
											`Config API note: equal to ` + "`" + `displayName` + "`" + `.`,
									},
									"options": schema.ListNestedAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.Standard),
										},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"default_value": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
													},
													MarkdownDescription: `Parsed as JSON.` + "\n" +
														`An optional default value for the field.`,
													Validators: []validator.String{
														validators.IsValidJSON(),
													},
												},
												"description": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
													},
													Description: `An optional short text description of the field.`,
												},
												"label": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
													},
													Description: `An optional label for this field.`,
												},
												"name": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
													},
													Description: `The name identifying this option in the context of a Segment Integration.`,
												},
												"required": schema.BoolAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.Bool{
														speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.Standard),
													},
													Description: `Whether this is a required option when setting up the Integration.`,
												},
												"type": schema.StringAttribute{
													Computed: true,
													PlanModifiers: []planmodifier.String{
														speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
													},
													MarkdownDescription: `Defines the type for this option in the schema. Types are most commonly strings, but may also represent other` + "\n" +
														`primitive types, such as booleans, and numbers, as well as complex types, such as objects and arrays.`,
												},
											},
										},
										Description: `Options for this Source.`,
									},
									"slug": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
										},
										MarkdownDescription: `The slug that identifies this Source in the Segment app.` + "\n" +
											`` + "\n" +
											`Config API note: equal to ` + "`" + `name` + "`" + `.`,
									},
								},
								MarkdownDescription: `The metadata for the Source.` + "\n" +
									`` + "\n" +
									`Config API note: includes ` + "`" + `catalogName` + "`" + ` and ` + "`" + `catalogId` + "`" + `.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
								},
								MarkdownDescription: `The name of the Source.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `displayName` + "`" + `.`,
							},
							"settings": schema.MapAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Map{
									speakeasy_mapplanmodifier.SuppressDiff(speakeasy_mapplanmodifier.Standard),
								},
								ElementType: types.StringType,
								Description: `The settings associated with the Source.`,
								Validators: []validator.Map{
									mapvalidator.ValueStringsAre(validators.IsValidJSON()),
								},
							},
							"slug": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
								},
								MarkdownDescription: `The slug used to identify the Source in the Segment app.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `name` + "`" + `.`,
							},
							"workspace_id": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.Standard),
								},
								MarkdownDescription: `The id of the Workspace that owns the Source.` + "\n" +
									`` + "\n" +
									`Config API note: equal to ` + "`" + `parent` + "`" + `.`,
							},
							"write_keys": schema.ListAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.List{
									speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.Standard),
								},
								ElementType: types.StringType,
								MarkdownDescription: `The write keys used to send data from the Source. This field is left empty when the current token does not have the` + "\n" +
									`'source admin' permission.`,
							},
						},
						Description: `The newly created Source.`,
					},
				},
				Description: `Returns a newly created Source.`,
			},
			"enabled": schema.BoolAttribute{
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `Enable to allow this Source to send data. Defaults to true.`,
			},
			"metadata_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required: true,
				MarkdownDescription: `The id of the Source metadata from which this instance of the Source derives.` + "\n" +
					`` + "\n" +
					`All Source metadata is available under ` + "`" + `/catalog/sources` + "`" + `.`,
			},
			"settings": schema.MapAttribute{
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `A key-value object that contains instance-specific settings for the Source.`,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"slug": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The slug by which to identify the Source in the Segment app.`,
			},
		},
	}
}

func (r *CreateSourceV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CreateSourceV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CreateSourceV1InputResourceModel
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

	request := *data.ToSharedCreateSourceV1Input()
	res, err := r.client.Sources.CreateSource(ctx, request)
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
	data.RefreshFromOperationsCreateSourceResponseBody(res.TwoHundredApplicationJSONObject)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CreateSourceV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CreateSourceV1InputResourceModel
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

func (r *CreateSourceV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CreateSourceV1InputResourceModel
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

func (r *CreateSourceV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CreateSourceV1InputResourceModel
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

func (r *CreateSourceV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource create_source_v1_input.")
}
