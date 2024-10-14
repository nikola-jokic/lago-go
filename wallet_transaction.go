package lago

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type WalletTransactionStatus string

const (
	WalletTransactionStatusPending WalletTransactionStatus = "pending"
	WalletTransactionStatusSettled WalletTransactionStatus = "settled"
)

type TransactionStatus string

const (
	Purchased TransactionStatus = "purchased"
	Granted   TransactionStatus = "granted"
	Voided    TransactionStatus = "voided"
	Invoiced  TransactionStatus = "invoiced"
)

type TransactionType string

const (
	Outbound TransactionType = "outbound"
	Inbound  TransactionType = "inbound"
)

type WalletTransactionListInput struct {
	PerPage           int                     `json:"per_page,omitempty,string"`
	Page              int                     `json:"page,omitempty,string"`
	WalletID          string                  `json:"wallet_id,omitempty"`
	Status            WalletTransactionStatus `json:"status,omitempty"`
	TransactionStatus TransactionStatus       `json:"transaction_status,omitempty"`
	TransactionType   TransactionType         `json:"transaction_type,omitempty"`
}

func (i *WalletTransactionListInput) query() url.Values {
	q := make(url.Values)

	if i.PerPage > 0 {
		q.Add("per_page", strconv.Itoa(i.PerPage))
	}

	if i.Page > 0 {
		q.Add("page", strconv.Itoa(i.Page))
	}

	if i.WalletID != "" {
		q.Add("wallet_id", i.WalletID)
	}

	if i.Status != "" {
		q.Add("status", string(i.Status))
	}

	if i.TransactionStatus != "" {
		q.Add("transaction_status", string(i.TransactionStatus))
	}

	if i.TransactionType != "" {
		q.Add("transaction_type", string(i.TransactionType))
	}

	return q
}

type WalletTransactionParams struct {
	WalletTransactionInput *WalletTransactionInput `json:"wallet_transaction"`
}

type WalletTransactionInput struct {
	WalletID                         string                       `json:"wallet_id,omitempty"`
	PaidCredits                      string                       `json:"paid_credits,omitempty"`
	GrantedCredits                   string                       `json:"granted_credits,omitempty"`
	VoidedCredits                    string                       `json:"voided_credits,omitempty"`
	InvoiceRequiresSuccessfulPayment bool                         `json:"invoice_requires_successful_payment,omitempty"`
	Metadata                         []*WalletTransactionMetadata `json:"metadata,omitempty"`
}

type WalletTransactionMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type WalletTransactionResult struct {
	WalletTransactions []*WalletTransaction `json:"wallet_transactions,omitempty"`
	Meta               Metadata             `json:"meta,omitempty"`
}

type WalletTransaction struct {
	LagoID                           uuid.UUID                    `json:"lago_id,omitempty"`
	LagoWalletID                     uuid.UUID                    `json:"lago_wallet_id,omitempty"`
	Status                           WalletTransactionStatus      `json:"status,omitempty"`
	TransactionType                  TransactionType              `json:"transaction_type,omitempty"`
	Amount                           string                       `json:"amount,omitempty"`
	CreditAmount                     string                       `json:"credit_amount,omitempty"`
	InvoiceRequiresSuccessfulPayment bool                         `json:"invoice_requires_successful_payment,omitempty"`
	CreatedAt                        time.Time                    `json:"created_at,omitempty"`
	SettledAt                        time.Time                    `json:"settled_at,omitempty"`
	Metadata                         []*WalletTransactionMetadata `json:"metadata,omitempty"`
}

func (c *Client) CreateWalletTransaction(ctx context.Context, walletTransactionInput *WalletTransactionInput) (*WalletTransactionResult, *Error) {
	u := c.url("wallet_transactions", nil)
	return post[WalletTransactionParams, WalletTransactionResult](ctx, c, u, &WalletTransactionParams{WalletTransactionInput: walletTransactionInput})
}

func (c *Client) ListWalletTransactions(ctx context.Context, walletTransactionListInput *WalletTransactionListInput) (*WalletTransactionResult, *Error) {
	u := c.url("wallets/"+walletTransactionListInput.WalletID+"/wallet_transactions", walletTransactionListInput.query())
	return get[WalletTransactionResult](ctx, c, u)
}
