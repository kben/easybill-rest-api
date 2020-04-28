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

// DocumentPosition document position
//
// swagger:model DocumentPosition
type DocumentPosition struct {

	// booking account
	BookingAccount *string `json:"booking_account,omitempty"`

	// cost price charge
	// Read Only: true
	CostPriceCharge float32 `json:"cost_price_charge,omitempty"`

	// cost price charge type
	// Read Only: true
	// Enum: [PERCENT AMOUNT]
	CostPriceChargeType string `json:"cost_price_charge_type,omitempty"`

	// cost price net
	CostPriceNet *float32 `json:"cost_price_net,omitempty"`

	// cost price total
	// Read Only: true
	CostPriceTotal float32 `json:"cost_price_total,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// discount
	Discount *float32 `json:"discount,omitempty"`

	// discount type
	// Enum: [PERCENT AMOUNT]
	DiscountType *string `json:"discount_type,omitempty"`

	// export cost 1
	ExportCost1 *string `json:"export_cost_1,omitempty"`

	// export cost 2
	ExportCost2 *string `json:"export_cost_2,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// item type
	// Enum: [PRODUCT SERVICE UNDEFINED]
	ItemType *string `json:"itemType,omitempty"`

	// number
	Number *string `json:"number,omitempty"`

	// Automatic by default (first item: 1, second item: 2, ...)
	Position int64 `json:"position,omitempty"`

	// If set, values are copied from the referenced position
	PositionID *int64 `json:"position_id,omitempty"`

	// quantity
	// Max Length: 10
	Quantity *float32 `json:"quantity,omitempty"`

	// Use quantity_str if you want to set a quantity like: 1:30 h or 3x5 m. quantity_str overwrites quantity.
	// Max Length: 10
	QuantityStr string `json:"quantity_str,omitempty"`

	// serial number
	// Read Only: true
	SerialNumber string `json:"serial_number,omitempty"`

	// serial number id
	// Read Only: true
	SerialNumberID string `json:"serial_number_id,omitempty"`

	// single price gross
	// Read Only: true
	SinglePriceGross float32 `json:"single_price_gross,omitempty"`

	// single price net
	SinglePriceNet *float32 `json:"single_price_net,omitempty"`

	// total price gross
	// Read Only: true
	TotalPriceGross float32 `json:"total_price_gross,omitempty"`

	// total price net
	// Read Only: true
	TotalPriceNet float32 `json:"total_price_net,omitempty"`

	// total vat
	// Read Only: true
	TotalVat float32 `json:"total_vat,omitempty"`

	// type
	// Enum: [POSITION POSITION_NOCALC TEXT]
	Type *string `json:"type,omitempty"`

	// unit
	Unit *string `json:"unit,omitempty"`

	// vat percent
	VatPercent float32 `json:"vat_percent,omitempty"`
}

// Validate validates this document position
func (m *DocumentPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostPriceChargeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantityStr(formats); err != nil {
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

var documentPositionTypeCostPriceChargeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PERCENT","AMOUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentPositionTypeCostPriceChargeTypePropEnum = append(documentPositionTypeCostPriceChargeTypePropEnum, v)
	}
}

const (

	// DocumentPositionCostPriceChargeTypePERCENT captures enum value "PERCENT"
	DocumentPositionCostPriceChargeTypePERCENT string = "PERCENT"

	// DocumentPositionCostPriceChargeTypeAMOUNT captures enum value "AMOUNT"
	DocumentPositionCostPriceChargeTypeAMOUNT string = "AMOUNT"
)

// prop value enum
func (m *DocumentPosition) validateCostPriceChargeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentPositionTypeCostPriceChargeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DocumentPosition) validateCostPriceChargeType(formats strfmt.Registry) error {

	if swag.IsZero(m.CostPriceChargeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCostPriceChargeTypeEnum("cost_price_charge_type", "body", m.CostPriceChargeType); err != nil {
		return err
	}

	return nil
}

var documentPositionTypeDiscountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PERCENT","AMOUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentPositionTypeDiscountTypePropEnum = append(documentPositionTypeDiscountTypePropEnum, v)
	}
}

const (

	// DocumentPositionDiscountTypePERCENT captures enum value "PERCENT"
	DocumentPositionDiscountTypePERCENT string = "PERCENT"

	// DocumentPositionDiscountTypeAMOUNT captures enum value "AMOUNT"
	DocumentPositionDiscountTypeAMOUNT string = "AMOUNT"
)

// prop value enum
func (m *DocumentPosition) validateDiscountTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentPositionTypeDiscountTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DocumentPosition) validateDiscountType(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiscountTypeEnum("discount_type", "body", *m.DiscountType); err != nil {
		return err
	}

	return nil
}

var documentPositionTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRODUCT","SERVICE","UNDEFINED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentPositionTypeItemTypePropEnum = append(documentPositionTypeItemTypePropEnum, v)
	}
}

const (

	// DocumentPositionItemTypePRODUCT captures enum value "PRODUCT"
	DocumentPositionItemTypePRODUCT string = "PRODUCT"

	// DocumentPositionItemTypeSERVICE captures enum value "SERVICE"
	DocumentPositionItemTypeSERVICE string = "SERVICE"

	// DocumentPositionItemTypeUNDEFINED captures enum value "UNDEFINED"
	DocumentPositionItemTypeUNDEFINED string = "UNDEFINED"
)

// prop value enum
func (m *DocumentPosition) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentPositionTypeItemTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DocumentPosition) validateItemType(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", *m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *DocumentPosition) validateQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.Quantity) { // not required
		return nil
	}

	if err := validate.MaxLength("quantity", "body", string(*m.Quantity), 10); err != nil {
		return err
	}

	return nil
}

func (m *DocumentPosition) validateQuantityStr(formats strfmt.Registry) error {

	if swag.IsZero(m.QuantityStr) { // not required
		return nil
	}

	if err := validate.MaxLength("quantity_str", "body", string(m.QuantityStr), 10); err != nil {
		return err
	}

	return nil
}

var documentPositionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["POSITION","POSITION_NOCALC","TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentPositionTypeTypePropEnum = append(documentPositionTypeTypePropEnum, v)
	}
}

const (

	// DocumentPositionTypePOSITION captures enum value "POSITION"
	DocumentPositionTypePOSITION string = "POSITION"

	// DocumentPositionTypePOSITIONNOCALC captures enum value "POSITION_NOCALC"
	DocumentPositionTypePOSITIONNOCALC string = "POSITION_NOCALC"

	// DocumentPositionTypeTEXT captures enum value "TEXT"
	DocumentPositionTypeTEXT string = "TEXT"
)

// prop value enum
func (m *DocumentPosition) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentPositionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DocumentPosition) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocumentPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentPosition) UnmarshalBinary(b []byte) error {
	var res DocumentPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}