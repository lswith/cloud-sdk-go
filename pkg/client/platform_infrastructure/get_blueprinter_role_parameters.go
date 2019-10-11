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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBlueprinterRoleParams creates a new GetBlueprinterRoleParams object
// with the default values initialized.
func NewGetBlueprinterRoleParams() *GetBlueprinterRoleParams {
	var ()
	return &GetBlueprinterRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlueprinterRoleParamsWithTimeout creates a new GetBlueprinterRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlueprinterRoleParamsWithTimeout(timeout time.Duration) *GetBlueprinterRoleParams {
	var ()
	return &GetBlueprinterRoleParams{

		timeout: timeout,
	}
}

// NewGetBlueprinterRoleParamsWithContext creates a new GetBlueprinterRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlueprinterRoleParamsWithContext(ctx context.Context) *GetBlueprinterRoleParams {
	var ()
	return &GetBlueprinterRoleParams{

		Context: ctx,
	}
}

// NewGetBlueprinterRoleParamsWithHTTPClient creates a new GetBlueprinterRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlueprinterRoleParamsWithHTTPClient(client *http.Client) *GetBlueprinterRoleParams {
	var ()
	return &GetBlueprinterRoleParams{
		HTTPClient: client,
	}
}

/*GetBlueprinterRoleParams contains all the parameters to send to the API endpoint
for the get blueprinter role operation typically these are written to a http.Request
*/
type GetBlueprinterRoleParams struct {

	/*BlueprinterRoleID
	  User-specified Blueprinter role ID.

	*/
	BlueprinterRoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get blueprinter role params
func (o *GetBlueprinterRoleParams) WithTimeout(timeout time.Duration) *GetBlueprinterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blueprinter role params
func (o *GetBlueprinterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blueprinter role params
func (o *GetBlueprinterRoleParams) WithContext(ctx context.Context) *GetBlueprinterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blueprinter role params
func (o *GetBlueprinterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blueprinter role params
func (o *GetBlueprinterRoleParams) WithHTTPClient(client *http.Client) *GetBlueprinterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blueprinter role params
func (o *GetBlueprinterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlueprinterRoleID adds the blueprinterRoleID to the get blueprinter role params
func (o *GetBlueprinterRoleParams) WithBlueprinterRoleID(blueprinterRoleID string) *GetBlueprinterRoleParams {
	o.SetBlueprinterRoleID(blueprinterRoleID)
	return o
}

// SetBlueprinterRoleID adds the blueprinterRoleId to the get blueprinter role params
func (o *GetBlueprinterRoleParams) SetBlueprinterRoleID(blueprinterRoleID string) {
	o.BlueprinterRoleID = blueprinterRoleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlueprinterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param blueprinter_role_id
	if err := r.SetPathParam("blueprinter_role_id", o.BlueprinterRoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}