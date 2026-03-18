package models

import (
	"github.com/google/uuid"
)

// ========== Store Format DTOs ==========

// StoreFormatDto represents store format data
type StoreFormatDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"isDeleted"`
}

// StoreFormatListDto represents list of store formats
type StoreFormatListDto struct {
	Total  int64            `json:"total"`
	Values []StoreFormatDto `json:"values"`
}

// ========== Store Format Requests ==========

// CreateStoreFormatRequest represents request for creating a store format
type CreateStoreFormatRequest struct {
	BaseCommand
	Name string `json:"name"`
}
