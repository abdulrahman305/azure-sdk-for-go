// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinksServer is a fake server for instances of the armmongocluster.PrivateLinksClient type.
type PrivateLinksServer struct {
	// NewListByMongoClusterPager is the fake for method PrivateLinksClient.NewListByMongoClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMongoClusterPager func(resourceGroupName string, mongoClusterName string, options *armmongocluster.PrivateLinksClientListByMongoClusterOptions) (resp azfake.PagerResponder[armmongocluster.PrivateLinksClientListByMongoClusterResponse])
}

// NewPrivateLinksServerTransport creates a new instance of PrivateLinksServerTransport with the provided implementation.
// The returned PrivateLinksServerTransport instance is connected to an instance of armmongocluster.PrivateLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinksServerTransport(srv *PrivateLinksServer) *PrivateLinksServerTransport {
	return &PrivateLinksServerTransport{
		srv:                        srv,
		newListByMongoClusterPager: newTracker[azfake.PagerResponder[armmongocluster.PrivateLinksClientListByMongoClusterResponse]](),
	}
}

// PrivateLinksServerTransport connects instances of armmongocluster.PrivateLinksClient to instances of PrivateLinksServer.
// Don't use this type directly, use NewPrivateLinksServerTransport instead.
type PrivateLinksServerTransport struct {
	srv                        *PrivateLinksServer
	newListByMongoClusterPager *tracker[azfake.PagerResponder[armmongocluster.PrivateLinksClientListByMongoClusterResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinksServerTransport.
func (p *PrivateLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinksServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinksClient.NewListByMongoClusterPager":
		resp, err = p.dispatchNewListByMongoClusterPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (p *PrivateLinksServerTransport) dispatchNewListByMongoClusterPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByMongoClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMongoClusterPager not implemented")}
	}
	newListByMongoClusterPager := p.newListByMongoClusterPager.get(req)
	if newListByMongoClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByMongoClusterPager(resourceGroupNameParam, mongoClusterNameParam, nil)
		newListByMongoClusterPager = &resp
		p.newListByMongoClusterPager.add(req, newListByMongoClusterPager)
		server.PagerResponderInjectNextLinks(newListByMongoClusterPager, req, func(page *armmongocluster.PrivateLinksClientListByMongoClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMongoClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByMongoClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMongoClusterPager) {
		p.newListByMongoClusterPager.remove(req)
	}
	return resp, nil
}
