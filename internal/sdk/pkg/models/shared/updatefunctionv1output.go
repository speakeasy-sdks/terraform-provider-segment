// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateFunctionV1OutputResourceType - The Function type.
//
// Config API note: equal to `type`.
type UpdateFunctionV1OutputResourceType string

const (
	UpdateFunctionV1OutputResourceTypeDestination       UpdateFunctionV1OutputResourceType = "DESTINATION"
	UpdateFunctionV1OutputResourceTypeInsertDestination UpdateFunctionV1OutputResourceType = "INSERT_DESTINATION"
	UpdateFunctionV1OutputResourceTypeSource            UpdateFunctionV1OutputResourceType = "SOURCE"
)

func (e UpdateFunctionV1OutputResourceType) ToPointer() *UpdateFunctionV1OutputResourceType {
	return &e
}

func (e *UpdateFunctionV1OutputResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DESTINATION":
		fallthrough
	case "INSERT_DESTINATION":
		fallthrough
	case "SOURCE":
		*e = UpdateFunctionV1OutputResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateFunctionV1OutputResourceType: %v", v)
	}
}

// UpdateFunctionV1OutputFunction - The updated Function object.
type UpdateFunctionV1OutputFunction struct {
	// The max count of the batch for this Function.
	BatchMaxCount *float64 `json:"batchMaxCount,omitempty"`
	// The catalog id of this Function.
	CatalogID *string `json:"catalogId,omitempty"`
	// The Function code.
	Code *string `json:"code,omitempty"`
	// The time this Function was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The id of the user who created this Function.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The time of this Function's last deployment.
	DeployedAt *string `json:"deployedAt,omitempty"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
	// A display name for this Function.
	DisplayName *string `json:"displayName,omitempty"`
	// An identifier for this Function.
	ID *string `json:"id,omitempty"`
	// Whether the deployment of this Function is the latest version.
	IsLatestVersion *bool `json:"isLatestVersion,omitempty"`
	// The URL of the logo for this Function.
	LogoURL *string `json:"logoUrl,omitempty"`
	// The preview webhook URL for this Function.
	PreviewWebhookURL *string `json:"previewWebhookUrl,omitempty"`
	// The Function type.
	//
	// Config API note: equal to `type`.
	ResourceType *UpdateFunctionV1OutputResourceType `json:"resourceType,omitempty"`
	// The list of settings for this Function.
	Settings []FunctionSettingV1 `json:"settings,omitempty"`
}

func (o *UpdateFunctionV1OutputFunction) GetBatchMaxCount() *float64 {
	if o == nil {
		return nil
	}
	return o.BatchMaxCount
}

func (o *UpdateFunctionV1OutputFunction) GetCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.CatalogID
}

func (o *UpdateFunctionV1OutputFunction) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *UpdateFunctionV1OutputFunction) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UpdateFunctionV1OutputFunction) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *UpdateFunctionV1OutputFunction) GetDeployedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *UpdateFunctionV1OutputFunction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateFunctionV1OutputFunction) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *UpdateFunctionV1OutputFunction) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFunctionV1OutputFunction) GetIsLatestVersion() *bool {
	if o == nil {
		return nil
	}
	return o.IsLatestVersion
}

func (o *UpdateFunctionV1OutputFunction) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}

func (o *UpdateFunctionV1OutputFunction) GetPreviewWebhookURL() *string {
	if o == nil {
		return nil
	}
	return o.PreviewWebhookURL
}

func (o *UpdateFunctionV1OutputFunction) GetResourceType() *UpdateFunctionV1OutputResourceType {
	if o == nil {
		return nil
	}
	return o.ResourceType
}

func (o *UpdateFunctionV1OutputFunction) GetSettings() []FunctionSettingV1 {
	if o == nil {
		return nil
	}
	return o.Settings
}

// UpdateFunctionV1Output - Create a Function.
type UpdateFunctionV1Output struct {
	// The updated Function object.
	Function UpdateFunctionV1OutputFunction `json:"function"`
}

func (o *UpdateFunctionV1Output) GetFunction() UpdateFunctionV1OutputFunction {
	if o == nil {
		return UpdateFunctionV1OutputFunction{}
	}
	return o.Function
}
