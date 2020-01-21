// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JSONRegexFeature JSON regex feature
// swagger:model JSONRegexFeature
type JSONRegexFeature struct {

	// activated
	Activated bool `json:"activated,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`
}

// Validate validates this JSON regex feature
func (m *JSONRegexFeature) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSONRegexFeature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONRegexFeature) UnmarshalBinary(b []byte) error {
	var res JSONRegexFeature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
