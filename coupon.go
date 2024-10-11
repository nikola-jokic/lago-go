package lago

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CouponExpiration string

const (
	CouponExpirationTimeLimit    CouponExpiration = "time_limit"
	CouponExpirationNoExpiration CouponExpiration = "no_expiration"
)

type CouponCalculationType string

const (
	CouponTypeFixedAmount CouponCalculationType = "fixed_amount"
	CouponTypePercentage  CouponCalculationType = "percentage"
)

type CouponFrequency string

const (
	CouponFrequencyOnce      CouponFrequency = "once"
	CouponFrequencyRecurring CouponFrequency = "recurring"
)

type AppliedCouponStatus string

const (
	AppliedCouponStatusActive     AppliedCouponStatus = "active"
	AppliedCouponStatusTerminated AppliedCouponStatus = "terminated"
)

type CouponRequest struct {
	client *Client
}

type AppliedCouponRequest struct {
	client *Client
}

type CouponResult struct {
	Coupon  *Coupon   `json:"coupon,omitempty"`
	Coupons []*Coupon `json:"coupons,omitempty"`
	Meta    Metadata  `json:"meta,omitempty"`
}

type CouponParams struct {
	Coupon *CouponInput `json:"coupon"`
}

type LimitationInput struct {
	PlanCodes           []string `json:"plan_codes,omitempty"`
	BillableMetricCodes []string `json:"billable_metric_codes,omitempty"`
}

type CouponInput struct {
	Name              string                `json:"name,omitempty"`
	Code              string                `json:"code,omitempty"`
	Description       string                `json:"description,omitempty"`
	AmountCents       int                   `json:"amount_cents,omitempty"`
	AmountCurrency    Currency              `json:"amount_currency,omitempty"`
	Expiration        CouponExpiration      `json:"expiration,omitempty"`
	ExpirationAt      *time.Time            `json:"expiration_at,omitempty"`
	PercentageRate    float64               `json:"percentage_rate,omitempty,string"`
	CouponType        CouponCalculationType `json:"coupon_type,omitempty"`
	Frequency         CouponFrequency       `json:"frequency,omitempty"`
	Reusable          bool                  `json:"reusable,omitempty"`
	FrequencyDuration int                   `json:"frequency_duration,omitempty"`
	AppliesTo         LimitationInput       `json:"applies_to,omitempty"`
}

type CouponListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *CouponListInput) query() url.Values {
	q := make(url.Values)
	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	return q
}

type Coupon struct {
	LagoID                 uuid.UUID             `json:"lago_id,omitempty"`
	Name                   string                `json:"name,omitempty"`
	Code                   string                `json:"code,omitempty"`
	Description            string                `json:"description,omitempty"`
	AmountCents            int                   `json:"amount_cents,omitempty"`
	AmountCurrency         Currency              `json:"amount_currency,omitempty"`
	Expiration             CouponExpiration      `json:"expiration,omitempty"`
	ExpirationAt           *time.Time            `json:"expiration_at,omitempty"`
	PercentageRate         float64               `json:"percentage_rate,omitempty,string"`
	CouponType             CouponCalculationType `json:"coupon_type,omitempty"`
	Frequency              CouponFrequency       `json:"frequency,omitempty"`
	Reusable               bool                  `json:"reusable,omitempty"`
	LimitedPlans           bool                  `json:"limited_plans,omitempty"`
	PlanCodes              []string              `json:"plan_codes,omitempty"`
	LimitedBillableMetrics bool                  `json:"limited_billable_metrics,omitempty"`
	BillableMetricCodes    []string              `json:"billable_metric_codes,omitempty"`
	FrequencyDuration      int                   `json:"frequency_duration,omitempty"`
	CreatedAt              time.Time             `json:"created_at,omitempty"`
	TerminatedAt           *time.Time            `json:"terminated_at,omitempty"`
}

type AppliedCouponResult struct {
	AppliedCoupon  *AppliedCoupon   `json:"applied_coupon,omitempty"`
	AppliedCoupons []*AppliedCoupon `json:"applied_coupons,omitempty"`
	Meta           Metadata         `json:"meta,omitempty"`
}

type AppliedCouponListInput struct {
	PerPage            int                 `json:"per_page,omitempty,string"`
	Page               int                 `json:"page,omitempty,string"`
	Status             AppliedCouponStatus `json:"status,omitempty"`
	ExternalCustomerID string              `json:"external_customer_id,omitempty"`
}

func (i *AppliedCouponListInput) query() url.Values {
	q := make(url.Values)
	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	if i.Status != "" {
		q.Add("status", string(i.Status))
	}

	if i.ExternalCustomerID != "" {
		q.Add("external_customer_id", i.ExternalCustomerID)
	}

	return q
}

type ApplyCouponParams struct {
	AppliedCoupon *ApplyCouponInput `json:"applied_coupon"`
}

type ApplyCouponInput struct {
	ExternalCustomerID string          `json:"external_customer_id,omitempty"`
	CouponCode         string          `json:"coupon_code,omitempty"`
	AmountCents        int             `json:"amount_cents,omitempty"`
	AmountCurrency     Currency        `json:"amount_currency,omitempty"`
	PercentageRate     float64         `json:"percentage_rate,omitempty,string"`
	Frequency          CouponFrequency `json:"frequency,omitempty"`
	FrequencyDuration  int             `json:"frequency_duration,omitempty"`
}

type AppliedCoupon struct {
	LagoID             uuid.UUID           `json:"lago_id,omitempty"`
	LagoCouponID       uuid.UUID           `json:"lago_coupon_id,omitempty"`
	ExternalCustomerID string              `json:"external_customer_id,omitempty"`
	LagoCustomerID     uuid.UUID           `json:"lago_customer_id,omitempty"`
	Status             AppliedCouponStatus `json:"status,omitempty"`

	CouponName     string   `json:"coupon_name,omitempty"`
	CouponCode     string   `json:"coupon_code,omitempty"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"amount_currency,omitempty"`

	ExpirationAt time.Time `json:"expiration_at,omitempty"`
	TerminatedAt time.Time `json:"terminated_at,omitempty"`

	PercentageRate    float64         `json:"percentage_rate,omitempty,string"`
	Frequency         CouponFrequency `json:"frequency,omitempty"`
	FrequencyDuration int             `json:"frequency_duration,omitempty"`

	AmountCentsRemaining       int `json:"amount_cents_remaining,omitempty"`
	FrequencyDurationRemaining int `json:"frequency_duration_remaining,omitempty"`

	Credits []*InvoiceCredit `json:"credits,omitempty"`
}

func (c *Client) Coupon() *CouponRequest {
	return &CouponRequest{
		client: c,
	}
}

func (c *Client) AppliedCoupon() *AppliedCouponRequest {
	return &AppliedCouponRequest{
		client: c,
	}
}

func (cr *CouponRequest) Get(ctx context.Context, couponCode string) (*Coupon, *Error) {
	u := cr.client.url("coupons/"+couponCode, nil)

	result, err := get[CouponResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (cr *CouponRequest) GetList(ctx context.Context, couponListInput *CouponListInput) (*CouponResult, *Error) {
	u := cr.client.url("coupons", couponListInput.query())
	return get[CouponResult](ctx, cr.client, u)
}

func (cr *CouponRequest) Create(ctx context.Context, couponInput *CouponInput) (*Coupon, *Error) {
	u := cr.client.url("coupons", nil)

	result, err := post[CouponParams, CouponResult](ctx, cr.client, u, &CouponParams{Coupon: couponInput})
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (cr *CouponRequest) Update(ctx context.Context, couponInput *CouponInput) (*Coupon, *Error) {
	u := cr.client.url("coupons/"+couponInput.Code, nil)

	result, err := put[CouponParams, CouponResult](ctx, cr.client, u, &CouponParams{Coupon: couponInput})
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (cr *CouponRequest) Delete(ctx context.Context, couponCode string) (*Coupon, *Error) {
	u := cr.client.url("coupons/"+couponCode, nil)

	result, err := delete[CouponResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (cr *AppliedCouponRequest) GetList(ctx context.Context, appliedCouponListInput *AppliedCouponListInput) (*AppliedCouponResult, *Error) {
	u := cr.client.url("applied_coupons", appliedCouponListInput.query())
	return get[AppliedCouponResult](ctx, cr.client, u)
}

func (cr *CouponRequest) ApplyToCustomer(ctx context.Context, applyCouponInput *ApplyCouponInput) (*AppliedCoupon, *Error) {
	u := cr.client.url("applied_coupons", nil)
	result, err := post[ApplyCouponParams, AppliedCouponResult](ctx, cr.client, u, &ApplyCouponParams{AppliedCoupon: applyCouponInput})
	if err != nil {
		return nil, err
	}

	return result.AppliedCoupon, nil
}

func (acr *AppliedCouponRequest) AppliedCouponDelete(ctx context.Context, externalCustomerID string, appliedCouponID string) (*AppliedCoupon, *Error) {
	subPath := fmt.Sprintf("customers/%s/applied_coupons/%s", externalCustomerID, appliedCouponID)
	u := acr.client.url(subPath, nil)

	result, err := delete[AppliedCouponResult](ctx, acr.client, u)
	if err != nil {
		return nil, err
	}

	return result.AppliedCoupon, nil
}
