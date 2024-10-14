package lago

import (
	"context"
	"net/url"
	"strconv"
)

type OverdueBalanceListInput struct {
	AmountCurrency     string `json:"currency,omitempty"`
	ExternalCustomerId string `json:"external_customer_id,omitempty"`
	Months             int    `json:"months,omitempty,string"`
}

func (i *OverdueBalanceListInput) query() url.Values {
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

type OverdueBalanceList struct {
	OverdueBalances []*OverdueBalance `json:"overdue_balances,omitempty"`
}

type OverdueBalance struct {
	Month          string   `json:"month,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"currency,omitempty"`
}

func (c *Client) ListOverdueBalances(ctx context.Context, OverdueBalanceListInput *OverdueBalanceListInput) (*OverdueBalanceList, *Error) {
	u := c.url("analytics/overdue_balance", OverdueBalanceListInput.query())
	return get[OverdueBalanceList](ctx, c, u)
}
