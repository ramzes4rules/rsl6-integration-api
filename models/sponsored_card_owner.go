package models

import (
	"github.com/google/uuid"
)

// ========== Sponsored Card Owner DTOs ==========

// SponsoredCardOwnerDto represents sponsored card owner data
type SponsoredCardOwnerDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IsDeleted  bool      `json:"isDeleted"`
	IsArchived bool      `json:"isArchived"`
}

// SponsoredCardOwnerListDto represents list of sponsored card owners
type SponsoredCardOwnerListDto struct {
	Total  int64                   `json:"total"`
	Values []SponsoredCardOwnerDto `json:"values"`
}

// ========== Sponsored Card Owner Requests ==========

// CreateSponsoredCardOwnerRequest represents request for creating a sponsored card owner
type CreateSponsoredCardOwnerRequest struct {
	BaseCommand
	Name string `json:"name"`
}
