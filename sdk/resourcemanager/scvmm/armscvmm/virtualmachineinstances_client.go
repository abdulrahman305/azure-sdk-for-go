//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// VirtualMachineInstancesClient contains the methods for the VirtualMachineInstances group.
// Don't use this type directly, use NewVirtualMachineInstancesClient() instead.
type VirtualMachineInstancesClient struct {
	internal *arm.Client
}

// NewVirtualMachineInstancesClient creates a new instance of VirtualMachineInstancesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineInstancesClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateCheckpoint - Creates a checkpoint in virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - VirtualMachineInstancesClientBeginCreateCheckpointOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginCreateCheckpoint
//     method.
func (client *VirtualMachineInstancesClient) BeginCreateCheckpoint(ctx context.Context, resourceURI string, body VirtualMachineCreateCheckpoint, options *VirtualMachineInstancesClientBeginCreateCheckpointOptions) (*runtime.Poller[VirtualMachineInstancesClientCreateCheckpointResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createCheckpoint(ctx, resourceURI, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientCreateCheckpointResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientCreateCheckpointResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateCheckpoint - Creates a checkpoint in virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) createCheckpoint(ctx context.Context, resourceURI string, body VirtualMachineCreateCheckpoint, options *VirtualMachineInstancesClientBeginCreateCheckpointOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginCreateCheckpoint"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCheckpointCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCheckpointCreateRequest creates the CreateCheckpoint request.
func (client *VirtualMachineInstancesClient) createCheckpointCreateRequest(ctx context.Context, resourceURI string, body VirtualMachineCreateCheckpoint, options *VirtualMachineInstancesClientBeginCreateCheckpointOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default/createCheckpoint"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - The operation to create or update a virtual machine instance. Please note some properties can be
// set only during virtual machine instance creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - resource - Resource create parameters.
//   - options - VirtualMachineInstancesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualMachineInstancesClient) BeginCreateOrUpdate(ctx context.Context, resourceURI string, resource VirtualMachineInstance, options *VirtualMachineInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineInstancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceURI, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation to create or update a virtual machine instance. Please note some properties can be set only
// during virtual machine instance creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) createOrUpdate(ctx context.Context, resourceURI string, resource VirtualMachineInstance, options *VirtualMachineInstancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, resource, options)
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
func (client *VirtualMachineInstancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, resource VirtualMachineInstance, options *VirtualMachineInstancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
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

// BeginDelete - The operation to delete a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - VirtualMachineInstancesClientBeginDeleteOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginDelete
//     method.
func (client *VirtualMachineInstancesClient) BeginDelete(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) deleteOperation(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, options)
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
func (client *VirtualMachineInstancesClient) deleteCreateRequest(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	if options != nil && options.DeleteFromHost != nil {
		reqQP.Set("deleteFromHost", string(*options.DeleteFromHost))
	}
	if options != nil && options.Force != nil {
		reqQP.Set("force", string(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeleteCheckpoint - Deletes a checkpoint in virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - VirtualMachineInstancesClientBeginDeleteCheckpointOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginDeleteCheckpoint
//     method.
func (client *VirtualMachineInstancesClient) BeginDeleteCheckpoint(ctx context.Context, resourceURI string, body VirtualMachineDeleteCheckpoint, options *VirtualMachineInstancesClientBeginDeleteCheckpointOptions) (*runtime.Poller[VirtualMachineInstancesClientDeleteCheckpointResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteCheckpoint(ctx, resourceURI, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientDeleteCheckpointResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientDeleteCheckpointResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DeleteCheckpoint - Deletes a checkpoint in virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) deleteCheckpoint(ctx context.Context, resourceURI string, body VirtualMachineDeleteCheckpoint, options *VirtualMachineInstancesClientBeginDeleteCheckpointOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginDeleteCheckpoint"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCheckpointCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCheckpointCreateRequest creates the DeleteCheckpoint request.
func (client *VirtualMachineInstancesClient) deleteCheckpointCreateRequest(ctx context.Context, resourceURI string, body VirtualMachineDeleteCheckpoint, options *VirtualMachineInstancesClientBeginDeleteCheckpointOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default/deleteCheckpoint"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Retrieves information about a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - VirtualMachineInstancesClientGetOptions contains the optional parameters for the VirtualMachineInstancesClient.Get
//     method.
func (client *VirtualMachineInstancesClient) Get(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientGetOptions) (VirtualMachineInstancesClientGetResponse, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return VirtualMachineInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineInstancesClient) getCreateRequest(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
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
func (client *VirtualMachineInstancesClient) getHandleResponse(resp *http.Response) (VirtualMachineInstancesClientGetResponse, error) {
	result := VirtualMachineInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineInstance); err != nil {
		return VirtualMachineInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the virtual machine instances within the specified parent resource.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - VirtualMachineInstancesClientListOptions contains the optional parameters for the VirtualMachineInstancesClient.NewListPager
//     method.
func (client *VirtualMachineInstancesClient) NewListPager(resourceURI string, options *VirtualMachineInstancesClientListOptions) *runtime.Pager[VirtualMachineInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineInstancesClientListResponse]{
		More: func(page VirtualMachineInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineInstancesClientListResponse) (VirtualMachineInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualMachineInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return VirtualMachineInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineInstancesClient) listCreateRequest(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
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

// listHandleResponse handles the List response.
func (client *VirtualMachineInstancesClient) listHandleResponse(resp *http.Response) (VirtualMachineInstancesClientListResponse, error) {
	result := VirtualMachineInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineInstanceListResult); err != nil {
		return VirtualMachineInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginRestart - The operation to restart a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - VirtualMachineInstancesClientBeginRestartOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginRestart
//     method.
func (client *VirtualMachineInstancesClient) BeginRestart(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginRestartOptions) (*runtime.Poller[VirtualMachineInstancesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientRestartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - The operation to restart a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) restart(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *VirtualMachineInstancesClient) restartCreateRequest(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default/restart"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginRestoreCheckpoint - Restores to a checkpoint in virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - VirtualMachineInstancesClientBeginRestoreCheckpointOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginRestoreCheckpoint
//     method.
func (client *VirtualMachineInstancesClient) BeginRestoreCheckpoint(ctx context.Context, resourceURI string, body VirtualMachineRestoreCheckpoint, options *VirtualMachineInstancesClientBeginRestoreCheckpointOptions) (*runtime.Poller[VirtualMachineInstancesClientRestoreCheckpointResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restoreCheckpoint(ctx, resourceURI, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientRestoreCheckpointResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientRestoreCheckpointResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RestoreCheckpoint - Restores to a checkpoint in virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) restoreCheckpoint(ctx context.Context, resourceURI string, body VirtualMachineRestoreCheckpoint, options *VirtualMachineInstancesClientBeginRestoreCheckpointOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginRestoreCheckpoint"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restoreCheckpointCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restoreCheckpointCreateRequest creates the RestoreCheckpoint request.
func (client *VirtualMachineInstancesClient) restoreCheckpointCreateRequest(ctx context.Context, resourceURI string, body VirtualMachineRestoreCheckpoint, options *VirtualMachineInstancesClientBeginRestoreCheckpointOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default/restoreCheckpoint"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginStart - The operation to start a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - VirtualMachineInstancesClientBeginStartOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginStart
//     method.
func (client *VirtualMachineInstancesClient) BeginStart(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginStartOptions) (*runtime.Poller[VirtualMachineInstancesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceURI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - The operation to start a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) start(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *VirtualMachineInstancesClient) startCreateRequest(ctx context.Context, resourceURI string, options *VirtualMachineInstancesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default/start"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - The operation to power off (stop) a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - VirtualMachineInstancesClientBeginStopOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginStop
//     method.
func (client *VirtualMachineInstancesClient) BeginStop(ctx context.Context, resourceURI string, body StopVirtualMachineOptions, options *VirtualMachineInstancesClientBeginStopOptions) (*runtime.Poller[VirtualMachineInstancesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceURI, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientStopResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Stop - The operation to power off (stop) a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) stop(ctx context.Context, resourceURI string, body StopVirtualMachineOptions, options *VirtualMachineInstancesClientBeginStopOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginStop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopCreateRequest creates the Stop request.
func (client *VirtualMachineInstancesClient) stopCreateRequest(ctx context.Context, resourceURI string, body StopVirtualMachineOptions, options *VirtualMachineInstancesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default/stop"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - The operation to update a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - properties - The resource properties to be updated.
//   - options - VirtualMachineInstancesClientBeginUpdateOptions contains the optional parameters for the VirtualMachineInstancesClient.BeginUpdate
//     method.
func (client *VirtualMachineInstancesClient) BeginUpdate(ctx context.Context, resourceURI string, properties VirtualMachineInstanceUpdate, options *VirtualMachineInstancesClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineInstancesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceURI, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineInstancesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachineInstancesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation to update a virtual machine instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-07
func (client *VirtualMachineInstancesClient) update(ctx context.Context, resourceURI string, properties VirtualMachineInstanceUpdate, options *VirtualMachineInstancesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineInstancesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceURI, properties, options)
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
func (client *VirtualMachineInstancesClient) updateCreateRequest(ctx context.Context, resourceURI string, properties VirtualMachineInstanceUpdate, options *VirtualMachineInstancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ScVmm/virtualMachineInstances/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
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
