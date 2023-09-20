// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment/internal/sdk/pkg/models/shared"
)

type RestoreFunctionVersionRequest struct {
	RestoreFunctionVersionAlphaInput shared.RestoreFunctionVersionAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	FunctionID                       string                                  `pathParam:"style=simple,explode=false,name=functionId"`
}

// RestoreFunctionVersion200ApplicationVndSegmentV1alphaPlusJSON - OK
type RestoreFunctionVersion200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Restore version output.
	Data *shared.RestoreFunctionVersionAlphaOutput `json:"data,omitempty"`
}

type RestoreFunctionVersionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	RestoreFunctionVersion200ApplicationVndSegmentV1alphaPlusJSONObject *RestoreFunctionVersion200ApplicationVndSegmentV1alphaPlusJSON
}