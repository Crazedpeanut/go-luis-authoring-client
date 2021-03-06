// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HierarchicalModelCreateObject hierarchical model create object
// swagger:model HierarchicalModelCreateObject
type HierarchicalModelCreateObject struct {

	// children
	Children []string `json:"Children"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this hierarchical model create object
func (m *HierarchicalModelCreateObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HierarchicalModelCreateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HierarchicalModelCreateObject) UnmarshalBinary(b []byte) error {
	var res HierarchicalModelCreateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
