package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AggregationType string

const (
	CountAggregation          AggregationType = "count_agg"
	SumAggregation            AggregationType = "sum_agg"
	MaxAggregation            AggregationType = "max_agg"
	UniqueCountAggregation    AggregationType = "unique_count_agg"
	RecurringCountAggregation AggregationType = "recurring_count_agg"
	WeightedSumAggregation    AggregationType = "weighted_sum_agg"
)

type WeightedInterval string

const (
	SecondsInterval WeightedInterval = "seconds"
)

type billableMetricParams struct {
	BillableMetricInput *BillableMetricInput `json:"billable_metric,omitempty"`
}

type BillableMetricInput struct {
	Name             string                  `json:"name,omitempty"`
	Code             string                  `json:"code,omitempty"`
	Description      string                  `json:"description,omitempty"`
	AggregationType  AggregationType         `json:"aggregation_type,omitempty"`
	Recurring        bool                    `json:"recurring,omitempty"`
	FieldName        string                  `json:"field_name"`
	WeightedInterval WeightedInterval        `json:"weighted_interval,omitempty"`
	Filters          []*BillableMetricFilter `json:"filters,omitempty"`
}

type BillableMetricListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *BillableMetricListInput) query() url.Values {
	q := make(url.Values)
	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}
	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	return q
}

type BillableMetricList struct {
	BillableMetrics []*BillableMetric `json:"billable_metrics,omitempty"`
	Meta            Metadata          `json:"meta,omitempty"`
}

type billableMetricResult struct {
	BillableMetric *BillableMetric `json:"billable_metric,omitempty"`
}

type BillableMetricFilter struct {
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}

type BillableMetric struct {
	LagoID                   uuid.UUID               `json:"lago_id"`
	Name                     string                  `json:"name,omitempty"`
	Code                     string                  `json:"code,omitempty"`
	Description              string                  `json:"description,omitempty"`
	Recurring                bool                    `json:"recurring,omitempty"`
	AggregationType          AggregationType         `json:"aggregation_type,omitempty"`
	FieldName                string                  `json:"field_name"`
	CreatedAt                time.Time               `json:"created_at,omitempty"`
	WeightedInterval         *WeightedInterval       `json:"weighted_interval,omitempty"`
	Filters                  []*BillableMetricFilter `json:"filters,omitempty"`
	ActiveSubscriptionsCount int                     `json:"active_subscriptions_count,omitempty"`
	DraftInvoicesCount       int                     `json:"draft_invoices_count,omitempty"`
	PlansCount               int                     `json:"plans_count,omitempty"`
}

func (c *Client) GetBillableMetric(ctx context.Context, billableMetricCode string) (*BillableMetric, *Error) {
	u := c.url("billable_metrics/"+billableMetricCode, nil)

	result, err := get[billableMetricResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}

func (c *Client) ListBillableMetrics(ctx context.Context, billableMetricListInput *BillableMetricListInput) (*BillableMetricList, *Error) {
	u := c.url("billable_metrics", billableMetricListInput.query())
	return get[BillableMetricList](ctx, c, u)
}

func (c *Client) CreateBillableMetric(ctx context.Context, billableMetricInput *BillableMetricInput) (*BillableMetric, *Error) {
	u := c.url("billable_metrics", nil)

	result, err := post[billableMetricParams, billableMetricResult](
		ctx,
		c,
		u,
		&billableMetricParams{BillableMetricInput: billableMetricInput},
	)
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}

func (c *Client) UpdateBillableMetric(ctx context.Context, billableMetricInput *BillableMetricInput) (*BillableMetric, *Error) {
	u := c.url("billable_metrics/"+billableMetricInput.Code, nil)

	result, err := put[billableMetricParams, billableMetricResult](
		ctx,
		c,
		u,
		&billableMetricParams{BillableMetricInput: billableMetricInput},
	)
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}

func (c *Client) DeleteBillableMetric(ctx context.Context, billableMetricCode string) (*BillableMetric, *Error) {
	u := c.url("billable_metrics/"+billableMetricCode, nil)

	result, err := delete[billableMetricResult](
		ctx,
		c,
		u,
	)
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}
