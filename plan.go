package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type PlanRequest struct {
	client *Client
}

type PlanResult struct {
	Plan  *Plan    `json:"plan,omitempty"`
	Plans []*Plan  `json:"plans,omitempty"`
	Meta  Metadata `json:"meta,omitempty"`
}

type PlanParams struct {
	Plan *PlanInput `json:"plan"`
}

type PlanInterval string

const (
	PlanWeekly    PlanInterval = "weekly"
	PlanMonthly   PlanInterval = "monthly"
	PlanQuarterly PlanInterval = "quarterly"
	PlanYearly    PlanInterval = "yearly"
)

type PlanChargeInput struct {
	LagoID           *uuid.UUID             `json:"id,omitempty"`
	BillableMetricID uuid.UUID              `json:"billable_metric_id,omitempty"`
	AmountCurrency   Currency               `json:"amount_currency,omitempty"`
	ChargeModel      ChargeModel            `json:"charge_model,omitempty"`
	PayInAdvance     bool                   `json:"pay_in_advance,omitempty"`
	Invoiceable      bool                   `json:"invoiceable,omitempty"`
	RegroupPaidFees  string                 `json:"regroup_paid_fees,omitempty"`
	Prorated         bool                   `json:"prorated,omitempty"`
	MinAmountCents   int                    `json:"min_amount_cents,omitempty"`
	Properties       map[string]interface{} `json:"properties"`
	Filters          []*ChargeFilter        `json:"filters,omitempty"`

	TaxCodes []string `json:"tax_codes,omitempty"`
}

type MinimumCommitmentInput struct {
	AmountCents        int      `json:"amount_cents,omitempty"`
	InvoiceDisplayName string   `json:"invoice_display_name,omitempty"`
	TaxCodes           []string `json:"tax_codes,omitempty"`
}

type PlanInput struct {
	Name               string                  `json:"name,omitempty"`
	InvoiceDisplayName string                  `json:"invoice_display_name,omitempty"`
	Code               string                  `json:"code,omitempty"`
	Interval           PlanInterval            `json:"interval,omitempty"`
	Description        string                  `json:"description,omitempty"`
	AmountCents        int                     `json:"amount_cents"`
	AmountCurrency     Currency                `json:"amount_currency,omitempty"`
	PayInAdvance       bool                    `json:"pay_in_advance"`
	BillChargeMonthly  bool                    `json:"bill_charge_monthly"`
	TrialPeriod        float32                 `json:"trial_period"`
	Charges            []*PlanChargeInput      `json:"charges,omitempty"`
	MinimumCommitment  *MinimumCommitmentInput `json:"minimum_commitment,omitempty"`
	TaxCodes           []string                `json:"tax_codes,omitempty"`
	UsageThresholds    []*UsageThresholdInput  `json:"usage_thresholds,omitempty"`
}

type PlanListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *PlanListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	return q
}

type MinimumCommitment struct {
	LagoID             uuid.UUID    `json:"lago_id"`
	PlanCode           string       `json:"plan_code,omitempty"`
	InvoiceDisplayName string       `json:"invoice_display_name,omitempty"`
	AmountCents        int          `json:"amount_cents"`
	Interval           PlanInterval `json:"interval,omitempty"`
	CreatedAt          time.Time    `json:"created_at,omitempty"`
	UpdatedAt          time.Time    `json:"updated_at,omitempty"`

	Taxes []*Tax `json:"tax,omitempty"`
}

type Plan struct {
	LagoID                   uuid.UUID          `json:"lago_id"`
	Name                     string             `json:"name,omitempty"`
	InvoiceDisplayName       string             `json:"invoice_display_name,omitempty"`
	Code                     string             `json:"code,omitempty"`
	Interval                 PlanInterval       `json:"interval,omitempty"`
	Description              string             `json:"description,omitempty"`
	AmountCents              int                `json:"amount_cents,omitempty"`
	AmountCurrency           Currency           `json:"amount_currency,omitempty"`
	PayInAdvance             bool               `json:"pay_in_advance,omitempty"`
	BillChargeMonthly        bool               `json:"bill_charge_monthly,omitempty"`
	ActiveSubscriptionsCount int                `json:"active_subscriptions_count,omitempty"`
	DraftInvoicesCount       int                `json:"draft_invoices_count,omitempty"`
	Charges                  []*Charge          `json:"charges,omitempty"`
	MinimumCommitment        *MinimumCommitment `json:"minimum_commitment"`

	Taxes           []*Tax            `json:"taxes,omitempty"`
	UsageThresholds []*UsageThreshold `json:"usage_thresholds,omitempty"`
}

func (c *Client) Plan() *PlanRequest {
	return &PlanRequest{
		client: c,
	}
}

func (pr *PlanRequest) Get(ctx context.Context, planCode string) (*Plan, *Error) {
	u := pr.client.url("plans/"+planCode, nil)
	result, err := get[PlanResult](ctx, pr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Plan, nil
}

func (pr *PlanRequest) GetList(ctx context.Context, planListInput *PlanListInput) (*PlanResult, *Error) {
	u := pr.client.url("plans", planListInput.query())
	return get[PlanResult](ctx, pr.client, u)
}

func (pr *PlanRequest) Create(ctx context.Context, planInput *PlanInput) (*Plan, *Error) {
	u := pr.client.url("plans", nil)
	result, err := post[PlanParams, PlanResult](ctx, pr.client, u, &PlanParams{Plan: planInput})
	if err != nil {
		return nil, err
	}

	return result.Plan, nil
}

func (pr *PlanRequest) Update(ctx context.Context, planInput *PlanInput) (*Plan, *Error) {
	u := pr.client.url("plans/"+planInput.Code, nil)
	result, err := put[PlanParams, PlanResult](ctx, pr.client, u, &PlanParams{Plan: planInput})
	if err != nil {
		return nil, err
	}

	return result.Plan, nil
}

func (pr *PlanRequest) Delete(ctx context.Context, planCode string) (*Plan, *Error) {
	u := pr.client.url("plans/"+planCode, nil)

	result, err := delete[PlanResult](ctx, pr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Plan, nil
}