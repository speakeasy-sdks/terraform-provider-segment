// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BatchQueryMessagingSubscriptionsForSpaceAlphaInput - Batch get request.
type BatchQueryMessagingSubscriptionsForSpaceAlphaInput struct {
	// A list of subscriptions to retrieve subscription status.
	Subscriptions []GetSubscriptionRequest `json:"subscriptions"`
}