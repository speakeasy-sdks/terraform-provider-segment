// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PropertyValueTransformationBeta struct {
	// The property paths. The maximum number of paths is 10.
	PropertyPaths []string `json:"propertyPaths"`
	// The new value of the property paths.
	PropertyValue string `json:"propertyValue"`
}
