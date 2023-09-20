// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetSpaceAlphaOutputSpace - Space matching the given id.
type GetSpaceAlphaOutputSpace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// GetSpaceAlphaOutput - Response for the getSpaceById endpoint.
type GetSpaceAlphaOutput struct {
	// Space matching the given id.
	Space *GetSpaceAlphaOutputSpace `json:"space"`
}