package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type PaymentRequestResult struct {
	PaymentRequest  *PaymentRequest   `json:"payment_request,omitempty"`
	PaymentRequests []*PaymentRequest `json:"payment_requests,omitempty"`
}

type PaymentRequestListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`

	ExternalCustomerID string `json:"external_customer_id,omitempty"`
}

func (i *PaymentRequestListInput) query() url.Values {
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

type PaymentRequest struct {
	LagoID         uuid.UUID `json:"lago_id,omitempty"`
	Email          string    `json:"email,omitempty"`
	AmountCurrency Currency  `json:"amount_currency,omitempty"`
	AmountCents    int       `json:"amount_cents,omitempty"`
	PaymentStatus  string    `json:"payment_status,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`

	Customer *Customer  `json:"customer,omitempty"`
	Invoices []*Invoice `json:"fees,omitempty"`
}

type PaymentRequestParams struct {
	PaymentRequest *PaymentRequestInput `json:"payment_request"`
}

type PaymentRequestInput struct {
	Email              string   `json:"email,omitempty"`
	CustomerExternalId string   `json:"customer_external_id,omitempty"`
	LagoInvoiceIds     []string `json:"lago_invoice_ids,omitempty"`
}

func (c *Client) ListPaymentRequests(ctx context.Context, paymentRequestListInput *PaymentRequestListInput) (*PaymentRequestResult, *Error) {
	u := c.url("payment_requests", paymentRequestListInput.query())
	return get[PaymentRequestResult](ctx, c, u)
}

func (c *Client) CreatePaymentRequest(ctx context.Context, paymentRequestInput *PaymentRequestInput) (*PaymentRequest, *Error) {
	u := c.url("payment_requests", nil)
	result, err := post[PaymentRequestParams, PaymentRequestResult](ctx, c, u, &PaymentRequestParams{PaymentRequest: paymentRequestInput})
	if err != nil {
		return nil, err
	}

	return result.PaymentRequest, nil
}
