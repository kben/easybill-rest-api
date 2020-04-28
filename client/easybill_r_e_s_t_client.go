// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/client/attachment"
	"github.com/kben/easybill-rest-api/client/contact"
	"github.com/kben/easybill-rest-api/client/customer"
	"github.com/kben/easybill-rest-api/client/customer_group"
	"github.com/kben/easybill-rest-api/client/discount"
	"github.com/kben/easybill-rest-api/client/document"
	"github.com/kben/easybill-rest-api/client/document_payment"
	"github.com/kben/easybill-rest-api/client/logins"
	"github.com/kben/easybill-rest-api/client/pdf_templates"
	"github.com/kben/easybill-rest-api/client/position"
	"github.com/kben/easybill-rest-api/client/position_group"
	"github.com/kben/easybill-rest-api/client/post_box"
	"github.com/kben/easybill-rest-api/client/project"
	"github.com/kben/easybill-rest-api/client/sepa_payment"
	"github.com/kben/easybill-rest-api/client/serial_number"
	"github.com/kben/easybill-rest-api/client/stock"
	"github.com/kben/easybill-rest-api/client/task"
	"github.com/kben/easybill-rest-api/client/text_template"
	"github.com/kben/easybill-rest-api/client/time_tracking"
	"github.com/kben/easybill-rest-api/client/webhook"
)

// Default easybill r e s t HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.easybill.de"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/rest/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new easybill r e s t HTTP client.
func NewHTTPClient(formats strfmt.Registry) *EasybillREST {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new easybill r e s t HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *EasybillREST {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new easybill r e s t client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *EasybillREST {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(EasybillREST)
	cli.Transport = transport
	cli.Attachment = attachment.New(transport, formats)
	cli.Contact = contact.New(transport, formats)
	cli.Customer = customer.New(transport, formats)
	cli.CustomerGroup = customer_group.New(transport, formats)
	cli.Discount = discount.New(transport, formats)
	cli.Document = document.New(transport, formats)
	cli.DocumentPayment = document_payment.New(transport, formats)
	cli.Logins = logins.New(transport, formats)
	cli.PdfTemplates = pdf_templates.New(transport, formats)
	cli.Position = position.New(transport, formats)
	cli.PositionGroup = position_group.New(transport, formats)
	cli.PostBox = post_box.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.SepaPayment = sepa_payment.New(transport, formats)
	cli.SerialNumber = serial_number.New(transport, formats)
	cli.Stock = stock.New(transport, formats)
	cli.Task = task.New(transport, formats)
	cli.TextTemplate = text_template.New(transport, formats)
	cli.TimeTracking = time_tracking.New(transport, formats)
	cli.Webhook = webhook.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// EasybillREST is a client for easybill r e s t
type EasybillREST struct {
	Attachment attachment.ClientService

	Contact contact.ClientService

	Customer customer.ClientService

	CustomerGroup customer_group.ClientService

	Discount discount.ClientService

	Document document.ClientService

	DocumentPayment document_payment.ClientService

	Logins logins.ClientService

	PdfTemplates pdf_templates.ClientService

	Position position.ClientService

	PositionGroup position_group.ClientService

	PostBox post_box.ClientService

	Project project.ClientService

	SepaPayment sepa_payment.ClientService

	SerialNumber serial_number.ClientService

	Stock stock.ClientService

	Task task.ClientService

	TextTemplate text_template.ClientService

	TimeTracking time_tracking.ClientService

	Webhook webhook.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *EasybillREST) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Attachment.SetTransport(transport)
	c.Contact.SetTransport(transport)
	c.Customer.SetTransport(transport)
	c.CustomerGroup.SetTransport(transport)
	c.Discount.SetTransport(transport)
	c.Document.SetTransport(transport)
	c.DocumentPayment.SetTransport(transport)
	c.Logins.SetTransport(transport)
	c.PdfTemplates.SetTransport(transport)
	c.Position.SetTransport(transport)
	c.PositionGroup.SetTransport(transport)
	c.PostBox.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.SepaPayment.SetTransport(transport)
	c.SerialNumber.SetTransport(transport)
	c.Stock.SetTransport(transport)
	c.Task.SetTransport(transport)
	c.TextTemplate.SetTransport(transport)
	c.TimeTracking.SetTransport(transport)
	c.Webhook.SetTransport(transport)
}
