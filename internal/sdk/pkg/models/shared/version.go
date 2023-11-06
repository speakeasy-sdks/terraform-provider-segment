// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Version - Represents a Function Version in a list.
type Version struct {
	// Author of this version.
	Author *string `json:"author,omitempty"`
	// Source code of this version.
	Code string `json:"code"`
	// The time of this Version's creation.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time of this Version's last deployment.
	DeployedAt *string `json:"deployedAt,omitempty"`
	// An identifier for this version.
	ID string `json:"id"`
	// The deployed status of this version.
	IsDeployed *bool `json:"isDeployed,omitempty"`
	// The time of this Version's latest update.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (o *Version) GetAuthor() *string {
	if o == nil {
		return nil
	}
	return o.Author
}

func (o *Version) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *Version) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Version) GetDeployedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *Version) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Version) GetIsDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeployed
}

func (o *Version) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
