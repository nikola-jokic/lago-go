package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type TaxRequest struct {
	client *Client
}

type TaxParams struct {
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

type TaxResult struct {
	Tax   *Tax     `json:"tax,omitempty"`
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

func (c *Client) Tax() *TaxRequest {
	return &TaxRequest{
		client: c,
	}
}

func (adr *TaxRequest) Get(ctx context.Context, taxCode string) (*Tax, *Error) {
	u := adr.client.url("taxes/"+taxCode, nil)
	result, err := get[TaxResult](ctx, adr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}

func (adr *TaxRequest) GetList(ctx context.Context, taxListInput *TaxListInput) (*TaxResult, *Error) {
	u := adr.client.url("taxes", taxListInput.query())
	return get[TaxResult](ctx, adr.client, u)
}

func (adr *TaxRequest) Create(ctx context.Context, taxInput *TaxInput) (*Tax, *Error) {
	u := adr.client.url("taxes", nil)

	result, err := post[TaxParams, TaxResult](ctx, adr.client, u, &TaxParams{Tax: taxInput})
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}

func (adr *TaxRequest) Update(ctx context.Context, taxInput *TaxInput) (*Tax, *Error) {
	u := adr.client.url("taxes/"+taxInput.Code, nil)

	result, err := put[TaxParams, TaxResult](ctx, adr.client, u, &TaxParams{Tax: taxInput})
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}

func (adr *TaxRequest) Delete(ctx context.Context, taxCode string) (*Tax, *Error) {
	u := adr.client.url("taxes/"+taxCode, nil)

	result, err := delete[TaxResult](ctx, adr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Tax, nil
}
