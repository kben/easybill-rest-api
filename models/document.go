// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Document document
//
// swagger:model Document
type Document struct {

	// This information comes from the customer which can be set with customer_id.
	// Read Only: true
	Address struct {
		DocumentAddress
	} `json:"address,omitempty"`

	// Amount in cents  (e.g. "150" = 1.50€)
	// Read Only: true
	Amount int64 `json:"amount,omitempty"`

	// Amount in cents  (e.g. "150" = 1.50€)
	// Read Only: true
	AmountNet int64 `json:"amount_net,omitempty"`

	// attachment ids
	// Read Only: true
	AttachmentIds []int64 `json:"attachment_ids"`

	// bank debit form
	BankDebitForm *string `json:"bank_debit_form,omitempty"`

	// billing country
	// Read Only: true
	BillingCountry string `json:"billing_country,omitempty"`

	// 0 === Net, 1 === Gross.
	// Enum: [0 1]
	CalcVatFrom int64 `json:"calc_vat_from,omitempty"`

	// ID from the cancel document. Only for document type INVOICE.
	// Read Only: true
	CancelID int64 `json:"cancel_id,omitempty"`

	// cash allowance
	CashAllowance *float32 `json:"cash_allowance,omitempty"`

	// cash allowance days
	CashAllowanceDays *int64 `json:"cash_allowance_days,omitempty"`

	// cash allowance text
	CashAllowanceText *string `json:"cash_allowance_text,omitempty"`

	// contact id
	ContactID *int64 `json:"contact_id,omitempty"`

	// contact label
	ContactLabel string `json:"contact_label,omitempty"`

	// contact text
	ContactText string `json:"contact_text,omitempty"`

	// created at
	// Read Only: true
	// Format: date-time
	// XXX: remove for now
	//CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// currency
	Currency *string `json:"currency,omitempty"`

	// customer id
	CustomerID *int64 `json:"customer_id,omitempty"`

	// customer snapshot
	// Read Only: true
	CustomerSnapshot struct {
		CustomerSnapshot
	} `json:"customer_snapshot,omitempty"`

	// discount
	Discount *string `json:"discount,omitempty"`

	// discount type
	// Enum: [PERCENT AMOUNT]
	DiscountType *string `json:"discount_type,omitempty"`

	// document date
	// Format: date
	DocumentDate strfmt.Date `json:"document_date,omitempty"`

	// To change the value use grace_period.
	// Read Only: true
	// Format: date
	DueDate strfmt.Date `json:"due_date,omitempty"`

	// due date in days.
	DueInDays *int64 `json:"due_in_days,omitempty"`

	// edited at
	// Read Only: true
	// Format: date-time
	// XXX: remove for now
	//EditedAt strfmt.DateTime `json:"edited_at,omitempty"`

	// external id
	ExternalID *string `json:"external_id,omitempty"`

	// fulfillment country
	FulfillmentCountry *string `json:"fulfillment_country,omitempty"`

	// will be replaced by its alias due_in_days.
	GracePeriod *int64 `json:"grace_period,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// is archive
	IsArchive *bool `json:"is_archive,omitempty"`

	// This property is read only. To finish the document call /documents/{id}/done.
	// Read Only: true
	IsDraft *bool `json:"is_draft,omitempty"`

	// Marks a document as a replica from another software.
	IsReplica *bool `json:"is_replica,omitempty"`

	// items
	Items []*DocumentPosition `json:"items"`

	// This information comes from the customer which can be set with customer_id.
	// Read Only: true
	LabelAddress struct {
		DocumentAddress
	} `json:"label_address,omitempty"`

	// last postbox id
	// Read Only: true
	LastPostboxID int64 `json:"last_postbox_id,omitempty"`

	// If omitted or null, the currently active login is used.
	LoginID int64 `json:"login_id,omitempty"`

	// number
	Number *string `json:"number,omitempty"`

	// order number
	OrderNumber string `json:"order_number,omitempty"`

	// paid amount
	// Read Only: true
	PaidAmount int64 `json:"paid_amount,omitempty"`

	// paid at
	// Read Only: true
	// Format: date
	PaidAt strfmt.Date `json:"paid_at,omitempty"`

	// pdf pages
	// Read Only: true
	PdfPages int64 `json:"pdf_pages,omitempty"`

	// Default template is null or 'DE', default english is 'EN' and for all others use the numeric template ID.
	PdfTemplate string `json:"pdf_template,omitempty"`

	// project id
	ProjectID *int64 `json:"project_id,omitempty"`

	// This object is only available in document type RECURRING
	RecurringOptions struct {
		DocumentRecurring
	} `json:"recurring_options,omitempty"`

	// Reference document id
	RefID *int64 `json:"ref_id,omitempty"`

	// replica url
	ReplicaURL *string `json:"replica_url,omitempty"`

	// This object is only available in document type INVOICE or CREDIT.
	ServiceDate struct {
		ServiceDate
	} `json:"service_date,omitempty"`

	// shipping country
	ShippingCountry *string `json:"shipping_country,omitempty"`

	// This value can only be used in document type DELIVERY, ORDER, CHARGE or OFFER. NULL is default = not set.
	// Enum: [ACCEPT DONE DROPSHIPPING CANCEL]
	Status *string `json:"status,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// text prefix
	TextPrefix string `json:"text_prefix,omitempty"`

	// Overwrites the default vat-option text from the document layout. It is only displayed in documents with the type other than: Delivery, Dunning, Reminder or Letter and a different vat-option than null
	TextTax *string `json:"text_tax,omitempty"`

	// title
	Title *string `json:"title,omitempty"`

	// Can only set on create.
	// Enum: [INVOICE RECURRING CREDIT OFFER REMINDER DUNNING STORNO STORNO_CREDIT DELIVERY PDF CHARGE CHARGE_CONFIRM LETTER ORDER]
	Type *string `json:"type,omitempty"`

	// If true and customer has shipping address then it will be used.
	UseShippingAddress *bool `json:"use_shipping_address,omitempty"`

	// vat country
	VatCountry *string `json:"vat_country,omitempty"`

	// NULL: Normal steuerbar<br/> nStb: Nicht steuerbar (Drittland)<br/> nStbUstID: Nicht steuerbar (EU mit USt-IdNr.)<br/> nStbNoneUstID: Nicht steuerbar (EU ohne USt-IdNr.)<br/> nStbIm: Nicht steuerbarer Innenumsatz<br/> revc: Steuerschuldwechsel §13b (Inland)<br/> IG: Innergemeinschaftliche Lieferung<br/> AL: Ausfuhrlieferung<br/> sStfr: sonstige Steuerbefreiung<br/> smallBusiness: Kleinunternehmen (Keine MwSt.)
	// Enum: [NULL nStb nStbUstID nStbNoneUstID nStbIm revc IG AL sStfr smallBusiness]
	VatOption *string `json:"vat_option,omitempty"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachmentIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalcVatFrom(formats); err != nil {
		res = append(res, err)
	}

	/*
		if err := m.validateCreatedAt(formats); err != nil {
			res = append(res, err)
		}
	*/

	if err := m.validateCustomerSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	/*
		if err := m.validateEditedAt(formats); err != nil {
			res = append(res, err)
		}
	*/

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaidAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	return nil
}

func (m *Document) validateAttachmentIds(formats strfmt.Registry) error {

	if swag.IsZero(m.AttachmentIds) { // not required
		return nil
	}

	for i := 0; i < len(m.AttachmentIds); i++ {

	}

	return nil
}

var documentTypeCalcVatFromPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeCalcVatFromPropEnum = append(documentTypeCalcVatFromPropEnum, v)
	}
}

// prop value enum
func (m *Document) validateCalcVatFromEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, documentTypeCalcVatFromPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateCalcVatFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.CalcVatFrom) { // not required
		return nil
	}

	// value enum
	if err := m.validateCalcVatFromEnum("calc_vat_from", "body", m.CalcVatFrom); err != nil {
		return err
	}

	return nil
}

/*
func (m *Document) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}
*/

func (m *Document) validateCustomerSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerSnapshot) { // not required
		return nil
	}

	return nil
}

var documentTypeDiscountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PERCENT","AMOUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeDiscountTypePropEnum = append(documentTypeDiscountTypePropEnum, v)
	}
}

const (

	// DocumentDiscountTypePERCENT captures enum value "PERCENT"
	DocumentDiscountTypePERCENT string = "PERCENT"

	// DocumentDiscountTypeAMOUNT captures enum value "AMOUNT"
	DocumentDiscountTypeAMOUNT string = "AMOUNT"
)

// prop value enum
func (m *Document) validateDiscountTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentTypeDiscountTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateDiscountType(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiscountTypeEnum("discount_type", "body", *m.DiscountType); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateDocumentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DocumentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("document_date", "body", "date", m.DocumentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateDueDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("due_date", "body", "date", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

/*
func (m *Document) validateEditedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EditedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("edited_at", "body", "date-time", m.EditedAt.String(), formats); err != nil {
		return err
	}

	return nil
}
*/

func (m *Document) validateItems(formats strfmt.Registry) error {

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

func (m *Document) validateLabelAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.LabelAddress) { // not required
		return nil
	}

	return nil
}

func (m *Document) validatePaidAt(formats strfmt.Registry) error {

	if swag.IsZero(m.PaidAt) { // not required
		return nil
	}

	if err := validate.FormatOf("paid_at", "body", "date", m.PaidAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateRecurringOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.RecurringOptions) { // not required
		return nil
	}

	return nil
}

func (m *Document) validateServiceDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceDate) { // not required
		return nil
	}

	return nil
}

var documentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCEPT","DONE","DROPSHIPPING","CANCEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeStatusPropEnum = append(documentTypeStatusPropEnum, v)
	}
}

const (

	// DocumentStatusACCEPT captures enum value "ACCEPT"
	DocumentStatusACCEPT string = "ACCEPT"

	// DocumentStatusDONE captures enum value "DONE"
	DocumentStatusDONE string = "DONE"

	// DocumentStatusDROPSHIPPING captures enum value "DROPSHIPPING"
	DocumentStatusDROPSHIPPING string = "DROPSHIPPING"

	// DocumentStatusCANCEL captures enum value "CANCEL"
	DocumentStatusCANCEL string = "CANCEL"
)

// prop value enum
func (m *Document) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var documentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INVOICE","RECURRING","CREDIT","OFFER","REMINDER","DUNNING","STORNO","STORNO_CREDIT","DELIVERY","PDF","CHARGE","CHARGE_CONFIRM","LETTER","ORDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeTypePropEnum = append(documentTypeTypePropEnum, v)
	}
}

const (

	// DocumentTypeINVOICE captures enum value "INVOICE"
	DocumentTypeINVOICE string = "INVOICE"

	// DocumentTypeRECURRING captures enum value "RECURRING"
	DocumentTypeRECURRING string = "RECURRING"

	// DocumentTypeCREDIT captures enum value "CREDIT"
	DocumentTypeCREDIT string = "CREDIT"

	// DocumentTypeOFFER captures enum value "OFFER"
	DocumentTypeOFFER string = "OFFER"

	// DocumentTypeREMINDER captures enum value "REMINDER"
	DocumentTypeREMINDER string = "REMINDER"

	// DocumentTypeDUNNING captures enum value "DUNNING"
	DocumentTypeDUNNING string = "DUNNING"

	// DocumentTypeSTORNO captures enum value "STORNO"
	DocumentTypeSTORNO string = "STORNO"

	// DocumentTypeSTORNOCREDIT captures enum value "STORNO_CREDIT"
	DocumentTypeSTORNOCREDIT string = "STORNO_CREDIT"

	// DocumentTypeDELIVERY captures enum value "DELIVERY"
	DocumentTypeDELIVERY string = "DELIVERY"

	// DocumentTypePDF captures enum value "PDF"
	DocumentTypePDF string = "PDF"

	// DocumentTypeCHARGE captures enum value "CHARGE"
	DocumentTypeCHARGE string = "CHARGE"

	// DocumentTypeCHARGECONFIRM captures enum value "CHARGE_CONFIRM"
	DocumentTypeCHARGECONFIRM string = "CHARGE_CONFIRM"

	// DocumentTypeLETTER captures enum value "LETTER"
	DocumentTypeLETTER string = "LETTER"

	// DocumentTypeORDER captures enum value "ORDER"
	DocumentTypeORDER string = "ORDER"
)

// prop value enum
func (m *Document) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var documentTypeVatOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NULL","nStb","nStbUstID","nStbNoneUstID","nStbIm","revc","IG","AL","sStfr","smallBusiness"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeVatOptionPropEnum = append(documentTypeVatOptionPropEnum, v)
	}
}

const (

	// DocumentVatOptionNULL captures enum value "NULL"
	DocumentVatOptionNULL string = "NULL"

	// DocumentVatOptionNStb captures enum value "nStb"
	DocumentVatOptionNStb string = "nStb"

	// DocumentVatOptionNStbUstID captures enum value "nStbUstID"
	DocumentVatOptionNStbUstID string = "nStbUstID"

	// DocumentVatOptionNStbNoneUstID captures enum value "nStbNoneUstID"
	DocumentVatOptionNStbNoneUstID string = "nStbNoneUstID"

	// DocumentVatOptionNStbIm captures enum value "nStbIm"
	DocumentVatOptionNStbIm string = "nStbIm"

	// DocumentVatOptionRevc captures enum value "revc"
	DocumentVatOptionRevc string = "revc"

	// DocumentVatOptionIG captures enum value "IG"
	DocumentVatOptionIG string = "IG"

	// DocumentVatOptionAL captures enum value "AL"
	DocumentVatOptionAL string = "AL"

	// DocumentVatOptionSStfr captures enum value "sStfr"
	DocumentVatOptionSStfr string = "sStfr"

	// DocumentVatOptionSmallBusiness captures enum value "smallBusiness"
	DocumentVatOptionSmallBusiness string = "smallBusiness"
)

// prop value enum
func (m *Document) validateVatOptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, documentTypeVatOptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateVatOption(formats strfmt.Registry) error {

	if swag.IsZero(m.VatOption) { // not required
		return nil
	}

	// value enum
	if err := m.validateVatOptionEnum("vat_option", "body", *m.VatOption); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
