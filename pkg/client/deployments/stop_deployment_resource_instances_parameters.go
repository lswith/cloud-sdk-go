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

package deployments

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
)

// NewStopDeploymentResourceInstancesParams creates a new StopDeploymentResourceInstancesParams object
// with the default values initialized.
func NewStopDeploymentResourceInstancesParams() *StopDeploymentResourceInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopDeploymentResourceInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStopDeploymentResourceInstancesParamsWithTimeout creates a new StopDeploymentResourceInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopDeploymentResourceInstancesParamsWithTimeout(timeout time.Duration) *StopDeploymentResourceInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopDeploymentResourceInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,

		timeout: timeout,
	}
}

// NewStopDeploymentResourceInstancesParamsWithContext creates a new StopDeploymentResourceInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopDeploymentResourceInstancesParamsWithContext(ctx context.Context) *StopDeploymentResourceInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopDeploymentResourceInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,

		Context: ctx,
	}
}

// NewStopDeploymentResourceInstancesParamsWithHTTPClient creates a new StopDeploymentResourceInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopDeploymentResourceInstancesParamsWithHTTPClient(client *http.Client) *StopDeploymentResourceInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopDeploymentResourceInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,
		HTTPClient:    client,
	}
}

/*StopDeploymentResourceInstancesParams contains all the parameters to send to the API endpoint
for the stop deployment resource instances operation typically these are written to a http.Request
*/
type StopDeploymentResourceInstancesParams struct {

	/*DeploymentID
	  Identifier for the Deployment

	*/
	DeploymentID string
	/*IgnoreMissing
	  If true and the instance does not exist then quietly proceed to the next instance, otherwise treated as an error

	*/
	IgnoreMissing *bool
	/*InstanceIds
	  Comma-delimited list of instance identifiers of the Resource

	*/
	InstanceIds []string
	/*RefID
	  User-specified RefId for the Resource

	*/
	RefID string
	/*ResourceKind
	  The kind of resource (one of elasticsearch, kibana or apm)

	*/
	ResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithTimeout(timeout time.Duration) *StopDeploymentResourceInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithContext(ctx context.Context) *StopDeploymentResourceInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithHTTPClient(client *http.Client) *StopDeploymentResourceInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithDeploymentID(deploymentID string) *StopDeploymentResourceInstancesParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithIgnoreMissing adds the ignoreMissing to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithIgnoreMissing(ignoreMissing *bool) *StopDeploymentResourceInstancesParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithInstanceIds(instanceIds []string) *StopDeploymentResourceInstancesParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WithRefID adds the refID to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithRefID(refID string) *StopDeploymentResourceInstancesParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) WithResourceKind(resourceKind string) *StopDeploymentResourceInstancesParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the stop deployment resource instances params
func (o *StopDeploymentResourceInstancesParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *StopDeploymentResourceInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool
		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {
			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}

	}

	valuesInstanceIds := o.InstanceIds

	joinedInstanceIds := swag.JoinByFormat(valuesInstanceIds, "csv")
	// path array param instance_ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedInstanceIds) > 0 {
		if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
			return err
		}
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	// path param resource_kind
	if err := r.SetPathParam("resource_kind", o.ResourceKind); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}