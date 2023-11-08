// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteSourceV1OutputStatus - The status of the Source deletion operation.
type DeleteSourceV1OutputStatus string

const (
	DeleteSourceV1OutputStatusSuccess DeleteSourceV1OutputStatus = "SUCCESS"
)

func (e DeleteSourceV1OutputStatus) ToPointer() *DeleteSourceV1OutputStatus {
	return &e
}

func (e *DeleteSourceV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteSourceV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteSourceV1OutputStatus: %v", v)
	}
}

// DeleteSourceV1Output - Returns the status of a Source deletion.
type DeleteSourceV1Output struct {
	// The status of the Source deletion operation.
	Status DeleteSourceV1OutputStatus `json:"status"`
}

func (o *DeleteSourceV1Output) GetStatus() DeleteSourceV1OutputStatus {
	if o == nil {
		return DeleteSourceV1OutputStatus("")
	}
	return o.Status
}
