package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type SignatureAlgo string

const (
	JWT  SignatureAlgo = "jwt"
	HMac SignatureAlgo = "hmac"
)

type webhookEndpointParams struct {
	WebhookEndpointInput *WebhookEndpointInput `json:"webhook_endpoint"`
}

type WebhookEndpointInput struct {
	WebhookURL    string        `json:"webhook_url,omitempty"`
	SignatureAlgo SignatureAlgo `json:"signature_algo,omitempty"`
}

type WebhookEndpointListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *WebhookEndpointListInput) query() url.Values {
	v := url.Values{}
	if i.PerPage != 0 {
		v.Add("per_page", strconv.Itoa(i.PerPage))
	}
	if i.Page != 0 {
		v.Add("page", strconv.Itoa(i.Page))
	}
	return v
}

type webhookEndpointResult struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

type WebhookEndpointList struct {
	WebhookEndpoints []*WebhookEndpoint `json:"webhook_endpoints,omitempty"`
	Meta             Metadata           `json:"meta,omitempty"`
}

type WebhookEndpoint struct {
	LagoID             uuid.UUID     `json:"lago_id,omitempty"`
	LagoOrganizationID uuid.UUID     `json:"lago_organization_id,omitempty"`
	WebhookURL         string        `json:"webhook_url,omitempty"`
	SignatureAlgo      SignatureAlgo `json:"signature_algo,omitempty"`
	CreatedAt          time.Time     `json:"created_at,omitempty"`
}

func (c *Client) GetWebhookEndpoint(ctx context.Context, webhookEndpointID string) (*WebhookEndpoint, *Error) {
	u := c.url("webhook_endpoints/"+webhookEndpointID, nil)
	result, err := get[webhookEndpointResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}

func (c *Client) ListWebhookEndpoints(ctx context.Context, webhookEndpointListInput *WebhookEndpointListInput) (*WebhookEndpointList, *Error) {
	u := c.url("webhook_endpoints", webhookEndpointListInput.query())
	return get[WebhookEndpointList](ctx, c, u)
}

func (c *Client) CreateWebhookEndpoint(ctx context.Context, webhookEndpointInput *WebhookEndpointInput) (*WebhookEndpoint, *Error) {
	u := c.url("webhook_endpoints", nil)

	result, err := post[webhookEndpointParams, webhookEndpointResult](
		ctx,
		c,
		u,
		&webhookEndpointParams{WebhookEndpointInput: webhookEndpointInput},
	)
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}

func (c *Client) UpdateWebhookEndpoint(ctx context.Context, webhookEndpointInput *WebhookEndpointInput, webhookEndpointID string) (*WebhookEndpoint, *Error) {
	u := c.url("webhook_endpoints/"+webhookEndpointID, nil)
	result, err := put[webhookEndpointParams, webhookEndpointResult](
		ctx,
		c,
		u,
		&webhookEndpointParams{WebhookEndpointInput: webhookEndpointInput},
	)
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}

func (c *Client) DeleteWebhookEndpoint(ctx context.Context, webhookEndpointID string) (*WebhookEndpoint, *Error) {
	u := c.url("webhook_endpoints/"+webhookEndpointID, nil)
	result, err := delete[webhookEndpointResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}
