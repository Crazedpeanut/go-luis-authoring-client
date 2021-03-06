// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationCreateObject application create object
// swagger:model ApplicationCreateObject
type ApplicationCreateObject struct {

	// culture
	Culture string `json:"Culture,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// domain
	Domain string `json:"Domain,omitempty"`

	// initial version Id
	InitialVersionID string `json:"InitialVersionId,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// usage scenario
	UsageScenario string `json:"UsageScenario,omitempty"`
}

// Validate validates this application create object
func (m *ApplicationCreateObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationCreateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationCreateObject) UnmarshalBinary(b []byte) error {
	var res ApplicationCreateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
