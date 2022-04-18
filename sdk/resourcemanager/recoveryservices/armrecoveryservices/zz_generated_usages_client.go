//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

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

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsagesClient, error) {
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
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByVaultsPager - Fetches the usages of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// vaultName - The name of the recovery services vault.
// options - UsagesClientListByVaultsOptions contains the optional parameters for the UsagesClient.ListByVaults method.
func (client *UsagesClient) NewListByVaultsPager(resourceGroupName string, vaultName string, options *UsagesClientListByVaultsOptions) *runtime.Pager[UsagesClientListByVaultsResponse] {
	return runtime.NewPager(runtime.PageProcessor[UsagesClientListByVaultsResponse]{
		More: func(page UsagesClientListByVaultsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UsagesClientListByVaultsResponse) (UsagesClientListByVaultsResponse, error) {
			req, err := client.listByVaultsCreateRequest(ctx, resourceGroupName, vaultName, options)
			if err != nil {
				return UsagesClientListByVaultsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return UsagesClientListByVaultsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsagesClientListByVaultsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVaultsHandleResponse(resp)
		},
	})
}

// listByVaultsCreateRequest creates the ListByVaults request.
func (client *UsagesClient) listByVaultsCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *UsagesClientListByVaultsOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVaultsHandleResponse handles the ListByVaults response.
func (client *UsagesClient) listByVaultsHandleResponse(resp *http.Response) (UsagesClientListByVaultsResponse, error) {
	result := UsagesClientListByVaultsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultUsageList); err != nil {
		return UsagesClientListByVaultsResponse{}, err
	}
	return result, nil
}
