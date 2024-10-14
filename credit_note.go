package lago

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CreditNoteCreditStatus string
type CreditNoteRefundStatus string
type CreditNoteReason string

const (
	CreditNoteCreditStatusAvailable CreditNoteCreditStatus = "available"
	CreditNoteCreditStatusConsumed  CreditNoteCreditStatus = "consumed"
)

const (
	CreditNoteRefundStatusPending  CreditNoteRefundStatus = "pending"
	CreditNoteRefundStatusRefunded CreditNoteRefundStatus = "refunded"
)

const (
	CreditNoteReasonDuplicatedCharge      CreditNoteReason = "duplicated_charge"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "product_unsatisfactory"
	CreditNoteReasonOrderChange           CreditNoteReason = "order_change"
	CreditNoteReasonOrderCancellation     CreditNoteReason = "order_cancellation"
	CreditNoteReasonFraudulentCharge      CreditNoteReason = "fraudulent_charge"
	CreditNoteReasonOther                 CreditNoteReason = "other"
)

type creditNoteParams struct {
	CreditNote *CreditNoteInput `json:"credit_note"`
}

type creditNoteResult struct {
	CreditNote *CreditNote `json:"credit_note,omitempty"`
}

type CreditNoteList struct {
	CreditNotes []*CreditNote `json:"credit_notes,omitempty"`
	Meta        Metadata      `json:"meta,omitempty"`
}

type CreditNoteEstimatedResult struct {
	CreditNoteEstimated *CreditNoteEstimated `json:"credit_note_estimated"`
}

type CreditListInput struct {
	PerPage            int    `json:"per_page,omitempty,string"`
	Page               int    `json:"page,omitempty,string"`
	ExternalCustomerID string `json:"external_customer_id,omitempty"`
}

func (i *CreditListInput) query() url.Values {
	q := make(url.Values)
	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}
	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}
	if i.ExternalCustomerID != "" {
		q.Add("external_customer_id", i.ExternalCustomerID)
	}
	return q
}

type CreditNoteItem struct {
	LagoID         uuid.UUID `json:"lago_id,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
	Fee            Fee       `json:"fee,omitempty"`
}

type CreditNoteAppliedTax struct {
	LagoId           uuid.UUID `json:"lago_id,omitempty"`
	LagoCreditNoteId uuid.UUID `json:"lago_credit_note_id,omitempty"`
	LagoTaxId        uuid.UUID `json:"lago_tax_id,omitempty"`
	TaxName          string    `json:"tax_name,omitempty"`
	TaxCode          string    `json:"tax_code,omitempty"`
	TaxRate          float32   `json:"tax_rate,omitempty"`
	TaxDescription   string    `json:"tax_description,omitempty"`
	AmountCents      int       `json:"amount_cents,omitempty"`
	AmountCurrency   Currency  `json:"amount_currency,omitempty"`
	BaseAmountCents  int       `json:"base_amount_cents,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}

type CreditNote struct {
	LagoID        uuid.UUID        `json:"lago_id,omitempty"`
	SequentialID  int              `json:"sequential_id,omitempty"`
	Number        string           `json:"number,omitempty"`
	LagoInvoiceID uuid.UUID        `json:"lago_invoice_id,omitempty"`
	InvoiceNumber string           `json:"invoice_number,omitempty"`
	Reason        CreditNoteReason `json:"reason,omitempty"`

	CreditStatus CreditNoteCreditStatus `json:"credit_status,omitempty"`
	RefundStatus CreditNoteRefundStatus `json:"refund_status,omitempty"`

	Currency                          Currency `json:"currency,omitempty"`
	TotalAmountCents                  int      `json:"total_amount_cents,omitempty"`
	CreditAmountCents                 int      `json:"credit_amount_cents,omitempty"`
	BalanceAmountCents                int      `json:"balance_amount_cents,omitempty"`
	RefundAmountCents                 int      `json:"refund_amount_cents,omitempty"`
	TaxesAmountCents                  int      `json:"taxes_amount_cents,omitempty"`
	TaxesRate                         float32  `json:"taxes_rate,omitempty"`
	SubTotalExcludingTaxesAmountCents int      `json:"sub_total_excluding_taxes_amount_cents,omitempty"`
	CouponsAdjustmentAmountCents      int      `json:"coupons_adjustment_amount_cents,omitempty"`

	FileURL string `json:"file_url,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Items []*CreditNoteItem `json:"items,omitempty"`
}

type CreditNoteEstimated struct {
	LagoInvoiceID uuid.UUID `json:"lago_invoice_id,omitempty"`
	InvoiceNumber string    `json:"invoice_number,omitempty"`

	Currency                          Currency `json:"currency,omitempty"`
	MaxCreditableAmountCents          int      `json:"max_creditable_amount_cents,omitempty"`
	MaxRefundableAmountCents          int      `json:"max_refundable_amount_cents,omitempty"`
	TaxesAmountCents                  int      `json:"taxes_amount_cents,omitempty"`
	TaxesRate                         float32  `json:"taxes_rate,omitempty"`
	SubTotalExcludingTaxesAmountCents int      `json:"sub_total_excluding_taxes_amount_cents,omitempty"`
	CouponsAdjustmentAmountCents      int      `json:"coupons_adjustment_amount_cents,omitempty"`

	Items []*CreditNoteEstimatedItem `json:"items,omitempty"`

	AppliedTaxes []*CreditNoteEstimatedAppliedTax `json:"applied_taxes,omitempty"`
}

type CreditNoteEstimatedItem struct {
	AmountCents int       `json:"amount_cents,omitempty"`
	LagoFeeID   uuid.UUID `json:"lago_fee_id,omitempty"`
}

type CreditNoteEstimatedAppliedTax struct {
	LagoTaxId      uuid.UUID `json:"lago_tax_id,omitempty"`
	TaxName        string    `json:"tax_name,omitempty"`
	TaxCode        string    `json:"tax_code,omitempty"`
	TaxRate        float32   `json:"tax_rate,omitempty"`
	TaxDescription string    `json:"tax_description,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
}

type CreditNoteItemInput struct {
	LagoFeeID   uuid.UUID `json:"fee_id,omitempty"`
	AmountCents int       `json:"amount_cents,omitempty"`
}

type CreditNoteInput struct {
	LagoInvoiceID     uuid.UUID              `json:"invoice_id,omitempty"`
	Reason            CreditNoteReason       `json:"reason,omitempty"`
	Items             []*CreditNoteItemInput `json:"items,omitempty"`
	CreditAmountCents int                    `json:"refund_amount_cents,omitempty"`
	RefundAmountCents int                    `json:"credit_amount_cents,omitempty"`
}

type CreditNoteUpdateInput struct {
	LagoID       string                 `json:"id,omitempty"`
	RefundStatus CreditNoteRefundStatus `json:"refund_status,omitempty"`
}

type creditNoteUpdateParams struct {
	CreditNote *CreditNoteUpdateInput `json:"credit_note"`
}

type CreditNoteEstimateInput struct {
	LagoInvoiceID uuid.UUID              `json:"invoice_id,omitempty"`
	Items         []*CreditNoteItemInput `json:"items,omitempty"`
}

type CreditNoteEstimateParams struct {
	CreditNote *CreditNoteEstimateInput `json:"credit_note"`
}

func (c *Client) GetCreditNote(ctx context.Context, creditNoteID uuid.UUID) (*CreditNote, *Error) {
	u := fmt.Sprintf("credit_notes/%s", creditNoteID)
	result, err := get[creditNoteResult](ctx, c, u)
	if err != nil {
		return nil, err
	}
	return result.CreditNote, nil
}

func (c *Client) DownloadCreditNote(ctx context.Context, creditNoteID string) (*CreditNote, *Error) {
	u := c.url("credit_notes/"+creditNoteID+"/download", nil)
	result, err := postWithoutBody[creditNoteResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.CreditNote, nil
}

func (c *Client) ListCreditNotes(ctx context.Context, creditNoteListInput *CreditListInput) (*CreditNoteList, *Error) {
	u := c.url("credit_notes", creditNoteListInput.query())
	return get[CreditNoteList](ctx, c, u)
}

func (c *Client) CreateCreditNote(ctx context.Context, creditNoteInput *CreditNoteInput) (*CreditNote, *Error) {
	u := c.url("credit_notes", nil)
	result, err := post[creditNoteParams, creditNoteResult](
		ctx,
		c,
		u,
		&creditNoteParams{CreditNote: creditNoteInput},
	)
	if err != nil {
		return nil, err
	}

	return result.CreditNote, nil
}

func (c *Client) UpdateCreditNote(ctx context.Context, creditNoteUpdateInput *CreditNoteUpdateInput) (*CreditNote, *Error) {
	u := c.url("credit_notes/"+creditNoteUpdateInput.LagoID, nil)

	result, err := put[creditNoteUpdateParams, creditNoteResult](
		ctx,
		c,
		u,
		&creditNoteUpdateParams{CreditNote: creditNoteUpdateInput},
	)
	if err != nil {
		return nil, err
	}

	return result.CreditNote, nil
}

func (c *Client) VoidCreditNote(ctx context.Context, creditNoteID string) (*CreditNote, *Error) {
	u := c.url("credit_notes/"+creditNoteID+"/void", nil)
	result, err := putWithoutBody[creditNoteResult](ctx, c, u)

	if err != nil {
		return nil, err
	}

	return result.CreditNote, nil
}

func (c *Client) EstimateCreditNote(ctx context.Context, creditNoteEstimateInput *CreditNoteEstimateInput) (*CreditNoteEstimated, *Error) {
	u := c.url("credit_notes/estimate", nil)

	result, err := post[CreditNoteEstimateParams, CreditNoteEstimatedResult](
		ctx,
		c,
		u,
		&CreditNoteEstimateParams{CreditNote: creditNoteEstimateInput},
	)
	if err != nil {
		return nil, err
	}

	return result.CreditNoteEstimated, nil
}
