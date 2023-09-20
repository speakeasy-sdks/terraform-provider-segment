// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateSourceV1Input - Create a new Source based on a set of parameters.
type CreateSourceV1Input struct {
	// Enable to allow this Source to send data. Defaults to true.
	Enabled bool `json:"enabled"`
	// The id of the Source metadata from which this instance of the Source derives.
	//
	// All Source metadata is available under `/catalog/sources`.
	MetadataID string `json:"metadataId"`
	// A key-value object that contains instance-specific settings for the Source.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// The slug by which to identify the Source in the Segment app.
	Slug string `json:"slug"`
}
