package lago

import (
	"context"
	"net/url"
	"strconv"
)

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

type GrossRevenueList struct {
	GrossRevenues []*GrossRevenue `json:"gross_revenues,omitempty"`
}

type GrossRevenue struct {
	Month          string   `json:"month,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"currency,omitempty"`
	InvoicesCount  int      `json:"invoices_count,omitempty"`
}

func (c *Client) ListGrossRevenues(ctx context.Context, GrossRevenueListInput *GrossRevenueListInput) (*GrossRevenueList, error) {
	u := c.url("analytics/gross_revenue", GrossRevenueListInput.query())
	return get[GrossRevenueList](ctx, c, u)
}
