// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteWarehouseV1OutputStatus - The status of the Warehouse deletion operation.
type DeleteWarehouseV1OutputStatus string

const (
	DeleteWarehouseV1OutputStatusSuccess DeleteWarehouseV1OutputStatus = "SUCCESS"
)

func (e DeleteWarehouseV1OutputStatus) ToPointer() *DeleteWarehouseV1OutputStatus {
	return &e
}

func (e *DeleteWarehouseV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteWarehouseV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteWarehouseV1OutputStatus: %v", v)
	}
}

// DeleteWarehouseV1Output - Returns the status of a Warehouse deletion.
type DeleteWarehouseV1Output struct {
	// The status of the Warehouse deletion operation.
	Status DeleteWarehouseV1OutputStatus `json:"status"`
}
