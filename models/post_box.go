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

// PostBox post box
//
// swagger:model PostBox
type PostBox struct {

	// cc
	Cc string `json:"cc,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// date
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// document id
	DocumentID int64 `json:"document_id,omitempty"`

	// from
	From string `json:"from,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// login id
	// Read Only: true
	LoginID int64 `json:"login_id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// processed at
	// Format: date-time
	ProcessedAt strfmt.DateTime `json:"processed_at,omitempty"`

	// send by self
	SendBySelf bool `json:"send_by_self,omitempty"`

	// send with attachment
	SendWithAttachment bool `json:"send_with_attachment,omitempty"`

	// status
	// Enum: [WAITING PREPARE ERROR OK PROCESSING]
	Status string `json:"status,omitempty"`

	// status msg
	StatusMsg string `json:"status_msg,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// to
	To string `json:"to,omitempty"`

	// type
	// Enum: [FAX MAIL POST]
	Type string `json:"type,omitempty"`
}

// Validate validates this post box
func (m *PostBox) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *PostBox) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostBox) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostBox) validateProcessedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("processed_at", "body", "date-time", m.ProcessedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var postBoxTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WAITING","PREPARE","ERROR","OK","PROCESSING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postBoxTypeStatusPropEnum = append(postBoxTypeStatusPropEnum, v)
	}
}

const (

	// PostBoxStatusWAITING captures enum value "WAITING"
	PostBoxStatusWAITING string = "WAITING"

	// PostBoxStatusPREPARE captures enum value "PREPARE"
	PostBoxStatusPREPARE string = "PREPARE"

	// PostBoxStatusERROR captures enum value "ERROR"
	PostBoxStatusERROR string = "ERROR"

	// PostBoxStatusOK captures enum value "OK"
	PostBoxStatusOK string = "OK"

	// PostBoxStatusPROCESSING captures enum value "PROCESSING"
	PostBoxStatusPROCESSING string = "PROCESSING"
)

// prop value enum
func (m *PostBox) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postBoxTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostBox) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var postBoxTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FAX","MAIL","POST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postBoxTypeTypePropEnum = append(postBoxTypeTypePropEnum, v)
	}
}

const (

	// PostBoxTypeFAX captures enum value "FAX"
	PostBoxTypeFAX string = "FAX"

	// PostBoxTypeMAIL captures enum value "MAIL"
	PostBoxTypeMAIL string = "MAIL"

	// PostBoxTypePOST captures enum value "POST"
	PostBoxTypePOST string = "POST"
)

// prop value enum
func (m *PostBox) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postBoxTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostBox) validateType(formats strfmt.Registry) error {

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
func (m *PostBox) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostBox) UnmarshalBinary(b []byte) error {
	var res PostBox
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
