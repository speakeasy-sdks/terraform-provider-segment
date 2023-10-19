// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ListFunctionItemV1ResourceType - The Function type.
//
// Config API note: equal to `type`.
type ListFunctionItemV1ResourceType string

const (
	ListFunctionItemV1ResourceTypeDestination       ListFunctionItemV1ResourceType = "DESTINATION"
	ListFunctionItemV1ResourceTypeInsertDestination ListFunctionItemV1ResourceType = "INSERT_DESTINATION"
	ListFunctionItemV1ResourceTypeSource            ListFunctionItemV1ResourceType = "SOURCE"
)

func (e ListFunctionItemV1ResourceType) ToPointer() *ListFunctionItemV1ResourceType {
	return &e
}

func (e *ListFunctionItemV1ResourceType) UnmarshalJSON(data []byte) error {
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
		*e = ListFunctionItemV1ResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListFunctionItemV1ResourceType: %v", v)
	}
}

// ListFunctionItemV1 - Represents a Function in a list.
type ListFunctionItemV1 struct {
	// The catalog id of this Function.
	CatalogID *string `json:"catalogId,omitempty"`
	// The time this Function was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The id of the user who created this Function.
	CreatedBy *string `json:"createdBy,omitempty"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
	// A display name for this Function.
	DisplayName *string `json:"displayName,omitempty"`
	// An identifier for this Function.
	ID *string `json:"id,omitempty"`
	// The URL of the logo for this Function.
	LogoURL *string `json:"logoUrl,omitempty"`
	// The Function type.
	//
	// Config API note: equal to `type`.
	ResourceType *ListFunctionItemV1ResourceType `json:"resourceType,omitempty"`
}

func (o *ListFunctionItemV1) GetCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.CatalogID
}

func (o *ListFunctionItemV1) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ListFunctionItemV1) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *ListFunctionItemV1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ListFunctionItemV1) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ListFunctionItemV1) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListFunctionItemV1) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}

func (o *ListFunctionItemV1) GetResourceType() *ListFunctionItemV1ResourceType {
	if o == nil {
		return nil
	}
	return o.ResourceType
}
