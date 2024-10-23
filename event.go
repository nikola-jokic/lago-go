package lago

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type eventParams struct {
	Event *EventInput `json:"event"`
}

type batchEventParams struct {
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

type eventEstimateFeesParams struct {
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

func (c *Client) CreateEvent(ctx context.Context, eventInput *EventInput) (*Event, *Error) {
	u := c.url("events", nil)
	result, err := post[eventParams, EventResult](
		ctx,
		c,
		u,
		&eventParams{Event: eventInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Event, nil
}

func (c *Client) EstimateEventFees(ctx context.Context, estimateInput *EventEstimateFeesInput) (*feeResult, *Error) {
	u := c.url("events/estimate_fees", nil)
	return post[eventEstimateFeesParams, feeResult](
		ctx,
		c,
		u,
		&eventEstimateFeesParams{Event: estimateInput},
	)
}

func (c *Client) GetEvent(ctx context.Context, eventID string) (*Event, *Error) {
	u := c.url("events/"+eventID, nil)
	result, err := get[EventResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Event, nil
}

func (c *Client) BatchEvents(ctx context.Context, batchInput *[]*EventInput) (*[]*Event, *Error) {
	u := c.url("events/batch", nil)
	result, err := post[batchEventParams, BatchEventResult](
		ctx,
		c,
		u,
		&batchEventParams{Events: batchInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Events, nil
}
