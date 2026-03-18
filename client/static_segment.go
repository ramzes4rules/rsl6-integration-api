// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// StaticSegmentService handles static segment operations
type StaticSegmentService struct {
	client *Client
}

// Create creates a new static segment
func (s *StaticSegmentService) Create(ctx context.Context, req *models.CreateStaticSegmentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/create", req, headers)
}

// Rename renames a static segment
func (s *StaticSegmentService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/rename", req, headers)
}

// SetDescription sets the description of a static segment
func (s *StaticSegmentService) SetDescription(ctx context.Context, req *models.SetDescriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/set_description", req, headers)
}

// SetColor sets the color of a static segment
func (s *StaticSegmentService) SetColor(ctx context.Context, req *models.SetColorRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/set_color", req, headers)
}

// SetGroup sets the group of a static segment
func (s *StaticSegmentService) SetGroup(ctx context.Context, req *models.SetStaticSegmentGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/set_group", req, headers)
}

// Delete deletes a static segment
func (s *StaticSegmentService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/delete", req, headers)
}

// Restore restores a deleted static segment
func (s *StaticSegmentService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/restore", req, headers)
}

// MoveToArchive moves a static segment to archive
func (s *StaticSegmentService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/move_to_archive", req, headers)
}

// Unarchive restores a static segment from archive
func (s *StaticSegmentService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *StaticSegmentService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/static_segments/batch", req, headers)
}

// GetByID retrieves a static segment by ID
func (s *StaticSegmentService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.StaticSegmentDto, error) {
	var response models.StaticSegmentDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/static_segments/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of static segments
func (s *StaticSegmentService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.StaticSegmentListDto, error) {
	var response models.StaticSegmentListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/static_segments/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
