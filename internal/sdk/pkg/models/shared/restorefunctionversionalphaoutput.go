// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RestoreFunctionVersionAlphaOutputVersion - Represents a Function Version in a list.
type RestoreFunctionVersionAlphaOutputVersion struct {
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

func (o *RestoreFunctionVersionAlphaOutputVersion) GetAuthor() *string {
	if o == nil {
		return nil
	}
	return o.Author
}

func (o *RestoreFunctionVersionAlphaOutputVersion) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *RestoreFunctionVersionAlphaOutputVersion) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RestoreFunctionVersionAlphaOutputVersion) GetDeployedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *RestoreFunctionVersionAlphaOutputVersion) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RestoreFunctionVersionAlphaOutputVersion) GetIsDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeployed
}

func (o *RestoreFunctionVersionAlphaOutputVersion) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// RestoreFunctionVersionAlphaOutput - Restore version output.
type RestoreFunctionVersionAlphaOutput struct {
	// Restored version.
	Version RestoreFunctionVersionAlphaOutputVersion `json:"version"`
}

func (o *RestoreFunctionVersionAlphaOutput) GetVersion() RestoreFunctionVersionAlphaOutputVersion {
	if o == nil {
		return RestoreFunctionVersionAlphaOutputVersion{}
	}
	return o.Version
}
