// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Routing Status of routing
//
// +k8s:deepcopy-gen=true
//
// swagger:model Routing
type Routing struct {

	// Datapath routing mode for cross-cluster connectivity
	// Enum: [Native Tunnel]
	InterHostRoutingMode string `json:"inter-host-routing-mode,omitempty"`

	// Datapath routing mode for connectivity within the host
	// Enum: [BPF Legacy]
	IntraHostRoutingMode string `json:"intra-host-routing-mode,omitempty"`

	// Tunnel protocol in use for cross-cluster connectivity
	TunnelProtocol string `json:"tunnel-protocol,omitempty"`
}

// Validate validates this routing
func (m *Routing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterHostRoutingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntraHostRoutingMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var routingTypeInterHostRoutingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Native","Tunnel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routingTypeInterHostRoutingModePropEnum = append(routingTypeInterHostRoutingModePropEnum, v)
	}
}

const (

	// RoutingInterHostRoutingModeNative captures enum value "Native"
	RoutingInterHostRoutingModeNative string = "Native"

	// RoutingInterHostRoutingModeTunnel captures enum value "Tunnel"
	RoutingInterHostRoutingModeTunnel string = "Tunnel"
)

// prop value enum
func (m *Routing) validateInterHostRoutingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routingTypeInterHostRoutingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Routing) validateInterHostRoutingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.InterHostRoutingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateInterHostRoutingModeEnum("inter-host-routing-mode", "body", m.InterHostRoutingMode); err != nil {
		return err
	}

	return nil
}

var routingTypeIntraHostRoutingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BPF","Legacy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routingTypeIntraHostRoutingModePropEnum = append(routingTypeIntraHostRoutingModePropEnum, v)
	}
}

const (

	// RoutingIntraHostRoutingModeBPF captures enum value "BPF"
	RoutingIntraHostRoutingModeBPF string = "BPF"

	// RoutingIntraHostRoutingModeLegacy captures enum value "Legacy"
	RoutingIntraHostRoutingModeLegacy string = "Legacy"
)

// prop value enum
func (m *Routing) validateIntraHostRoutingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routingTypeIntraHostRoutingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Routing) validateIntraHostRoutingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.IntraHostRoutingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateIntraHostRoutingModeEnum("intra-host-routing-mode", "body", m.IntraHostRoutingMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this routing based on context it is used
func (m *Routing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Routing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Routing) UnmarshalBinary(b []byte) error {
	var res Routing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}