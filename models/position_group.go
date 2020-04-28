// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PositionGroup position group
//
// swagger:model PositionGroup
type PositionGroup struct {

	// description
	Description *string `json:"description,omitempty"`

	// display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// login id
	// Read Only: true
	LoginID int64 `json:"login_id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// number
	// Required: true
	Number *string `json:"number"`
}

// Validate validates this position group
func (m *PositionGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PositionGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PositionGroup) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PositionGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PositionGroup) UnmarshalBinary(b []byte) error {
	var res PositionGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
