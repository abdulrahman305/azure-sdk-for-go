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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
	"net/http"
	"net/url"
	"regexp"
)

// HardwareComponentGroupsServer is a fake server for instances of the armstorsimple8000series.HardwareComponentGroupsClient type.
type HardwareComponentGroupsServer struct {
	// BeginChangeControllerPowerState is the fake for method HardwareComponentGroupsClient.BeginChangeControllerPowerState
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginChangeControllerPowerState func(ctx context.Context, deviceName string, hardwareComponentGroupName string, resourceGroupName string, managerName string, parameters armstorsimple8000series.ControllerPowerStateChangeRequest, options *armstorsimple8000series.HardwareComponentGroupsClientBeginChangeControllerPowerStateOptions) (resp azfake.PollerResponder[armstorsimple8000series.HardwareComponentGroupsClientChangeControllerPowerStateResponse], errResp azfake.ErrorResponder)

	// NewListByDevicePager is the fake for method HardwareComponentGroupsClient.NewListByDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDevicePager func(deviceName string, resourceGroupName string, managerName string, options *armstorsimple8000series.HardwareComponentGroupsClientListByDeviceOptions) (resp azfake.PagerResponder[armstorsimple8000series.HardwareComponentGroupsClientListByDeviceResponse])
}

// NewHardwareComponentGroupsServerTransport creates a new instance of HardwareComponentGroupsServerTransport with the provided implementation.
// The returned HardwareComponentGroupsServerTransport instance is connected to an instance of armstorsimple8000series.HardwareComponentGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHardwareComponentGroupsServerTransport(srv *HardwareComponentGroupsServer) *HardwareComponentGroupsServerTransport {
	return &HardwareComponentGroupsServerTransport{
		srv:                             srv,
		beginChangeControllerPowerState: newTracker[azfake.PollerResponder[armstorsimple8000series.HardwareComponentGroupsClientChangeControllerPowerStateResponse]](),
		newListByDevicePager:            newTracker[azfake.PagerResponder[armstorsimple8000series.HardwareComponentGroupsClientListByDeviceResponse]](),
	}
}

// HardwareComponentGroupsServerTransport connects instances of armstorsimple8000series.HardwareComponentGroupsClient to instances of HardwareComponentGroupsServer.
// Don't use this type directly, use NewHardwareComponentGroupsServerTransport instead.
type HardwareComponentGroupsServerTransport struct {
	srv                             *HardwareComponentGroupsServer
	beginChangeControllerPowerState *tracker[azfake.PollerResponder[armstorsimple8000series.HardwareComponentGroupsClientChangeControllerPowerStateResponse]]
	newListByDevicePager            *tracker[azfake.PagerResponder[armstorsimple8000series.HardwareComponentGroupsClientListByDeviceResponse]]
}

// Do implements the policy.Transporter interface for HardwareComponentGroupsServerTransport.
func (h *HardwareComponentGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HardwareComponentGroupsClient.BeginChangeControllerPowerState":
		resp, err = h.dispatchBeginChangeControllerPowerState(req)
	case "HardwareComponentGroupsClient.NewListByDevicePager":
		resp, err = h.dispatchNewListByDevicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HardwareComponentGroupsServerTransport) dispatchBeginChangeControllerPowerState(req *http.Request) (*http.Response, error) {
	if h.srv.BeginChangeControllerPowerState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginChangeControllerPowerState not implemented")}
	}
	beginChangeControllerPowerState := h.beginChangeControllerPowerState.get(req)
	if beginChangeControllerPowerState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hardwareComponentGroups/(?P<hardwareComponentGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/changeControllerPowerState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple8000series.ControllerPowerStateChangeRequest](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		hardwareComponentGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hardwareComponentGroupName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginChangeControllerPowerState(req.Context(), deviceNameParam, hardwareComponentGroupNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginChangeControllerPowerState = &respr
		h.beginChangeControllerPowerState.add(req, beginChangeControllerPowerState)
	}

	resp, err := server.PollerResponderNext(beginChangeControllerPowerState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		h.beginChangeControllerPowerState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginChangeControllerPowerState) {
		h.beginChangeControllerPowerState.remove(req)
	}

	return resp, nil
}

func (h *HardwareComponentGroupsServerTransport) dispatchNewListByDevicePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDevicePager not implemented")}
	}
	newListByDevicePager := h.newListByDevicePager.get(req)
	if newListByDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hardwareComponentGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListByDevicePager(deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByDevicePager = &resp
		h.newListByDevicePager.add(req, newListByDevicePager)
	}
	resp, err := server.PagerResponderNext(newListByDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDevicePager) {
		h.newListByDevicePager.remove(req)
	}
	return resp, nil
}
