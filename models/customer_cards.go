package models

import (
	"github.com/google/uuid"
)

// ========== Customer Cards Requests ==========

// BindLoyaltyCardRequest represents request for binding loyalty card to customer
type BindLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID      uuid.UUID `json:"loyaltyCardId"`
	PinCode            *string   `json:"pinCode,omitempty"`
	InteractionChannel *string   `json:"interactionChannel,omitempty"`
}

// UnbindLoyaltyCardRequest represents request for unbinding loyalty card from customer
type UnbindLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID uuid.UUID `json:"loyaltyCardId"`
}

// CustomerCardsBatchRequest represents batch request for customer cards commands
type CustomerCardsBatchRequest struct {
	Commands []interface{} `json:"commands"`
}
