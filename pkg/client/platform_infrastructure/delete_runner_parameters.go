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

// NewDeleteRunnerParams creates a new DeleteRunnerParams object
// with the default values initialized.
func NewDeleteRunnerParams() *DeleteRunnerParams {
	var ()
	return &DeleteRunnerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunnerParamsWithTimeout creates a new DeleteRunnerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRunnerParamsWithTimeout(timeout time.Duration) *DeleteRunnerParams {
	var ()
	return &DeleteRunnerParams{

		timeout: timeout,
	}
}

// NewDeleteRunnerParamsWithContext creates a new DeleteRunnerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRunnerParamsWithContext(ctx context.Context) *DeleteRunnerParams {
	var ()
	return &DeleteRunnerParams{

		Context: ctx,
	}
}

// NewDeleteRunnerParamsWithHTTPClient creates a new DeleteRunnerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRunnerParamsWithHTTPClient(client *http.Client) *DeleteRunnerParams {
	var ()
	return &DeleteRunnerParams{
		HTTPClient: client,
	}
}

/*DeleteRunnerParams contains all the parameters to send to the API endpoint
for the delete runner operation typically these are written to a http.Request
*/
type DeleteRunnerParams struct {

	/*RunnerID
	  The identifier for the runner

	*/
	RunnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete runner params
func (o *DeleteRunnerParams) WithTimeout(timeout time.Duration) *DeleteRunnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runner params
func (o *DeleteRunnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runner params
func (o *DeleteRunnerParams) WithContext(ctx context.Context) *DeleteRunnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runner params
func (o *DeleteRunnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runner params
func (o *DeleteRunnerParams) WithHTTPClient(client *http.Client) *DeleteRunnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runner params
func (o *DeleteRunnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunnerID adds the runnerID to the delete runner params
func (o *DeleteRunnerParams) WithRunnerID(runnerID string) *DeleteRunnerParams {
	o.SetRunnerID(runnerID)
	return o
}

// SetRunnerID adds the runnerId to the delete runner params
func (o *DeleteRunnerParams) SetRunnerID(runnerID string) {
	o.RunnerID = runnerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param runner_id
	if err := r.SetPathParam("runner_id", o.RunnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}