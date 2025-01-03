package lago

import (
	"context"
	"net/url"
	"strconv"
)

type InvoicedUsageListInput struct {
	AmountCurrency string `json:"currency,omitempty"`
	Months         int    `json:"months,omitempty,string"`
}

func (i *InvoicedUsageListInput) query() url.Values {
	q := make(url.Values)

	if i.AmountCurrency != "" {
		q.Add("currency", i.AmountCurrency)
	}

	if i.Months > 0 {
		q.Add("months", strconv.Itoa(i.Months))
	}

	return q
}

type InvoicedUsageList struct {
	InvoicedUsages []*InvoicedUsage `json:"invoiced_usages,omitempty"`
}

type InvoicedUsage struct {
	Month          string   `json:"month,omitempty"`
	Code           string   `json:"code,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"currency,omitempty"`
}

func (c *Client) ListInvoiceUsages(ctx context.Context, invoicedUsageListInput *InvoicedUsageListInput) (*InvoicedUsageList, error) {
	u := c.url("analytics/invoiced_usage", invoicedUsageListInput.query())
	return get[InvoicedUsageList](ctx, c, u)
}
