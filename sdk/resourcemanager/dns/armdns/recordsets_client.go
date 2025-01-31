//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// RecordSetsClient contains the methods for the RecordSets group.
// Don't use this type directly, use NewRecordSetsClient() instead.
type RecordSetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRecordSetsClient creates a new instance of RecordSetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecordSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecordSetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecordSetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a record set within a DNS zone. Record sets of type SOA can be updated but not created
// (they are created when the DNS zone is created).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - relativeRecordSetName - The name of the record set, relative to the name of the zone.
//   - recordType - The type of DNS record in this record set.
//   - parameters - Parameters supplied to the CreateOrUpdate operation.
//   - options - RecordSetsClientCreateOrUpdateOptions contains the optional parameters for the RecordSetsClient.CreateOrUpdate
//     method.
func (client *RecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, options *RecordSetsClientCreateOrUpdateOptions) (RecordSetsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "RecordSetsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, zoneName, relativeRecordSetName, recordType, parameters, options)
	if err != nil {
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RecordSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, options *RecordSetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RecordSetsClient) createOrUpdateHandleResponse(resp *http.Response) (RecordSetsClientCreateOrUpdateResponse, error) {
	result := RecordSetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSet); err != nil {
		return RecordSetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a record set from a DNS zone. This operation cannot be undone. Record sets of type SOA cannot be deleted
// (they are deleted when the DNS zone is deleted).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - relativeRecordSetName - The name of the record set, relative to the name of the zone.
//   - recordType - The type of DNS record in this record set.
//   - options - RecordSetsClientDeleteOptions contains the optional parameters for the RecordSetsClient.Delete method.
func (client *RecordSetsClient) Delete(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, options *RecordSetsClientDeleteOptions) (RecordSetsClientDeleteResponse, error) {
	var err error
	const operationName = "RecordSetsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, zoneName, relativeRecordSetName, recordType, options)
	if err != nil {
		return RecordSetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecordSetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RecordSetsClientDeleteResponse{}, err
	}
	return RecordSetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RecordSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, options *RecordSetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// Get - Gets a record set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - relativeRecordSetName - The name of the record set, relative to the name of the zone.
//   - recordType - The type of DNS record in this record set.
//   - options - RecordSetsClientGetOptions contains the optional parameters for the RecordSetsClient.Get method.
func (client *RecordSetsClient) Get(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, options *RecordSetsClientGetOptions) (RecordSetsClientGetResponse, error) {
	var err error
	const operationName = "RecordSetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, zoneName, relativeRecordSetName, recordType, options)
	if err != nil {
		return RecordSetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecordSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecordSetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecordSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, options *RecordSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecordSetsClient) getHandleResponse(resp *http.Response) (RecordSetsClientGetResponse, error) {
	result := RecordSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSet); err != nil {
		return RecordSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListAllByDNSZonePager - Lists all record sets in a DNS zone.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - RecordSetsClientListAllByDNSZoneOptions contains the optional parameters for the RecordSetsClient.NewListAllByDNSZonePager
//     method.
func (client *RecordSetsClient) NewListAllByDNSZonePager(resourceGroupName string, zoneName string, options *RecordSetsClientListAllByDNSZoneOptions) *runtime.Pager[RecordSetsClientListAllByDNSZoneResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecordSetsClientListAllByDNSZoneResponse]{
		More: func(page RecordSetsClientListAllByDNSZoneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecordSetsClientListAllByDNSZoneResponse) (RecordSetsClientListAllByDNSZoneResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecordSetsClient.NewListAllByDNSZonePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllByDNSZoneCreateRequest(ctx, resourceGroupName, zoneName, options)
			}, nil)
			if err != nil {
				return RecordSetsClientListAllByDNSZoneResponse{}, err
			}
			return client.listAllByDNSZoneHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllByDNSZoneCreateRequest creates the ListAllByDNSZone request.
func (client *RecordSetsClient) listAllByDNSZoneCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *RecordSetsClientListAllByDNSZoneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/all"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.RecordSetNameSuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.RecordSetNameSuffix)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllByDNSZoneHandleResponse handles the ListAllByDNSZone response.
func (client *RecordSetsClient) listAllByDNSZoneHandleResponse(resp *http.Response) (RecordSetsClientListAllByDNSZoneResponse, error) {
	result := RecordSetsClientListAllByDNSZoneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSetListResult); err != nil {
		return RecordSetsClientListAllByDNSZoneResponse{}, err
	}
	return result, nil
}

// NewListByDNSZonePager - Lists all record sets in a DNS zone.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - RecordSetsClientListByDNSZoneOptions contains the optional parameters for the RecordSetsClient.NewListByDNSZonePager
//     method.
func (client *RecordSetsClient) NewListByDNSZonePager(resourceGroupName string, zoneName string, options *RecordSetsClientListByDNSZoneOptions) *runtime.Pager[RecordSetsClientListByDNSZoneResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecordSetsClientListByDNSZoneResponse]{
		More: func(page RecordSetsClientListByDNSZoneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecordSetsClientListByDNSZoneResponse) (RecordSetsClientListByDNSZoneResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecordSetsClient.NewListByDNSZonePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDNSZoneCreateRequest(ctx, resourceGroupName, zoneName, options)
			}, nil)
			if err != nil {
				return RecordSetsClientListByDNSZoneResponse{}, err
			}
			return client.listByDNSZoneHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDNSZoneCreateRequest creates the ListByDNSZone request.
func (client *RecordSetsClient) listByDNSZoneCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *RecordSetsClientListByDNSZoneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Recordsetnamesuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.Recordsetnamesuffix)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDNSZoneHandleResponse handles the ListByDNSZone response.
func (client *RecordSetsClient) listByDNSZoneHandleResponse(resp *http.Response) (RecordSetsClientListByDNSZoneResponse, error) {
	result := RecordSetsClientListByDNSZoneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSetListResult); err != nil {
		return RecordSetsClientListByDNSZoneResponse{}, err
	}
	return result, nil
}

// NewListByTypePager - Lists the record sets of a specified type in a DNS zone.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - recordType - The type of DNS record in this record set.
//   - options - RecordSetsClientListByTypeOptions contains the optional parameters for the RecordSetsClient.NewListByTypePager
//     method.
func (client *RecordSetsClient) NewListByTypePager(resourceGroupName string, zoneName string, recordType RecordType, options *RecordSetsClientListByTypeOptions) *runtime.Pager[RecordSetsClientListByTypeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecordSetsClientListByTypeResponse]{
		More: func(page RecordSetsClientListByTypeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecordSetsClientListByTypeResponse) (RecordSetsClientListByTypeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecordSetsClient.NewListByTypePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTypeCreateRequest(ctx, resourceGroupName, zoneName, recordType, options)
			}, nil)
			if err != nil {
				return RecordSetsClientListByTypeResponse{}, err
			}
			return client.listByTypeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTypeCreateRequest creates the ListByType request.
func (client *RecordSetsClient) listByTypeCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, recordType RecordType, options *RecordSetsClientListByTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Recordsetnamesuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.Recordsetnamesuffix)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTypeHandleResponse handles the ListByType response.
func (client *RecordSetsClient) listByTypeHandleResponse(resp *http.Response) (RecordSetsClientListByTypeResponse, error) {
	result := RecordSetsClientListByTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSetListResult); err != nil {
		return RecordSetsClientListByTypeResponse{}, err
	}
	return result, nil
}

// Update - Updates a record set within a DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - relativeRecordSetName - The name of the record set, relative to the name of the zone.
//   - recordType - The type of DNS record in this record set.
//   - parameters - Parameters supplied to the Update operation.
//   - options - RecordSetsClientUpdateOptions contains the optional parameters for the RecordSetsClient.Update method.
func (client *RecordSetsClient) Update(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, options *RecordSetsClientUpdateOptions) (RecordSetsClientUpdateResponse, error) {
	var err error
	const operationName = "RecordSetsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, zoneName, relativeRecordSetName, recordType, parameters, options)
	if err != nil {
		return RecordSetsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecordSetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecordSetsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *RecordSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType RecordType, parameters RecordSet, options *RecordSetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *RecordSetsClient) updateHandleResponse(resp *http.Response) (RecordSetsClientUpdateResponse, error) {
	result := RecordSetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecordSet); err != nil {
		return RecordSetsClientUpdateResponse{}, err
	}
	return result, nil
}
