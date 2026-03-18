// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// SegmentGroupService handles segment group operations
type SegmentGroupService struct {
	client *Client
}

// Create creates a new segment group
func (s *SegmentGroupService) Create(ctx context.Context, req *models.CreateGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/create", req, headers)
}

// Rename renames a segment group
func (s *SegmentGroupService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/rename", req, headers)
}

// SetDescription sets the description of a segment group
func (s *SegmentGroupService) SetDescription(ctx context.Context, req *models.SetDescriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/set_description", req, headers)
}

// Delete deletes a segment group
func (s *SegmentGroupService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/delete", req, headers)
}

// Restore restores a deleted segment group
func (s *SegmentGroupService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/restore", req, headers)
}

// MoveToArchive moves a segment group to archive
func (s *SegmentGroupService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/move_to_archive", req, headers)
}

// Unarchive restores a segment group from archive
func (s *SegmentGroupService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *SegmentGroupService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/segment_groups/batch", req, headers)
}

// GetByID retrieves a segment group by ID
func (s *SegmentGroupService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.SegmentGroupDto, error) {
	var response models.SegmentGroupDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/segment_groups/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of segment groups
func (s *SegmentGroupService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.SegmentGroupListDto, error) {
	var response models.SegmentGroupListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/segment_groups/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
