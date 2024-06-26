//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationdiscoverysap

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// SapDiscoverySitesClientCreateResponse contains the response from method SapDiscoverySitesClient.BeginCreate.
type SapDiscoverySitesClientCreateResponse struct {
	// Define the SAP Migration discovery site resource.
	SAPDiscoverySite
}

// SapDiscoverySitesClientDeleteResponse contains the response from method SapDiscoverySitesClient.BeginDelete.
type SapDiscoverySitesClientDeleteResponse struct {
	// placeholder for future response values
}

// SapDiscoverySitesClientGetResponse contains the response from method SapDiscoverySitesClient.Get.
type SapDiscoverySitesClientGetResponse struct {
	// Define the SAP Migration discovery site resource.
	SAPDiscoverySite
}

// SapDiscoverySitesClientImportEntitiesResponse contains the response from method SapDiscoverySitesClient.BeginImportEntities.
type SapDiscoverySitesClientImportEntitiesResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapDiscoverySitesClientListByResourceGroupResponse contains the response from method SapDiscoverySitesClient.NewListByResourceGroupPager.
type SapDiscoverySitesClientListByResourceGroupResponse struct {
	// The response of a SAPDiscoverySite list operation.
	SAPDiscoverySiteListResult
}

// SapDiscoverySitesClientListBySubscriptionResponse contains the response from method SapDiscoverySitesClient.NewListBySubscriptionPager.
type SapDiscoverySitesClientListBySubscriptionResponse struct {
	// The response of a SAPDiscoverySite list operation.
	SAPDiscoverySiteListResult
}

// SapDiscoverySitesClientUpdateResponse contains the response from method SapDiscoverySitesClient.Update.
type SapDiscoverySitesClientUpdateResponse struct {
	// Define the SAP Migration discovery site resource.
	SAPDiscoverySite
}

// SapInstancesClientCreateResponse contains the response from method SapInstancesClient.BeginCreate.
type SapInstancesClientCreateResponse struct {
	// Define the SAP Instance resource.
	SAPInstance
}

// SapInstancesClientDeleteResponse contains the response from method SapInstancesClient.BeginDelete.
type SapInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SapInstancesClientGetResponse contains the response from method SapInstancesClient.Get.
type SapInstancesClientGetResponse struct {
	// Define the SAP Instance resource.
	SAPInstance
}

// SapInstancesClientListBySapDiscoverySiteResponse contains the response from method SapInstancesClient.NewListBySapDiscoverySitePager.
type SapInstancesClientListBySapDiscoverySiteResponse struct {
	// The response of a SAPInstance list operation.
	SAPInstanceListResult
}

// SapInstancesClientUpdateResponse contains the response from method SapInstancesClient.Update.
type SapInstancesClientUpdateResponse struct {
	// Define the SAP Instance resource.
	SAPInstance
}

// ServerInstancesClientCreateResponse contains the response from method ServerInstancesClient.BeginCreate.
type ServerInstancesClientCreateResponse struct {
	// Define the Server Instance resource.
	ServerInstance
}

// ServerInstancesClientDeleteResponse contains the response from method ServerInstancesClient.BeginDelete.
type ServerInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerInstancesClientGetResponse contains the response from method ServerInstancesClient.Get.
type ServerInstancesClientGetResponse struct {
	// Define the Server Instance resource.
	ServerInstance
}

// ServerInstancesClientListBySapInstanceResponse contains the response from method ServerInstancesClient.NewListBySapInstancePager.
type ServerInstancesClientListBySapInstanceResponse struct {
	// The response of a ServerInstance list operation.
	ServerInstanceListResult
}

// ServerInstancesClientUpdateResponse contains the response from method ServerInstancesClient.Update.
type ServerInstancesClientUpdateResponse struct {
	// Define the Server Instance resource.
	ServerInstance
}
