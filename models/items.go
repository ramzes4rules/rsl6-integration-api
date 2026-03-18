package models

import (
	"github.com/google/uuid"
)

// ========== Item DTOs ==========

// ItemDto represents item data
type ItemDto struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	PublicName  *string    `json:"publicName,omitempty"`
	Description *string    `json:"description,omitempty"`
	Article     *string    `json:"article,omitempty"`
	ExternalID  *string    `json:"externalId,omitempty"`
	ItemGroupID *uuid.UUID `json:"itemGroupId,omitempty"`
	IsDeleted   bool       `json:"isDeleted"`
}

// ItemListDto represents list of items
type ItemListDto struct {
	Total  int64     `json:"total"`
	Values []ItemDto `json:"values"`
}

// ========== Item Requests ==========

// CreateItemRequest represents request for creating an item
type CreateItemRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetArticleRequest represents request for setting article
type SetArticleRequest struct {
	BaseCommand
	Article string `json:"article"`
}

// SetExternalIDRequest represents request for setting external ID
type SetExternalIDRequest struct {
	BaseCommand
	ExternalID string `json:"externalId"`
}

// SetItemGroupIDRequest represents request for setting item group ID
type SetItemGroupIDRequest struct {
	BaseCommand
	ItemGroupID *uuid.UUID `json:"itemGroupId,omitempty"`
}

// SetRestrictionRequest represents request for setting restriction
type SetRestrictionRequest struct {
	BaseCommand
	Restriction string `json:"restriction"`
}

// AddSaleItemRequest represents request for adding sale item
type AddSaleItemRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
	ExternalID string    `json:"externalId"`
}

// DeleteSaleItemRequest represents request for deleting sale item
type DeleteSaleItemRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
}

// RestoreSaleItemRequest represents request for restoring sale item
type RestoreSaleItemRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
}

// SetSaleItemExternalIDRequest represents request for setting sale item external ID
type SetSaleItemExternalIDRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
	ExternalID string    `json:"externalId"`
}

// SetSaleItemPropertyValueRequest represents request for setting sale item property value
type SetSaleItemPropertyValueRequest struct {
	BaseCommand
	SaleItemID uuid.UUID   `json:"saleItemId"`
	PropertyID uuid.UUID   `json:"propertyId"`
	Value      interface{} `json:"value"`
}
