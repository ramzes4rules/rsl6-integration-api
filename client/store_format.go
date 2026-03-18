// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// StoreFormatService handles store format operations
type StoreFormatService struct {
	client *Client
}

// Create creates a new store format
func (s *StoreFormatService) Create(ctx context.Context, req *models.CreateStoreFormatRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_formats/create", req, headers)
}

// Rename renames a store format
func (s *StoreFormatService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_formats/rename", req, headers)
}

// Delete deletes a store format
func (s *StoreFormatService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_formats/delete", req, headers)
}

// Restore restores a deleted store format
func (s *StoreFormatService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_formats/restore", req, headers)
}

// Batch executes multiple commands in a batch
func (s *StoreFormatService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_formats/batch", req, headers)
}

// GetByID retrieves a store format by ID
func (s *StoreFormatService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.StoreFormatDto, error) {
	var response models.StoreFormatDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/store_formats/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of store formats
func (s *StoreFormatService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.StoreFormatListDto, error) {
	var response models.StoreFormatListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/store_formats/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
