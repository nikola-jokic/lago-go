package lago

import (
	"context"
	"net/url"
	"strconv"
)

type MrrListInput struct {
	AmountCurrency string `json:"currency,omitempty"`
	Months         int    `json:"months,omitempty,string"`
}

func (i *MrrListInput) query() url.Values {
	q := make(url.Values)

	if i.AmountCurrency != "" {
		q.Add("currency", i.AmountCurrency)
	}

	if i.Months > 0 {
		q.Add("months", strconv.Itoa(i.Months))
	}

	return q
}

type MrrList struct {
	Mrrs []*Mrr `json:"mrrs,omitempty"`
}

type Mrr struct {
	Month          string   `json:"month,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"currency,omitempty"`
}

func (c *Client) ListMrrs(ctx context.Context, MrrListInput *MrrListInput) (*MrrList, *Error) {
	u := c.url("analytics/mrr", MrrListInput.query())
	return get[MrrList](ctx, c, u)
}
