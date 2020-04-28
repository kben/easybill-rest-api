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

// Attachment If customer_id, project_id and document_id are null then attachment has global access from web ui.
//
// swagger:model Attachment
type Attachment struct {

	// created at
	// Read Only: true
	// Format: date
	CreatedAt strfmt.Date `json:"created_at,omitempty"`

	// customer id
	CustomerID *int64 `json:"customer_id,omitempty"`

	// document id
	DocumentID *int64 `json:"document_id,omitempty"`

	// file name
	// Read Only: true
	FileName string `json:"file_name,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// project id
	ProjectID *int64 `json:"project_id,omitempty"`

	// In byte
	// Read Only: true
	Size int64 `json:"size,omitempty"`
}

// Validate validates this attachment
func (m *Attachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attachment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Attachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attachment) UnmarshalBinary(b []byte) error {
	var res Attachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}