package models

import (
	"time"

	"github.com/google/uuid"
)

// ========== Loyalty Card Enums ==========

// LoyaltyCardStatuses represents loyalty card status
type LoyaltyCardStatuses string

const (
	LoyaltyCardStatusActive                               LoyaltyCardStatuses = "Active"
	LoyaltyCardStatusBlocked                              LoyaltyCardStatuses = "Blocked"
	LoyaltyCardStatusCreated                              LoyaltyCardStatuses = "Created"
	LoyaltyCardStatusIssued                               LoyaltyCardStatuses = "Issued"
	LoyaltyCardStatusReadyToIssued                        LoyaltyCardStatuses = "ReadyToIssued"
	LoyaltyCardStatusExpired                              LoyaltyCardStatuses = "Expired"
	LoyaltyCardStatusBlockedByCustomerWasBroken           LoyaltyCardStatuses = "BlockedByCustomerWasBroken"
	LoyaltyCardStatusBlockedByCustomerWasLost             LoyaltyCardStatuses = "BlockedByCustomerWasLost"
	LoyaltyCardStatusBlockedByCustomerDeletedPersonalData LoyaltyCardStatuses = "BlockedByCustomerDeletedPersonalData"
)

// ========== Loyalty Card DTOs ==========

// LoyaltyCardDto represents loyalty card data
type LoyaltyCardDto struct {
	ID           uuid.UUID           `json:"id"`
	Number       string              `json:"number"`
	SeriesID     uuid.UUID           `json:"seriesId"`
	Status       LoyaltyCardStatuses `json:"status"`
	CustomerID   *uuid.UUID          `json:"customerId,omitempty"`
	IssueDate    *time.Time          `json:"issueDate,omitempty"`
	ExpireDate   *time.Time          `json:"expireDate,omitempty"`
	ActivateDate *time.Time          `json:"activateDate,omitempty"`
	IsDeleted    bool                `json:"isDeleted"`
}

// LoyaltyCardListDto represents list of loyalty cards
type LoyaltyCardListDto struct {
	Total  int64            `json:"total"`
	Values []LoyaltyCardDto `json:"values"`
}
