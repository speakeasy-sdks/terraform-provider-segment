// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"segment/internal/sdk/pkg/models/operations"
	"segment/internal/sdk/pkg/models/shared"
	"segment/internal/sdk/pkg/utils"
	"strings"
)

// trackingPlans - A Tracking Plan is a data spec outlining the events and properties you intend to collect across your Segment Sources.
//
// The Segment Tracking Plan feature allows you to validate your expected events against the live events that Segment receives. Segment generates violations when an event doesn’t match the spec’d event in the Tracking Plan.
//
// You can store your Tracking Plans in your Workspace, and connect them to one or more Sources.
//
// Using the Segment Public API, you can create, delete, update, list, and connect Tracking Plans and their Rules.
//
// ## Migrate from the Config API
//
// The [getTrackingPlan endpoint](https://reference.segmentapis.com/#1092fe01-379b-4ca1-8b1d-9f42b33d2899) returns the following fields:
//
// | Config API     | Public API  |
// | -------------- | ----------- |
// | `display_name` | `name`      |
// | `name`         | `slug`      |
// | `updateTime`   | `updatedAt` |
// | `createTime`   | `createdAt` |
//
// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
type trackingPlans struct {
	sdkConfiguration sdkConfiguration
}

func newTrackingPlans(sdkConfig sdkConfiguration) *trackingPlans {
	return &trackingPlans{
		sdkConfiguration: sdkConfig,
	}
}

// AddSourceToTrackingPlan - Add Source to Tracking Plan
// Connects a Source to a Tracking Plan.
//
// • When called, this endpoint may generate the `Source Modified` event in the [audit trail](/tag/Audit-Trail).
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
func (s *trackingPlans) AddSourceToTrackingPlan(ctx context.Context, request operations.AddSourceToTrackingPlanRequest) (*operations.AddSourceToTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}/sources", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "AddSourceToTrackingPlanV1Input", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.AddSourceToTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AddSourceToTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.AddSourceToTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.AddSourceToTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.AddSourceToTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.AddSourceToTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.AddSourceToTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.AddSourceToTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.AddSourceToTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// CreateTrackingPlan - Create Tracking Plan
// Creates a Tracking Plan.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
func (s *trackingPlans) CreateTrackingPlan(ctx context.Context, request shared.CreateTrackingPlanV1Input) (*operations.CreateTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/tracking-plans"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.CreateTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.CreateTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.CreateTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.CreateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.CreateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.CreateTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.CreateTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// DeleteTrackingPlan - Delete Tracking Plan
// Deletes a Tracking Plan.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
func (s *trackingPlans) DeleteTrackingPlan(ctx context.Context, request operations.DeleteTrackingPlanRequest) (*operations.DeleteTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.DeleteTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// GetTrackingPlan - Get Tracking Plan
// Returns a Tracking Plan.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
func (s *trackingPlans) GetTrackingPlan(ctx context.Context, request operations.GetTrackingPlanRequest) (*operations.GetTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.GetTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// ListRulesFromTrackingPlan - List Rules from Tracking Plan
// Lists Tracking Plan rules.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
//
// The rate limit for this endpoint is 20 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *trackingPlans) ListRulesFromTrackingPlan(ctx context.Context, request operations.ListRulesFromTrackingPlanRequest) (*operations.ListRulesFromTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}/rules", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListRulesFromTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListRulesFromTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRulesFromTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ListRulesFromTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRulesFromTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ListRulesFromTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRulesFromTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ListRulesFromTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRulesFromTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// ListSourcesFromTrackingPlan - List Sources from Tracking Plan
// Lists Sources connected to a Tracking Plan.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
//
// This endpoint requires the user to have at least the following permission(s):
//   - Source Read-only
//   - Tracking Plan Read-only
func (s *trackingPlans) ListSourcesFromTrackingPlan(ctx context.Context, request operations.ListSourcesFromTrackingPlanRequest) (*operations.ListSourcesFromTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}/sources", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSourcesFromTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSourcesFromTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSourcesFromTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ListSourcesFromTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSourcesFromTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ListSourcesFromTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSourcesFromTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ListSourcesFromTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSourcesFromTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// ReplaceRulesInTrackingPlan - Replace Rules in Tracking Plan
// Replaces Tracking Plan rules.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
func (s *trackingPlans) ReplaceRulesInTrackingPlan(ctx context.Context, request operations.ReplaceRulesInTrackingPlanRequest) (*operations.ReplaceRulesInTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}/rules", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ReplaceRulesInTrackingPlanV1Input", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ReplaceRulesInTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ReplaceRulesInTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceRulesInTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// UpdateRulesInTrackingPlan - Update Rules in Tracking Plan
// Updates Tracking Plan rules.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
func (s *trackingPlans) UpdateRulesInTrackingPlan(ctx context.Context, request operations.UpdateRulesInTrackingPlanRequest) (*operations.UpdateRulesInTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}/rules", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "UpdateRulesInTrackingPlanV1Input", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateRulesInTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateRulesInTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateRulesInTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// UpdateTrackingPlan - Update Tracking Plan
// Updates a Tracking Plan.
//
// • In order to successfully call this endpoint, the specified Workspace needs to have the Protocols feature enabled. Please reach out to your customer success manager for more information.
//
// Config API omitted fields:
// - `updateMask`
func (s *trackingPlans) UpdateTrackingPlan(ctx context.Context, request operations.UpdateTrackingPlanRequest) (*operations.UpdateTrackingPlanResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tracking-plans/{trackingPlanId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "UpdateTrackingPlanV1Input", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateTrackingPlanResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateTrackingPlan200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateTrackingPlan200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.UpdateTrackingPlan200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateTrackingPlan200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.UpdateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.UpdateTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}
