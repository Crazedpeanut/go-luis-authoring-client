// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PatternUpdateObject pattern update object
// swagger:model PatternUpdateObject
type PatternUpdateObject struct {

	// is active
	IsActive *bool `json:"IsActive,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// pattern
	Pattern string `json:"Pattern,omitempty"`
}

// Validate validates this pattern update object
func (m *PatternUpdateObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatternUpdateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatternUpdateObject) UnmarshalBinary(b []byte) error {
	var res PatternUpdateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
