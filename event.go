package lago

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type EventRequest struct {
	client *Client
}

type EventParams struct {
	Event *EventInput `json:"event"`
}

type BatchEventParams struct {
	Events *[]*EventInput `json:"events"`
}

type EventInput struct {
	TransactionID           string                 `json:"transaction_id,omitempty"`
	ExternalSubscriptionID  string                 `json:"external_subscription_id,omitempty"`
	Code                    string                 `json:"code,omitempty"`
	Timestamp               string                 `json:"timestamp,omitempty"`
	PreciseTotalAmountCents string                 `json:"precise_total_amount_cents,omitempty"`
	Properties              map[string]interface{} `json:"properties,omitempty"`
}

type EventEstimateFeesParams struct {
	Event *EventEstimateFeesInput `json:"event"`
}

type EventEstimateFeesInput struct {
	ExternalSubscriptionID string            `json:"external_subscription_id,omitempty"`
	Code                   string            `json:"code,omitempty"`
	Properties             map[string]string `json:"properties,omitempty"`
}

type BatchEventResult struct {
	Events *[]*Event `json:"events"`
}

type EventResult struct {
	Event *Event `json:"event"`
}

type Event struct {
	LagoID                  uuid.UUID              `json:"lago_id"`
	TransactionID           string                 `json:"transaction_id"`
	LagoCustomerID          *uuid.UUID             `json:"lago_customer_id,omitempty"`
	Code                    string                 `json:"code,omitempty"`
	Timestamp               time.Time              `json:"timestamp"`
	PreciseTotalAmountCents string                 `json:"precise_total_amount_cents,omitempty"`
	Properties              map[string]interface{} `json:"properties,omitempty"`
	LagoSubscriptionID      *uuid.UUID             `json:"lago_subscription_id,omitempty"`
	ExternalSubscriptionID  string                 `json:"external_subscription_id,omitempty"`
	CreatedAt               time.Time              `json:"created_at"`
}

func (c *Client) Event() *EventRequest {
	return &EventRequest{
		client: c,
	}
}

func (er *EventRequest) Create(ctx context.Context, eventInput *EventInput) (*Event, *Error) {
	u := er.client.url("events", nil)
	result, err := post[EventParams, EventResult](ctx, er.client, u, &EventParams{Event: eventInput})
	if err != nil {
		return nil, err
	}

	return result.Event, nil
}

func (er *EventRequest) EstimateFees(ctx context.Context, estimateInput EventEstimateFeesInput) (*FeeResult, *Error) {
	u := er.client.url("events/estimate_fees", nil)
	return post[EventEstimateFeesParams, FeeResult](ctx, er.client, u, &EventEstimateFeesParams{Event: &estimateInput})
}

func (er *EventRequest) Get(ctx context.Context, eventID string) (*Event, *Error) {
	u := er.client.url("events/"+eventID, nil)
	result, err := get[EventResult](ctx, er.client, u)
	if err != nil {
		return nil, err
	}

	return result.Event, nil
}

func (er *EventRequest) Batch(ctx context.Context, batchInput *[]*EventInput) (*[]*Event, *Error) {
	u := er.client.url("events/batch", nil)
	result, err := post[BatchEventParams, BatchEventResult](ctx, er.client, u, &BatchEventParams{Events: batchInput})
	if err != nil {
		return nil, err
	}

	return result.Events, nil
}
