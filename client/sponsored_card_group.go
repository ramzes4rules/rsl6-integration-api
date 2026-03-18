// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// SponsoredCardGroupService handles sponsored card group operations
type SponsoredCardGroupService struct {
	client *Client
}

// Create creates a new sponsored card group
func (s *SponsoredCardGroupService) Create(ctx context.Context, req *models.CreateGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/create", req, headers)
}

// Rename renames a sponsored card group
func (s *SponsoredCardGroupService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/rename", req, headers)
}

// SetColor sets the color of a sponsored card group
func (s *SponsoredCardGroupService) SetColor(ctx context.Context, req *models.SetColorRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/set_color", req, headers)
}

// Delete deletes a sponsored card group
func (s *SponsoredCardGroupService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/delete", req, headers)
}

// Restore restores a deleted sponsored card group
func (s *SponsoredCardGroupService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/restore", req, headers)
}

// MoveToArchive moves a sponsored card group to archive
func (s *SponsoredCardGroupService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/move_to_archive", req, headers)
}

// Unarchive restores a sponsored card group from archive
func (s *SponsoredCardGroupService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *SponsoredCardGroupService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_groups/batch", req, headers)
}

// GetByID retrieves a sponsored card group by ID
func (s *SponsoredCardGroupService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.GroupDto, error) {
	var response models.GroupDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_groups/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of sponsored card groups
func (s *SponsoredCardGroupService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.GroupListDto, error) {
	var response models.GroupListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_groups/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
