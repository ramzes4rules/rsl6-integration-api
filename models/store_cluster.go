package models

import (
	"github.com/google/uuid"
)

// ========== Store Cluster DTOs ==========

// StoreClusterDto represents store cluster data
type StoreClusterDto struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	Color      *string     `json:"color,omitempty"`
	IsDeleted  bool        `json:"isDeleted"`
	IsArchived bool        `json:"isArchived"`
	Stores     []uuid.UUID `json:"stores,omitempty"`
}

// StoreClusterListDto represents list of store clusters
type StoreClusterListDto struct {
	Total  int64             `json:"total"`
	Values []StoreClusterDto `json:"values"`
}

// ========== Store Cluster Requests ==========

// AddStoreToClusterRequest represents request for adding store to cluster
type AddStoreToClusterRequest struct {
	BaseCommand
	StoreID uuid.UUID `json:"storeId"`
}

// RemoveStoreFromClusterRequest represents request for removing store from cluster
type RemoveStoreFromClusterRequest struct {
	BaseCommand
	StoreID uuid.UUID `json:"storeId"`
}
