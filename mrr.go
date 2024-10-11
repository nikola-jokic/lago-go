package lago

import (
	"context"
	"net/url"
	"strconv"
)

type MrrRequest struct {
	client *Client
}

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

type MrrResult struct {
	Mrr  *Mrr   `json:"mrr,omitempty"`
	Mrrs []*Mrr `json:"mrrs,omitempty"`
}

type Mrr struct {
	Month          string   `json:"month,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"currency,omitempty"`
}

func (c *Client) Mrr() *MrrRequest {
	return &MrrRequest{
		client: c,
	}
}

func (adr *MrrRequest) GetList(ctx context.Context, MrrListInput *MrrListInput) (*MrrResult, *Error) {
	u := adr.client.url("analytics/mrr", MrrListInput.query())
	return get[MrrResult](ctx, adr.client, u)
}
