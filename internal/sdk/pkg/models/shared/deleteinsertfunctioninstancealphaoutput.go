// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteInsertFunctionInstanceAlphaOutputStatus - The status of the operation.
type DeleteInsertFunctionInstanceAlphaOutputStatus string

const (
	DeleteInsertFunctionInstanceAlphaOutputStatusSuccess DeleteInsertFunctionInstanceAlphaOutputStatus = "SUCCESS"
)

func (e DeleteInsertFunctionInstanceAlphaOutputStatus) ToPointer() *DeleteInsertFunctionInstanceAlphaOutputStatus {
	return &e
}

func (e *DeleteInsertFunctionInstanceAlphaOutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteInsertFunctionInstanceAlphaOutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteInsertFunctionInstanceAlphaOutputStatus: %v", v)
	}
}

// DeleteInsertFunctionInstanceAlphaOutput - Delete an insert Function instance.
type DeleteInsertFunctionInstanceAlphaOutput struct {
	// The status of the operation.
	Status DeleteInsertFunctionInstanceAlphaOutputStatus `json:"status"`
}

func (o *DeleteInsertFunctionInstanceAlphaOutput) GetStatus() DeleteInsertFunctionInstanceAlphaOutputStatus {
	if o == nil {
		return DeleteInsertFunctionInstanceAlphaOutputStatus("")
	}
	return o.Status
}
