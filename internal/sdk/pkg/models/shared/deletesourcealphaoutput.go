// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteSourceAlphaOutputStatus - The status of the Source deletion operation.
type DeleteSourceAlphaOutputStatus string

const (
	DeleteSourceAlphaOutputStatusSuccess DeleteSourceAlphaOutputStatus = "SUCCESS"
)

func (e DeleteSourceAlphaOutputStatus) ToPointer() *DeleteSourceAlphaOutputStatus {
	return &e
}

func (e *DeleteSourceAlphaOutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteSourceAlphaOutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteSourceAlphaOutputStatus: %v", v)
	}
}

// DeleteSourceAlphaOutput - Returns the status of a Source deletion.
type DeleteSourceAlphaOutput struct {
	// The status of the Source deletion operation.
	Status DeleteSourceAlphaOutputStatus `json:"status"`
}

func (o *DeleteSourceAlphaOutput) GetStatus() DeleteSourceAlphaOutputStatus {
	if o == nil {
		return DeleteSourceAlphaOutputStatus("")
	}
	return o.Status
}
