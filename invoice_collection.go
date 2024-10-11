package lago

import (
	"context"
	"net/url"
	"strconv"
)

type PaymentStatus string

type InvoiceCollectionRequest struct {
	client *Client
}

type InvoiceCollectionListInput struct {
	AmountCurrency string `json:"currency,omitempty"`
	Months         int    `json:"months,omitempty,string"`
}

func (i *InvoiceCollectionListInput) query() url.Values {
	q := make(url.Values)

	if i.AmountCurrency != "" {
		q.Add("currency", i.AmountCurrency)
	}

	if i.Months > 0 {
		q.Add("months", strconv.Itoa(i.Months))
	}

	return q
}

type InvoiceCollectionResult struct {
	InvoiceCollection  *InvoiceCollection   `json:"invoice_collection,omitempty"`
	InvoiceCollections []*InvoiceCollection `json:"invoice_collections,omitempty"`
}

type InvoiceCollection struct {
	Month          string               `json:"month,omitempty"`
	PaymentStatus  InvoicePaymentStatus `json:"payment_status,omitempty"`
	InvoicesCount  int                  `json:"invoices_count,omitempty"`
	AmountCents    int                  `json:"amount_cents,omitempty"`
	AmountCurrency Currency             `json:"currency,omitempty"`
}

func (c *Client) InvoiceCollection() *InvoiceCollectionRequest {
	return &InvoiceCollectionRequest{
		client: c,
	}
}

func (adr *InvoiceCollectionRequest) GetList(ctx context.Context, InvoiceCollectionListInput *InvoiceCollectionListInput) (*InvoiceCollectionResult, *Error) {
	u := adr.client.url("analytics/invoice_collection", InvoiceCollectionListInput.query())
	return get[InvoiceCollectionResult](ctx, adr.client, u)
}
