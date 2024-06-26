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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// LongRunningBackupServer is a fake server for instances of the armmysqlflexibleservers.LongRunningBackupClient type.
type LongRunningBackupServer struct {
	// BeginCreate is the fake for method LongRunningBackupClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, serverName string, backupName string, options *armmysqlflexibleservers.LongRunningBackupClientBeginCreateOptions) (resp azfake.PollerResponder[armmysqlflexibleservers.LongRunningBackupClientCreateResponse], errResp azfake.ErrorResponder)
}

// NewLongRunningBackupServerTransport creates a new instance of LongRunningBackupServerTransport with the provided implementation.
// The returned LongRunningBackupServerTransport instance is connected to an instance of armmysqlflexibleservers.LongRunningBackupClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLongRunningBackupServerTransport(srv *LongRunningBackupServer) *LongRunningBackupServerTransport {
	return &LongRunningBackupServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armmysqlflexibleservers.LongRunningBackupClientCreateResponse]](),
	}
}

// LongRunningBackupServerTransport connects instances of armmysqlflexibleservers.LongRunningBackupClient to instances of LongRunningBackupServer.
// Don't use this type directly, use NewLongRunningBackupServerTransport instead.
type LongRunningBackupServerTransport struct {
	srv         *LongRunningBackupServer
	beginCreate *tracker[azfake.PollerResponder[armmysqlflexibleservers.LongRunningBackupClientCreateResponse]]
}

// Do implements the policy.Transporter interface for LongRunningBackupServerTransport.
func (l *LongRunningBackupServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LongRunningBackupClient.BeginCreate":
		resp, err = l.dispatchBeginCreate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LongRunningBackupServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := l.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupsV2/(?P<backupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmysqlflexibleservers.ServerBackupV2](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		backupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupName")])
		if err != nil {
			return nil, err
		}
		var options *armmysqlflexibleservers.LongRunningBackupClientBeginCreateOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armmysqlflexibleservers.LongRunningBackupClientBeginCreateOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := l.srv.BeginCreate(req.Context(), resourceGroupNameParam, serverNameParam, backupNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		l.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		l.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		l.beginCreate.remove(req)
	}

	return resp, nil
}
