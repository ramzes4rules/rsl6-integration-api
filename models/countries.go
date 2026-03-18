package models

import (
	"github.com/google/uuid"
)

// ========== Country DTOs ==========

// CountryDto represents country data
type CountryDto struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	IsDeleted bool      `json:"isDeleted"`
}

// CountryListDto represents list of countries
type CountryListDto struct {
	Total  int64        `json:"total"`
	Values []CountryDto `json:"values"`
}

// ========== Country Requests ==========

// CreateCountryRequest represents request for creating country
type CreateCountryRequest struct {
	BaseCommand
	Code string `json:"code"`
}

// SetCountryCodeRequest represents request for setting country code
type SetCountryCodeRequest struct {
	BaseCommand
	Code string `json:"code"`
}

// DeleteCountryRequest represents request for deleting country
type DeleteCountryRequest struct {
	BaseCommand
}

// RestoreCountryRequest represents request for restoring country
type RestoreCountryRequest struct {
	BaseCommand
}
