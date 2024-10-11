package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type BillableMetricRequest struct {
	client *Client
}

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

type BillableMetricParams struct {
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

type BillableMetricResult struct {
	BillableMetric  *BillableMetric   `json:"billable_metric,omitempty"`
	BillableMetrics []*BillableMetric `json:"billable_metrics,omitempty"`
	Meta            Metadata          `json:"meta,omitempty"`
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

func (c *Client) BillableMetric() *BillableMetricRequest {
	return &BillableMetricRequest{
		client: c,
	}
}

func (bmr *BillableMetricRequest) Get(ctx context.Context, billableMetricCode string) (*BillableMetric, *Error) {
	u := bmr.client.url("billable_metrics/"+billableMetricCode, nil)

	result, err := get[BillableMetricResult](ctx, bmr.client, u)
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}

func (bmr *BillableMetricRequest) GetList(ctx context.Context, billableMetricListInput *BillableMetricListInput) (*BillableMetricResult, *Error) {

	u := bmr.client.url("billable_metrics", billableMetricListInput.query())
	return get[BillableMetricResult](ctx, bmr.client, u)
}

func (bmr *BillableMetricRequest) Create(ctx context.Context, billableMetricInput *BillableMetricInput) (*BillableMetric, *Error) {
	u := bmr.client.url("billable_metrics", nil)

	result, err := post[BillableMetricParams, BillableMetricResult](ctx, bmr.client, u, &BillableMetricParams{BillableMetricInput: billableMetricInput})
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}

func (bmr *BillableMetricRequest) Update(ctx context.Context, billableMetricInput *BillableMetricInput) (*BillableMetric, *Error) {
	u := bmr.client.url("billable_metrics/"+billableMetricInput.Code, nil)

	result, err := put[BillableMetricParams, BillableMetricResult](ctx, bmr.client, u, &BillableMetricParams{BillableMetricInput: billableMetricInput})
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}

func (bmr *BillableMetricRequest) Delete(ctx context.Context, billableMetricCode string) (*BillableMetric, *Error) {
	u := bmr.client.url("billable_metrics/"+billableMetricCode, nil)

	result, err := delete[BillableMetricResult](ctx, bmr.client, u)
	if err != nil {
		return nil, err
	}

	return result.BillableMetric, nil
}