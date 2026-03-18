package models

import (
	"github.com/google/uuid"
)

// ========== Static Segment DTOs ==========

// StaticSegmentDto represents static segment data
type StaticSegmentDto struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Color       *string    `json:"color,omitempty"`
	GroupID     *uuid.UUID `json:"groupId,omitempty"`
	IsDeleted   bool       `json:"isDeleted"`
	IsArchived  bool       `json:"isArchived"`
}

// StaticSegmentListDto represents list of static segments
type StaticSegmentListDto struct {
	Total  int64              `json:"total"`
	Values []StaticSegmentDto `json:"values"`
}

// ========== Static Segment Requests ==========

// CreateStaticSegmentRequest represents request for creating a static segment
type CreateStaticSegmentRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetStaticSegmentGroupRequest represents request for setting static segment group
type SetStaticSegmentGroupRequest struct {
	BaseCommand
	GroupID *uuid.UUID `json:"groupId,omitempty"`
}
