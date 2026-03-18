package models

// ========== Store Property Requests ==========

// CreateStorePropertyRequest represents request for creating a store property
type CreateStorePropertyRequest struct {
	BaseCommand
	Name string        `json:"name"`
	Type PropertyTypes `json:"type"`
}
