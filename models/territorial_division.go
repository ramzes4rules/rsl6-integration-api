package models

import (
	"github.com/google/uuid"
)

// ========== Territorial Division DTOs ==========

// TerritorialDivisionDto represents territorial division data
type TerritorialDivisionDto struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	ParentID   *uuid.UUID `json:"parentId,omitempty"`
	IsDeleted  bool       `json:"isDeleted"`
	IsArchived bool       `json:"isArchived"`
}

// TerritorialDivisionListDto represents list of territorial divisions
type TerritorialDivisionListDto struct {
	Total  int64                    `json:"total"`
	Values []TerritorialDivisionDto `json:"values"`
}

// ========== Territorial Division Requests ==========

// CreateTerritorialDivisionRequest represents request for creating a territorial division
type CreateTerritorialDivisionRequest struct {
	BaseCommand
	Name     string     `json:"name"`
	ParentID *uuid.UUID `json:"parentId,omitempty"`
}
