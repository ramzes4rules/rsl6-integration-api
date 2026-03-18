// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// SponsoredCardBlockReasonService handles sponsored card block reason operations
type SponsoredCardBlockReasonService struct {
	client *Client
}

// Create creates a new sponsored card block reason
func (s *SponsoredCardBlockReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/create", req, headers)
}

// Rename renames a sponsored card block reason
func (s *SponsoredCardBlockReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/rename", req, headers)
}

// Delete deletes a sponsored card block reason
func (s *SponsoredCardBlockReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/delete", req, headers)
}

// Restore restores a deleted sponsored card block reason
func (s *SponsoredCardBlockReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/restore", req, headers)
}

// MoveToArchive moves a sponsored card block reason to archive
func (s *SponsoredCardBlockReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/move_to_archive", req, headers)
}

// Unarchive restores a sponsored card block reason from archive
func (s *SponsoredCardBlockReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *SponsoredCardBlockReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_block_reasons/batch", req, headers)
}

// GetByID retrieves a sponsored card block reason by ID
func (s *SponsoredCardBlockReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_block_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of sponsored card block reasons
func (s *SponsoredCardBlockReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_block_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
