//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CloudsClient contains the methods for the Clouds group.
// Don't use this type directly, use NewCloudsClient() instead.
type CloudsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCloudsClient creates a new instance of CloudsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCloudsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CloudsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Onboards the ScVmm fabric cloud as an Azure cloud resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudResourceName - Name of the Cloud.
//   - resource - Resource create parameters.
//   - options - CloudsClientBeginCreateOrUpdateOptions contains the optional parameters for the CloudsClient.BeginCreateOrUpdate
//     method.
func (client *CloudsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, cloudResourceName string, resource Cloud, options *CloudsClientBeginCreateOrUpdateOptions) (*runtime.Poller[CloudsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, cloudResourceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Onboards the ScVmm fabric cloud as an Azure cloud resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *CloudsClient) createOrUpdate(ctx context.Context, resourceGroupName string, cloudResourceName string, resource Cloud, options *CloudsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, cloudResourceName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CloudsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, cloudResourceName string, resource Cloud, options *CloudsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/clouds/{cloudResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudResourceName == "" {
		return nil, errors.New("parameter cloudResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudResourceName}", url.PathEscape(cloudResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deregisters the ScVmm fabric cloud from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudResourceName - Name of the Cloud.
//   - options - CloudsClientBeginDeleteOptions contains the optional parameters for the CloudsClient.BeginDelete method.
func (client *CloudsClient) BeginDelete(ctx context.Context, resourceGroupName string, cloudResourceName string, options *CloudsClientBeginDeleteOptions) (*runtime.Poller[CloudsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cloudResourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deregisters the ScVmm fabric cloud from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *CloudsClient) deleteOperation(ctx context.Context, resourceGroupName string, cloudResourceName string, options *CloudsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cloudResourceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CloudsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cloudResourceName string, options *CloudsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/clouds/{cloudResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudResourceName == "" {
		return nil, errors.New("parameter cloudResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudResourceName}", url.PathEscape(cloudResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	if options != nil && options.Force != nil {
		reqQP.Set("force", string(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements Cloud GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudResourceName - Name of the Cloud.
//   - options - CloudsClientGetOptions contains the optional parameters for the CloudsClient.Get method.
func (client *CloudsClient) Get(ctx context.Context, resourceGroupName string, cloudResourceName string, options *CloudsClientGetOptions) (CloudsClientGetResponse, error) {
	var err error
	const operationName = "CloudsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, cloudResourceName, options)
	if err != nil {
		return CloudsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CloudsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CloudsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CloudsClient) getCreateRequest(ctx context.Context, resourceGroupName string, cloudResourceName string, options *CloudsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/clouds/{cloudResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudResourceName == "" {
		return nil, errors.New("parameter cloudResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudResourceName}", url.PathEscape(cloudResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudsClient) getHandleResponse(resp *http.Response) (CloudsClientGetResponse, error) {
	result := CloudsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cloud); err != nil {
		return CloudsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of Clouds in a resource group.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CloudsClientListByResourceGroupOptions contains the optional parameters for the CloudsClient.NewListByResourceGroupPager
//     method.
func (client *CloudsClient) NewListByResourceGroupPager(resourceGroupName string, options *CloudsClientListByResourceGroupOptions) *runtime.Pager[CloudsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudsClientListByResourceGroupResponse]{
		More: func(page CloudsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudsClientListByResourceGroupResponse) (CloudsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CloudsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CloudsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CloudsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CloudsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/clouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CloudsClient) listByResourceGroupHandleResponse(resp *http.Response) (CloudsClientListByResourceGroupResponse, error) {
	result := CloudsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudListResult); err != nil {
		return CloudsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List of Clouds in a subscription.
//
// Generated from API version 2023-10-07
//   - options - CloudsClientListBySubscriptionOptions contains the optional parameters for the CloudsClient.NewListBySubscriptionPager
//     method.
func (client *CloudsClient) NewListBySubscriptionPager(options *CloudsClientListBySubscriptionOptions) *runtime.Pager[CloudsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudsClientListBySubscriptionResponse]{
		More: func(page CloudsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudsClientListBySubscriptionResponse) (CloudsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CloudsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CloudsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CloudsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CloudsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ScVmm/clouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CloudsClient) listBySubscriptionHandleResponse(resp *http.Response) (CloudsClientListBySubscriptionResponse, error) {
	result := CloudsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudListResult); err != nil {
		return CloudsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the Clouds resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudResourceName - Name of the Cloud.
//   - properties - The resource properties to be updated.
//   - options - CloudsClientBeginUpdateOptions contains the optional parameters for the CloudsClient.BeginUpdate method.
func (client *CloudsClient) BeginUpdate(ctx context.Context, resourceGroupName string, cloudResourceName string, properties CloudTagsUpdate, options *CloudsClientBeginUpdateOptions) (*runtime.Poller[CloudsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, cloudResourceName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CloudsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the Clouds resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *CloudsClient) update(ctx context.Context, resourceGroupName string, cloudResourceName string, properties CloudTagsUpdate, options *CloudsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CloudsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, cloudResourceName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *CloudsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, cloudResourceName string, properties CloudTagsUpdate, options *CloudsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/clouds/{cloudResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudResourceName == "" {
		return nil, errors.New("parameter cloudResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudResourceName}", url.PathEscape(cloudResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
