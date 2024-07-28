//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CloudServicesUpdateDomainServer is a fake server for instances of the armcompute.CloudServicesUpdateDomainClient type.
type CloudServicesUpdateDomainServer struct {
	// GetUpdateDomain is the fake for method CloudServicesUpdateDomainClient.GetUpdateDomain
	// HTTP status codes to indicate success: http.StatusOK
	GetUpdateDomain func(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, options *armcompute.CloudServicesUpdateDomainClientGetUpdateDomainOptions) (resp azfake.Responder[armcompute.CloudServicesUpdateDomainClientGetUpdateDomainResponse], errResp azfake.ErrorResponder)

	// NewListUpdateDomainsPager is the fake for method CloudServicesUpdateDomainClient.NewListUpdateDomainsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListUpdateDomainsPager func(resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesUpdateDomainClientListUpdateDomainsOptions) (resp azfake.PagerResponder[armcompute.CloudServicesUpdateDomainClientListUpdateDomainsResponse])

	// BeginWalkUpdateDomain is the fake for method CloudServicesUpdateDomainClient.BeginWalkUpdateDomain
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginWalkUpdateDomain func(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, parameters armcompute.UpdateDomain, options *armcompute.CloudServicesUpdateDomainClientBeginWalkUpdateDomainOptions) (resp azfake.PollerResponder[armcompute.CloudServicesUpdateDomainClientWalkUpdateDomainResponse], errResp azfake.ErrorResponder)
}

// NewCloudServicesUpdateDomainServerTransport creates a new instance of CloudServicesUpdateDomainServerTransport with the provided implementation.
// The returned CloudServicesUpdateDomainServerTransport instance is connected to an instance of armcompute.CloudServicesUpdateDomainClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudServicesUpdateDomainServerTransport(srv *CloudServicesUpdateDomainServer) *CloudServicesUpdateDomainServerTransport {
	return &CloudServicesUpdateDomainServerTransport{
		srv:                       srv,
		newListUpdateDomainsPager: newTracker[azfake.PagerResponder[armcompute.CloudServicesUpdateDomainClientListUpdateDomainsResponse]](),
		beginWalkUpdateDomain:     newTracker[azfake.PollerResponder[armcompute.CloudServicesUpdateDomainClientWalkUpdateDomainResponse]](),
	}
}

// CloudServicesUpdateDomainServerTransport connects instances of armcompute.CloudServicesUpdateDomainClient to instances of CloudServicesUpdateDomainServer.
// Don't use this type directly, use NewCloudServicesUpdateDomainServerTransport instead.
type CloudServicesUpdateDomainServerTransport struct {
	srv                       *CloudServicesUpdateDomainServer
	newListUpdateDomainsPager *tracker[azfake.PagerResponder[armcompute.CloudServicesUpdateDomainClientListUpdateDomainsResponse]]
	beginWalkUpdateDomain     *tracker[azfake.PollerResponder[armcompute.CloudServicesUpdateDomainClientWalkUpdateDomainResponse]]
}

// Do implements the policy.Transporter interface for CloudServicesUpdateDomainServerTransport.
func (c *CloudServicesUpdateDomainServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudServicesUpdateDomainClient.GetUpdateDomain":
		resp, err = c.dispatchGetUpdateDomain(req)
	case "CloudServicesUpdateDomainClient.NewListUpdateDomainsPager":
		resp, err = c.dispatchNewListUpdateDomainsPager(req)
	case "CloudServicesUpdateDomainClient.BeginWalkUpdateDomain":
		resp, err = c.dispatchBeginWalkUpdateDomain(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudServicesUpdateDomainServerTransport) dispatchGetUpdateDomain(req *http.Request) (*http.Response, error) {
	if c.srv.GetUpdateDomain == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetUpdateDomain not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDomains/(?P<updateDomain>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	updateDomainUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("updateDomain")])
	if err != nil {
		return nil, err
	}
	updateDomainParam, err := parseWithCast(updateDomainUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetUpdateDomain(req.Context(), resourceGroupNameParam, cloudServiceNameParam, updateDomainParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UpdateDomain, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServicesUpdateDomainServerTransport) dispatchNewListUpdateDomainsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListUpdateDomainsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListUpdateDomainsPager not implemented")}
	}
	newListUpdateDomainsPager := c.newListUpdateDomainsPager.get(req)
	if newListUpdateDomainsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDomains`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListUpdateDomainsPager(resourceGroupNameParam, cloudServiceNameParam, nil)
		newListUpdateDomainsPager = &resp
		c.newListUpdateDomainsPager.add(req, newListUpdateDomainsPager)
		server.PagerResponderInjectNextLinks(newListUpdateDomainsPager, req, func(page *armcompute.CloudServicesUpdateDomainClientListUpdateDomainsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListUpdateDomainsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListUpdateDomainsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListUpdateDomainsPager) {
		c.newListUpdateDomainsPager.remove(req)
	}
	return resp, nil
}

func (c *CloudServicesUpdateDomainServerTransport) dispatchBeginWalkUpdateDomain(req *http.Request) (*http.Response, error) {
	if c.srv.BeginWalkUpdateDomain == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginWalkUpdateDomain not implemented")}
	}
	beginWalkUpdateDomain := c.beginWalkUpdateDomain.get(req)
	if beginWalkUpdateDomain == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDomains/(?P<updateDomain>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.UpdateDomain](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		updateDomainUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("updateDomain")])
		if err != nil {
			return nil, err
		}
		updateDomainParam, err := parseWithCast(updateDomainUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginWalkUpdateDomain(req.Context(), resourceGroupNameParam, cloudServiceNameParam, updateDomainParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginWalkUpdateDomain = &respr
		c.beginWalkUpdateDomain.add(req, beginWalkUpdateDomain)
	}

	resp, err := server.PollerResponderNext(beginWalkUpdateDomain, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginWalkUpdateDomain.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginWalkUpdateDomain) {
		c.beginWalkUpdateDomain.remove(req)
	}

	return resp, nil
}
