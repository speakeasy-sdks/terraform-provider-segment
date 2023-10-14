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
)

// selectiveSync - Warehouse Selective Sync allows you to manage the data that you send to your Warehouses. You can use this feature to stop syncing specific events (also known as collections) or properties that aren’t relevant, and may slow down your Warehouse syncs.
type selectiveSync struct {
	sdkConfiguration sdkConfiguration
}

func newSelectiveSync(sdkConfig sdkConfiguration) *selectiveSync {
	return &selectiveSync{
		sdkConfiguration: sdkConfig,
	}
}

// GetAdvancedSyncScheduleFromWarehouse - Get Advanced Sync Schedule from Warehouse
// Returns the advanced sync schedule for a Warehouse.
//
// The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *selectiveSync) GetAdvancedSyncScheduleFromWarehouse(ctx context.Context, request operations.GetAdvancedSyncScheduleFromWarehouseRequest) (*operations.GetAdvancedSyncScheduleFromWarehouseResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/warehouses/{warehouseId}/advanced-sync-schedule", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.GetAdvancedSyncScheduleFromWarehouseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetAdvancedSyncScheduleFromWarehouse200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetAdvancedSyncScheduleFromWarehouse200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1betaPlusJSONObject = out
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

// ListSelectiveSyncsFromWarehouseAndSource - List Selective Syncs from Warehouse And Source
// Returns the schema for a Warehouse, including Sources, Collections, and Properties.
//
// The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *selectiveSync) ListSelectiveSyncsFromWarehouseAndSource(ctx context.Context, request operations.ListSelectiveSyncsFromWarehouseAndSourceRequest) (*operations.ListSelectiveSyncsFromWarehouseAndSourceResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/warehouses/{warehouseId}/connected-sources/{sourceId}/selective-syncs", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.ListSelectiveSyncsFromWarehouseAndSourceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSelectiveSyncsFromWarehouseAndSource200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSelectiveSyncsFromWarehouseAndSource200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ListSelectiveSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSelectiveSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ListSelectiveSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSelectiveSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ListSelectiveSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSelectiveSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSONObject = out
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

// ListSyncsFromWarehouse - List Syncs from Warehouse
// Returns the overview of syncs for every Source connected to a Warehouse.
//
// The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *selectiveSync) ListSyncsFromWarehouse(ctx context.Context, request operations.ListSyncsFromWarehouseRequest) (*operations.ListSyncsFromWarehouseResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/warehouses/{warehouseId}/syncs", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.ListSyncsFromWarehouseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSyncsFromWarehouse200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouse200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ListSyncsFromWarehouse200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouse200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ListSyncsFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ListSyncsFromWarehouse200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouse200ApplicationVndSegmentV1betaPlusJSONObject = out
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

// ListSyncsFromWarehouseAndSource - List Syncs from Warehouse And Source
// Returns the overview of syncs for a Source connected to a Warehouse.
//
// The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *selectiveSync) ListSyncsFromWarehouseAndSource(ctx context.Context, request operations.ListSyncsFromWarehouseAndSourceRequest) (*operations.ListSyncsFromWarehouseAndSourceResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/warehouses/{warehouseId}/connected-sources/{sourceId}/syncs", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.ListSyncsFromWarehouseAndSourceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSyncsFromWarehouseAndSource200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouseAndSource200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSONObject = out
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

// ReplaceAdvancedSyncScheduleForWarehouse - Replace Advanced Sync Schedule for Warehouse
// Updates the advanced sync schedule for a Warehouse, replacing the sync schedule with a new schedule.
//
// The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *selectiveSync) ReplaceAdvancedSyncScheduleForWarehouse(ctx context.Context, request operations.ReplaceAdvancedSyncScheduleForWarehouseRequest) (*operations.ReplaceAdvancedSyncScheduleForWarehouseResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/warehouses/{warehouseId}/advanced-sync-schedule", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ReplaceAdvancedSyncScheduleForWarehouseV1Input", "json")
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
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.ReplaceAdvancedSyncScheduleForWarehouseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject = out
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

// UpdateSelectiveSyncForWarehouse - Update Selective Sync for Warehouse
// Configures the schema for a Warehouse, including Sources, Collections, and Properties.
//
// • When called, this endpoint may generate the `Storage Destination Modified` event in the [audit trail](/tag/Audit-Trail).
//
// The rate limit for this endpoint is 2 requests per minute, which is lower than the default due to access pattern restrictions. Once reached, this endpoint will respond with the 429 HTTP status code with headers indicating the limit parameters. See [Rate Limiting](/#tag/Rate-Limits) for more information.
func (s *selectiveSync) UpdateSelectiveSyncForWarehouse(ctx context.Context, request operations.UpdateSelectiveSyncForWarehouseRequest) (*operations.UpdateSelectiveSyncForWarehouseResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/warehouses/{warehouseId}/selective-sync", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "UpdateSelectiveSyncForWarehouseV1Input", "json")
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
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.UpdateSelectiveSyncForWarehouseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.UpdateSelectiveSyncForWarehouse200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateSelectiveSyncForWarehouse200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject = out
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
