// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// StorePropertyService handles store property operations
type StorePropertyService struct {
	client *Client
}

// Create creates a new store property
func (s *StorePropertyService) Create(ctx context.Context, req *models.CreateStorePropertyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/create", req, headers)
}

// Rename renames a store property
func (s *StorePropertyService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/rename", req, headers)
}

// Delete deletes a store property
func (s *StorePropertyService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/delete", req, headers)
}

// Restore restores a deleted store property
func (s *StorePropertyService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/restore", req, headers)
}

// AddEnum adds an enum value to a store property
func (s *StorePropertyService) AddEnum(ctx context.Context, req *models.AddEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/add_enum", req, headers)
}

// RenameEnum renames an enum value of a store property
func (s *StorePropertyService) RenameEnum(ctx context.Context, req *models.RenameEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/rename_enum", req, headers)
}

// DeleteEnum deletes an enum value from a store property
func (s *StorePropertyService) DeleteEnum(ctx context.Context, req *models.DeleteEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/delete_enum", req, headers)
}

// RestoreEnum restores an enum value of a store property
func (s *StorePropertyService) RestoreEnum(ctx context.Context, req *models.RestoreEnumRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/restore_enum", req, headers)
}

// Batch executes multiple commands in a batch
func (s *StorePropertyService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/store_properties/batch", req, headers)
}

// GetByID retrieves a store property by ID
func (s *StorePropertyService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.PropertyDefinitionDto, error) {
	var response models.PropertyDefinitionDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/store_properties/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of store properties
func (s *StorePropertyService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.PropertyDefinitionListDto, error) {
	var response models.PropertyDefinitionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/store_properties/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
