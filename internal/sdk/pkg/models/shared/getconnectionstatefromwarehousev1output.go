// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetConnectionStateFromWarehouseV1OutputConnectionState - Represents the status for the current connection settings.
type GetConnectionStateFromWarehouseV1OutputConnectionState string

const (
	GetConnectionStateFromWarehouseV1OutputConnectionStateConnected GetConnectionStateFromWarehouseV1OutputConnectionState = "CONNECTED"
	GetConnectionStateFromWarehouseV1OutputConnectionStateFailed    GetConnectionStateFromWarehouseV1OutputConnectionState = "FAILED"
)

func (e GetConnectionStateFromWarehouseV1OutputConnectionState) ToPointer() *GetConnectionStateFromWarehouseV1OutputConnectionState {
	return &e
}

func (e *GetConnectionStateFromWarehouseV1OutputConnectionState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONNECTED":
		fallthrough
	case "FAILED":
		*e = GetConnectionStateFromWarehouseV1OutputConnectionState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionStateFromWarehouseV1OutputConnectionState: %v", v)
	}
}

// GetConnectionStateFromWarehouseV1Output - Returns the status of a Warehouse connection settings after an attempt to connect to it.
type GetConnectionStateFromWarehouseV1Output struct {
	// Represents the status for the current connection settings.
	ConnectionState GetConnectionStateFromWarehouseV1OutputConnectionState `json:"connectionState"`
}

func (o *GetConnectionStateFromWarehouseV1Output) GetConnectionState() GetConnectionStateFromWarehouseV1OutputConnectionState {
	if o == nil {
		return GetConnectionStateFromWarehouseV1OutputConnectionState("")
	}
	return o.ConnectionState
}
