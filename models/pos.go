package models

import (
	"github.com/google/uuid"
)

// ========== POS DTOs ==========

// PosDto represents POS data
type PosDto struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	StoreID    uuid.UUID  `json:"storeId"`
	PosTypeID  *uuid.UUID `json:"posTypeId,omitempty"`
	IsBlocked  bool       `json:"isBlocked"`
	IsDeleted  bool       `json:"isDeleted"`
	IsArchived bool       `json:"isArchived"`
}

// PosListDto represents list of POSes
type PosListDto struct {
	Total  int64    `json:"total"`
	Values []PosDto `json:"values"`
}

// ========== POS Requests ==========

// CreatePosRequest represents request for creating a POS
type CreatePosRequest struct {
	BaseCommand
	Name    string    `json:"name"`
	StoreID uuid.UUID `json:"storeId"`
}

// SetPosTypeRequest represents request for setting POS type
type SetPosTypeRequest struct {
	BaseCommand
	PosTypeID *uuid.UUID `json:"posTypeId,omitempty"`
}

// SetPosAuthorizationRequest represents request for setting POS authorization
type SetPosAuthorizationRequest struct {
	BaseCommand
	Login    string `json:"login"`
	Password string `json:"password"`
}

// MoveToStoreRequest represents request for moving POS to store
type MoveToStoreRequest struct {
	BaseCommand
	StoreID uuid.UUID `json:"storeId"`
}

// BlockPosRequest represents request for blocking POS
type BlockPosRequest struct {
	BaseCommand
}

// UnblockPosRequest represents request for unblocking POS
type UnblockPosRequest struct {
	BaseCommand
}
