// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExampleLabelObject example label object
// swagger:model ExampleLabelObject
type ExampleLabelObject struct {

	// entity labels
	EntityLabels []*EntityLabelObject `json:"EntityLabels"`

	// intent name
	IntentName string `json:"IntentName,omitempty"`

	// text
	Text string `json:"Text,omitempty"`
}

// Validate validates this example label object
func (m *ExampleLabelObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExampleLabelObject) validateEntityLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityLabels) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityLabels); i++ {
		if swag.IsZero(m.EntityLabels[i]) { // not required
			continue
		}

		if m.EntityLabels[i] != nil {
			if err := m.EntityLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EntityLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExampleLabelObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExampleLabelObject) UnmarshalBinary(b []byte) error {
	var res ExampleLabelObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
