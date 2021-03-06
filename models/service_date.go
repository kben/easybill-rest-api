// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDate service date
//
// swagger:model ServiceDate
type ServiceDate struct {

	// date
	// Format: date
	Date *strfmt.Date `json:"date,omitempty"`

	// date from
	// Format: date
	DateFrom *strfmt.Date `json:"date_from,omitempty"`

	// date to
	// Format: date
	DateTo *strfmt.Date `json:"date_to,omitempty"`

	// text
	Text *string `json:"text,omitempty"`

	// With DEFAULT no other fields are required and this message will be printed: 'Invoice date coincides with the time of supply'.<br/> For SERVICE or DELIVERY exactly one of the following fields must be set: date, date_from and date_to or text.
	// Enum: [DEFAULT SERVICE DELIVERY]
	Type string `json:"type,omitempty"`
}

// Validate validates this service date
func (m *ServiceDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDate) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDate) validateDateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.DateFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("date_from", "body", "date", m.DateFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDate) validateDateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTo) { // not required
		return nil
	}

	if err := validate.FormatOf("date_to", "body", "date", m.DateTo.String(), formats); err != nil {
		return err
	}

	return nil
}

var serviceDateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","SERVICE","DELIVERY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDateTypeTypePropEnum = append(serviceDateTypeTypePropEnum, v)
	}
}

const (

	// ServiceDateTypeDEFAULT captures enum value "DEFAULT"
	ServiceDateTypeDEFAULT string = "DEFAULT"

	// ServiceDateTypeSERVICE captures enum value "SERVICE"
	ServiceDateTypeSERVICE string = "SERVICE"

	// ServiceDateTypeDELIVERY captures enum value "DELIVERY"
	ServiceDateTypeDELIVERY string = "DELIVERY"
)

// prop value enum
func (m *ServiceDate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDate) UnmarshalBinary(b []byte) error {
	var res ServiceDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
