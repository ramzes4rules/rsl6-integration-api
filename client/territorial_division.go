// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// TerritorialDivisionService handles territorial division operations
type TerritorialDivisionService struct {
	client *Client
}

// Create creates a new territorial division
func (s *TerritorialDivisionService) Create(ctx context.Context, req *models.CreateTerritorialDivisionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/create", req, headers)
}

// Rename renames a territorial division
func (s *TerritorialDivisionService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/rename", req, headers)
}

// SetParent sets the parent of a territorial division
func (s *TerritorialDivisionService) SetParent(ctx context.Context, req *models.SetParentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/set_parent", req, headers)
}

// Delete deletes a territorial division
func (s *TerritorialDivisionService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/delete", req, headers)
}

// Restore restores a deleted territorial division
func (s *TerritorialDivisionService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/restore", req, headers)
}

// MoveToArchive moves a territorial division to archive
func (s *TerritorialDivisionService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/move_to_archive", req, headers)
}

// Unarchive restores a territorial division from archive
func (s *TerritorialDivisionService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *TerritorialDivisionService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/territorial_divisions/batch", req, headers)
}

// GetByID retrieves a territorial division by ID
func (s *TerritorialDivisionService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.TerritorialDivisionDto, error) {
	var response models.TerritorialDivisionDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/territorial_divisions/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of territorial divisions
func (s *TerritorialDivisionService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.TerritorialDivisionListDto, error) {
	var response models.TerritorialDivisionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/territorial_divisions/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
