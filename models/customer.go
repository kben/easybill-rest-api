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

// Customer customer
//
// swagger:model Customer
type Customer struct {

	// 1 = Empfehlung eines anderen Kunden, 2 = Zeitungsanzeige, 3 = Eigene Akquisition, 4 = Mitarbeiter Akquisition, 5 = Google, 6 = Gelbe Seiten, 7 = Kostenlose Internet Plattform, 8 = Bezahlte Internet Plattform
	// Enum: [1 2 3 4 5 6 7 8]
	AcquireOptions *int64 `json:"acquire_options,omitempty"`

	// additional groups ids
	// Read Only: true
	AdditionalGroupsIds []int64 `json:"additional_groups_ids"`

	// bank account
	BankAccount *string `json:"bank_account,omitempty"`

	// bank account owner
	BankAccountOwner *string `json:"bank_account_owner,omitempty"`

	// bank bic
	BankBic *string `json:"bank_bic,omitempty"`

	// bank code
	BankCode *string `json:"bank_code,omitempty"`

	// bank iban
	BankIban *string `json:"bank_iban,omitempty"`

	// bank name
	BankName *string `json:"bank_name,omitempty"`

	// birth date
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// cash allowance
	// Maximum: 100
	// Minimum: 0
	CashAllowance *float32 `json:"cash_allowance,omitempty"`

	// cash allowance days
	CashAllowanceDays *int64 `json:"cash_allowance_days,omitempty"`

	// cash discount
	CashDiscount *float32 `json:"cash_discount,omitempty"`

	// cash discount type
	// Enum: [PERCENT AMOUNT]
	CashDiscountType *string `json:"cash_discount_type,omitempty"`

	// city
	City *string `json:"city,omitempty"`

	// company name
	// Required: true
	CompanyName *string `json:"company_name"`

	// country
	Country string `json:"country,omitempty"`

	// court
	Court *string `json:"court,omitempty"`

	// court registry number
	CourtRegistryNumber *string `json:"court_registry_number,omitempty"`

	// created at
	// Read Only: true
	// Format: date
	CreatedAt strfmt.Date `json:"created_at,omitempty"`

	// delivery city
	DeliveryCity *string `json:"delivery_city,omitempty"`

	// delivery company name
	DeliveryCompanyName *string `json:"delivery_company_name,omitempty"`

	// delivery country
	DeliveryCountry *string `json:"delivery_country,omitempty"`

	// delivery first name
	DeliveryFirstName *string `json:"delivery_first_name,omitempty"`

	// delivery last name
	DeliveryLastName *string `json:"delivery_last_name,omitempty"`

	// delivery personal
	DeliveryPersonal bool `json:"delivery_personal,omitempty"`

	// 0 = nothing, 1 = Mr, 2 = Mrs, 3 = Company, 4 = Mr & Mrs, 5 = Married couple, 6 = Family
	// Enum: [0 1 2 3 4 5 6]
	DeliverySalutation int64 `json:"delivery_salutation,omitempty"`

	// delivery street
	DeliveryStreet *string `json:"delivery_street,omitempty"`

	// delivery suffix 1
	DeliverySuffix1 *string `json:"delivery_suffix_1,omitempty"`

	// delivery suffix 2
	DeliverySuffix2 *string `json:"delivery_suffix_2,omitempty"`

	// delivery zip code
	DeliveryZipCode *string `json:"delivery_zip_code,omitempty"`

	// display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// due date in days
	DueInDays *int64 `json:"due_in_days,omitempty"`

	// emails
	Emails []string `json:"emails"`

	// fax
	Fax *string `json:"fax,omitempty"`

	// first name
	FirstName *string `json:"first_name,omitempty"`

	// will be replaced by its alias due_in_days.
	GracePeriod *int64 `json:"grace_period,omitempty"`

	// group id
	GroupID *int64 `json:"group_id,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// info 1
	Info1 *string `json:"info_1,omitempty"`

	// info 2
	Info2 *string `json:"info_2,omitempty"`

	// internet
	Internet *string `json:"internet,omitempty"`

	// last name
	// Required: true
	LastName *string `json:"last_name"`

	// login id
	LoginID int64 `json:"login_id,omitempty"`

	// mobile
	Mobile *string `json:"mobile,omitempty"`

	// note
	Note *string `json:"note,omitempty"`

	// Automatically generated if empty.
	Number string `json:"number,omitempty"`

	// 1 = Stets pünktliche Zahlung, 2 = überwiegend pünktliche Zahlung, 3 = überwiegend verspätete Zahlung, 5 = Grundsätzlich verspätete Zahlung
	// Enum: [1 2 3 5]
	PaymentOptions *int64 `json:"payment_options,omitempty"`

	// personal
	Personal *bool `json:"personal,omitempty"`

	// phone 1
	Phone1 *string `json:"phone_1,omitempty"`

	// phone 2
	Phone2 *string `json:"phone_2,omitempty"`

	// postbox
	Postbox *string `json:"postbox,omitempty"`

	// postbox city
	PostboxCity *string `json:"postbox_city,omitempty"`

	// postbox country
	PostboxCountry *string `json:"postbox_country,omitempty"`

	// postbox zip code
	PostboxZipCode *string `json:"postbox_zip_code,omitempty"`

	// sale price level
	// Enum: [SALEPRICE2 SALEPRICE3 SALEPRICE4 SALEPRICE5 SALEPRICE6 SALEPRICE7 SALEPRICE8 SALEPRICE9 SALEPRICE10]
	SalePriceLevel *string `json:"sale_price_level,omitempty"`

	// 0 = nothing, 1 = Mr, 2 = Mrs, 3 = Company, 4 = Mr & Mrs, 5 = Married couple, 6 = Family
	// Enum: [0 1 2 3 4 5 6]
	Salutation int64 `json:"salutation,omitempty"`

	// BASIC = SEPA-Basislastschrift, COR1 = SEPA-Basislastschrift COR1, COMPANY = SEPA-Firmenlastschrift, NULL = Noch kein Mandat erteilt
	// Enum: [BASIC COR1 COMPANY NULL]
	SepaAgreement *string `json:"sepa_agreement,omitempty"`

	// sepa agreement date
	// Format: date
	SepaAgreementDate *strfmt.Date `json:"sepa_agreement_date,omitempty"`

	// sepa mandate reference
	SepaMandateReference *string `json:"sepa_mandate_reference,omitempty"`

	// since date
	// Format: date
	SinceDate *strfmt.Date `json:"since_date,omitempty"`

	// street
	Street *string `json:"street,omitempty"`

	// suffix 1
	Suffix1 *string `json:"suffix_1,omitempty"`

	// suffix 2
	Suffix2 *string `json:"suffix_2,omitempty"`

	// tax number
	TaxNumber *string `json:"tax_number,omitempty"`

	// nStb = Nicht steuerbar (Drittland), nStbUstID = Nicht steuerbar (EU mit USt-IdNr.), nStbNoneUstID = Nicht steuerbar (EU ohne USt-IdNr.), revc = Steuerschuldwechsel §13b (Inland), IG = Innergemeinschaftliche Lieferung, AL = Ausfuhrlieferung, sStfr = sonstige Steuerbefreiung, NULL = Umsatzsteuerpflichtig
	// Enum: [nStb nStbUstID nStbNoneUstID nStbIm revc IG AL sStfr NULL]
	TaxOptions *string `json:"tax_options,omitempty"`

	// title
	Title *string `json:"title,omitempty"`

	// updated at
	// Read Only: true
	// Format: datetime
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// vat identifier
	VatIdentifier *string `json:"vat_identifier,omitempty"`

	// zip code
	ZipCode *string `json:"zip_code,omitempty"`
}

// Validate validates this customer
func (m *Customer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcquireOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalGroupsIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashAllowance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashDiscountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompanyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliverySalutation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalePriceLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalutation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSepaAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSepaAgreementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSinceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customerTypeAcquireOptionsPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeAcquireOptionsPropEnum = append(customerTypeAcquireOptionsPropEnum, v)
	}
}

// prop value enum
func (m *Customer) validateAcquireOptionsEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, customerTypeAcquireOptionsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateAcquireOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.AcquireOptions) { // not required
		return nil
	}

	// value enum
	if err := m.validateAcquireOptionsEnum("acquire_options", "body", *m.AcquireOptions); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateAdditionalGroupsIds(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalGroupsIds) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalGroupsIds); i++ {

	}

	return nil
}

func (m *Customer) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateCashAllowance(formats strfmt.Registry) error {

	if swag.IsZero(m.CashAllowance) { // not required
		return nil
	}

	if err := validate.Minimum("cash_allowance", "body", float64(*m.CashAllowance), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("cash_allowance", "body", float64(*m.CashAllowance), 100, false); err != nil {
		return err
	}

	return nil
}

var customerTypeCashDiscountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PERCENT","AMOUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeCashDiscountTypePropEnum = append(customerTypeCashDiscountTypePropEnum, v)
	}
}

const (

	// CustomerCashDiscountTypePERCENT captures enum value "PERCENT"
	CustomerCashDiscountTypePERCENT string = "PERCENT"

	// CustomerCashDiscountTypeAMOUNT captures enum value "AMOUNT"
	CustomerCashDiscountTypeAMOUNT string = "AMOUNT"
)

// prop value enum
func (m *Customer) validateCashDiscountTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeCashDiscountTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateCashDiscountType(formats strfmt.Registry) error {

	if swag.IsZero(m.CashDiscountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCashDiscountTypeEnum("cash_discount_type", "body", *m.CashDiscountType); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateCompanyName(formats strfmt.Registry) error {

	if err := validate.Required("company_name", "body", m.CompanyName); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var customerTypeDeliverySalutationPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeDeliverySalutationPropEnum = append(customerTypeDeliverySalutationPropEnum, v)
	}
}

// prop value enum
func (m *Customer) validateDeliverySalutationEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, customerTypeDeliverySalutationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateDeliverySalutation(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliverySalutation) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeliverySalutationEnum("delivery_salutation", "body", m.DeliverySalutation); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {

	}

	return nil
}

func (m *Customer) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

var customerTypePaymentOptionsPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypePaymentOptionsPropEnum = append(customerTypePaymentOptionsPropEnum, v)
	}
}

// prop value enum
func (m *Customer) validatePaymentOptionsEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, customerTypePaymentOptionsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validatePaymentOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentOptions) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentOptionsEnum("payment_options", "body", *m.PaymentOptions); err != nil {
		return err
	}

	return nil
}

var customerTypeSalePriceLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SALEPRICE2","SALEPRICE3","SALEPRICE4","SALEPRICE5","SALEPRICE6","SALEPRICE7","SALEPRICE8","SALEPRICE9","SALEPRICE10"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeSalePriceLevelPropEnum = append(customerTypeSalePriceLevelPropEnum, v)
	}
}

const (

	// CustomerSalePriceLevelSALEPRICE2 captures enum value "SALEPRICE2"
	CustomerSalePriceLevelSALEPRICE2 string = "SALEPRICE2"

	// CustomerSalePriceLevelSALEPRICE3 captures enum value "SALEPRICE3"
	CustomerSalePriceLevelSALEPRICE3 string = "SALEPRICE3"

	// CustomerSalePriceLevelSALEPRICE4 captures enum value "SALEPRICE4"
	CustomerSalePriceLevelSALEPRICE4 string = "SALEPRICE4"

	// CustomerSalePriceLevelSALEPRICE5 captures enum value "SALEPRICE5"
	CustomerSalePriceLevelSALEPRICE5 string = "SALEPRICE5"

	// CustomerSalePriceLevelSALEPRICE6 captures enum value "SALEPRICE6"
	CustomerSalePriceLevelSALEPRICE6 string = "SALEPRICE6"

	// CustomerSalePriceLevelSALEPRICE7 captures enum value "SALEPRICE7"
	CustomerSalePriceLevelSALEPRICE7 string = "SALEPRICE7"

	// CustomerSalePriceLevelSALEPRICE8 captures enum value "SALEPRICE8"
	CustomerSalePriceLevelSALEPRICE8 string = "SALEPRICE8"

	// CustomerSalePriceLevelSALEPRICE9 captures enum value "SALEPRICE9"
	CustomerSalePriceLevelSALEPRICE9 string = "SALEPRICE9"

	// CustomerSalePriceLevelSALEPRICE10 captures enum value "SALEPRICE10"
	CustomerSalePriceLevelSALEPRICE10 string = "SALEPRICE10"
)

// prop value enum
func (m *Customer) validateSalePriceLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeSalePriceLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateSalePriceLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.SalePriceLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateSalePriceLevelEnum("sale_price_level", "body", *m.SalePriceLevel); err != nil {
		return err
	}

	return nil
}

var customerTypeSalutationPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeSalutationPropEnum = append(customerTypeSalutationPropEnum, v)
	}
}

// prop value enum
func (m *Customer) validateSalutationEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, customerTypeSalutationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateSalutation(formats strfmt.Registry) error {

	if swag.IsZero(m.Salutation) { // not required
		return nil
	}

	// value enum
	if err := m.validateSalutationEnum("salutation", "body", m.Salutation); err != nil {
		return err
	}

	return nil
}

var customerTypeSepaAgreementPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BASIC","COR1","COMPANY","NULL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeSepaAgreementPropEnum = append(customerTypeSepaAgreementPropEnum, v)
	}
}

const (

	// CustomerSepaAgreementBASIC captures enum value "BASIC"
	CustomerSepaAgreementBASIC string = "BASIC"

	// CustomerSepaAgreementCOR1 captures enum value "COR1"
	CustomerSepaAgreementCOR1 string = "COR1"

	// CustomerSepaAgreementCOMPANY captures enum value "COMPANY"
	CustomerSepaAgreementCOMPANY string = "COMPANY"

	// CustomerSepaAgreementNULL captures enum value "NULL"
	CustomerSepaAgreementNULL string = "NULL"
)

// prop value enum
func (m *Customer) validateSepaAgreementEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeSepaAgreementPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateSepaAgreement(formats strfmt.Registry) error {

	if swag.IsZero(m.SepaAgreement) { // not required
		return nil
	}

	// value enum
	if err := m.validateSepaAgreementEnum("sepa_agreement", "body", *m.SepaAgreement); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateSepaAgreementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SepaAgreementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sepa_agreement_date", "body", "date", m.SepaAgreementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateSinceDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SinceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("since_date", "body", "date", m.SinceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var customerTypeTaxOptionsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nStb","nStbUstID","nStbNoneUstID","nStbIm","revc","IG","AL","sStfr","NULL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerTypeTaxOptionsPropEnum = append(customerTypeTaxOptionsPropEnum, v)
	}
}

const (

	// CustomerTaxOptionsNStb captures enum value "nStb"
	CustomerTaxOptionsNStb string = "nStb"

	// CustomerTaxOptionsNStbUstID captures enum value "nStbUstID"
	CustomerTaxOptionsNStbUstID string = "nStbUstID"

	// CustomerTaxOptionsNStbNoneUstID captures enum value "nStbNoneUstID"
	CustomerTaxOptionsNStbNoneUstID string = "nStbNoneUstID"

	// CustomerTaxOptionsNStbIm captures enum value "nStbIm"
	CustomerTaxOptionsNStbIm string = "nStbIm"

	// CustomerTaxOptionsRevc captures enum value "revc"
	CustomerTaxOptionsRevc string = "revc"

	// CustomerTaxOptionsIG captures enum value "IG"
	CustomerTaxOptionsIG string = "IG"

	// CustomerTaxOptionsAL captures enum value "AL"
	CustomerTaxOptionsAL string = "AL"

	// CustomerTaxOptionsSStfr captures enum value "sStfr"
	CustomerTaxOptionsSStfr string = "sStfr"

	// CustomerTaxOptionsNULL captures enum value "NULL"
	CustomerTaxOptionsNULL string = "NULL"
)

// prop value enum
func (m *Customer) validateTaxOptionsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerTypeTaxOptionsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Customer) validateTaxOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxOptions) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaxOptionsEnum("tax_options", "body", *m.TaxOptions); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "datetime", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Customer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customer) UnmarshalBinary(b []byte) error {
	var res Customer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
