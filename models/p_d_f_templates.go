// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PDFTemplates p d f templates
//
// swagger:model PDFTemplates
type PDFTemplates struct {

	// items
	Items []*PDFTemplate `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PDFTemplates) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Items []*PDFTemplate `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Items = dataAO0.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PDFTemplates) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Items []*PDFTemplate `json:"items"`
	}

	dataAO0.Items = m.Items

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this p d f templates
func (m *PDFTemplates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PDFTemplates) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PDFTemplates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PDFTemplates) UnmarshalBinary(b []byte) error {
	var res PDFTemplates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}