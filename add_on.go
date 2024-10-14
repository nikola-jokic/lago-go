package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type addOnParams struct {
	AddOn *AddOnInput `json:"add_on"`
}

type AddOnInput struct {
	Name               string `json:"name,omitempty"`
	InvoiceDisplayName string `json:"invoice_display_name,omitempty"`
	Code               string `json:"code,omitempty"`
	Description        string `json:"description,omitempty"`

	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"amount_currency,omitempty"`

	TaxCodes []string `json:"tax_codes,omitempty"`
}

type AddOnListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *AddOnListInput) query() url.Values {
	q := make(url.Values)
	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}
	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}
	return q
}

type AddOnList struct {
	AddOns []*AddOn `json:"add_ons,omitempty"`
	Meta   Metadata `json:"meta,omitempty"`
}

type addOnResult struct {
	AddOn *AddOn `json:"add_on,omitempty"`
}

type AddOn struct {
	LagoID             uuid.UUID `json:"lago_id,omitempty"`
	Name               string    `json:"name,omitempty"`
	InvoiceDisplayName string    `json:"invoice_display_name,omitempty"`
	Code               string    `json:"code,omitempty"`
	Description        string    `json:"description,omitempty"`

	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"amount_currency,omitempty"`

	Taxes []*Tax `json:"tax,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (c *Client) GetAddOn(ctx context.Context, addOnCode string) (*AddOn, *Error) {
	u := c.url("add_ons/"+addOnCode, nil)
	result, err := get[addOnResult](ctx, c, u)
	if err != nil {
		return nil, err
	}
	return result.AddOn, nil
}

func (c *Client) ListAddOns(ctx context.Context, addOnListInput *AddOnListInput) (*AddOnList, *Error) {
	u := c.url("add_ons", addOnListInput.query())

	return get[AddOnList](ctx, c, u)
}

func (c *Client) CreateAddOn(ctx context.Context, addOnInput *AddOnInput) (*AddOn, *Error) {
	u := c.url("add_ons", nil)
	result, err := post[addOnParams, addOnResult](
		ctx,
		c,
		u,
		&addOnParams{AddOn: addOnInput},
	)
	if err != nil {
		return nil, err
	}

	return result.AddOn, nil
}

func (c *Client) UpdateAddOn(ctx context.Context, addOnInput *AddOnInput) (*AddOn, *Error) {
	u := c.url("add_ons/"+addOnInput.Code, nil)

	result, err := put[addOnParams, addOnResult](
		ctx,
		c,
		u,
		&addOnParams{AddOn: addOnInput},
	)
	if err != nil {
		return nil, err
	}

	return result.AddOn, nil
}

func (c *Client) DeleteAddOn(ctx context.Context, addOnCode string) (*AddOn, *Error) {
	u := c.url("add_ons/"+addOnCode, nil)
	result, err := delete[addOnResult](ctx, c, u)
	if err != nil {
		return nil, err
	}
	return result.AddOn, nil
}
