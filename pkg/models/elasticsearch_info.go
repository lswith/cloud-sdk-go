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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchInfo Information about the Elasticsearch cluster.
// swagger:model ElasticsearchInfo
type ElasticsearchInfo struct {

	// blocking issues
	BlockingIssues *ElasticsearchClusterBlockingIssues `json:"blocking_issues,omitempty"`

	// Whether the Elasticsearch cluster is healthy (check the sub-objects for more details if not)
	// Required: true
	Healthy *bool `json:"healthy"`

	// master info
	// Required: true
	MasterInfo *ElasticsearchMasterInfo `json:"master_info"`

	// shard info
	// Required: true
	ShardInfo *ElasticsearchShardsInfo `json:"shard_info"`
}

// Validate validates this elasticsearch info
func (m *ElasticsearchInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockingIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShardInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchInfo) validateBlockingIssues(formats strfmt.Registry) error {

	if swag.IsZero(m.BlockingIssues) { // not required
		return nil
	}

	if m.BlockingIssues != nil {
		if err := m.BlockingIssues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blocking_issues")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchInfo) validateMasterInfo(formats strfmt.Registry) error {

	if err := validate.Required("master_info", "body", m.MasterInfo); err != nil {
		return err
	}

	if m.MasterInfo != nil {
		if err := m.MasterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master_info")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchInfo) validateShardInfo(formats strfmt.Registry) error {

	if err := validate.Required("shard_info", "body", m.ShardInfo); err != nil {
		return err
	}

	if m.ShardInfo != nil {
		if err := m.ShardInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shard_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}