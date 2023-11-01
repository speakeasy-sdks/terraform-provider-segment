// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetMessagingSubscriptionSuccessResponseStatus - The user subscribed, unsubscribed, or on initial status. This is absent if the phone number or email is not found.
type GetMessagingSubscriptionSuccessResponseStatus string

const (
	GetMessagingSubscriptionSuccessResponseStatusDidNotSubscribe GetMessagingSubscriptionSuccessResponseStatus = "DID_NOT_SUBSCRIBE"
	GetMessagingSubscriptionSuccessResponseStatusSubscribed      GetMessagingSubscriptionSuccessResponseStatus = "SUBSCRIBED"
	GetMessagingSubscriptionSuccessResponseStatusUnsubscribed    GetMessagingSubscriptionSuccessResponseStatus = "UNSUBSCRIBED"
)

func (e GetMessagingSubscriptionSuccessResponseStatus) ToPointer() *GetMessagingSubscriptionSuccessResponseStatus {
	return &e
}

func (e *GetMessagingSubscriptionSuccessResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DID_NOT_SUBSCRIBE":
		fallthrough
	case "SUBSCRIBED":
		fallthrough
	case "UNSUBSCRIBED":
		*e = GetMessagingSubscriptionSuccessResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetMessagingSubscriptionSuccessResponseStatus: %v", v)
	}
}

// GetMessagingSubscriptionSuccessResponseType - Type is communication medium used.
type GetMessagingSubscriptionSuccessResponseType string

const (
	GetMessagingSubscriptionSuccessResponseTypeAndroidPush GetMessagingSubscriptionSuccessResponseType = "ANDROID_PUSH"
	GetMessagingSubscriptionSuccessResponseTypeEmail       GetMessagingSubscriptionSuccessResponseType = "EMAIL"
	GetMessagingSubscriptionSuccessResponseTypeIosPush     GetMessagingSubscriptionSuccessResponseType = "IOS_PUSH"
	GetMessagingSubscriptionSuccessResponseTypeSms         GetMessagingSubscriptionSuccessResponseType = "SMS"
	GetMessagingSubscriptionSuccessResponseTypeWhatsapp    GetMessagingSubscriptionSuccessResponseType = "WHATSAPP"
)

func (e GetMessagingSubscriptionSuccessResponseType) ToPointer() *GetMessagingSubscriptionSuccessResponseType {
	return &e
}

func (e *GetMessagingSubscriptionSuccessResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ANDROID_PUSH":
		fallthrough
	case "EMAIL":
		fallthrough
	case "IOS_PUSH":
		fallthrough
	case "SMS":
		fallthrough
	case "WHATSAPP":
		*e = GetMessagingSubscriptionSuccessResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetMessagingSubscriptionSuccessResponseType: %v", v)
	}
}

type GetMessagingSubscriptionSuccessResponse struct {
	// Optional subscription groups.
	Groups []GroupSubscriptionStatusResponse `json:"groups,omitempty"`
	// Key is the phone number or email.
	Key string `json:"key"`
	// The user subscribed, unsubscribed, or on initial status. This is absent if the phone number or email is not found.
	Status *GetMessagingSubscriptionSuccessResponseStatus `json:"status,omitempty"`
	// Type is communication medium used.
	Type GetMessagingSubscriptionSuccessResponseType `json:"type"`
	// The timestamp of this subscription status's last change.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (o *GetMessagingSubscriptionSuccessResponse) GetGroups() []GroupSubscriptionStatusResponse {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *GetMessagingSubscriptionSuccessResponse) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *GetMessagingSubscriptionSuccessResponse) GetStatus() *GetMessagingSubscriptionSuccessResponseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetMessagingSubscriptionSuccessResponse) GetType() GetMessagingSubscriptionSuccessResponseType {
	if o == nil {
		return GetMessagingSubscriptionSuccessResponseType("")
	}
	return o.Type
}

func (o *GetMessagingSubscriptionSuccessResponse) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
