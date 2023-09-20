// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DeliveryMetricsSummaryBeta - Defines the summary of delivery metrics for a Destination.
type DeliveryMetricsSummaryBeta struct {
	// The Destination metadata id.
	DestinationMetadataID string `json:"destinationMetadataId"`
	// The summary of event delivery metrics for the requested Destination.
	Metrics []MetricBeta `json:"metrics"`
	// The Source id.
	//
	// Config API note: analogous to `parent`.
	SourceID string `json:"sourceId"`
}
