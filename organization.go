package lago

import (
	"context"
	"time"
)

type OrganizationDocumentNumbering string

const (
	DocumentNumberingPerCustomer     OrganizationDocumentNumbering = "per_customer"
	DocumentNumberingPerOrganization OrganizationDocumentNumbering = "per_organization"
)

type organizationParams struct {
	Organization *OrganizationInput `json:"organization"`
}

type OrganizationBillingConfigurationInput struct {
	InvoiceGracePeriod int    `json:"invoice_grace_period,omitempty"`
	InvoiceFooter      string `json:"invoice_footer,omitempty"`
	DocumentLocale     string `json:"document_locale,omitempty"`
}

type OrganizationBillingConfiguration struct {
	InvoiceGracePeriod int    `json:"invoice_grace_period,omitempty"`
	InvoiceFooter      string `json:"invoice_footer,omitempty"`
	DocumentLocale     string `json:"document_locale,omitempty"`
}

type OrganizationInput struct {
	Name string `json:"name,omitempty"`

	Email                     string                        `json:"email,omitempty"`
	AddressLine1              string                        `json:"address_line1,omitempty"`
	AddressLine2              string                        `json:"address_line2,omitempty"`
	City                      string                        `json:"city,omitempty"`
	Zipcode                   string                        `json:"zipcode,omitempty"`
	State                     string                        `json:"state,omitempty"`
	Country                   string                        `json:"country,omitempty"`
	DefaultCurrency           Currency                      `json:"default_currency,omitempty"`
	LegalName                 string                        `json:"legal_name,omitempty"`
	LegalNumber               string                        `json:"legal_number,omitempty"`
	DocumentNumbering         OrganizationDocumentNumbering `json:"document_numbering,omitempty"`
	DocumentNumberPrefix      string                        `json:"document_number_prefix,omitempty"`
	NetPaymentTerm            int                           `json:"net_payment_term,omitempty"`
	TaxIdentificationNumber   string                        `json:"tax_identification_number,omitempty"`
	WebhookURL                string                        `json:"webhook_url,omitempty"`
	Timezone                  string                        `json:"timezone,omitempty"`
	EmailSettings             []string                      `json:"email_settings,omitempty"`
	FinalizeZeroAmountInvoice bool                          `json:"finalize_zero_amount_invoice,omitempty"`

	BillingConfiguration OrganizationBillingConfigurationInput `json:"billing_configuration,omitempty"`
}

type OrganizationResult struct {
	Organization *Organization `json:"organization,omitempty"`
}

type Organization struct {
	Name string `json:"name,omitempty"`

	Email                     string                        `json:"email,omitempty"`
	AddressLine1              string                        `json:"address_line1,omitempty"`
	AddressLine2              string                        `json:"address_line2,omitempty"`
	City                      string                        `json:"city,omitempty"`
	Zipcode                   string                        `json:"zipcode,omitempty"`
	State                     string                        `json:"state,omitempty"`
	Country                   string                        `json:"country,omitempty"`
	DefaultCurrency           Currency                      `json:"default_currency,omitempty"`
	LegalName                 string                        `json:"legal_name,omitempty"`
	LegalNumber               string                        `json:"legal_number,omitempty"`
	DocumentNumbering         OrganizationDocumentNumbering `json:"document_numbering,omitempty"`
	DocumentNumberPrefix      string                        `json:"document_number_prefix,omitempty"`
	NetPaymentTerm            int                           `json:"net_payment_term,omitempty"`
	TaxIdentificationNumber   string                        `json:"tax_identification_number,omitempty"`
	WebhookURL                string                        `json:"webhook_url,omitempty"`
	WebhookURLs               []string                      `json:"webhook_urls,omitempty"`
	Timezone                  string                        `json:"timezone,omitempty"`
	EmailSettings             []string                      `json:"email_settings,omitempty"`
	FinalizeZeroAmountInvoice bool                          `json:"finalize_zero_amount_invoice,omitempty"`

	BillingConfiguration OrganizationBillingConfiguration `json:"billing_configuration,omitempty"`

	Taxes []*Tax `json:"taxes,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (c *Client) UpdateOrganization(ctx context.Context, organizationInput *OrganizationInput) (*Organization, error) {
	u := c.url("organizations", nil)
	result, err := put[organizationParams, OrganizationResult](
		ctx,
		c,
		u,
		&organizationParams{Organization: organizationInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Organization, nil
}
