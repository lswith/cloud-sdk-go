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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Extension An API extension. It represents clusters' plugins and bundles
// swagger:model Extension
type Extension struct {

	// The extension type.
	// Required: true
	// Enum: [plugin bundle]
	ExtensionType *string `json:"extension_type"`

	// The extension ID
	// Required: true
	ID *string `json:"id"`

	// The extension name.
	// Required: true
	Name *string `json:"name"`

	// The extension URL.
	URL string `json:"url,omitempty"`

	// The Elasticsearch version.
	Version string `json:"version,omitempty"`
}

// Validate validates this extension
func (m *Extension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var extensionTypeExtensionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["plugin","bundle"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extensionTypeExtensionTypePropEnum = append(extensionTypeExtensionTypePropEnum, v)
	}
}

const (

	// ExtensionExtensionTypePlugin captures enum value "plugin"
	ExtensionExtensionTypePlugin string = "plugin"

	// ExtensionExtensionTypeBundle captures enum value "bundle"
	ExtensionExtensionTypeBundle string = "bundle"
)

// prop value enum
func (m *Extension) validateExtensionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, extensionTypeExtensionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Extension) validateExtensionType(formats strfmt.Registry) error {

	if err := validate.Required("extension_type", "body", m.ExtensionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateExtensionTypeEnum("extension_type", "body", *m.ExtensionType); err != nil {
		return err
	}

	return nil
}

func (m *Extension) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Extension) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Extension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Extension) UnmarshalBinary(b []byte) error {
	var res Extension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}