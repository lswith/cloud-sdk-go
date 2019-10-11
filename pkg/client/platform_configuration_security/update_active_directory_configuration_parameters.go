// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateActiveDirectoryConfigurationParams creates a new UpdateActiveDirectoryConfigurationParams object
// with the default values initialized.
func NewUpdateActiveDirectoryConfigurationParams() *UpdateActiveDirectoryConfigurationParams {
	var ()
	return &UpdateActiveDirectoryConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActiveDirectoryConfigurationParamsWithTimeout creates a new UpdateActiveDirectoryConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateActiveDirectoryConfigurationParamsWithTimeout(timeout time.Duration) *UpdateActiveDirectoryConfigurationParams {
	var ()
	return &UpdateActiveDirectoryConfigurationParams{

		timeout: timeout,
	}
}

// NewUpdateActiveDirectoryConfigurationParamsWithContext creates a new UpdateActiveDirectoryConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateActiveDirectoryConfigurationParamsWithContext(ctx context.Context) *UpdateActiveDirectoryConfigurationParams {
	var ()
	return &UpdateActiveDirectoryConfigurationParams{

		Context: ctx,
	}
}

// NewUpdateActiveDirectoryConfigurationParamsWithHTTPClient creates a new UpdateActiveDirectoryConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateActiveDirectoryConfigurationParamsWithHTTPClient(client *http.Client) *UpdateActiveDirectoryConfigurationParams {
	var ()
	return &UpdateActiveDirectoryConfigurationParams{
		HTTPClient: client,
	}
}

/*UpdateActiveDirectoryConfigurationParams contains all the parameters to send to the API endpoint
for the update active directory configuration operation typically these are written to a http.Request
*/
type UpdateActiveDirectoryConfigurationParams struct {

	/*Body
	  The Active Directory configuration

	*/
	Body *models.ActiveDirectorySettings
	/*RealmID
	  The Elasticsearch Security realm identifier.

	*/
	RealmID string
	/*Version
	  When specified, checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) WithTimeout(timeout time.Duration) *UpdateActiveDirectoryConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) WithContext(ctx context.Context) *UpdateActiveDirectoryConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) WithHTTPClient(client *http.Client) *UpdateActiveDirectoryConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) WithBody(body *models.ActiveDirectorySettings) *UpdateActiveDirectoryConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) SetBody(body *models.ActiveDirectorySettings) {
	o.Body = body
}

// WithRealmID adds the realmID to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) WithRealmID(realmID string) *UpdateActiveDirectoryConfigurationParams {
	o.SetRealmID(realmID)
	return o
}

// SetRealmID adds the realmId to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) SetRealmID(realmID string) {
	o.RealmID = realmID
}

// WithVersion adds the version to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) WithVersion(version *int64) *UpdateActiveDirectoryConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update active directory configuration params
func (o *UpdateActiveDirectoryConfigurationParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActiveDirectoryConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param realm_id
	if err := r.SetPathParam("realm_id", o.RealmID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}