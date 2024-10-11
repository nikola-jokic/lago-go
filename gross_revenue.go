package lago

import (
	"context"
	"net/url"
	"strconv"
)

type GrossRevenueRequest struct {
	client *Client
}

type GrossRevenueListInput struct {
	AmountCurrency     string `json:"currency,omitempty"`
	ExternalCustomerId string `json:"external_customer_id,omitempty"`
	Months             int    `json:"months,omitempty,string"`
}

func (i *GrossRevenueListInput) query() url.Values {
	q := make(url.Values)

	if i.AmountCurrency != "" {
		q.Add("currency", i.AmountCurrency)
	}

	if i.ExternalCustomerId != "" {
		q.Add("external_customer_id", i.ExternalCustomerId)
	}

	if i.Months > 0 {
		q.Add("months", strconv.Itoa(i.Months))
	}

	return q
}

type GrossRevenueResult struct {
	GrossRevenue  *GrossRevenue   `json:"gross_revenue,omitempty"`
	GrossRevenues []*GrossRevenue `json:"gross_revenues,omitempty"`
}

type GrossRevenue struct {
	Month          string   `json:"month,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"currency,omitempty"`
	InvoicesCount  int      `json:"invoices_count,omitempty"`
}

func (c *Client) GrossRevenue() *GrossRevenueRequest {
	return &GrossRevenueRequest{
		client: c,
	}
}

func (adr *GrossRevenueRequest) GetList(ctx context.Context, GrossRevenueListInput *GrossRevenueListInput) (*GrossRevenueResult, *Error) {
	u := adr.client.url("analytics/gross_revenue", GrossRevenueListInput.query())
	return get[GrossRevenueResult](ctx, adr.client, u)
}
