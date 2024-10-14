package lago

import (
	"context"
	"net/url"
	"strconv"
)

type PaymentStatus string

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

type InvoiceCollectionList struct {
	InvoiceCollections []*InvoiceCollection `json:"invoice_collections,omitempty"`
}

type InvoiceCollection struct {
	Month          string               `json:"month,omitempty"`
	PaymentStatus  InvoicePaymentStatus `json:"payment_status,omitempty"`
	InvoicesCount  int                  `json:"invoices_count,omitempty"`
	AmountCents    int                  `json:"amount_cents,omitempty"`
	AmountCurrency Currency             `json:"currency,omitempty"`
}

func (c *Client) ListInvoiceCollections(ctx context.Context, InvoiceCollectionListInput *InvoiceCollectionListInput) (*InvoiceCollectionList, *Error) {
	u := c.url("analytics/invoice_collection", InvoiceCollectionListInput.query())
	return get[InvoiceCollectionList](ctx, c, u)
}
