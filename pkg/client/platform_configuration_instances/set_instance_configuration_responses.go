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

package platform_configuration_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetInstanceConfigurationReader is a Reader for the SetInstanceConfiguration structure.
type SetInstanceConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetInstanceConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetInstanceConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewSetInstanceConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetInstanceConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetInstanceConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetInstanceConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewSetInstanceConfigurationRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetInstanceConfigurationOK creates a SetInstanceConfigurationOK with default headers values
func NewSetInstanceConfigurationOK() *SetInstanceConfigurationOK {
	return &SetInstanceConfigurationOK{}
}

/*SetInstanceConfigurationOK handles this case with default header values.

The instance configuration was updated successfully.
*/
type SetInstanceConfigurationOK struct {
	Payload *models.IDResponse
}

func (o *SetInstanceConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/instances/{id}][%d] setInstanceConfigurationOK  %+v", 200, o.Payload)
}

func (o *SetInstanceConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetInstanceConfigurationCreated creates a SetInstanceConfigurationCreated with default headers values
func NewSetInstanceConfigurationCreated() *SetInstanceConfigurationCreated {
	return &SetInstanceConfigurationCreated{}
}

/*SetInstanceConfigurationCreated handles this case with default header values.

The instance configuration was created successfully.
*/
type SetInstanceConfigurationCreated struct {
	Payload *models.IDResponse
}

func (o *SetInstanceConfigurationCreated) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/instances/{id}][%d] setInstanceConfigurationCreated  %+v", 201, o.Payload)
}

func (o *SetInstanceConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetInstanceConfigurationBadRequest creates a SetInstanceConfigurationBadRequest with default headers values
func NewSetInstanceConfigurationBadRequest() *SetInstanceConfigurationBadRequest {
	return &SetInstanceConfigurationBadRequest{}
}

/*SetInstanceConfigurationBadRequest handles this case with default header values.

cluster_type in the InstanceConfiguration model is invalid (code: 'configuration.invalid_cluster_type') or the ZooKeeper operation failed due to bad version, etc. (code: 'configuration.instance_configuration_update_failed') or the id in the InstanceConfiguration model is reserved/prohibited (code: 'configuration.configuration_id_reserved') or the discrete_sizes in the InstanceConfiguration model is invalid (code: 'configuration.invalid_discrete_sizes') or the metadata in the InstanceConfiguration model has empty keys or values (code: 'configuration.bad_meta_data')
*/
type SetInstanceConfigurationBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *SetInstanceConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/instances/{id}][%d] setInstanceConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *SetInstanceConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetInstanceConfigurationForbidden creates a SetInstanceConfigurationForbidden with default headers values
func NewSetInstanceConfigurationForbidden() *SetInstanceConfigurationForbidden {
	return &SetInstanceConfigurationForbidden{}
}

/*SetInstanceConfigurationForbidden handles this case with default header values.

system_owned or deleted_on cannot be set externally (code: 'configuration.system_owned')
*/
type SetInstanceConfigurationForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *SetInstanceConfigurationForbidden) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/instances/{id}][%d] setInstanceConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *SetInstanceConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetInstanceConfigurationNotFound creates a SetInstanceConfigurationNotFound with default headers values
func NewSetInstanceConfigurationNotFound() *SetInstanceConfigurationNotFound {
	return &SetInstanceConfigurationNotFound{}
}

/*SetInstanceConfigurationNotFound handles this case with default header values.

Instance configuration specified by {id} cannot be found or the operation failed (code: 'configuration.instance_configuration_not_found')
*/
type SetInstanceConfigurationNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *SetInstanceConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/instances/{id}][%d] setInstanceConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *SetInstanceConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetInstanceConfigurationRetryWith creates a SetInstanceConfigurationRetryWith with default headers values
func NewSetInstanceConfigurationRetryWith() *SetInstanceConfigurationRetryWith {
	return &SetInstanceConfigurationRetryWith{}
}

/*SetInstanceConfigurationRetryWith handles this case with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type SetInstanceConfigurationRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *SetInstanceConfigurationRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/instances/{id}][%d] setInstanceConfigurationRetryWith  %+v", 449, o.Payload)
}

func (o *SetInstanceConfigurationRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}