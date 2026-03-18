package models

// ========== Item Property Requests ==========

// CreateItemPropertyRequest represents request for creating an item property
type CreateItemPropertyRequest struct {
	BaseCommand
	Name string        `json:"name"`
	Type PropertyTypes `json:"type"`
}
