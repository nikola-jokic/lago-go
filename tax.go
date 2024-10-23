package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type taxParams struct {
	Tax *TaxInput `json:"tax"`
}

type TaxInput struct {
	Name                  string   `json:"name,omitempty"`
	Code                  string   `json:"code,omitempty"`
	Rate                  *float32 `json:"rate,omitempty"`
	Description           string   `json:"description,omitempty"`
	AppliedToOrganization bool     `json:"applied_to_organization,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}

type TaxListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *TaxListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	return q
}

type taxResult struct {
	Tax *Tax `json:"tax,omitempty"`
}

type TaxList struct {
	Taxes []*Tax   `json:"taxes,omitempty"`
	Meta  Metadata `json:"meta,omitempty"`
}

type Tax struct {
	LagoID                uuid.UUID `json:"lago_id,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Code                  string    `json:"code,omitempty"`
	Rate                  float32   `json:"rate,omitempty"`
	Description           string    `json:"description,omitempty"`
	AddOnsCount           int       `json:"add_ons_count,omitempty"`
	CustomersCount        int       `json:"customers_count,omitempty"`
	PlansCount            int       `json:"plans_count,omitempty"`
	ChargesCount          int       `json:"charges_count,omitempty"`
	AppliedToOrganization bool      `json:"applied_to_organization,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

func (c *Client) GetTax(ctx context.Context, taxCode string) (*Tax, error) {
	u := c.url("taxes/"+taxCode, nil)
	result, err := get[taxResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}

func (c *Client) ListTaxes(ctx context.Context, taxListInput *TaxListInput) (*TaxList, error) {
	u := c.url("taxes", taxListInput.query())
	return get[TaxList](ctx, c, u)
}

func (c *Client) CreateTax(ctx context.Context, taxInput *TaxInput) (*Tax, error) {
	u := c.url("taxes", nil)

	result, err := post[taxParams, taxResult](
		ctx,
		c,
		u,
		&taxParams{Tax: taxInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}

func (c *Client) UpdateTax(ctx context.Context, taxInput *TaxInput) (*Tax, error) {
	u := c.url("taxes/"+taxInput.Code, nil)

	result, err := put[taxParams, taxResult](
		ctx,
		c,
		u,
		&taxParams{Tax: taxInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}

func (c *Client) DeleteTax(ctx context.Context, taxCode string) (*Tax, error) {
	u := c.url("taxes/"+taxCode, nil)

	result, err := delete[taxResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}
