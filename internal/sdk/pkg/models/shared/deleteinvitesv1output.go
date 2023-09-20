// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteInvitesV1OutputStatus - The status of the invite deletion operation.
type DeleteInvitesV1OutputStatus string

const (
	DeleteInvitesV1OutputStatusSuccess DeleteInvitesV1OutputStatus = "SUCCESS"
)

func (e DeleteInvitesV1OutputStatus) ToPointer() *DeleteInvitesV1OutputStatus {
	return &e
}

func (e *DeleteInvitesV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteInvitesV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteInvitesV1OutputStatus: %v", v)
	}
}

// DeleteInvitesV1Output - Returns the status of the removal operation.
type DeleteInvitesV1Output struct {
	// The status of the invite deletion operation.
	Status DeleteInvitesV1OutputStatus `json:"status"`
}
