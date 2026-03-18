package models

import (
	"github.com/google/uuid"
)

// ========== Customer Segment Requests ==========

// AddToStaticSegmentRequest represents request for adding customer to static segment
type AddToStaticSegmentRequest struct {
	BaseCommand
	StaticSegmentID uuid.UUID `json:"staticSegmentId"`
}

// RemoveFromStaticSegmentRequest represents request for removing customer from static segment
type RemoveFromStaticSegmentRequest struct {
	BaseCommand
	StaticSegmentID uuid.UUID `json:"staticSegmentId"`
}

// CustomerSegmentBatchRequest represents batch request for customer segment commands
type CustomerSegmentBatchRequest struct {
	Commands []interface{} `json:"commands"`
}
