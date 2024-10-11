package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type FeeType string
type FeePaymentStatus string
type FeeItemType string

const (
	FeeItemSubscription FeeType = "subscription"
	FeeItemCharge       FeeType = "charge"
	FeeItemAddOn        FeeType = "add_on"
)

const (
	FeePaymentStatusPending   FeePaymentStatus = "pending"
	FeePaymentStatusSucceeded FeePaymentStatus = "succeeded"
	FeePaymentStatusFailed    FeePaymentStatus = "failed"
	FeePaymentStatusRefunded  FeePaymentStatus = "refunded"
)

const (
	FeeAddOnType         FeeItemType = "AddOn"
	FeeBillableMetric    FeeItemType = "BillableMetric"
	FeeSubscription      FeeItemType = "Subscription"
	FeeWalletTransaction FeeItemType = "WalletTransaction"
)

type FeeRequest struct {
	client *Client
}

type FeeResult struct {
	Fee  *Fee     `json:"fee,omitempty"`
	Fees []*Fee   `json:"fees,omitempty"`
	Meta Metadata `json:"meta,omitempty"`
}

type FeeUpdateParams struct {
	Fee *FeeUpdateInput `json:"fee"`
}

type FeeUpdateInput struct {
	LagoID        uuid.UUID        `json:"lago_id,omitempty"`
	PaymentStatus FeePaymentStatus `json:"payment_status,omitempty"`
}

type FeeListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`

	FeeType       FeeType          `json:"fee_type,omitempty"`
	PaymentStatus FeePaymentStatus `json:"payment_status,omitempty"`

	ExternalSubscriptionID string `json:"external_subscription_id,omitempty"`
	ExternalCustomerID     string `json:"external_customer_id,omitempty"`

	BillableMetricCode string `json:"billable_metric_code,omitempty"`

	Currency Currency `json:"currency"`

	CreatedAtFrom   string `json:"created_at_from,omitempty"`
	CreatedAtTo     string `json:"created_at_to,omitempty"`
	FailedAtFrom    string `json:"failed_at_from,omitempty"`
	FailedAtTo      string `json:"failed_at_to,omitempty"`
	SucceededAtFrom string `json:"succeeded_at_from,omitempty"`
	SucceededAtTo   string `json:"succeeded_at_to,omitempty"`
	RefundedAtFrom  string `json:"refunded_at_from,omitempty"`
	RefundedAtTo    string `json:"refunded_at_to,omitempty"`
}

func (i *FeeListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	if i.FeeType != "" {
		q.Add("fee_type", string(i.FeeType))
	}

	if i.PaymentStatus != "" {
		q.Add("payment_status", string(i.PaymentStatus))
	}

	if i.ExternalSubscriptionID != "" {
		q.Add("external_subscription_id", i.ExternalSubscriptionID)
	}

	if i.ExternalCustomerID != "" {
		q.Add("external_customer_id", i.ExternalCustomerID)
	}

	if i.BillableMetricCode != "" {
		q.Add("billable_metric_code", i.BillableMetricCode)
	}

	if i.Currency != "" {
		q.Add("currency", string(i.Currency))
	}

	if i.CreatedAtFrom != "" {
		q.Add("created_at_from", i.CreatedAtFrom)
	}

	if i.CreatedAtTo != "" {
		q.Add("created_at_to", i.CreatedAtTo)
	}

	if i.FailedAtFrom != "" {
		q.Add("failed_at_from", i.FailedAtFrom)
	}

	if i.FailedAtTo != "" {
		q.Add("failed_at_to", i.FailedAtTo)
	}

	if i.SucceededAtFrom != "" {
		q.Add("succeeded_at_from", i.SucceededAtFrom)
	}

	if i.SucceededAtTo != "" {
		q.Add("succeeded_at_to", i.SucceededAtTo)
	}

	if i.RefundedAtFrom != "" {
		q.Add("refunded_at_from", i.RefundedAtFrom)
	}

	if i.RefundedAtTo != "" {
		q.Add("refunded_at_to", i.RefundedAtTo)
	}

	return q
}

type FeeItem struct {
	Type                     FeeType                `json:"type,omitempty"`
	Code                     string                 `json:"code,omitempty"`
	Name                     string                 `json:"name,omitempty"`
	InvoiceDisplayName       string                 `json:"invoice_display_name,omitempty"`
	FilterInvoiceDisplayName string                 `json:"filter_invoice_display_name,omitempty"`
	Filters                  map[string]interface{} `json:"filters,omitempty"`
	LagoItemID               uuid.UUID              `json:"lago_item_id,omitempty"`
	ItemType                 FeeItemType            `json:"item_type,omitempty"`
	GroupedBy                map[string]interface{} `json:"grouped_by,omitempty"`
}

type FeeAppliedTax struct {
	LagoId         uuid.UUID `json:"lago_id,omitempty"`
	LagoFeeId      uuid.UUID `json:"lago_fee_id,omitempty"`
	LagoTaxId      uuid.UUID `json:"lago_tax_id,omitempty"`
	TaxName        string    `json:"tax_name,omitempty"`
	TaxCode        string    `json:"tax_code,omitempty"`
	TaxRate        float32   `json:"tax_rate,omitempty"`
	TaxDescription string    `json:"tax_description,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type Fee struct {
	LagoID                 uuid.UUID `json:"lago_id,omitempty"`
	LagoChargeFilterID     uuid.UUID `json:"lago_charge_filter_id,omitempty"`
	LagoInvoiceID          uuid.UUID `json:"lago_invoice_id,omitempty"`
	LagoTrueUpFeeID        uuid.UUID `json:"lago_true_up_fee_id,omitempty"`
	LagoTrueUpParentFeeID  uuid.UUID `json:"lago_true_up_parent_fee_id,omitempty"`
	ExternalSubscriptionID string    `json:"external_subscription_id,omitempty"`

	AmountCents         int                    `json:"amount_cents,omitempty"`
	AmountDetails       map[string]interface{} `json:"amount_details,omitempty"`
	PreciseUnitAmount   string                 `json:"precise_unit_amount,omitempty"`
	PreciseAmount       string                 `json:"precise_amount,omitempty"`
	PreciseTotalAmount  string                 `json:"precise_total_amount,omitempty"`
	AmountCurrency      string                 `json:"amount_currency,omitempty"`
	TaxesAmountCents    int                    `json:"taxes_amount_cents,omitempty"`
	TaxesPreciseAmount  string                 `json:"taxes_precise_amount,omitempty"`
	TaxesRate           float32                `json:"taxes_rate,omitempty"`
	TotalAmountCents    int                    `json:"total_amount_cents,omitempty"`
	TotalAmountCurrency string                 `json:"total_amount_currency,omitempty"`
	PayInAdvance        bool                   `json:"pay_in_advance,omitempty"`
	Invoiceable         bool                   `json:"invoiceable,omitempty"`
	FromDate            string                 `json:"from_date,omitempty"`
	ToDate              string                 `json:"to_date,omitempty"`
	InvoiceDisplayName  string                 `json:"invoice_display_name,omitempty"`

	Units       string `json:"units,omitempty"`
	Description string `json:"description,omitempty"`
	EventsCount int    `json:"events_count,omitempty"`

	PaymentStatus FeePaymentStatus `json:"payment_status,omitempty"`

	CreatedAt   time.Time `json:"created_at,omitempty"`
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
	FailedAt    time.Time `json:"failed_at,omitempty"`
	RefundedAt  time.Time `json:"refunded_at,omitempty"`

	Item         FeeItem          `json:"item,omitempty"`
	AppliedTaxes []*FeeAppliedTax `json:"applied_taxes,omitempty"`
}

func (c *Client) Fee() *FeeRequest {
	return &FeeRequest{
		client: c,
	}
}

func (fr *FeeRequest) Get(ctx context.Context, feeID string) (*Fee, *Error) {
	u := fr.client.url("fees/"+feeID, nil)
	result, err := get[FeeResult](ctx, fr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Fee, nil
}

func (fr *FeeRequest) Update(ctx context.Context, feeInput *FeeUpdateInput) (*Fee, *Error) {
	u := fr.client.url("fees/"+feeInput.LagoID.String(), nil)
	result, err := put[FeeUpdateParams, FeeResult](ctx, fr.client, u, &FeeUpdateParams{Fee: feeInput})
	if err != nil {
		return nil, err
	}

	return result.Fee, nil
}

func (fr *FeeRequest) GetList(ctx context.Context, feeListInput *FeeListInput) (*FeeResult, *Error) {
	u := fr.client.url("fees", feeListInput.query())
	return get[FeeResult](ctx, fr.client, u)
}

func (fr *FeeRequest) Delete(ctx context.Context, feeID string) (*Fee, *Error) {
	u := fr.client.url("fees/"+feeID, nil)
	result, err := delete[FeeResult](ctx, fr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Fee, nil
}
