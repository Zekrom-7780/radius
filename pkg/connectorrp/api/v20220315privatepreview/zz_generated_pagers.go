//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DaprInvokeHTTPRoutesListBySubscriptionPager provides operations for iterating over paged responses.
type DaprInvokeHTTPRoutesListBySubscriptionPager struct {
	client *DaprInvokeHTTPRoutesClient
	current DaprInvokeHTTPRoutesListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, DaprInvokeHTTPRoutesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DaprInvokeHTTPRoutesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DaprInvokeHTTPRoutesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DaprInvokeHTTPRouteList.NextLink == nil || len(*p.current.DaprInvokeHTTPRouteList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DaprInvokeHTTPRoutesListBySubscriptionResponse page.
func (p *DaprInvokeHTTPRoutesListBySubscriptionPager) PageResponse() DaprInvokeHTTPRoutesListBySubscriptionResponse {
	return p.current
}

// DaprInvokeHTTPRoutesListPager provides operations for iterating over paged responses.
type DaprInvokeHTTPRoutesListPager struct {
	client *DaprInvokeHTTPRoutesClient
	current DaprInvokeHTTPRoutesListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, DaprInvokeHTTPRoutesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DaprInvokeHTTPRoutesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DaprInvokeHTTPRoutesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DaprInvokeHTTPRouteList.NextLink == nil || len(*p.current.DaprInvokeHTTPRouteList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DaprInvokeHTTPRoutesListResponse page.
func (p *DaprInvokeHTTPRoutesListPager) PageResponse() DaprInvokeHTTPRoutesListResponse {
	return p.current
}

// DaprSecretStoresListBySubscriptionPager provides operations for iterating over paged responses.
type DaprSecretStoresListBySubscriptionPager struct {
	client *DaprSecretStoresClient
	current DaprSecretStoresListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, DaprSecretStoresListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DaprSecretStoresListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DaprSecretStoresListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DaprSecretStoreList.NextLink == nil || len(*p.current.DaprSecretStoreList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DaprSecretStoresListBySubscriptionResponse page.
func (p *DaprSecretStoresListBySubscriptionPager) PageResponse() DaprSecretStoresListBySubscriptionResponse {
	return p.current
}

// DaprSecretStoresListPager provides operations for iterating over paged responses.
type DaprSecretStoresListPager struct {
	client *DaprSecretStoresClient
	current DaprSecretStoresListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, DaprSecretStoresListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DaprSecretStoresListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DaprSecretStoresListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DaprSecretStoreList.NextLink == nil || len(*p.current.DaprSecretStoreList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DaprSecretStoresListResponse page.
func (p *DaprSecretStoresListPager) PageResponse() DaprSecretStoresListResponse {
	return p.current
}

// DaprStateStoresListBySubscriptionPager provides operations for iterating over paged responses.
type DaprStateStoresListBySubscriptionPager struct {
	client *DaprStateStoresClient
	current DaprStateStoresListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, DaprStateStoresListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DaprStateStoresListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DaprStateStoresListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DaprStateStoreList.NextLink == nil || len(*p.current.DaprStateStoreList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DaprStateStoresListBySubscriptionResponse page.
func (p *DaprStateStoresListBySubscriptionPager) PageResponse() DaprStateStoresListBySubscriptionResponse {
	return p.current
}

// DaprStateStoresListPager provides operations for iterating over paged responses.
type DaprStateStoresListPager struct {
	client *DaprStateStoresClient
	current DaprStateStoresListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, DaprStateStoresListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DaprStateStoresListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DaprStateStoresListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DaprStateStoreList.NextLink == nil || len(*p.current.DaprStateStoreList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DaprStateStoresListResponse page.
func (p *DaprStateStoresListPager) PageResponse() DaprStateStoresListResponse {
	return p.current
}

// MongoDatabasesListBySubscriptionPager provides operations for iterating over paged responses.
type MongoDatabasesListBySubscriptionPager struct {
	client *MongoDatabasesClient
	current MongoDatabasesListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, MongoDatabasesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MongoDatabasesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MongoDatabasesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MongoDatabaseList.NextLink == nil || len(*p.current.MongoDatabaseList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MongoDatabasesListBySubscriptionResponse page.
func (p *MongoDatabasesListBySubscriptionPager) PageResponse() MongoDatabasesListBySubscriptionResponse {
	return p.current
}

// MongoDatabasesListPager provides operations for iterating over paged responses.
type MongoDatabasesListPager struct {
	client *MongoDatabasesClient
	current MongoDatabasesListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, MongoDatabasesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MongoDatabasesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MongoDatabasesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MongoDatabaseList.NextLink == nil || len(*p.current.MongoDatabaseList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MongoDatabasesListResponse page.
func (p *MongoDatabasesListPager) PageResponse() MongoDatabasesListResponse {
	return p.current
}

// RabbitMQMessageQueuesListBySubscriptionPager provides operations for iterating over paged responses.
type RabbitMQMessageQueuesListBySubscriptionPager struct {
	client *RabbitMQMessageQueuesClient
	current RabbitMQMessageQueuesListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, RabbitMQMessageQueuesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RabbitMQMessageQueuesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RabbitMQMessageQueuesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RabbitMQMessageQueueList.NextLink == nil || len(*p.current.RabbitMQMessageQueueList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RabbitMQMessageQueuesListBySubscriptionResponse page.
func (p *RabbitMQMessageQueuesListBySubscriptionPager) PageResponse() RabbitMQMessageQueuesListBySubscriptionResponse {
	return p.current
}

// RabbitMQMessageQueuesListPager provides operations for iterating over paged responses.
type RabbitMQMessageQueuesListPager struct {
	client *RabbitMQMessageQueuesClient
	current RabbitMQMessageQueuesListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, RabbitMQMessageQueuesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RabbitMQMessageQueuesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RabbitMQMessageQueuesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RabbitMQMessageQueueList.NextLink == nil || len(*p.current.RabbitMQMessageQueueList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RabbitMQMessageQueuesListResponse page.
func (p *RabbitMQMessageQueuesListPager) PageResponse() RabbitMQMessageQueuesListResponse {
	return p.current
}

// RedisCachesListBySubscriptionPager provides operations for iterating over paged responses.
type RedisCachesListBySubscriptionPager struct {
	client *RedisCachesClient
	current RedisCachesListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, RedisCachesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RedisCachesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RedisCachesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RedisCacheList.NextLink == nil || len(*p.current.RedisCacheList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RedisCachesListBySubscriptionResponse page.
func (p *RedisCachesListBySubscriptionPager) PageResponse() RedisCachesListBySubscriptionResponse {
	return p.current
}

// RedisCachesListPager provides operations for iterating over paged responses.
type RedisCachesListPager struct {
	client *RedisCachesClient
	current RedisCachesListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, RedisCachesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RedisCachesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RedisCachesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RedisCacheList.NextLink == nil || len(*p.current.RedisCacheList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RedisCachesListResponse page.
func (p *RedisCachesListPager) PageResponse() RedisCachesListResponse {
	return p.current
}

// SQLDatabasesListBySubscriptionPager provides operations for iterating over paged responses.
type SQLDatabasesListBySubscriptionPager struct {
	client *SQLDatabasesClient
	current SQLDatabasesListBySubscriptionResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, SQLDatabasesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLDatabasesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLDatabasesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLDatabaseList.NextLink == nil || len(*p.current.SQLDatabaseList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLDatabasesListBySubscriptionResponse page.
func (p *SQLDatabasesListBySubscriptionPager) PageResponse() SQLDatabasesListBySubscriptionResponse {
	return p.current
}

// SQLDatabasesListPager provides operations for iterating over paged responses.
type SQLDatabasesListPager struct {
	client *SQLDatabasesClient
	current SQLDatabasesListResponse
	err error
	requester func(context.Context) (*policy.Request, error)
	advancer func(context.Context, SQLDatabasesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLDatabasesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLDatabasesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLDatabaseList.NextLink == nil || len(*p.current.SQLDatabaseList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.	client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLDatabasesListResponse page.
func (p *SQLDatabasesListPager) PageResponse() SQLDatabasesListResponse {
	return p.current
}
