package models

import (
	"github.com/google/uuid"
)

// ========== Item Category DTOs ==========

// ItemCategoryDto represents item category data
type ItemCategoryDto struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Color       *string     `json:"color,omitempty"`
	IsDeleted   bool        `json:"isDeleted"`
	IsArchived  bool        `json:"isArchived"`
	Items       []uuid.UUID `json:"items,omitempty"`
	ItemGroups  []uuid.UUID `json:"itemGroups,omitempty"`
}

// ItemCategoryListDto represents list of item categories
type ItemCategoryListDto struct {
	Total  int64             `json:"total"`
	Values []ItemCategoryDto `json:"values"`
}

// ========== Item Category Requests ==========

// CreateItemCategoryRequest represents request for creating an item category
type CreateItemCategoryRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// AddItemToCategoryRequest represents request for adding item to category
type AddItemToCategoryRequest struct {
	BaseCommand
	ItemID uuid.UUID `json:"itemId"`
}

// RemoveItemFromCategoryRequest represents request for removing item from category
type RemoveItemFromCategoryRequest struct {
	BaseCommand
	ItemID uuid.UUID `json:"itemId"`
}

// AddItemGroupToCategoryRequest represents request for adding item group to category
type AddItemGroupToCategoryRequest struct {
	BaseCommand
	ItemGroupID uuid.UUID `json:"itemGroupId"`
}

// RemoveItemGroupFromCategoryRequest represents request for removing item group from category
type RemoveItemGroupFromCategoryRequest struct {
	BaseCommand
	ItemGroupID uuid.UUID `json:"itemGroupId"`
}
