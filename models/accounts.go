package models

import (
	"time"

	"github.com/google/uuid"
)

// ========== Accounts Requests ==========

// AccrualToCustomerRequest represents request for accrual to customer
type AccrualToCustomerRequest struct {
	BaseCommand
	CustomerID     uuid.UUID  `json:"customerId"`
	CurrencyID     uuid.UUID  `json:"currencyId"`
	Amount         float64    `json:"amount"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	ActivationDate *time.Time `json:"activationDate,omitempty"`
	Comment        *string    `json:"comment,omitempty"`
}

// AccrualToLoyaltyCardRequest represents request for accrual to loyalty card
type AccrualToLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID  uuid.UUID  `json:"loyaltyCardId"`
	CurrencyID     uuid.UUID  `json:"currencyId"`
	Amount         float64    `json:"amount"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	ActivationDate *time.Time `json:"activationDate,omitempty"`
	Comment        *string    `json:"comment,omitempty"`
}

// SubtractFromCustomerRequest represents request for subtract from customer
type SubtractFromCustomerRequest struct {
	BaseCommand
	CustomerID uuid.UUID `json:"customerId"`
	CurrencyID uuid.UUID `json:"currencyId"`
	Amount     float64   `json:"amount"`
	Comment    *string   `json:"comment,omitempty"`
}

// SubtractFromLoyaltyCardRequest represents request for subtract from loyalty card
type SubtractFromLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID uuid.UUID `json:"loyaltyCardId"`
	CurrencyID    uuid.UUID `json:"currencyId"`
	Amount        float64   `json:"amount"`
	Comment       *string   `json:"comment,omitempty"`
}

// ========== Transaction DTOs ==========

// TransactionDto represents transaction data
type TransactionDto struct {
	TransactionID   uuid.UUID        `json:"transactionId"`
	OperationDate   time.Time        `json:"operationDate"`
	TransactionType TransactionTypes `json:"transactionType"`
	Amount          float64          `json:"amount"`
	CurrencyID      uuid.UUID        `json:"currencyId"`
	ExpirationDate  *time.Time       `json:"expirationDate,omitempty"`
	ActivationDate  *time.Time       `json:"activationDate,omitempty"`
}

// TransactionDetailDto represents detailed transaction
type TransactionDetailDto struct {
	TransactionDto
	AccountID uuid.UUID `json:"accountId"`
}

// TransactionListDto represents list of transactions
type TransactionListDto struct {
	Transactions []TransactionDetailDto `json:"transactions"`
}

// GetTransactionsRequest represents request for getting transactions
type GetTransactionsRequest struct {
	FromDate *time.Time `json:"fromDate,omitempty"`
	ToDate   *time.Time `json:"toDate,omitempty"`
	Take     *int       `json:"take,omitempty"`
	Skip     *int       `json:"skip,omitempty"`
}
