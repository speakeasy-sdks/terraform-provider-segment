// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RemoveProfilesWarehouseFromSpaceAlphaOutputStatus - The status of the Warehouse deletion operation.
type RemoveProfilesWarehouseFromSpaceAlphaOutputStatus string

const (
	RemoveProfilesWarehouseFromSpaceAlphaOutputStatusSuccess RemoveProfilesWarehouseFromSpaceAlphaOutputStatus = "SUCCESS"
)

func (e RemoveProfilesWarehouseFromSpaceAlphaOutputStatus) ToPointer() *RemoveProfilesWarehouseFromSpaceAlphaOutputStatus {
	return &e
}

func (e *RemoveProfilesWarehouseFromSpaceAlphaOutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = RemoveProfilesWarehouseFromSpaceAlphaOutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RemoveProfilesWarehouseFromSpaceAlphaOutputStatus: %v", v)
	}
}

// RemoveProfilesWarehouseFromSpaceAlphaOutput - Returns the status of a Warehouse deletion.
type RemoveProfilesWarehouseFromSpaceAlphaOutput struct {
	// The status of the Warehouse deletion operation.
	Status RemoveProfilesWarehouseFromSpaceAlphaOutputStatus `json:"status"`
}

func (o *RemoveProfilesWarehouseFromSpaceAlphaOutput) GetStatus() RemoveProfilesWarehouseFromSpaceAlphaOutputStatus {
	if o == nil {
		return RemoveProfilesWarehouseFromSpaceAlphaOutputStatus("")
	}
	return o.Status
}
