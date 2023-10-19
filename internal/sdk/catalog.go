// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"segment/internal/sdk/pkg/models/operations"
	"segment/internal/sdk/pkg/models/sdkerrors"
	"segment/internal/sdk/pkg/models/shared"
	"segment/internal/sdk/pkg/utils"
	"strings"
)

// catalog - These endpoints let you list all the Sources, Destinations, and Destination settings that Segment supports.
//
// ## Migrate from the Config API
//
// Like the Segment Public API, the Config API has endpoints to navigate the catalog of Integrations supported by Segment. See the sections below for some key differences:
//
// ### Source catalog items
//
// | Config API  | Public API                                      |
// | ----------- | ----------------------------------------------- |
// | name        | See note on names vs IDs in the migration guide |
// | displayName | `name`                                          |
// | description | `description`                                   |
// | categories  | `categories`                                    |
// | logos       | `logos`                                         |
// | type        | Removed                                         |
//
// ### Destination catalog items
//
// | Config API                 | Public API                                      |
// | -------------------------- | ----------------------------------------------- |
// | browserUnbundlingPublic    | `supportedFeatures.browserUnbundlingPublic`     |
// | browserUnbundlingSupported | `supportedFeatures.browserUnbundling`           |
// | categories                 | `categories`                                    |
// | components                 | `components`                                    |
// | description                | `description`                                   |
// | displayName                | `name`,                                         |
// | id                         | `id`                                            |
// | logos                      | `logos`                                         |
// | methods                    | `supportedMethods`                              |
// | name                       | See note on names vs IDs in the migration guide |
// | platforms                  | `supportedPlatforms`                            |
// | previousNames              | `previousNames`                                 |
// | settings                   | `options`                                       |
// | slug                       | `slug`                                          |
// | status                     | `status`                                        |
// | type                       | Not returned                                    |
// | website                    | `website`                                       |
//
// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
type catalog struct {
	sdkConfiguration sdkConfiguration
}

func newCatalog(sdkConfig sdkConfiguration) *catalog {
	return &catalog{
		sdkConfiguration: sdkConfig,
	}
}

// GetDestinationMetadata - Get Destination Metadata
// Returns a Destination catalog item by its id.
func (s *catalog) GetDestinationMetadata(ctx context.Context, request operations.GetDestinationMetadataRequest, opts ...operations.Option) (*operations.GetDestinationMetadataResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/catalog/destinations/{destinationMetadataId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

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

	res := &operations.GetDestinationMetadataResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetDestinationMetadata200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationMetadata200ApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetDestinationMetadata200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationMetadata200ApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetDestinationMetadata200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationMetadata200ApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetDestinationMetadata200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationMetadata200ApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetDestinationsCatalog - Get Destinations Catalog
// Returns a list of all available Destinations in the Segment catalog.
func (s *catalog) GetDestinationsCatalog(ctx context.Context, request operations.GetDestinationsCatalogRequest, opts ...operations.Option) (*operations.GetDestinationsCatalogResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/catalog/destinations"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

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

	res := &operations.GetDestinationsCatalogResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetDestinationsCatalog200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationsCatalog200ApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetDestinationsCatalog200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationsCatalog200ApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetDestinationsCatalog200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationsCatalog200ApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetDestinationsCatalog200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetDestinationsCatalog200ApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetSourceMetadata - Get Source Metadata
// Returns a Source catalog item by its id.
func (s *catalog) GetSourceMetadata(ctx context.Context, request operations.GetSourceMetadataRequest, opts ...operations.Option) (*operations.GetSourceMetadataResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/catalog/sources/{sourceMetadataId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

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

	res := &operations.GetSourceMetadataResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetSourceMetadata200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourceMetadata200ApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetSourceMetadata200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourceMetadata200ApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetSourceMetadata200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourceMetadata200ApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetSourceMetadata200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourceMetadata200ApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetSourcesCatalog - Get Sources Catalog
// Returns a list of all available Sources in the Segment catalog.
func (s *catalog) GetSourcesCatalog(ctx context.Context, request operations.GetSourcesCatalogRequest, opts ...operations.Option) (*operations.GetSourcesCatalogResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/catalog/sources"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

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

	res := &operations.GetSourcesCatalogResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetSourcesCatalog200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourcesCatalog200ApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetSourcesCatalog200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourcesCatalog200ApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetSourcesCatalog200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourcesCatalog200ApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetSourcesCatalog200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetSourcesCatalog200ApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetWarehouseMetadata - Get Warehouse Metadata
// Returns a Warehouse catalog item by its id.
func (s *catalog) GetWarehouseMetadata(ctx context.Context, request operations.GetWarehouseMetadataRequest, opts ...operations.Option) (*operations.GetWarehouseMetadataResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/catalog/warehouses/{warehouseMetadataId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

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

	res := &operations.GetWarehouseMetadataResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetWarehouseMetadata200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehouseMetadata200ApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehouseMetadata200ApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehouseMetadata200ApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehouseMetadata200ApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetWarehousesCatalog - Get Warehouses Catalog
// Returns a list of all available Warehouses in the Segment catalog.
func (s *catalog) GetWarehousesCatalog(ctx context.Context, request operations.GetWarehousesCatalogRequest, opts ...operations.Option) (*operations.GetWarehousesCatalogResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/catalog/warehouses"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

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

	res := &operations.GetWarehousesCatalogResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetWarehousesCatalog200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehousesCatalog200ApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetWarehousesCatalog200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehousesCatalog200ApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetWarehousesCatalog200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehousesCatalog200ApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetWarehousesCatalog200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetWarehousesCatalog200ApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
