package models

import (
	"time"

	"github.com/google/uuid"
)

// ========== Gift Card Enums ==========

// GiftCardStatuses represents gift card status
type GiftCardStatuses string

const (
	GiftCardStatusCreated       GiftCardStatuses = "Created"
	GiftCardStatusReadyToIssued GiftCardStatuses = "ReadyToIssued"
	GiftCardStatusIssued        GiftCardStatuses = "Issued"
	GiftCardStatusActive        GiftCardStatuses = "Active"
	GiftCardStatusBlocked       GiftCardStatuses = "Blocked"
	GiftCardStatusExpired       GiftCardStatuses = "Expired"
	GiftCardStatusFullyRedeemed GiftCardStatuses = "FullyRedeemed"
)

// ========== Gift Card DTOs ==========

// GiftCardDto represents gift card data
type GiftCardDto struct {
	ID             uuid.UUID        `json:"id"`
	Number         string           `json:"number"`
	Barcode        *string          `json:"barcode,omitempty"`
	SeriesID       uuid.UUID        `json:"seriesId"`
	Status         GiftCardStatuses `json:"status"`
	Balance        float64          `json:"balance"`
	InitialAmount  float64          `json:"initialAmount"`
	IssueDate      *time.Time       `json:"issueDate,omitempty"`
	ExpirationDate *time.Time       `json:"expirationDate,omitempty"`
	IsDeleted      bool             `json:"isDeleted"`
	Groups         []uuid.UUID      `json:"groups,omitempty"`
}

// GiftCardListDto represents list of gift cards
type GiftCardListDto struct {
	Total  int64         `json:"total"`
	Values []GiftCardDto `json:"values"`
}

// ========== Gift Card Requests ==========

// CreateGiftCardRequest represents request for creating a gift card
type CreateGiftCardRequest struct {
	BaseCommand
	SeriesID uuid.UUID `json:"seriesId"`
}

// SetGiftCardValidityPeriodRequest represents request for setting gift card validity period
type SetGiftCardValidityPeriodRequest struct {
	BaseCommand
	ValidityPeriod int `json:"validityPeriod"`
}

// IssueGiftCardRequest represents request for issuing a gift card
type IssueGiftCardRequest struct {
	BaseCommand
	Amount        float64    `json:"amount"`
	IssueReasonID *uuid.UUID `json:"issueReasonId,omitempty"`
}

// BlockGiftCardRequest represents request for blocking a gift card
type BlockGiftCardRequest struct {
	BaseCommand
	BlockReasonID *uuid.UUID `json:"blockReasonId,omitempty"`
}

// SubtractGiftCardRequest represents request for subtracting from gift card
type SubtractGiftCardRequest struct {
	BaseCommand
	Amount float64 `json:"amount"`
}

// RefundGiftCardRequest represents request for refunding to gift card
type RefundGiftCardRequest struct {
	BaseCommand
	Amount float64 `json:"amount"`
}
