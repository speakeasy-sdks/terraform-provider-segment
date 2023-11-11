// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/operations"
	"github.com/scentregroup/terraform-provider-segment/internal/sdk/pkg/models/shared"
)

func (r *CreateTransformationV1InputResourceModel) ToCreateSDKType() *shared.CreateTransformationV1Input {
	destinationMetadataID := new(string)
	if !r.DestinationMetadataID.IsUnknown() && !r.DestinationMetadataID.IsNull() {
		*destinationMetadataID = r.DestinationMetadataID.ValueString()
	} else {
		destinationMetadataID = nil
	}
	enabled := r.Enabled.ValueBool()
	var fqlDefinedProperties []shared.FQLDefinedPropertyV1 = nil
	for _, fqlDefinedPropertiesItem := range r.FqlDefinedProperties {
		fql := fqlDefinedPropertiesItem.Fql.ValueString()
		propertyName := fqlDefinedPropertiesItem.PropertyName.ValueString()
		fqlDefinedProperties = append(fqlDefinedProperties, shared.FQLDefinedPropertyV1{
			Fql:          fql,
			PropertyName: propertyName,
		})
	}
	ifVar := r.If.ValueString()
	name := r.Name.ValueString()
	newEventName := new(string)
	if !r.NewEventName.IsUnknown() && !r.NewEventName.IsNull() {
		*newEventName = r.NewEventName.ValueString()
	} else {
		newEventName = nil
	}
	var propertyRenames []shared.PropertyRenameV1 = nil
	for _, propertyRenamesItem := range r.PropertyRenames {
		newName := propertyRenamesItem.NewName.ValueString()
		oldName := propertyRenamesItem.OldName.ValueString()
		propertyRenames = append(propertyRenames, shared.PropertyRenameV1{
			NewName: newName,
			OldName: oldName,
		})
	}
	var propertyValueTransformations []shared.PropertyValueTransformationV1 = nil
	for _, propertyValueTransformationsItem := range r.PropertyValueTransformations {
		var propertyPaths []string = nil
		for _, propertyPathsItem := range propertyValueTransformationsItem.PropertyPaths {
			propertyPaths = append(propertyPaths, propertyPathsItem.ValueString())
		}
		propertyValue := propertyValueTransformationsItem.PropertyValue.ValueString()
		propertyValueTransformations = append(propertyValueTransformations, shared.PropertyValueTransformationV1{
			PropertyPaths: propertyPaths,
			PropertyValue: propertyValue,
		})
	}
	sourceID := r.SourceID.ValueString()
	out := shared.CreateTransformationV1Input{
		DestinationMetadataID:        destinationMetadataID,
		Enabled:                      enabled,
		FqlDefinedProperties:         fqlDefinedProperties,
		If:                           ifVar,
		Name:                         name,
		NewEventName:                 newEventName,
		PropertyRenames:              propertyRenames,
		PropertyValueTransformations: propertyValueTransformations,
		SourceID:                     sourceID,
	}
	return &out
}

func (r *CreateTransformationV1InputResourceModel) RefreshFromCreateResponse(resp *operations.CreateTransformationResponseBody) {
	if resp.Data == nil {
		r.Data = nil
	} else {
		r.Data = &CreateTransformationV1Output{}
		if resp.Data.Transformation.DestinationMetadataID != nil {
			r.Data.Transformation.DestinationMetadataID = types.StringValue(*resp.Data.Transformation.DestinationMetadataID)
		} else {
			r.Data.Transformation.DestinationMetadataID = types.StringNull()
		}
		r.Data.Transformation.Enabled = types.BoolValue(resp.Data.Transformation.Enabled)
		if len(r.Data.Transformation.FqlDefinedProperties) > len(resp.Data.Transformation.FqlDefinedProperties) {
			r.Data.Transformation.FqlDefinedProperties = r.Data.Transformation.FqlDefinedProperties[:len(resp.Data.Transformation.FqlDefinedProperties)]
		}
		for fqlDefinedPropertiesCount, fqlDefinedPropertiesItem := range resp.Data.Transformation.FqlDefinedProperties {
			var fqlDefinedProperties1 FQLDefinedPropertyV1
			fqlDefinedProperties1.Fql = types.StringValue(fqlDefinedPropertiesItem.Fql)
			fqlDefinedProperties1.PropertyName = types.StringValue(fqlDefinedPropertiesItem.PropertyName)
			if fqlDefinedPropertiesCount+1 > len(r.Data.Transformation.FqlDefinedProperties) {
				r.Data.Transformation.FqlDefinedProperties = append(r.Data.Transformation.FqlDefinedProperties, fqlDefinedProperties1)
			} else {
				r.Data.Transformation.FqlDefinedProperties[fqlDefinedPropertiesCount].Fql = fqlDefinedProperties1.Fql
				r.Data.Transformation.FqlDefinedProperties[fqlDefinedPropertiesCount].PropertyName = fqlDefinedProperties1.PropertyName
			}
		}
		r.Data.Transformation.ID = types.StringValue(resp.Data.Transformation.ID)
		r.Data.Transformation.If = types.StringValue(resp.Data.Transformation.If)
		r.Data.Transformation.Name = types.StringValue(resp.Data.Transformation.Name)
		if resp.Data.Transformation.NewEventName != nil {
			r.Data.Transformation.NewEventName = types.StringValue(*resp.Data.Transformation.NewEventName)
		} else {
			r.Data.Transformation.NewEventName = types.StringNull()
		}
		if len(r.Data.Transformation.PropertyRenames) > len(resp.Data.Transformation.PropertyRenames) {
			r.Data.Transformation.PropertyRenames = r.Data.Transformation.PropertyRenames[:len(resp.Data.Transformation.PropertyRenames)]
		}
		for propertyRenamesCount, propertyRenamesItem := range resp.Data.Transformation.PropertyRenames {
			var propertyRenames1 PropertyRenameV1
			propertyRenames1.NewName = types.StringValue(propertyRenamesItem.NewName)
			propertyRenames1.OldName = types.StringValue(propertyRenamesItem.OldName)
			if propertyRenamesCount+1 > len(r.Data.Transformation.PropertyRenames) {
				r.Data.Transformation.PropertyRenames = append(r.Data.Transformation.PropertyRenames, propertyRenames1)
			} else {
				r.Data.Transformation.PropertyRenames[propertyRenamesCount].NewName = propertyRenames1.NewName
				r.Data.Transformation.PropertyRenames[propertyRenamesCount].OldName = propertyRenames1.OldName
			}
		}
		if len(r.Data.Transformation.PropertyValueTransformations) > len(resp.Data.Transformation.PropertyValueTransformations) {
			r.Data.Transformation.PropertyValueTransformations = r.Data.Transformation.PropertyValueTransformations[:len(resp.Data.Transformation.PropertyValueTransformations)]
		}
		for propertyValueTransformationsCount, propertyValueTransformationsItem := range resp.Data.Transformation.PropertyValueTransformations {
			var propertyValueTransformations1 PropertyValueTransformationV1
			propertyValueTransformations1.PropertyPaths = nil
			for _, v := range propertyValueTransformationsItem.PropertyPaths {
				propertyValueTransformations1.PropertyPaths = append(propertyValueTransformations1.PropertyPaths, types.StringValue(v))
			}
			propertyValueTransformations1.PropertyValue = types.StringValue(propertyValueTransformationsItem.PropertyValue)
			if propertyValueTransformationsCount+1 > len(r.Data.Transformation.PropertyValueTransformations) {
				r.Data.Transformation.PropertyValueTransformations = append(r.Data.Transformation.PropertyValueTransformations, propertyValueTransformations1)
			} else {
				r.Data.Transformation.PropertyValueTransformations[propertyValueTransformationsCount].PropertyPaths = propertyValueTransformations1.PropertyPaths
				r.Data.Transformation.PropertyValueTransformations[propertyValueTransformationsCount].PropertyValue = propertyValueTransformations1.PropertyValue
			}
		}
		r.Data.Transformation.SourceID = types.StringValue(resp.Data.Transformation.SourceID)
	}
}
