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

type couponResult struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CouponList struct {
	Coupons []*Coupon `json:"coupons,omitempty"`
	Meta    Metadata  `json:"meta,omitempty"`
}

type couponParams struct {
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

type appliedCouponResult struct {
	AppliedCoupon *AppliedCoupon `json:"applied_coupon,omitempty"`
}

type AppliedCouponList struct {
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

type applyCouponParams struct {
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

func (c *Client) GetCoupon(ctx context.Context, couponCode string) (*Coupon, *Error) {
	u := c.url("coupons/"+couponCode, nil)

	result, err := get[couponResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (c *Client) ListCoupons(ctx context.Context, couponListInput *CouponListInput) (*CouponList, *Error) {
	u := c.url("coupons", couponListInput.query())
	return get[CouponList](ctx, c, u)
}

func (c *Client) CreateCoupon(ctx context.Context, couponInput *CouponInput) (*Coupon, *Error) {
	u := c.url("coupons", nil)

	result, err := post[couponParams, couponResult](
		ctx,
		c,
		u,
		&couponParams{Coupon: couponInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (c *Client) UpdateCoupon(ctx context.Context, couponInput *CouponInput) (*Coupon, *Error) {
	u := c.url("coupons/"+couponInput.Code, nil)

	result, err := put[couponParams, couponResult](
		ctx,
		c,
		u,
		&couponParams{Coupon: couponInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (c *Client) DeleteCoupon(ctx context.Context, couponCode string) (*Coupon, *Error) {
	u := c.url("coupons/"+couponCode, nil)

	result, err := delete[couponResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Coupon, nil
}

func (c *Client) ListAppliedCoupons(ctx context.Context, appliedCouponListInput *AppliedCouponListInput) (*AppliedCouponList, *Error) {
	u := c.url("applied_coupons", appliedCouponListInput.query())
	return get[AppliedCouponList](ctx, c, u)
}

func (c *Client) ApplyCouponToCustomer(ctx context.Context, applyCouponInput *ApplyCouponInput) (*AppliedCoupon, *Error) {
	u := c.url("applied_coupons", nil)
	result, err := post[applyCouponParams, appliedCouponResult](
		ctx,
		c,
		u,
		&applyCouponParams{AppliedCoupon: applyCouponInput},
	)
	if err != nil {
		return nil, err
	}

	return result.AppliedCoupon, nil
}

func (ac *Client) DeleteAppliedCoupon(ctx context.Context, externalCustomerID string, appliedCouponID string) (*AppliedCoupon, *Error) {
	subPath := fmt.Sprintf("customers/%s/applied_coupons/%s", externalCustomerID, appliedCouponID)
	u := ac.url(subPath, nil)

	result, err := delete[appliedCouponResult](ctx, ac, u)
	if err != nil {
		return nil, err
	}

	return result.AppliedCoupon, nil
}
