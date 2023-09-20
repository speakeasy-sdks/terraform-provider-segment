// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateReverseEtlModelOutputReverseEtlModel - Defines a Reverse ETL Model.
type UpdateReverseEtlModelOutputReverseEtlModel struct {
	// A longer, more descriptive explanation of the Model.
	Description string `json:"description"`
	// Indicates whether the Model should have syncs enabled. When disabled, no
	// syncs will be triggered, regardless of the enabled status of the attached
	// destinations/subscriptions.
	Enabled bool `json:"enabled"`
	// The id of the Model.
	ID string `json:"id"`
	// A short, human-readable description of the Model.
	Name string `json:"name"`
	// The SQL query that will be executed to extract data from the connected
	// Source.
	Query string `json:"query"`
	// Indicates the column named in `query` that should be used to uniquely
	// identify the extracted records.
	QueryIdentifierColumn string `json:"queryIdentifierColumn"`
	// Depending on the chosen strategy, configures the schedule for this model.
	ScheduleConfig map[string]interface{} `json:"scheduleConfig"`
	// Determines the strategy used for triggering syncs, which will be used in
	// conjunction with scheduleConfig.
	//
	// Possible values: "manual", "periodic", "specific_days".
	ScheduleStrategy string `json:"scheduleStrategy"`
	// The id for the attached Source.
	SourceID string `json:"sourceId"`
}

// UpdateReverseEtlModelOutput - Defines the results of updating a Model.
type UpdateReverseEtlModelOutput struct {
	// The created Model.
	ReverseEtlModel UpdateReverseEtlModelOutputReverseEtlModel `json:"reverseEtlModel"`
}
