// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/scentregroup/terraform-provider-segment/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PreviewDestinationFilterV1InputResource{}
var _ resource.ResourceWithImportState = &PreviewDestinationFilterV1InputResource{}

func NewPreviewDestinationFilterV1InputResource() resource.Resource {
	return &PreviewDestinationFilterV1InputResource{}
}

// PreviewDestinationFilterV1InputResource defines the resource implementation.
type PreviewDestinationFilterV1InputResource struct {
	client *sdk.Segment
}

// PreviewDestinationFilterV1InputResourceModel describes the resource data model.
type PreviewDestinationFilterV1InputResourceModel struct {
	Data    *PreviewDestinationFilterV1Output `tfsdk:"data"`
	Filter  PreviewDestinationFilterV1        `tfsdk:"filter"`
	Payload map[string]types.String           `tfsdk:"payload"`
}

func (r *PreviewDestinationFilterV1InputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_preview_destination_filter_v1_input"
}

func (r *PreviewDestinationFilterV1InputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PreviewDestinationFilterV1Input Resource",

		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"input_payload": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
						Description: `The pre-filter JSON input.`,
					},
					"result": schema.SingleNestedAttribute{
						Computed:    true,
						Attributes:  map[string]schema.Attribute{},
						Description: `The filtered JSON output.`,
					},
				},
				MarkdownDescription: `Preview output from applying the Destination filter.` + "\n" +
					`Segment modifies or nullifies payloads depending on the provided filter actions.`,
			},
			"filter": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"actions": schema.ListNestedAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplace(),
						},
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"fields": schema.MapAttribute{
									PlanModifiers: []planmodifier.Map{
										mapplanmodifier.RequiresReplace(),
									},
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.Map{
										mapvalidator.ValueStringsAre(validators.IsValidJSON()),
									},
									MarkdownDescription: `A dictionary of paths to object keys that this filter applies to.` + "\n" +
										`  The literal string '' represents the top level of the object.`,
								},
								"path": schema.StringAttribute{
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
									Optional: true,
									MarkdownDescription: `The JSON path to a property within a payload object from which Segment generates a deterministic` + "\n" +
										`sampling rate.`,
								},
								"percent": schema.NumberAttribute{
									PlanModifiers: []planmodifier.Number{
										numberplanmodifier.RequiresReplace(),
									},
									Optional: true,
									MarkdownDescription: `A decimal between 0 and 1 used for 'sample' type events and` + "\n" +
										`influences the likelihood of sampling to occur.`,
								},
								"type": schema.StringAttribute{
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
									Required: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"ALLOW_PROPERTIES",
											"DROP",
											"DROP_PROPERTIES",
											"SAMPLE",
										),
									},
									MarkdownDescription: `must be one of ["ALLOW_PROPERTIES", "DROP", "DROP_PROPERTIES", "SAMPLE"]` + "\n" +
										`The kind of Transformation to apply to any matched properties.`,
								},
							},
						},
						MarkdownDescription: `The filtering action to take on events that match the "if" statement.` + "\n" +
							`Action types must be one of: "drop", "allow_properties", "drop_properties" or "sample".`,
					},
					"if": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
						MarkdownDescription: `A FQL statement which determines if the provided filter's actions will apply to the provided JSON payload.` + "\n" +
							`The literal string "all" will result in this filter to all events.` + "\n" +
							`For guidance on using FQL, see the Segment documentation site.`,
					},
				},
				Description: `A simplified Destination filter that includes the if and actions for a DestinationFilterV1.`,
			},
			"payload": schema.MapAttribute{
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
				},
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
				Description: `The JSON payload to apply the filter to.`,
			},
		},
	}
}

func (r *PreviewDestinationFilterV1InputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PreviewDestinationFilterV1InputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PreviewDestinationFilterV1InputResourceModel
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
	res, err := r.client.DestinationFilters.PreviewDestinationFilter(ctx, request)
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

func (r *PreviewDestinationFilterV1InputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PreviewDestinationFilterV1InputResourceModel
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

func (r *PreviewDestinationFilterV1InputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PreviewDestinationFilterV1InputResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PreviewDestinationFilterV1InputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PreviewDestinationFilterV1InputResourceModel
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

func (r *PreviewDestinationFilterV1InputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource preview_destination_filter_v1_input.")
}
