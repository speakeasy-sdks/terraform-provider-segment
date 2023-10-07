// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type DeleteTransformationRequest struct {
	TransformationID string `pathParam:"style=simple,explode=false,name=transformationId"`
}

// DeleteTransformation200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteTransformation200ApplicationVndSegmentV1betaPlusJSON struct {
	// The output of delete Transformation.
	Data *shared.DeleteTransformationV1Output `json:"data,omitempty"`
}

// DeleteTransformation200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteTransformation200ApplicationVndSegmentV1alphaPlusJSON struct {
	// The output of delete Transformation.
	Data *shared.DeleteTransformationV1Output `json:"data,omitempty"`
}

// DeleteTransformation200ApplicationVndSegmentV1PlusJSON - OK
type DeleteTransformation200ApplicationVndSegmentV1PlusJSON struct {
	// The output of delete Transformation.
	Data *shared.DeleteTransformationV1Output `json:"data,omitempty"`
}

// DeleteTransformation200ApplicationJSON - OK
type DeleteTransformation200ApplicationJSON struct {
	// The output of delete Transformation.
	Data *shared.DeleteTransformationV1Output `json:"data,omitempty"`
}

type DeleteTransformationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteTransformation200ApplicationJSONObject *DeleteTransformation200ApplicationJSON
	// OK
	DeleteTransformation200ApplicationVndSegmentV1PlusJSONObject *DeleteTransformation200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteTransformation200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteTransformation200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteTransformation200ApplicationVndSegmentV1betaPlusJSONObject *DeleteTransformation200ApplicationVndSegmentV1betaPlusJSON
}
