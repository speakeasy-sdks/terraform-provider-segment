// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateDestinationSubscriptionAlphaOutputDestinationSubscription - The Destination subscription.
type CreateDestinationSubscriptionAlphaOutputDestinationSubscription struct {
	// The unique identifier for the Destination action to trigger.
	ActionID string `json:"actionId"`
	// The URL-friendly key for the associated Destination action.
	ActionSlug string `json:"actionSlug"`
	// The associated Destination instance id.
	DestinationID string `json:"destinationId"`
	// Is the subscription enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for the subscription.
	ID string `json:"id"`
	// The unique identifier for the linked ReverseETLModel, if this part of a
	// Reverse ETL connection.
	ModelID *string `json:"modelId,omitempty"`
	// The name of the subscription.
	Name string `json:"name"`
	// The customer settings for action fields.
	Settings map[string]interface{} `json:"settings"`
	// FQL string that describes what events should trigger a Destination action.
	Trigger string `json:"trigger"`
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetActionID() string {
	if o == nil {
		return ""
	}
	return o.ActionID
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetActionSlug() string {
	if o == nil {
		return ""
	}
	return o.ActionSlug
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetModelID() *string {
	if o == nil {
		return nil
	}
	return o.ModelID
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}

func (o *CreateDestinationSubscriptionAlphaOutputDestinationSubscription) GetTrigger() string {
	if o == nil {
		return ""
	}
	return o.Trigger
}

// CreateDestinationSubscriptionAlphaOutput - Returns a newly created Destination subscription.
type CreateDestinationSubscriptionAlphaOutput struct {
	// The Destination subscription.
	DestinationSubscription CreateDestinationSubscriptionAlphaOutputDestinationSubscription `json:"destinationSubscription"`
}

func (o *CreateDestinationSubscriptionAlphaOutput) GetDestinationSubscription() CreateDestinationSubscriptionAlphaOutputDestinationSubscription {
	if o == nil {
		return CreateDestinationSubscriptionAlphaOutputDestinationSubscription{}
	}
	return o.DestinationSubscription
}
