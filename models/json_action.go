// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JSONAction JSON action
// swagger:model JSONAction
type JSONAction struct {

	// action name
	ActionName string `json:"actionName,omitempty"`

	// action parameters
	ActionParameters []*JSONActionParam `json:"actionParameters"`

	// channel
	Channel *Channel `json:"channel,omitempty"`

	// intent name
	IntentName string `json:"intentName,omitempty"`
}

// Validate validates this JSON action
func (m *JSONAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONAction) validateActionParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.ActionParameters); i++ {
		if swag.IsZero(m.ActionParameters[i]) { // not required
			continue
		}

		if m.ActionParameters[i] != nil {
			if err := m.ActionParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actionParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JSONAction) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONAction) UnmarshalBinary(b []byte) error {
	var res JSONAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}