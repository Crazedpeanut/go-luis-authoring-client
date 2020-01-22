// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PrebuiltDomainModelCreateObject prebuilt domain model create object
// swagger:model PrebuiltDomainModelCreateObject
type PrebuiltDomainModelCreateObject struct {

	// domain name
	DomainName string `json:"DomainName,omitempty"`

	// model name
	ModelName string `json:"ModelName,omitempty"`
}

// Validate validates this prebuilt domain model create object
func (m *PrebuiltDomainModelCreateObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrebuiltDomainModelCreateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrebuiltDomainModelCreateObject) UnmarshalBinary(b []byte) error {
	var res PrebuiltDomainModelCreateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}