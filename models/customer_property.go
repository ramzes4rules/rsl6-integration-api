package models

import (
	"github.com/google/uuid"
)

// ========== Customer Property Requests ==========

// CreateCustomerPropertyRequest represents request for creating customer property
type CreateCustomerPropertyRequest struct {
	BaseCommand
	Name       *string       `json:"name,omitempty"`
	Type       PropertyTypes `json:"type"`
	ExternalID *string       `json:"externalId,omitempty"`
}

// RenameCustomerPropertyRequest represents request for renaming customer property
type RenameCustomerPropertyRequest struct {
	BaseCommand
	Name *string `json:"name,omitempty"`
}

// DeleteCustomerPropertyRequest represents request for deleting customer property
type DeleteCustomerPropertyRequest struct {
	BaseCommand
}

// RestoreCustomerPropertyRequest represents request for restoring customer property
type RestoreCustomerPropertyRequest struct {
	BaseCommand
}

// AddEnumCustomerPropertyRequest represents request for adding enum value to customer property
type AddEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
	Name        *string   `json:"name,omitempty"`
}

// RenameEnumCustomerPropertyRequest represents request for renaming enum value of customer property
type RenameEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
	Name        *string   `json:"name,omitempty"`
}

// DeleteEnumCustomerPropertyRequest represents request for deleting enum value from customer property
type DeleteEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
}

// RestoreEnumCustomerPropertyRequest represents request for restoring enum value of customer property
type RestoreEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
}

// CustomerPropertyBatchRequest represents batch request for customer property commands
type CustomerPropertyBatchRequest struct {
	Commands []interface{} `json:"commands"`
}
