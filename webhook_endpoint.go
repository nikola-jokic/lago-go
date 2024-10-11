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

type WebhookEndpointRequest struct {
	client *Client
}

type WebhookEndpointParams struct {
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

type WebhookEndpointResult struct {
	WebhookEndpoint  *WebhookEndpoint   `json:"webhook_endpoint,omitempty"`
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

func (c *Client) WebhookEndpoint() *WebhookEndpointRequest {
	return &WebhookEndpointRequest{
		client: c,
	}
}

func (wer *WebhookEndpointRequest) Get(ctx context.Context, webhookEndpointID string) (*WebhookEndpoint, *Error) {
	u := wer.client.url("webhook_endpoints/"+webhookEndpointID, nil)
	result, err := get[WebhookEndpointResult](ctx, wer.client, u)
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}

func (wer *WebhookEndpointRequest) GetList(ctx context.Context, webhookEndpointListInput *WebhookEndpointListInput) (*WebhookEndpointResult, *Error) {
	u := wer.client.url("webhook_endpoints", webhookEndpointListInput.query())
	return get[WebhookEndpointResult](ctx, wer.client, u)
}

func (wer *WebhookEndpointRequest) Create(ctx context.Context, webhookEndpointInput *WebhookEndpointInput) (*WebhookEndpoint, *Error) {
	u := wer.client.url("webhook_endpoints", nil)

	result, err := post[WebhookEndpointParams, WebhookEndpointResult](ctx, wer.client, u, &WebhookEndpointParams{WebhookEndpointInput: webhookEndpointInput})
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}

func (wer *WebhookEndpointRequest) Update(ctx context.Context, webhookEndpointInput *WebhookEndpointInput, webhookEndpointID string) (*WebhookEndpoint, *Error) {
	u := wer.client.url("webhook_endpoints/"+webhookEndpointID, nil)
	result, err := put[WebhookEndpointParams, WebhookEndpointResult](ctx, wer.client, u, &WebhookEndpointParams{WebhookEndpointInput: webhookEndpointInput})
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}

func (wer *WebhookEndpointRequest) Delete(ctx context.Context, webhookEndpointID string) (*WebhookEndpoint, *Error) {
	u := wer.client.url("webhook_endpoints/"+webhookEndpointID, nil)
	result, err := delete[WebhookEndpointResult](ctx, wer.client, u)
	if err != nil {
		return nil, err
	}

	return result.WebhookEndpoint, nil
}
