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

// DiscountPosition discount position
//
// swagger:model DiscountPosition
type DiscountPosition struct {
	Discount

	// position id
	// Required: true
	PositionID *int64 `json:"position_id"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DiscountPosition) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Discount
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Discount = aO0

	// now for regular properties
	var propsDiscountPosition struct {
		PositionID *int64 `json:"position_id"`
	}
	if err := swag.ReadJSON(raw, &propsDiscountPosition); err != nil {
		return err
	}
	m.PositionID = propsDiscountPosition.PositionID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DiscountPosition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Discount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsDiscountPosition struct {
		PositionID *int64 `json:"position_id"`
	}
	propsDiscountPosition.PositionID = m.PositionID

	jsonDataPropsDiscountPosition, errDiscountPosition := swag.WriteJSON(propsDiscountPosition)
	if errDiscountPosition != nil {
		return nil, errDiscountPosition
	}
	_parts = append(_parts, jsonDataPropsDiscountPosition)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this discount position
func (m *DiscountPosition) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Discount
	if err := m.Discount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscountPosition) validatePositionID(formats strfmt.Registry) error {

	if err := validate.Required("position_id", "body", m.PositionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscountPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscountPosition) UnmarshalBinary(b []byte) error {
	var res DiscountPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
