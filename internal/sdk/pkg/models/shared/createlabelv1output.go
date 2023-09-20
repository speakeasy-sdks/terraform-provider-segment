// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateLabelV1OutputLabelV1 - A label lets Workspace owners assign permissions to users, and grant these users access to groups.
//
// A Workspace owner may use labels to grant users access to groups of resources.
//
// When you add a label to a Source or Personas Spaces, any users granted access to that label gain access to those
// resources.
//
// All Workspaces include labels for Dev (development) and Prod (production) environments. On top of those, Free and
// Team plan customers may create up to five labels.
//
// Customers with the Enterprise pricing package may create an unlimited number of labels.
type CreateLabelV1OutputLabelV1 struct {
	// An optional description of the purpose of this label.
	Description *string `json:"description,omitempty"`
	// The key that represents the name of this label.
	Key string `json:"key"`
	// The value associated with the key of this label.
	Value string `json:"value"`
}

// CreateLabelV1Output - Result of creating a new label in the current Workspace.
type CreateLabelV1Output struct {
	// The newly created label.
	Label CreateLabelV1OutputLabelV1 `json:"label"`
}