// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// SponsoredCardIssueReasonService handles sponsored card issue reason operations
type SponsoredCardIssueReasonService struct {
	client *Client
}

// Create creates a new sponsored card issue reason
func (s *SponsoredCardIssueReasonService) Create(ctx context.Context, req *models.CreateReasonRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/create", req, headers)
}

// Rename renames a sponsored card issue reason
func (s *SponsoredCardIssueReasonService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/rename", req, headers)
}

// Delete deletes a sponsored card issue reason
func (s *SponsoredCardIssueReasonService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/delete", req, headers)
}

// Restore restores a deleted sponsored card issue reason
func (s *SponsoredCardIssueReasonService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/restore", req, headers)
}

// MoveToArchive moves a sponsored card issue reason to archive
func (s *SponsoredCardIssueReasonService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/move_to_archive", req, headers)
}

// Unarchive restores a sponsored card issue reason from archive
func (s *SponsoredCardIssueReasonService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *SponsoredCardIssueReasonService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_card_issue_reasons/batch", req, headers)
}

// GetByID retrieves a sponsored card issue reason by ID
func (s *SponsoredCardIssueReasonService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ReasonDto, error) {
	var response models.ReasonDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_issue_reasons/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of sponsored card issue reasons
func (s *SponsoredCardIssueReasonService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.ReasonListDto, error) {
	var response models.ReasonListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_card_issue_reasons/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
