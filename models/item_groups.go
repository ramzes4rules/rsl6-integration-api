package models

import (
	"github.com/google/uuid"
)

// ========== Item Group DTOs ==========

// ItemGroupDto represents item group data
type ItemGroupDto struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	ParentID  *uuid.UUID `json:"parentId,omitempty"`
	IsDeleted bool       `json:"isDeleted"`
}

// ItemGroupListDto represents list of item groups
type ItemGroupListDto struct {
	Total  int64          `json:"total"`
	Values []ItemGroupDto `json:"values"`
}

// ========== Item Group Requests ==========

// CreateItemGroupRequest represents request for creating an item group
type CreateItemGroupRequest struct {
	BaseCommand
	Name string `json:"name"`
}
