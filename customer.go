package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CustomerPaymentProvider string

const (
	PaymentProviderAdyen      CustomerPaymentProvider = "adyen"
	PaymentProviderStripe     CustomerPaymentProvider = "stripe"
	PaymentProviderGocardless CustomerPaymentProvider = "gocardless"
)

type FinalizeZeroAmountInvoice string

const (
	FinalizeInvoice FinalizeZeroAmountInvoice = "finalize"
	SkipInvoice     FinalizeZeroAmountInvoice = "skip"
	InheritInvoice  FinalizeZeroAmountInvoice = "inherit"
)

type IntegrationType string

const (
	IntegrationNetsuite IntegrationType = "netsuite"
	IntegrationAnrok    IntegrationType = "anrok"
	IntegrationXero     IntegrationType = "xero"
)

type CustomerType string

const (
	CompanyCustomerType    CustomerType = "company"
	IndividualCustomerType CustomerType = "individual"
)

type CustomerParams struct {
	Customer *CustomerInput `json:"customer"`
}

type CustomerResult struct {
	Customer  *Customer   `json:"customer"`
	Customers []*Customer `json:"customers,omitempty"`
	Meta      Metadata    `json:"metadata,omitempty"`
}

type CustomerUsageResult struct {
	CustomerUsage *CustomerUsage `json:"customer_usage"`
}

type CustomerPastUsageResult struct {
	UsagePeriods []*CustomerUsage `json:"usage_periods"`
	Meta         Metadata         `json:"metadata"`
}

type CustomerPortalUrlResult struct {
	CustomerPortalUrl *CustomerPortalUrl `json:"customer"`
}

type CustomerCheckoutUrlResult struct {
	CustomerCheckoutUrl *CustomerCheckoutUrl `json:"customer"`
}

type CustomerMetadataInput struct {
	LagoID           *uuid.UUID `json:"id,omitempty"`
	Key              string     `json:"key,omitempty"`
	Value            string     `json:"value,omitempty"`
	DisplayInInvoice bool       `json:"display_in_invoice,omitempty"`
}

type MetadataResponse struct {
	LagoID           uuid.UUID `json:"lago_id,omitempty"`
	Key              string    `json:"key,omitempty"`
	Value            string    `json:"value,omitempty"`
	DisplayInInvoice bool      `json:"display_in_invoice,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}

type CustomerInput struct {
	ExternalID                string                            `json:"external_id,omitempty"`
	Name                      string                            `json:"name,omitempty"`
	Firstname                 string                            `json:"firstname,omitempty"`
	Lastname                  string                            `json:"lastname,omitempty"`
	CustomerType              CustomerType                      `json:"customer_type,omitempty"`
	Email                     string                            `json:"email,omitempty"`
	AddressLine1              string                            `json:"address_line1,omitempty"`
	AddressLine2              string                            `json:"address_line2,omitempty"`
	City                      string                            `json:"city,omitempty"`
	Zipcode                   string                            `json:"zipcode,omitempty"`
	State                     string                            `json:"state,omitempty"`
	Country                   string                            `json:"country,omitempty"`
	LegalName                 string                            `json:"legal_name,omitempty"`
	LegalNumber               string                            `json:"legal_number,omitempty"`
	NetPaymentTerm            int                               `json:"net_payment_term,omitempty"`
	TaxIdentificationNumber   string                            `json:"tax_identification_number,omitempty"`
	Phone                     string                            `json:"phone,omitempty"`
	URL                       string                            `json:"url,omitempty"`
	Currency                  Currency                          `json:"currency,omitempty"`
	Timezone                  string                            `json:"timezone,omitempty"`
	Metadata                  []*CustomerMetadataInput          `json:"metadata,omitempty"`
	BillingConfiguration      CustomerBillingConfigurationInput `json:"billing_configuration,omitempty"`
	ShippingAddress           Address                           `json:"shipping_address,omitempty"`
	IntegrationCustomers      []*IntegrationCustomer            `json:"integration_customers,omitempty"`
	TaxCodes                  []string                          `json:"tax_codes,omitempty"`
	FinalizeZeroAmountInvoice FinalizeZeroAmountInvoice         `json:"finalize_zero_amount_invoice,omitempty"`
}

type CustomerListInput struct {
	PerPage int `json:"per_page,omitempty,string"`
	Page    int `json:"page,omitempty,string"`
}

func (i *CustomerListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}
	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	return q
}

type CustomerBillingConfigurationInput struct {
	InvoiceGracePeriod  int                     `json:"invoice_grace_period,omitempty"`
	PaymentProvider     CustomerPaymentProvider `json:"payment_provider,omitempty"`
	PaymentProviderCode string                  `json:"payment_provider_code,omitempty"`
	ProviderCustomerID  string                  `json:"provider_customer_id,omitempty"`
	Sync                bool                    `json:"sync,omitempty"`
	SyncWithProvider    bool                    `json:"sync_with_provider,omitempty"`
	DocumentLocale      string                  `json:"document_locale,omitempty"`
}

type CustomerBillingConfiguration struct {
	InvoiceGracePeriod  int                     `json:"invoice_grace_period,omitempty"`
	PaymentProvider     CustomerPaymentProvider `json:"payment_provider,omitempty"`
	PaymentProviderCode string                  `json:"payment_provider_code,omitempty"`
	ProviderCustomerID  string                  `json:"provider_customer_id,omitempty"`
	SyncWithProvider    bool                    `json:"sync_with_provider,omitempty"`
	DocumentLocale      string                  `json:"document_locale,omitempty"`
}

type Address struct {
	AddressLine1 string `json:"address_line1,omitempty"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city,omitempty"`
	Zipcode      string `json:"zipcode,omitempty"`
	State        string `json:"state,omitempty"`
	Country      string `json:"country,omitempty"`
}

type IntegrationCustomer struct {
	LagoID             uuid.UUID       `json:"id,omitempty"`
	ExternalCustomerId string          `json:"external_customer_id,omitempty"`
	IntegrationType    IntegrationType `json:"integration_type,omitempty"`
	IntegrationCode    string          `json:"integration_code,omitempty"`
	SubsidiaryId       string          `json:"subsidiary_id,omitempty"`
	SyncWithProvider   bool            `json:"sync_with_provider,omitempty"`
}

type IntegrationCustomersResponse struct {
	LagoID             uuid.UUID       `json:"lago_id,omitempty"`
	ExternalCustomerId string          `json:"external_customer_id,omitempty"`
	IntegrationType    IntegrationType `json:"type,omitempty"`
	IntegrationCode    string          `json:"integration_code,omitempty"`
	SubsidiaryId       string          `json:"subsidiary_id,omitempty"`
	SyncWithProvider   bool            `json:"sync_with_provider,omitempty"`
}

type CustomerChargeUsage struct {
	Units          string   `json:"units,omitempty"`
	EventsCount    int      `json:"events_count"`
	AmountCents    int      `json:"amount_cents,omitempty"`
	AmountCurrency Currency `json:"amount_currency,omitempty"`

	Charge         *Charge                       `json:"charge,omitempty"`
	BillableMetric *BillableMetric               `json:"billable_metric,omitempty"`
	Filters        []*CustomerChargeFilterUsage  `json:"filters,omitempty"`
	GroupedUsage   []*CustomerChargeGroupedUsage `json:"grouped_usage,omitempty"`
}

type CustomerChargeFilterUsage struct {
	InvoiceDisplayName string                 `json:"invoice_display_name,omitempty"`
	Values             map[string]interface{} `json:"value,omitempty"`
	AmountCents        int                    `json:"amount_cents,omitempty"`
	EventsCount        int                    `json:"events_count,omitempty"`
	Units              string                 `json:"units,omitempty"`
}

type CustomerChargeGroupedUsage struct {
	AmountCents int                          `json:"amount_cents,omitempty"`
	EventsCount int                          `json:"events_count,omitempty"`
	Units       string                       `json:"units,omitempty"`
	GroupedBy   map[string]interface{}       `json:"grouped_by,omitempty"`
	Filters     []*CustomerChargeFilterUsage `json:"filters,omitempty"`
}

type CustomerUsage struct {
	FromDatetime     time.Time `json:"from_datetime,omitempty"`
	ToDatetime       time.Time `json:"to_datetime,omitempty"`
	IssuingDate      string    `json:"issuing_date,omitempty"`
	LagoInvoiceID    string    `json:"lago_invoice_id,omitempty"`
	Currency         Currency  `json:"currency,omitempty"`
	AmountCents      int       `json:"amount_cents,omitempty"`
	TotalAmountCents int       `json:"total_amount_cents,omitempty"`
	TaxesAmountCents int       `json:"taxes_amount_cents,omitempty"`

	ChargesUsage []*CustomerChargeUsage `json:"charges_usage,omitempty"`
}

type CustomerPortalUrl struct {
	PortalUrl string `json:"portal_url,omitempty"`
}

type CustomerCheckoutUrl struct {
	CheckoutUrl string `json:"checkout_url,omitempty"`
}

type CustomerUsageInput struct {
	ExternalSubscriptionID string `json:"external_subscription_id,omitempty"`
}

func (i *CustomerUsageInput) query() url.Values {
	q := make(url.Values)
	if i.ExternalSubscriptionID != "" {
		q.Add("external_subscription_id", i.ExternalSubscriptionID)
	}

	return q
}

type CustomerPastUsageInput struct {
	ExternalSubscriptionID string `json:"external_subscription_id"`
	BillableMetricCode     string `json:"billable_metric_code,omitempty"`
	PeriodsCount           int    `json:"periods_count,omitempty"`
}

func (i *CustomerPastUsageInput) query() url.Values {
	q := make(url.Values)
	if i.ExternalSubscriptionID != "" {
		q.Add("external_subscription_id", i.ExternalSubscriptionID)
	}
	if i.BillableMetricCode != "" {
		q.Add("billable_metric_code", i.BillableMetricCode)
	}
	if i.PeriodsCount > 0 {
		q.Add("periods_count", strconv.Itoa(i.PeriodsCount))
	}

	return q
}

type Customer struct {
	LagoID       uuid.UUID `json:"lago_id,omitempty"`
	SequentialID int       `json:"sequential_id,omitempty"`
	ExternalID   string    `json:"external_id,omitempty"`
	Slug         string    `json:"slug,omitempty"`

	Name                      string                          `json:"name,omitempty"`
	Firstname                 string                          `json:"firstname,omitempty"`
	Lastname                  string                          `json:"lastname,omitempty"`
	CustomerType              string                          `json:"customer_type,omitempty"`
	Email                     string                          `json:"email,omitempty"`
	AddressLine1              string                          `json:"address_line1,omitempty"`
	AddressLine2              string                          `json:"address_line2,omitempty"`
	City                      string                          `json:"city,omitempty"`
	State                     string                          `json:"state,omitempty"`
	Zipcode                   string                          `json:"zipcode,omitempty"`
	Country                   string                          `json:"country,omitempty"`
	LegalName                 string                          `json:"legal_name,omitempty"`
	LegalNumber               string                          `json:"legal_number,omitempty"`
	NetPaymentTerm            int                             `json:"net_payment_term,omitempty"`
	TaxIdentificationNumber   string                          `json:"tax_identification_number,omitempty"`
	LogoURL                   string                          `json:"logo_url,omitempty"`
	Phone                     string                          `json:"phone,omitempty"`
	URL                       string                          `json:"url,omitempty"`
	FinalizeZeroAmountInvoice FinalizeZeroAmountInvoice       `json:"finalize_zero_amount_invoice,omitempty"`
	BillingConfiguration      CustomerBillingConfiguration    `json:"billing_configuration,omitempty"`
	ShippingAddress           Address                         `json:"shipping_address,omitempty"`
	IntegrationCustomers      []*IntegrationCustomersResponse `json:"integration_customers,omitempty"`
	Metadata                  []*MetadataResponse             `json:"metadata,omitempty"`
	Currency                  Currency                        `json:"currency,omitempty"`
	Timezone                  string                          `json:"timezone,omitempty"`
	ApplicableTimezone        string                          `json:"applicable_timezone,omitempty"`

	Taxes []*Tax `json:"taxes,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CustomerRequest struct {
	client *Client
}

func (c *Client) Customer() *CustomerRequest {
	return &CustomerRequest{
		client: c,
	}
}

func (cr *CustomerRequest) Create(ctx context.Context, customerInput *CustomerInput) (*Customer, *Error) {
	u := cr.client.url("customers", nil)
	result, err := post[CustomerParams, CustomerResult](ctx, cr.client, u, &CustomerParams{Customer: customerInput})
	if err != nil {
		return nil, err
	}

	return result.Customer, nil
}

// NOTE: Update endpoint does not exists, actually we use the create endpoint with the
// same externalID to update a customer
func (cr *CustomerRequest) Update(ctx context.Context, customerInput *CustomerInput) (*Customer, *Error) {
	return cr.Create(ctx, customerInput)
}

func (cr *CustomerRequest) CurrentUsage(ctx context.Context, externalCustomerID string, customerUsageInput *CustomerUsageInput) (*CustomerUsage, *Error) {
	u := cr.client.url("customers/"+externalCustomerID+"/current_usage", customerUsageInput.query())

	result, err := get[CustomerUsageResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.CustomerUsage, nil
}

func (cr *CustomerRequest) PastUsage(ctx context.Context, externalCustomerID string, customerPastUsageInput *CustomerPastUsageInput) (*CustomerPastUsageResult, *Error) {
	u := cr.client.url("customers/"+externalCustomerID+"/past_usage", customerPastUsageInput.query())

	return get[CustomerPastUsageResult](ctx, cr.client, u)
}

func (cr *CustomerRequest) PortalUrl(ctx context.Context, externalCustomerID string) (*CustomerPortalUrl, *Error) {
	u := cr.client.url("customers/"+externalCustomerID+"/portal_url", nil)
	result, err := get[CustomerPortalUrlResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.CustomerPortalUrl, nil
}

func (cr *CustomerRequest) CheckoutUrl(ctx context.Context, externalCustomerID string) (*CustomerCheckoutUrl, *Error) {
	u := cr.client.url("customers/"+externalCustomerID+"/checkout_url", nil)
	result, err := get[CustomerCheckoutUrlResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.CustomerCheckoutUrl, nil
}

func (cr *CustomerRequest) Delete(ctx context.Context, externalCustomerID string) (*Customer, *Error) {
	u := cr.client.url("customers/"+externalCustomerID, nil)
	result, err := delete[CustomerResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Customer, nil
}

func (cr *CustomerRequest) Get(ctx context.Context, externalCustomerID string) (*Customer, *Error) {
	u := cr.client.url("customers/"+externalCustomerID, nil)
	result, err := get[CustomerResult](ctx, cr.client, u)
	if err != nil {
		return nil, err
	}

	return result.Customer, nil
}

func (cr *CustomerRequest) GetList(ctx context.Context, customerListInput *CustomerListInput) (*CustomerResult, *Error) {
	u := cr.client.url("customers", customerListInput.query())
	return get[CustomerResult](ctx, cr.client, u)
}
