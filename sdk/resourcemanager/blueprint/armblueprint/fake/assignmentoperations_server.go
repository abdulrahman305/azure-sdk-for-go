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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
	"net/http"
	"net/url"
	"regexp"
)

// AssignmentOperationsServer is a fake server for instances of the armblueprint.AssignmentOperationsClient type.
type AssignmentOperationsServer struct {
	// Get is the fake for method AssignmentOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceScope string, assignmentName string, assignmentOperationName string, options *armblueprint.AssignmentOperationsClientGetOptions) (resp azfake.Responder[armblueprint.AssignmentOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AssignmentOperationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceScope string, assignmentName string, options *armblueprint.AssignmentOperationsClientListOptions) (resp azfake.PagerResponder[armblueprint.AssignmentOperationsClientListResponse])
}

// NewAssignmentOperationsServerTransport creates a new instance of AssignmentOperationsServerTransport with the provided implementation.
// The returned AssignmentOperationsServerTransport instance is connected to an instance of armblueprint.AssignmentOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssignmentOperationsServerTransport(srv *AssignmentOperationsServer) *AssignmentOperationsServerTransport {
	return &AssignmentOperationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armblueprint.AssignmentOperationsClientListResponse]](),
	}
}

// AssignmentOperationsServerTransport connects instances of armblueprint.AssignmentOperationsClient to instances of AssignmentOperationsServer.
// Don't use this type directly, use NewAssignmentOperationsServerTransport instead.
type AssignmentOperationsServerTransport struct {
	srv          *AssignmentOperationsServer
	newListPager *tracker[azfake.PagerResponder[armblueprint.AssignmentOperationsClientListResponse]]
}

// Do implements the policy.Transporter interface for AssignmentOperationsServerTransport.
func (a *AssignmentOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AssignmentOperationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AssignmentOperationsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AssignmentOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceScope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Blueprint/blueprintAssignments/(?P<assignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assignmentOperations/(?P<assignmentOperationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceScopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceScope")])
	if err != nil {
		return nil, err
	}
	assignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assignmentName")])
	if err != nil {
		return nil, err
	}
	assignmentOperationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assignmentOperationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceScopeParam, assignmentNameParam, assignmentOperationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssignmentOperation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssignmentOperationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceScope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Blueprint/blueprintAssignments/(?P<assignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assignmentOperations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceScopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceScope")])
		if err != nil {
			return nil, err
		}
		assignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assignmentName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceScopeParam, assignmentNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armblueprint.AssignmentOperationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
