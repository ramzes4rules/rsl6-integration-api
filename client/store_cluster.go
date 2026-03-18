// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// StoreClusterService handles store cluster operations
type StoreClusterService struct {
	client *Client
}

// Create creates a new store cluster
func (s *StoreClusterService) Create(ctx context.Context, req *models.CreateGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/create", req, headers)
}

// Rename renames a store cluster
func (s *StoreClusterService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/rename", req, headers)
}

// SetColor sets the color of a store cluster
func (s *StoreClusterService) SetColor(ctx context.Context, req *models.SetColorRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/set_color", req, headers)
}

// Delete deletes a store cluster
func (s *StoreClusterService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/delete", req, headers)
}

// Restore restores a deleted store cluster
func (s *StoreClusterService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/restore", req, headers)
}

// MoveToArchive moves a store cluster to archive
func (s *StoreClusterService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/move_to_archive", req, headers)
}

// Unarchive restores a store cluster from archive
func (s *StoreClusterService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/unarchive", req, headers)
}

// AddStore adds a store to the cluster
func (s *StoreClusterService) AddStore(ctx context.Context, req *models.AddStoreToClusterRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/add_store", req, headers)
}

// RemoveStore removes a store from the cluster
func (s *StoreClusterService) RemoveStore(ctx context.Context, req *models.RemoveStoreFromClusterRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/remove_store", req, headers)
}

// Batch executes multiple commands in a batch
func (s *StoreClusterService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_clusters/batch", req, headers)
}

// GetByID retrieves a store cluster by ID
func (s *StoreClusterService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.StoreClusterDto, error) {
	var response models.StoreClusterDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/store_clusters/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of store clusters
func (s *StoreClusterService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.StoreClusterListDto, error) {
	var response models.StoreClusterListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/store_clusters/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
