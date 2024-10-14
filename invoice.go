package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type InvoiceType string

const (
	SubscriptionInvoiceType InvoiceType = "subscription"
	AddOnInvoiceType        InvoiceType = "add_on"
	CreditInvoiceType       InvoiceType = "credit"
	OneOffInvoiceType       InvoiceType = "one_off"
)

type InvoiceStatus string

const (
	InvoiceStatusDraft     InvoiceStatus = "draft"
	InvoiceStatusFinalized InvoiceStatus = "finalized"
	InvoiceStatusFailed    InvoiceStatus = "failed"
)

type InvoicePaymentStatus string

const (
	InvoicePaymentStatusPending   InvoicePaymentStatus = "pending"
	InvoicePaymentStatusSucceeded InvoicePaymentStatus = "succeeded"
	InvoicePaymentStatusFailed    InvoicePaymentStatus = "failed"
)

type InvoiceCreditItemType string

const (
	InvoiceCreditItemCoupon     InvoiceCreditItemType = "coupon"
	InvoiceCreditItemCreditNote InvoiceCreditItemType = "credit_note"
	InvoiceCreditItemInvoice    InvoiceCreditItemType = "invoice"
)

type invoiceResult struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type InvoiceList struct {
	Invoices []*Invoice `json:"invoices,omitempty"`
	Meta     Metadata   `json:"meta,omitempty"`
}

type InvoicePaymentURLResult struct {
	InvoicePaymentURL *InvoicePaymentURL `json:"invoice_payment_url"`
}

type invoiceParams struct {
	Invoice *InvoiceInput `json:"invoice"`
}

type invoiceOneOffParams struct {
	Invoice *InvoiceOneOffInput `json:"invoice"`
}

type InvoiceMetadataInput struct {
	LagoID *uuid.UUID `json:"id,omitempty"`
	Key    string     `json:"key,omitempty"`
	Value  string     `json:"value,omitempty"`
}

type InvoiceFeesInput struct {
	AddOnCode          string   `json:"add_on_code,omitempty"`
	InvoiceDisplayName string   `json:"invoice_display_name,omitempty"`
	UnitAmountCents    int      `json:"unit_amount_cents,omitempty"`
	Description        string   `json:"description,omitempty"`
	Units              float32  `json:"units,omitempty"`
	TaxCodes           []string `json:"tax_codes,omitempty"`
}

type InvoiceMetadataResponse struct {
	LagoID    uuid.UUID `json:"lago_id,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type InvoiceInput struct {
	LagoID        uuid.UUID               `json:"lago_id,omitempty"`
	PaymentStatus InvoicePaymentStatus    `json:"payment_status,omitempty"`
	Metadata      []*InvoiceMetadataInput `json:"metadata,omitempty"`
}

type InvoiceOneOffInput struct {
	ExternalCustomerId string              `json:"external_customer_id,omitempty"`
	Currency           string              `json:"currency,omitempty"`
	Fees               []*InvoiceFeesInput `json:"fees,omitempty"`
}

type InvoiceListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`

	IssuingDateFrom string `json:"issuing_date_from,omitempty"`
	IssuingDateTo   string `json:"issuing_date_to,omitempty"`

	ExternalCustomerID string               `json:"external_customer_id,omitempty"`
	Status             InvoiceStatus        `json:"status,omitempty"`
	PaymentStatus      InvoicePaymentStatus `json:"payment_status,omitempty"`
	PaymentOverdue     bool                 `json:"payment_overdue,omitempty"`
}

func (i *InvoiceListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	if i.IssuingDateFrom != "" {
		q.Add("issuing_date_from", i.IssuingDateFrom)
	}

	if i.IssuingDateTo != "" {
		q.Add("issuing_date_to", i.IssuingDateTo)
	}

	if i.ExternalCustomerID != "" {
		q.Add("external_customer_id", i.ExternalCustomerID)
	}

	if i.Status != "" {
		q.Add("status", string(i.Status))
	}

	if i.PaymentStatus != "" {
		q.Add("payment_status", string(i.PaymentStatus))
	}

	if i.PaymentOverdue {
		q.Add("payment_overdue", strconv.FormatBool(i.PaymentOverdue))
	}

	return q
}

type InvoiceCreditItem struct {
	LagoID uuid.UUID             `json:"lago_id,omitempty"`
	Type   InvoiceCreditItemType `json:"type,omitempty"`
	Code   string                `json:"code,omitempty"`
	Name   string                `json:"name,omitempty"`
}

type InvoiceSummary struct {
	LagoID        uuid.UUID            `json:"lago_id,omitempty"`
	PaymentStatus InvoicePaymentStatus `json:"payment_status,omitempty"`
}

type InvoiceCredit struct {
	Item InvoiceCreditItem `json:"item,omitempty"`

	Invoice InvoiceSummary `json:"invoice,omitempty"`

	LagoItemID     uuid.UUID `json:"lago_item_id,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
	BeforeTaxes    bool      `json:"before_taxes,omitempty"`
}

type InvoiceAppliedTax struct {
	LagoId          uuid.UUID `json:"lago_id,omitempty"`
	LagoInvoiceId   uuid.UUID `json:"lago_invoice_id,omitempty"`
	LagoTaxId       uuid.UUID `json:"lago_tax_id,omitempty"`
	TaxName         string    `json:"tax_name,omitempty"`
	TaxCode         string    `json:"tax_code,omitempty"`
	TaxRate         float32   `json:"tax_rate,omitempty"`
	TaxDescription  string    `json:"tax_description,omitempty"`
	AmountCents     int       `json:"amount_cents,omitempty"`
	AmountCurrency  Currency  `json:"amount_currency,omitempty"`
	FeesAmountCents int       `json:"fees_amount_cents,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

type InvoiceErrorDetail struct {
	LagoId    uuid.UUID `json:"lago_id,omitempty"`
	ErrorCode string    `json:"error_code,omitempty"`
	Details   string    `json:"details,omitempty"`
}

type Invoice struct {
	LagoID       uuid.UUID `json:"lago_id,omitempty"`
	SequentialID int       `json:"sequential_id,omitempty"`
	Number       string    `json:"number,omitempty"`

	IssuingDate          string    `json:"issuing_date,omitempty"`
	PaymentDisputeLostAt time.Time `json:"payment_dispute_lost_at,omitempty"`
	PaymentDueDate       string    `json:"payment_due_date,omitempty"`
	PaymentOverdue       bool      `json:"payment_overdue,omitempty"`

	InvoiceType   InvoiceType          `json:"invoice_type,omitempty"`
	Status        InvoiceStatus        `json:"status,omitempty"`
	PaymentStatus InvoicePaymentStatus `json:"payment_status,omitempty"`

	Currency Currency `json:"currency,omitempty"`

	FeesAmountCents                     int `json:"fees_amount_cents,omitempty"`
	TaxesAmountCents                    int `json:"taxes_amount_cents,omitempty"`
	CouponsAmountCents                  int `json:"coupons_amount_cents,omitempty"`
	CreditNotesAmountCents              int `json:"credit_notes_amount_cents,omitempty"`
	SubTotalExcludingTaxesAmountCents   int `json:"sub_total_excluding_taxes_amount_cents,omitempty"`
	SubTotalIncludingTaxesAmountCents   int `json:"sub_total_including_taxes_amount_cents,omitempty"`
	TotalAmountCents                    int `json:"total_amount_cents,omitempty"`
	PrepaidCreditAmountCents            int `json:"prepaid_credit_amount_cents,omitempty"`
	ProgressiveBillingCreditAmountCents int `json:"progressive_billing_credit_amount_cents"`
	NetPaymentTerm                      int `json:"net_payment_term,omitempty"`

	FileURL       string                     `json:"file_url,omitempty"`
	Metadata      []*InvoiceMetadataResponse `json:"metadata,omitempty"`
	VersionNumber int                        `json:"version_number,omitempty"`

	Customer      *Customer       `json:"customer,omitempty"`
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`

	Fees                  []*Fee                   `json:"fees,omitempty"`
	Credits               []*InvoiceCredit         `json:"credits,omitempty"`
	AppliedTaxes          []*InvoiceAppliedTax     `json:"applied_taxes,omitempty"`
	ErrorDetails          []*InvoiceErrorDetail    `json:"error_details,omitempty"`
	AppliedUsageThreshold []*AppliedUsageThreshold `json:"applied_usage_threshold,omitempty"`
}

type InvoicePaymentURL struct {
	PaymentURL string `json:"payment_url,omitempty"`
}

func (c *Client) GetInvoice(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID, nil)
	result, err := get[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) ListInvoice(ctx context.Context, invoiceListInput *InvoiceListInput) (*InvoiceList, *Error) {
	u := c.url("invoices", invoiceListInput.query())
	return get[InvoiceList](ctx, c, u)
}

func (c *Client) CreateInvoice(ctx context.Context, oneOffInput *InvoiceOneOffInput) (*Invoice, *Error) {
	u := c.url("invoices", nil)
	result, err := post[invoiceOneOffParams, invoiceResult](ctx, c, u, &invoiceOneOffParams{Invoice: oneOffInput})
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) UpdateInvoice(ctx context.Context, invoiceInput *InvoiceInput) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceInput.LagoID.String(), nil)
	result, err := put[invoiceParams, invoiceResult](ctx, c, u, &invoiceParams{Invoice: invoiceInput})
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) DownloadInvoice(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID+"/download", nil)
	result, err := postWithoutBody[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) RefreshInvoice(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID+"/refresh", nil)
	result, err := putWithoutBody[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) RetryInvoice(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID+"/retry", nil)
	result, err := postWithoutBody[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) FinalizeInvoice(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID+"/finalize", nil)
	result, err := putWithoutBody[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) LoseInvoiceDispute(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID+"/lose_dispute", nil)
	result, err := putWithoutBody[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) RetryInvoicePayment(ctx context.Context, invoiceID string) (*Invoice, *Error) {
	u := c.url("invoices/"+invoiceID+"/retry_payment", nil)
	result, err := postWithoutBody[invoiceResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Invoice, nil
}

func (c *Client) GetInvoicePaymentURL(ctx context.Context, invoiceID string) (*InvoicePaymentURL, *Error) {
	u := c.url("invoices/"+invoiceID+"/payment_url", nil)
	result, err := postWithoutBody[InvoicePaymentURLResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.InvoicePaymentURL, nil
}
