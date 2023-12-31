// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateValidationInWarehouseV1OutputStatus - Represents the status for the current connection settings.
type CreateValidationInWarehouseV1OutputStatus string

const (
	CreateValidationInWarehouseV1OutputStatusConnected CreateValidationInWarehouseV1OutputStatus = "CONNECTED"
	CreateValidationInWarehouseV1OutputStatusFailed    CreateValidationInWarehouseV1OutputStatus = "FAILED"
)

func (e CreateValidationInWarehouseV1OutputStatus) ToPointer() *CreateValidationInWarehouseV1OutputStatus {
	return &e
}

func (e *CreateValidationInWarehouseV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONNECTED":
		fallthrough
	case "FAILED":
		*e = CreateValidationInWarehouseV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateValidationInWarehouseV1OutputStatus: %v", v)
	}
}

// CreateValidationInWarehouseV1Output - Returns the status of a Warehouse connection settings after an attempt to connect to it.
type CreateValidationInWarehouseV1Output struct {
	// Represents the status for the current connection settings.
	Status CreateValidationInWarehouseV1OutputStatus `json:"status"`
}

func (o *CreateValidationInWarehouseV1Output) GetStatus() CreateValidationInWarehouseV1OutputStatus {
	if o == nil {
		return CreateValidationInWarehouseV1OutputStatus("")
	}
	return o.Status
}
