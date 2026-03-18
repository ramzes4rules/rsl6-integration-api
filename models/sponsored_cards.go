package models

import (
	"time"

	"github.com/google/uuid"
)

// ========== Sponsored Card Enums ==========

// SponsoredCardStatuses represents sponsored card status
type SponsoredCardStatuses string

const (
	SponsoredCardStatusCreated       SponsoredCardStatuses = "Created"
	SponsoredCardStatusReadyToIssued SponsoredCardStatuses = "ReadyToIssued"
	SponsoredCardStatusIssued        SponsoredCardStatuses = "Issued"
	SponsoredCardStatusActive        SponsoredCardStatuses = "Active"
	SponsoredCardStatusBlocked       SponsoredCardStatuses = "Blocked"
	SponsoredCardStatusExpired       SponsoredCardStatuses = "Expired"
)

// ========== Sponsored Card DTOs ==========

// SponsoredCardDto represents sponsored card data
type SponsoredCardDto struct {
	ID             uuid.UUID             `json:"id"`
	Number         string                `json:"number"`
	Barcode        *string               `json:"barcode,omitempty"`
	SeriesID       uuid.UUID             `json:"seriesId"`
	Status         SponsoredCardStatuses `json:"status"`
	OwnerID        *uuid.UUID            `json:"ownerId,omitempty"`
	IssueDate      *time.Time            `json:"issueDate,omitempty"`
	ExpirationDate *time.Time            `json:"expirationDate,omitempty"`
	IsDeleted      bool                  `json:"isDeleted"`
	Groups         []uuid.UUID           `json:"groups,omitempty"`
	Balances       []BalanceDto          `json:"balances,omitempty"`
}

// SponsoredCardListDto represents list of sponsored cards
type SponsoredCardListDto struct {
	Total  int64              `json:"total"`
	Values []SponsoredCardDto `json:"values"`
}

// SponsoredCardTransactionDto represents sponsored card transaction data
type SponsoredCardTransactionDto struct {
	TransactionID   uuid.UUID        `json:"transactionId"`
	OperationDate   time.Time        `json:"operationDate"`
	TransactionType TransactionTypes `json:"transactionType"`
	Amount          float64          `json:"amount"`
	CurrencyID      uuid.UUID        `json:"currencyId"`
	ExpirationDate  *time.Time       `json:"expirationDate,omitempty"`
}

// SponsoredCardTransactionListDto represents list of sponsored card transactions
type SponsoredCardTransactionListDto struct {
	Transactions []SponsoredCardTransactionDto `json:"transactions"`
}

// ChequeDto represents cheque data
type ChequeDto struct {
	ChequeID      uuid.UUID  `json:"chequeId"`
	OperationDate time.Time  `json:"operationDate"`
	StoreID       *uuid.UUID `json:"storeId,omitempty"`
	PosID         *uuid.UUID `json:"posId,omitempty"`
	TotalAmount   float64    `json:"totalAmount"`
}

// ChequeListDto represents list of cheques
type ChequeListDto struct {
	Cheques []ChequeDto `json:"cheques"`
}

// ========== Sponsored Card Requests ==========

// CreateSponsoredCardRequest represents request for creating a sponsored card
type CreateSponsoredCardRequest struct {
	BaseCommand
	SeriesID uuid.UUID `json:"seriesId"`
}

// BlockSponsoredCardRequest represents request for blocking a sponsored card
type BlockSponsoredCardRequest struct {
	BaseCommand
	BlockReasonID *uuid.UUID `json:"blockReasonId,omitempty"`
}

// IssueSponsoredCardRequest represents request for issuing a sponsored card
type IssueSponsoredCardRequest struct {
	BaseCommand
	OwnerID       uuid.UUID  `json:"ownerId"`
	IssueReasonID *uuid.UUID `json:"issueReasonId,omitempty"`
}

// SponsoredCardAccrualRequest represents request for accrual to sponsored card
type SponsoredCardAccrualRequest struct {
	BaseCommand
	CurrencyID     uuid.UUID  `json:"currencyId"`
	Amount         float64    `json:"amount"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// SponsoredCardSubtractRequest represents request for subtract from sponsored card
type SponsoredCardSubtractRequest struct {
	BaseCommand
	CurrencyID uuid.UUID `json:"currencyId"`
	Amount     float64   `json:"amount"`
}
