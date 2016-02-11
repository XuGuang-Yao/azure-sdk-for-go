package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// VirtualNetworkGatewayConnectionsClient is the the Microsoft Azure Network
// management API provides a RESTful set of web services that interact with
// Microsoft Azure Networks service to manage your network resrources. The
// API has entities that capture the relationship between an end user and the
// Microsoft Azure Networks service.
type VirtualNetworkGatewayConnectionsClient struct {
	ManagementClient
}

// NewVirtualNetworkGatewayConnectionsClient creates an instance of the
// VirtualNetworkGatewayConnectionsClient client.
func NewVirtualNetworkGatewayConnectionsClient(subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return NewVirtualNetworkGatewayConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkGatewayConnectionsClientWithBaseURI creates an instance of
// the VirtualNetworkGatewayConnectionsClient client.
func NewVirtualNetworkGatewayConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return VirtualNetworkGatewayConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put VirtualNetworkGatewayConnection operation
// creates/updates a virtual network gateway connection in the specified
// resource group through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayConnectionName is the name of the virtual network
// gateway conenction. parameters is parameters supplied to the Begin Create
// or update Virtual Network Gateway connection operation through Network
// resource provider.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdate(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (result autorest.Response, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   url.QueryEscape(resourceGroupName),
		"subscriptionId":                      url.QueryEscape(client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": url.QueryEscape(virtualNetworkGatewayConnectionName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	resp, err := client.Send(req)
	if err == nil && azure.ResponseIsLongRunning(resp) {
		req, err := azure.NewAsyncPollingRequest(resp, client.Client)
		if err == nil {
			resp, err = autorest.SendWithSender(client, req,
				azure.WithAsyncPolling(autorest.DefaultPollingDelay))
		}
	}
	return resp, err
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result = autorest.Response{Response: resp}
	return
}

// Delete the Delete VirtualNetworkGatewayConnection operation deletes the
// specifed virtual network Gateway connection through Network resource
// provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayConnectionName is the name of the virtual network
// gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Delete(resourceGroupName string, virtualNetworkGatewayConnectionName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkGatewayConnectionsClient) DeletePreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   url.QueryEscape(resourceGroupName),
		"subscriptionId":                      url.QueryEscape(client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": url.QueryEscape(virtualNetworkGatewayConnectionName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	resp, err := client.Send(req)
	if err == nil && azure.ResponseIsLongRunning(resp) {
		req, err := azure.NewAsyncPollingRequest(resp, client.Client)
		if err == nil {
			resp, err = autorest.SendWithSender(client, req,
				azure.WithAsyncPolling(autorest.DefaultPollingDelay))
		}
	}
	return resp, err
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result = autorest.Response{Response: resp}
	return
}

// Get the Get VirtualNetworkGatewayConnection operation retrieves information
// about the specified virtual network gateway connection through Network
// resource provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayConnectionName is the name of the virtual network
// gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Get(resourceGroupName string, virtualNetworkGatewayConnectionName string) (result VirtualNetworkGatewayConnection, ae error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkGatewayConnectionsClient) GetPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   url.QueryEscape(resourceGroupName),
		"subscriptionId":                      url.QueryEscape(client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": url.QueryEscape(virtualNetworkGatewayConnectionName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) GetResponder(resp *http.Response) (result VirtualNetworkGatewayConnection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSharedKey the Get VirtualNetworkGatewayConnectionSharedKey operation
// retrieves information about the specified virtual network gateway
// connection shared key through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// connectionSharedKeyName is the virtual network gateway connection shared
// key name.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKey(resourceGroupName string, connectionSharedKeyName string) (result ConnectionSharedKeyResult, ae error) {
	req, err := client.GetSharedKeyPreparer(resourceGroupName, connectionSharedKeyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "GetSharedKey", nil, "Failure preparing request")
	}

	resp, err := client.GetSharedKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "GetSharedKey", resp, "Failure sending request")
	}

	result, err = client.GetSharedKeyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "GetSharedKey", resp, "Failure responding to request")
	}

	return
}

// GetSharedKeyPreparer prepares the GetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeyPreparer(resourceGroupName string, connectionSharedKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionSharedKeyName": url.QueryEscape(connectionSharedKeyName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{connectionSharedKeyName}/sharedkey"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSharedKeySender sends the GetSharedKey request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// GetSharedKeyResponder handles the response to the GetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeyResponder(resp *http.Response) (result ConnectionSharedKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List VirtualNetworkGatewayConnections operation retrieves all the
// virtual network gateways connections created.
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworkGatewayConnectionsClient) List(resourceGroupName string) (result VirtualNetworkGatewayConnectionListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworkGatewayConnectionsClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) ListResponder(resp *http.Response) (result VirtualNetworkGatewayConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client VirtualNetworkGatewayConnectionsClient) ListNextResults(lastResults VirtualNetworkGatewayConnectionListResult) (result VirtualNetworkGatewayConnectionListResult, ae error) {
	req, err := lastResults.VirtualNetworkGatewayConnectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ResetSharedKey the VirtualNetworkGatewayConnectionResetSharedKey operation
// resets the virtual network gateway connection shared key for passed
// virtual network gateway connection in the specified resource group through
// Network resource provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayConnectionName is the virtual network gateway
// connection reset shared key Name. parameters is parameters supplied to the
// Begin Reset Virtual Network Gateway connection shared key operation
// through Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKey(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (result autorest.Response, ae error) {
	req, err := client.ResetSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", nil, "Failure preparing request")
	}

	resp, err := client.ResetSharedKeySender(req)
	if err != nil {
		result = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", resp, "Failure sending request")
	}

	result, err = client.ResetSharedKeyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", resp, "Failure responding to request")
	}

	return
}

// ResetSharedKeyPreparer prepares the ResetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   url.QueryEscape(resourceGroupName),
		"subscriptionId":                      url.QueryEscape(client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": url.QueryEscape(virtualNetworkGatewayConnectionName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ResetSharedKeySender sends the ResetSharedKey request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeySender(req *http.Request) (*http.Response, error) {
	resp, err := client.Send(req)
	if err == nil && azure.ResponseIsLongRunning(resp) {
		req, err := azure.NewAsyncPollingRequest(resp, client.Client)
		if err == nil {
			resp, err = autorest.SendWithSender(client, req,
				azure.WithAsyncPolling(autorest.DefaultPollingDelay))
		}
	}
	return resp, err
}

// ResetSharedKeyResponder handles the response to the ResetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result = autorest.Response{Response: resp}
	return
}

// SetSharedKey the Put VirtualNetworkGatewayConnectionSharedKey operation
// sets the virtual network gateway connection shared key for passed virtual
// network gateway connection in the specified resource group through Network
// resource provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayConnectionName is the virtual network gateway
// connection name. parameters is parameters supplied to the Begin Set
// Virtual Network Gateway conection Shared key operation throughNetwork
// resource provider.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKey(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (result autorest.Response, ae error) {
	req, err := client.SetSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "SetSharedKey", nil, "Failure preparing request")
	}

	resp, err := client.SetSharedKeySender(req)
	if err != nil {
		result = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "SetSharedKey", resp, "Failure sending request")
	}

	result, err = client.SetSharedKeyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworkGatewayConnectionsClient", "SetSharedKey", resp, "Failure responding to request")
	}

	return
}

// SetSharedKeyPreparer prepares the SetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   url.QueryEscape(resourceGroupName),
		"subscriptionId":                      url.QueryEscape(client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": url.QueryEscape(virtualNetworkGatewayConnectionName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// SetSharedKeySender sends the SetSharedKey request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeySender(req *http.Request) (*http.Response, error) {
	resp, err := client.Send(req)
	if err == nil && azure.ResponseIsLongRunning(resp) {
		req, err := azure.NewAsyncPollingRequest(resp, client.Client)
		if err == nil {
			resp, err = autorest.SendWithSender(client, req,
				azure.WithAsyncPolling(autorest.DefaultPollingDelay))
		}
	}
	return resp, err
}

// SetSharedKeyResponder handles the response to the SetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result = autorest.Response{Response: resp}
	return
}
