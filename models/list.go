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

// List list
//
// swagger:model List
type List struct {

	// Items limitation. Max 1000
	// Required: true
	Limit *int64 `json:"limit"`

	// The current page
	// Required: true
	Page *int64 `json:"page"`

	// Max possible pages
	// Required: true
	Pages *int64 `json:"pages"`

	// Total Items
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this list
func (m *List) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *List) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *List) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

func (m *List) validatePages(formats strfmt.Registry) error {

	if err := validate.Required("pages", "body", m.Pages); err != nil {
		return err
	}

	return nil
}

func (m *List) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *List) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *List) UnmarshalBinary(b []byte) error {
	var res List
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
