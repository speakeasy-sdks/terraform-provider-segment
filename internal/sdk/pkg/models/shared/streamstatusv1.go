// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// StreamStatusV1 - StreamStatus represents status of each stream including all the Destinations corresponding to the stream.
type StreamStatusV1 struct {
	DestinationStatus []DestinationStatusV1 `json:"destinationStatus"`
	ID                string                `json:"id"`
}