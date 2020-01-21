// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JSONSubClosedList JSON sub closed list
// swagger:model JSONSubClosedList
type JSONSubClosedList struct {

	// canonical form
	CanonicalForm string `json:"CanonicalForm,omitempty"`

	// list
	List []string `json:"List"`
}

// Validate validates this JSON sub closed list
func (m *JSONSubClosedList) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSONSubClosedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONSubClosedList) UnmarshalBinary(b []byte) error {
	var res JSONSubClosedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
