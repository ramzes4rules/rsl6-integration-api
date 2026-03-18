package models

import (
	"github.com/google/uuid"
)

// ========== Currency DTOs ==========

// CurrencyDto represents currency data
type CurrencyDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	PublicName  *string   `json:"publicName,omitempty"`
	Rate        float64   `json:"rate"`
	IsActive    bool      `json:"isActive"`
	IsDeleted   bool      `json:"isDeleted"`
}

// CurrencyListDto represents list of currencies
type CurrencyListDto struct {
	Total  int64         `json:"total"`
	Values []CurrencyDto `json:"values"`
}

// ========== Currency Requests ==========

// CreateCurrencyRequest represents request for creating currency
type CreateCurrencyRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// RenameCurrencyRequest represents request for renaming currency
type RenameCurrencyRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetCurrencyDescriptionRequest represents request for setting currency description
type SetCurrencyDescriptionRequest struct {
	BaseCommand
	Description string `json:"description"`
}

// SetCurrencyRateRequest represents request for setting currency rate
type SetCurrencyRateRequest struct {
	BaseCommand
	Rate float64 `json:"rate"`
}

// ActivateCurrencyRequest represents request for activating currency
type ActivateCurrencyRequest struct {
	BaseCommand
}

// DeactivateCurrencyRequest represents request for deactivating currency
type DeactivateCurrencyRequest struct {
	BaseCommand
}

// SetCurrencyPublicNameRequest represents request for setting currency public name
type SetCurrencyPublicNameRequest struct {
	BaseCommand
	PublicName string `json:"publicName"`
}

// SetCurrencyCalculateRoundRuleRequest represents request for setting currency calculate round rule
type SetCurrencyCalculateRoundRuleRequest struct {
	BaseCommand
	CalculateRoundRule string `json:"calculateRoundRule"`
}

// SetCurrencyCaptionRequest represents request for setting currency caption
type SetCurrencyCaptionRequest struct {
	BaseCommand
	Caption string `json:"caption"`
}
