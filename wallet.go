package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Active     Status = "active"
	Terminated Status = "terminated"
)

type RecurringTransactionRuleInput struct {
	LagoID                           uuid.UUID                    `json:"lago_id,omitempty"`
	Interval                         string                       `json:"interval,omitempty"`
	Method                           string                       `json:"method,omitempty"`
	StartedAt                        *time.Time                   `json:"started_at,omitempty"`
	TargetOngoingBalance             string                       `json:"target_ongoing_balance,omitempty"`
	ThresholdCredits                 string                       `json:"threshold_credits,omitempty"`
	Trigger                          string                       `json:"trigger,omitempty"`
	PaidCredits                      string                       `json:"paid_credits,omitempty"`
	GrantedCredits                   string                       `json:"granted_credits,omitempty"`
	InvoiceRequiresSuccessfulPayment bool                         `json:"invoice_requires_successful_payment,omitempty"`
	TransactionMetadata              []*WalletTransactionMetadata `json:"transaction_metadata,omitempty"`
}

type RecurringTransactionRuleResponse struct {
	LagoID                           uuid.UUID                    `json:"lago_id,omitempty"`
	Interval                         string                       `json:"interval,omitempty"`
	Method                           string                       `json:"method,omitempty"`
	StartedAt                        *time.Time                   `json:"started_at,omitempty"`
	TargetOngoingBalance             string                       `json:"target_ongoing_balance,omitempty"`
	ThresholdCredits                 string                       `json:"threshold_credits,omitempty"`
	Trigger                          string                       `json:"trigger,omitempty"`
	PaidCredits                      string                       `json:"paid_credits,omitempty"`
	GrantedCredits                   string                       `json:"granted_credits,omitempty"`
	CreatedAt                        time.Time                    `json:"created_at,omitempty"`
	InvoiceRequiresSuccessfulPayment bool                         `json:"invoice_requires_successful_payment,omitempty"`
	TransactionMetadata              []*WalletTransactionMetadata `json:"transaction_metadata,omitempty"`
}

type walletParams struct {
	WalletInput *WalletInput `json:"wallet"`
}

type WalletInput struct {
	RateAmount                       string                           `json:"rate_amount,omitempty"`
	Currency                         Currency                         `json:"currency,omitempty"`
	Name                             string                           `json:"name,omitempty"`
	PaidCredits                      string                           `json:"paid_credits,omitempty"`
	GrantedCredits                   string                           `json:"granted_credits,omitempty"`
	ExpirationAt                     *time.Time                       `json:"expiration_at,omitempty"`
	ExternalCustomerID               string                           `json:"external_customer_id,omitempty"`
	InvoiceRequiresSuccessfulPayment bool                             `json:"invoice_requires_successful_payment,omitempty"`
	TransactionMetadata              []*WalletTransactionMetadata     `json:"transaction_metadata,omitempty"`
	RecurringTransactionRules        []*RecurringTransactionRuleInput `json:"recurring_transaction_rules,omitempty"`
}

type WalletListInput struct {
	PerPage            int    `json:"per_page,omitempty,string"`
	Page               int    `json:"page,omitempty,string"`
	ExternalCustomerID string `json:"external_customer_id,omitempty"`
}

func (i *WalletListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	if i.ExternalCustomerID != "" {
		q.Add("external_customer_id", i.ExternalCustomerID)
	}

	return q
}

type walletResult struct {
	Wallet *Wallet `json:"wallet,omitempty"`
}

type WalletList struct {
	Wallets []*Wallet `json:"wallets,omitempty"`
	Meta    Metadata  `json:"meta,omitempty"`
}

type Wallet struct {
	LagoID                           uuid.UUID                           `json:"lago_id,omitempty"`
	LagoCustomerID                   uuid.UUID                           `json:"lago_customer_id,omitempty"`
	ExternalCustomerID               string                              `json:"external_customer_id,omitempty"`
	Status                           Status                              `json:"status,omitempty"`
	Currency                         Currency                            `json:"currency,omitempty"`
	Name                             string                              `json:"name,omitempty"`
	RateAmount                       string                              `json:"rate_amount,omitempty"`
	CreditsBalance                   string                              `json:"credits_balance,omitempty"`
	BalanceCents                     int                                 `json:"balance_cents,omitempty"`
	ConsumedCredits                  string                              `json:"consumed_credits,omitempty"`
	InvoiceRequiresSuccessfulPayment bool                                `json:"invoice_requires_successful_payment,omitempty"`
	CreatedAt                        time.Time                           `json:"created_at,omitempty"`
	ExpirationAt                     time.Time                           `json:"expiration_at,omitempty"`
	LastBalanceSyncAt                time.Time                           `json:"last_balance_sync_at,omitempty"`
	LastConsumedCreditAt             time.Time                           `json:"last_consumed_credit_at,omitempty"`
	TerminatedAt                     time.Time                           `json:"terminated_at,omitempty"`
	RecurringTransactionRules        []*RecurringTransactionRuleResponse `json:"recurring_transaction_rules,omitempty"`
	OngoingBalanceCents              int                                 `json:"ongoing_balance_cents,omitempty"`
	OngoingUsageBalanceCents         int                                 `json:"ongoing_usage_balance_cents,omitempty"`
	CreditsOngoingBalance            string                              `json:"credits_ongoing_balance,omitempty"`
	CreditsOngoingUsageBalance       string                              `json:"credits_ongoing_usage_balance,omitempty"`
}

func (c *Client) GetWallet(ctx context.Context, walletID string) (*Wallet, *Error) {
	u := c.url("wallets/"+walletID, nil)
	result, err := get[walletResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Wallet, nil
}

func (c *Client) ListWallets(ctx context.Context, walletListInput *WalletListInput) (*WalletList, *Error) {
	u := c.url("wallets", walletListInput.query())
	return get[WalletList](ctx, c, u)
}

func (c *Client) CreateWallet(ctx context.Context, walletInput *WalletInput) (*Wallet, *Error) {
	u := c.url("wallets", nil)
	result, err := post[walletParams, walletResult](
		ctx,
		c,
		u,
		&walletParams{WalletInput: walletInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Wallet, nil
}

func (c *Client) UpdateWallet(ctx context.Context, walletInput *WalletInput, walletID string) (*Wallet, *Error) {
	u := c.url("wallets/"+walletID, nil)
	result, err := put[walletParams, walletResult](
		ctx,
		c,
		u,
		&walletParams{WalletInput: walletInput},
	)
	if err != nil {
		return nil, err
	}

	return result.Wallet, nil
}

func (c *Client) DeleteWallet(ctx context.Context, walletID string) (*Wallet, *Error) {
	u := c.url("wallets/"+walletID, nil)
	result, err := delete[walletResult](ctx, c, u)
	if err != nil {
		return nil, err
	}

	return result.Wallet, nil
}
