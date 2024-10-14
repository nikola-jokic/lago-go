package lago

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LifetimeUsage struct {
	LagoID                 uuid.UUID `json:"lago_id"`
	LagoSubscriptionID     uuid.UUID `json:"lago_subscription_id"`
	ExternalSubscriptionID string    `json:"external_subscription_id"`

	ExternalHistoricalUsageAmountCents *int `json:"external_historical_usage_amount_cents,omitempty"`
	InvoicedUsageAmountCents           *int `json:"invoiced_usage_amount_cents,omitempty"`
	CurrentUsageAmountCents            *int `json:"current_usage_amount_cents,omitempty"`

	FromDatetime time.Time `json:"from_datetime"`
	ToDatetime   time.Time `json:"to_datetime"`

	UsageThresholds []*LifetimeUsageThreshold `json:"usage_thresholds,omitempty"`
}

type LifetimeUsageThreshold struct {
	AmountCents     int        `json:"amount_cents"`
	CompletionRatio float32    `json:"completion_ratio"`
	ReachedAt       *time.Time `json:"reached_at"`
}

type lifetimeUsageResult struct {
	LifetimeUsage *LifetimeUsage `json:"lifetime_usage"`
}

type lifetimeUsageParams struct {
	LifetimeUsage *LifetimeUsageInput `json:"lifetime_usage"`
}

type LifetimeUsageInput struct {
	ExternalSubscriptionID             string `json:"external_subscription_id"`
	ExternalHistoricalUsageAmountCents int    `json:"external_historical_usage_amount_cents"`
}

func (c *Client) GetLifetimeUsage(ctx context.Context, externalSubscriptionID string) (*LifetimeUsage, *Error) {
	u := c.url("subscriptions/"+externalSubscriptionID+"/lifetime_usage", nil)
	result, err := get[lifetimeUsageResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.LifetimeUsage, nil
}

func (c *Client) UpdateLifetimeUsage(ctx context.Context, lifetimeUsageInput *LifetimeUsageInput) (*LifetimeUsage, *Error) {
	u := c.url("subscriptions/"+lifetimeUsageInput.ExternalSubscriptionID+"/lifetime_usage", nil)
	result, err := put[lifetimeUsageParams, lifetimeUsageResult](ctx, c, u, &lifetimeUsageParams{LifetimeUsage: lifetimeUsageInput})
	if err != nil {
		return nil, err
	}

	return result.LifetimeUsage, nil
}
