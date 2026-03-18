package models

import (
	"github.com/google/uuid"
)

// ========== External Card Series Requests ==========

// CreateExternalGiftCardSeriesRequest represents request for creating external gift card series
type CreateExternalGiftCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// CreateExternalLoyaltyCardSeriesRequest represents request for creating external loyalty card series
type CreateExternalLoyaltyCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// CreateExternalSponsoredCardSeriesRequest represents request for creating external sponsored card series
type CreateExternalSponsoredCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// RenameExternalCardSeriesRequest represents request for renaming external card series
type RenameExternalCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// AddCirculationRequest represents request for adding circulation to external card series
type AddCirculationRequest struct {
	BaseCommand
	CirculationID uuid.UUID `json:"circulationId"`
	Name          string    `json:"name"`
}

// RenameCirculationRequest represents request for renaming circulation
type RenameCirculationRequest struct {
	BaseCommand
	CirculationID uuid.UUID `json:"circulationId"`
	Name          string    `json:"name"`
}

// ========== External Card Series DTOs ==========

// ExternalCardSeriesDto represents external card series data
type ExternalCardSeriesDto struct {
	ID           uuid.UUID                    `json:"id"`
	Name         string                       `json:"name"`
	IsDeleted    bool                         `json:"isDeleted"`
	Circulations []ExternalCardCirculationDto `json:"circulations,omitempty"`
}

// ExternalCardCirculationDto represents external card circulation data
type ExternalCardCirculationDto struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// ExternalCardSeriesListDto represents list of external card series
type ExternalCardSeriesListDto struct {
	Total  int64                   `json:"total"`
	Values []ExternalCardSeriesDto `json:"values"`
}
