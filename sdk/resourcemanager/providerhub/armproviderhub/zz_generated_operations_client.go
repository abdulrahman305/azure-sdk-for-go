//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the operation supported by the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// operationsPutContent - The operations content properties supplied to the CreateOrUpdate operation.
// options - OperationsClientCreateOrUpdateOptions contains the optional parameters for the OperationsClient.CreateOrUpdate
// method.
func (client *OperationsClient) CreateOrUpdate(ctx context.Context, providerNamespace string, operationsPutContent OperationsPutContent, options *OperationsClientCreateOrUpdateOptions) (OperationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, operationsPutContent, options)
	if err != nil {
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OperationsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, operationsPutContent OperationsPutContent, options *OperationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/operations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, operationsPutContent)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *OperationsClient) createOrUpdateHandleResponse(resp *http.Response) (OperationsClientCreateOrUpdateResponse, error) {
	result := OperationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsContent); err != nil {
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an operation.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - OperationsClientDeleteOptions contains the optional parameters for the OperationsClient.Delete method.
func (client *OperationsClient) Delete(ctx context.Context, providerNamespace string, options *OperationsClientDeleteOptions) (OperationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return OperationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return OperationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return OperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OperationsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, options *OperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/operations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// NewListPager - Lists all the operations supported by Microsoft.ProviderHub.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
func (client *OperationsClient) NewListPager(options *OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[OperationsClientListResponse]{
		More: func(page OperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientListResponse) (OperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ProviderHub/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsDefinitionArrayResponseWithContinuation); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}

// ListByProviderRegistration - Gets the operations supported by the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - OperationsClientListByProviderRegistrationOptions contains the optional parameters for the OperationsClient.ListByProviderRegistration
// method.
func (client *OperationsClient) ListByProviderRegistration(ctx context.Context, providerNamespace string, options *OperationsClientListByProviderRegistrationOptions) (OperationsClientListByProviderRegistrationResponse, error) {
	req, err := client.listByProviderRegistrationCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsClientListByProviderRegistrationResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByProviderRegistrationHandleResponse(resp)
}

// listByProviderRegistrationCreateRequest creates the ListByProviderRegistration request.
func (client *OperationsClient) listByProviderRegistrationCreateRequest(ctx context.Context, providerNamespace string, options *OperationsClientListByProviderRegistrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/operations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProviderRegistrationHandleResponse handles the ListByProviderRegistration response.
func (client *OperationsClient) listByProviderRegistrationHandleResponse(resp *http.Response) (OperationsClientListByProviderRegistrationResponse, error) {
	result := OperationsClientListByProviderRegistrationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsDefinitionArray); err != nil {
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	return result, nil
}
