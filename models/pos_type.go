package models

import (
	"github.com/google/uuid"
)

// ========== POS Type DTOs ==========

// PosTypeDto represents POS type data
type PosTypeDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IsDeleted  bool      `json:"isDeleted"`
	IsArchived bool      `json:"isArchived"`
}

// PosTypeListDto represents list of POS types
type PosTypeListDto struct {
	Total  int64        `json:"total"`
	Values []PosTypeDto `json:"values"`
}

// ========== POS Type Requests ==========

// CreatePosTypeRequest represents request for creating a POS type
type CreatePosTypeRequest struct {
	BaseCommand
	Name string `json:"name"`
}
