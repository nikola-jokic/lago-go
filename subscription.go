package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type SubscriptionStatus string

const (
	SubscriptionStatusActive     SubscriptionStatus = "active"
	SubscriptionStatusPending    SubscriptionStatus = "pending"
	SubscriptionStatusTerminated SubscriptionStatus = "terminated"
	SubscriptionStatusCanceled   SubscriptionStatus = "canceled"
)

type BillingTime string

const (
	Anniversary BillingTime = "anniversary"
	Calendar    BillingTime = "calendar"
)

type subscriptionResult struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type SubscriptionList struct {
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	Meta          Metadata        `json:"meta,omitempty"`
}

type subscriptionParams struct {
	Subscription *SubscriptionInput `json:"subscription"`
}

type ChargeOverridesInput struct {
	ID                 *uuid.UUID             `json:"id,omitempty"`
	AmountCurrency     Currency               `json:"amount_currency,omitempty"`
	InvoiceDisplayName string                 `json:"invoice_display_name,omitempty"`
	MinAmountCents     int                    `json:"min_amount_cents,omitempty"`
	Properties         map[string]interface{} `json:"properties"`
	Filters            []*ChargeFilter        `json:"filters,omitempty"`
	TaxCodes           []string               `json:"tax_codes,omitempty"`
}

type MinimumCommitmentOverridesInput struct {
	AmountCents        int      `json:"amount_cents,omitempty"`
	InvoiceDisplayName string   `json:"invoice_display_name,omitempty"`
	TaxCodes           []string `json:"tax_codes,omitempty"`
}

type PlanOverridesInput struct {
	Name               string                           `json:"name,omitempty"`
	InvoiceDisplayName string                           `json:"invoice_display_name,omitempty"`
	Code               string                           `json:"code,omitempty"`
	Description        string                           `json:"description,omitempty"`
	AmountCents        int                              `json:"amount_cents"`
	AmountCurrency     Currency                         `json:"amount_currency,omitempty"`
	TrialPeriod        float32                          `json:"trial_period"`
	Charges            []*ChargeOverridesInput          `json:"charges,omitempty"`
	MinimumCommitment  *MinimumCommitmentOverridesInput `json:"minimum_commitment"`
	TaxCodes           []string                         `json:"tax_codes,omitempty"`
	UsageThresholds    []*UsageThreshold                `json:"usage_thresholds,omitempty"`
}

type SubscriptionInput struct {
	ExternalCustomerID string              `json:"external_customer_id,omitempty"`
	PlanCode           string              `json:"plan_code,omitempty"`
	SubscriptionAt     *time.Time          `json:"subscription_at,omitempty"`
	EndingAt           *time.Time          `json:"ending_at,omitempty"`
	BillingTime        BillingTime         `json:"billing_time,omitempty"`
	PlanOverrides      *PlanOverridesInput `json:"plan_overrides,omitempty"`
	ExternalID         string              `json:"external_id"`
	Name               string              `json:"name"`
}

type SubscriptionTerminateInput struct {
	ExternalID string `json:"external_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

func (i *SubscriptionTerminateInput) query() url.Values {
	q := make(url.Values)

	if i.ExternalID != "" {
		q.Add("external_id", i.ExternalID)
	}

	if i.Status != "" {
		q.Add("status", i.Status)
	}

	return q
}

type SubscriptionListInput struct {
	ExternalCustomerID string               `json:"external_customer_id,omitempty"`
	PlanCode           string               `json:"plan_code,omitempty"`
	PerPage            int                  `json:"per_page,omitempty,string"`
	Page               int                  `json:"page,omitempty,string"`
	Status             []SubscriptionStatus `json:"status,omitempty"`
}

func (i *SubscriptionListInput) query() url.Values {
	q := make(url.Values)

	if i.ExternalCustomerID != "" {
		q.Add("external_customer_id", i.ExternalCustomerID)
	}

	if i.PlanCode != "" {
		q.Add("plan_code", i.PlanCode)
	}

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	for _, status := range i.Status {
		q.Add("status[]", string(status))
	}

	return q
}

type Subscription struct {
	LagoID             uuid.UUID `json:"lago_id"`
	LagoCustomerID     uuid.UUID `json:"lago_customer_id"`
	ExternalCustomerID string    `json:"external_customer_id"`
	ExternalID         string    `json:"external_id"`

	PlanCode string `json:"plan_code"`

	Name string `json:"name"`

	Status         SubscriptionStatus `json:"status"`
	BillingTime    BillingTime        `json:"billing_time"`
	SubscriptionAt *time.Time         `json:"subscription_at"`
	EndingAt       *time.Time         `json:"ending_at"`
	TrialEndedAt   *time.Time         `json:"trial_ended_at"`

	PreviousPlanCode  string `json:"previous_plan_code"`
	NextPlanCode      string `json:"next_plan_code"`
	DowngradePlanDate string `json:"downgrade_plan_date"`

	CurrentBillingPeriodStartedAt *time.Time `json:"current_billing_period_started_at"`
	CurrentBillingPeriodEndingAt  *time.Time `json:"current_billing_period_ending_at"`

	Plan *Plan `json:"plan,omitempty"`

	CreatedAt    *time.Time `json:"created_at"`
	StartedAt    *time.Time `json:"started_at"`
	CanceledAt   *time.Time `json:"canceled_at"`
	TerminatedAt *time.Time `json:"terminated_at"`
}

func (c *Client) CreateSubscription(ctx context.Context, subscriptionInput *SubscriptionInput) (*Subscription, *Error) {
	u := c.url("subscriptions", nil)
	result, err := post[subscriptionParams, subscriptionResult](
		ctx,
		c,
		u,
		&subscriptionParams{Subscription: subscriptionInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Subscription, nil
}

func (c *Client) TerminateSubscription(ctx context.Context, subscriptionTerminateInput *SubscriptionTerminateInput) (*Subscription, *Error) {
	u := c.url("subscriptions/"+subscriptionTerminateInput.ExternalID, subscriptionTerminateInput.query())
	result, err := delete[subscriptionResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Subscription, nil
}

func (c *Client) GetSubscription(ctx context.Context, subscriptionExternalId string) (*Subscription, *Error) {
	u := c.url("subscriptions/"+subscriptionExternalId, nil)
	result, err := get[subscriptionResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Subscription, nil
}

func (c *Client) ListSubscriptions(ctx context.Context, subscriptionListInput *SubscriptionListInput) (*SubscriptionList, *Error) {
	u := c.url("subscriptions", subscriptionListInput.query())
	return get[SubscriptionList](ctx, c, u)
}

func (c *Client) UpdateSubscription(ctx context.Context, subscriptionInput *SubscriptionInput) (*Subscription, *Error) {
	u := c.url("subscriptions/"+subscriptionInput.ExternalID, nil)
	result, err := put[subscriptionParams, subscriptionResult](
		ctx,
		c,
		u,
		&subscriptionParams{Subscription: subscriptionInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Subscription, nil
}
