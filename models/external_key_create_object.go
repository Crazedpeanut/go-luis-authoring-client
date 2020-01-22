// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExternalKeyCreateObject external key create object
// swagger:model ExternalKeyCreateObject
type ExternalKeyCreateObject struct {

	// type
	Type string `json:"Type,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this external key create object
func (m *ExternalKeyCreateObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalKeyCreateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalKeyCreateObject) UnmarshalBinary(b []byte) error {
	var res ExternalKeyCreateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}