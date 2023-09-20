// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InsertFunctionInstanceAlpha struct {
	ClassID           string                 `json:"classId"`
	CreatedAt         string                 `json:"createdAt"`
	Enabled           bool                   `json:"enabled"`
	EncryptedSettings map[string]interface{} `json:"encryptedSettings"`
	ID                string                 `json:"id"`
	IntegrationID     string                 `json:"integrationId"`
	Name              *string                `json:"name,omitempty"`
	Settings          map[string]interface{} `json:"settings"`
	UpdatedAt         string                 `json:"updatedAt"`
}