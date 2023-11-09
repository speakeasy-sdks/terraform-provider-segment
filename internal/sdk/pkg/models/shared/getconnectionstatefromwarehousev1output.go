// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ConnectionState - Represents the status for the current connection settings.
type ConnectionState string

const (
	ConnectionStateConnected ConnectionState = "CONNECTED"
	ConnectionStateFailed    ConnectionState = "FAILED"
)

func (e ConnectionState) ToPointer() *ConnectionState {
	return &e
}

func (e *ConnectionState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONNECTED":
		fallthrough
	case "FAILED":
		*e = ConnectionState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionState: %v", v)
	}
}

// GetConnectionStateFromWarehouseV1Output - Returns the status of a Warehouse connection settings after an attempt to connect to it.
type GetConnectionStateFromWarehouseV1Output struct {
	// Represents the status for the current connection settings.
	ConnectionState ConnectionState `json:"connectionState"`
}

func (o *GetConnectionStateFromWarehouseV1Output) GetConnectionState() ConnectionState {
	if o == nil {
		return ConnectionState("")
	}
	return o.ConnectionState
}
