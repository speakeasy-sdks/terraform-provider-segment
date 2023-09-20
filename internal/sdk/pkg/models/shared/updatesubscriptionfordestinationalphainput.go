// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateSubscriptionForDestinationAlphaInputDestinationSubscriptionUpdateInput - The input parameters for updating a Destination subscription.
type UpdateSubscriptionForDestinationAlphaInputDestinationSubscriptionUpdateInput struct {
	// Is the subscription enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The user-defined name for the subscription.
	Name *string `json:"name,omitempty"`
	// The fields used for configuring this action.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// The fql statement.
	Trigger *string `json:"trigger,omitempty"`
}

// UpdateSubscriptionForDestinationAlphaInput - The basic input parameters for updating a Destination subscription.
type UpdateSubscriptionForDestinationAlphaInput struct {
	// A set of valid Destination input params required for updating.
	Input UpdateSubscriptionForDestinationAlphaInputDestinationSubscriptionUpdateInput `json:"input"`
}
