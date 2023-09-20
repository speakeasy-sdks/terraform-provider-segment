// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"segment/internal/sdk"

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
	"segment/internal/validators"
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
	Enabled    types.Bool              `tfsdk:"enabled"`
	Errors     []RequestError          `tfsdk:"errors"`
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
			"enabled": schema.BoolAttribute{
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Required:    true,
				Description: `Enable to allow this Source to send data. Defaults to true.`,
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
			"metadata_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
				MarkdownDescription: `The id of the Source metadata from which this instance of the Source derives.` + "\n" +
					`` + "\n" +
					`All Source metadata is available under ` + "`" + `/catalog/sources` + "`" + `.`,
			},
			"settings": schema.MapAttribute{
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
				Description: `A key-value object that contains instance-specific settings for the Source.`,
			},
			"slug": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
	if res.RequestErrorEnvelope == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.RequestErrorEnvelope)

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
