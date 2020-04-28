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

// Discount discount
//
// swagger:model Discount
type Discount struct {

	// customer id
	// Required: true
	CustomerID *int64 `json:"customer_id"`

	// The discount value depending on "discount_type"
	Discount int64 `json:"discount,omitempty"`

	// AMOUNT subtracts the value in "discount" from the total<br/> QUANTITY subtracts the value in "discount" multiplied by quantity<br/> PERCENT uses the value in "discount" as a percentage<br/> FIX sets the value in "discount" as the new price
	// Enum: [AMOUNT PERCENT QUANTITY FIX]
	DiscountType *string `json:"discount_type,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`
}

// Validate validates this discount
func (m *Discount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscountType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Discount) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

var discountTypeDiscountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AMOUNT","PERCENT","QUANTITY","FIX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discountTypeDiscountTypePropEnum = append(discountTypeDiscountTypePropEnum, v)
	}
}

const (

	// DiscountDiscountTypeAMOUNT captures enum value "AMOUNT"
	DiscountDiscountTypeAMOUNT string = "AMOUNT"

	// DiscountDiscountTypePERCENT captures enum value "PERCENT"
	DiscountDiscountTypePERCENT string = "PERCENT"

	// DiscountDiscountTypeQUANTITY captures enum value "QUANTITY"
	DiscountDiscountTypeQUANTITY string = "QUANTITY"

	// DiscountDiscountTypeFIX captures enum value "FIX"
	DiscountDiscountTypeFIX string = "FIX"
)

// prop value enum
func (m *Discount) validateDiscountTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, discountTypeDiscountTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Discount) validateDiscountType(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiscountTypeEnum("discount_type", "body", *m.DiscountType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Discount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Discount) UnmarshalBinary(b []byte) error {
	var res Discount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
