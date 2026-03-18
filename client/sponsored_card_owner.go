// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// SponsoredCardOwnerService handles sponsored card owner operations
type SponsoredCardOwnerService struct {
	client *Client
}

// Create creates a new sponsored card owner
func (s *SponsoredCardOwnerService) Create(ctx context.Context, req *models.CreateSponsoredCardOwnerRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/create", req, headers)
}

// Rename renames a sponsored card owner
func (s *SponsoredCardOwnerService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/rename", req, headers)
}

// Delete deletes a sponsored card owner
func (s *SponsoredCardOwnerService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/delete", req, headers)
}

// Restore restores a deleted sponsored card owner
func (s *SponsoredCardOwnerService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/restore", req, headers)
}

// MoveToArchive moves a sponsored card owner to archive
func (s *SponsoredCardOwnerService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/move_to_archive", req, headers)
}

// Unarchive restores a sponsored card owner from archive
func (s *SponsoredCardOwnerService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *SponsoredCardOwnerService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_owners/batch", req, headers)
}

// GetByID retrieves a sponsored card owner by ID
func (s *SponsoredCardOwnerService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.SponsoredCardOwnerDto, error) {
	var response models.SponsoredCardOwnerDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_owners/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of sponsored card owners
func (s *SponsoredCardOwnerService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.SponsoredCardOwnerListDto, error) {
	var response models.SponsoredCardOwnerListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_owners/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
