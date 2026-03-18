package models

import (
	"github.com/google/uuid"
)

// ========== Segment Group DTOs ==========

// SegmentGroupDto represents segment group data
type SegmentGroupDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDeleted   bool      `json:"isDeleted"`
	IsArchived  bool      `json:"isArchived"`
}

// SegmentGroupListDto represents list of segment groups
type SegmentGroupListDto struct {
	Total  int64             `json:"total"`
	Values []SegmentGroupDto `json:"values"`
}
