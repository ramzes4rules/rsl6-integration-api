// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// StoreService handles store operations
type StoreService struct {
	client *Client
}

// Create creates a new store
func (s *StoreService) Create(ctx context.Context, req *models.CreateStoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/create", req, headers)
}

// Rename renames a store
func (s *StoreService) Rename(ctx context.Context, req *models.RenameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/rename", req, headers)
}

// SetPublicName sets the public name of a store
func (s *StoreService) SetPublicName(ctx context.Context, req *models.SetPublicNameRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_public_name", req, headers)
}

// SetInternalNumber sets the internal number of a store
func (s *StoreService) SetInternalNumber(ctx context.Context, req *models.SetStoreInternalNumberRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_internal_number", req, headers)
}

// SetTimeZone sets the time zone of a store
func (s *StoreService) SetTimeZone(ctx context.Context, req *models.SetStoreTimeZoneRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_time_zone", req, headers)
}

// SetAddress sets the address of a store
func (s *StoreService) SetAddress(ctx context.Context, req *models.SetStoreAddressRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_address", req, headers)
}

// Close closes a store
func (s *StoreService) Close(ctx context.Context, req *models.CloseStoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/close", req, headers)
}

// Open opens a store
func (s *StoreService) Open(ctx context.Context, req *models.OpenStoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/open", req, headers)
}

// SetRetailSpace sets the retail space of a store
func (s *StoreService) SetRetailSpace(ctx context.Context, req *models.SetStoreRetailSpaceRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_retail_space", req, headers)
}

// AddMediaContent adds media content to a store
func (s *StoreService) AddMediaContent(ctx context.Context, req *models.AddMediaContentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/add_media_content", req, headers)
}

// RenameMediaContent renames media content of a store
func (s *StoreService) RenameMediaContent(ctx context.Context, req *models.RenameMediaContentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/rename_media_content", req, headers)
}

// SetMediaContentSource sets the source of media content
func (s *StoreService) SetMediaContentSource(ctx context.Context, req *models.SetMediaContentSourceRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_media_content_source", req, headers)
}

// SetMediaContentType sets the type of media content
func (s *StoreService) SetMediaContentType(ctx context.Context, req *models.SetMediaContentTypeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_media_content_type", req, headers)
}

// RemoveMediaContent removes media content from a store
func (s *StoreService) RemoveMediaContent(ctx context.Context, req *models.RemoveMediaContentRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/remove_media_content", req, headers)
}

// SetFormat sets the format of a store
func (s *StoreService) SetFormat(ctx context.Context, req *models.SetStoreFormatRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_format", req, headers)
}

// SetTerritorialDivision sets the territorial division of a store
func (s *StoreService) SetTerritorialDivision(ctx context.Context, req *models.SetStoreTerritorialDivisionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_territorial_division", req, headers)
}

// SetOpeningHours sets the opening hours of a store
func (s *StoreService) SetOpeningHours(ctx context.Context, req *models.SetStoreOpeningHoursRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_opening_hours", req, headers)
}

// SetIntPropertyValue sets an integer property value for a store
func (s *StoreService) SetIntPropertyValue(ctx context.Context, req *models.SetIntPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_int_property_value", req, headers)
}

// SetStringPropertyValue sets a string property value for a store
func (s *StoreService) SetStringPropertyValue(ctx context.Context, req *models.SetStringPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_string_property_value", req, headers)
}

// SetDatePropertyValue sets a date property value for a store
func (s *StoreService) SetDatePropertyValue(ctx context.Context, req *models.SetDatePropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_date_property_value", req, headers)
}

// SetBooleanPropertyValue sets a boolean property value for a store
func (s *StoreService) SetBooleanPropertyValue(ctx context.Context, req *models.SetBooleanPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_boolean_property_value", req, headers)
}

// SetEnumPropertyValue sets an enum property value for a store
func (s *StoreService) SetEnumPropertyValue(ctx context.Context, req *models.SetEnumPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_enum_property_value", req, headers)
}

// SetEnumsPropertyValue sets multiple enum property values for a store
func (s *StoreService) SetEnumsPropertyValue(ctx context.Context, req *models.SetEnumsPropertyValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_enums_property_value", req, headers)
}

// SetDefaultImage sets the default image for a store
func (s *StoreService) SetDefaultImage(ctx context.Context, req *models.SetDefaultImageRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_default_image", req, headers)
}

// ClearDefaultImage clears the default image for a store
func (s *StoreService) ClearDefaultImage(ctx context.Context, req *models.ClearDefaultImageRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/clear_default_image", req, headers)
}

// SetLocationCoordinates sets the location coordinates of a store
func (s *StoreService) SetLocationCoordinates(ctx context.Context, req *models.SetStoreLocationCoordinatesRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/set_location_coordinates", req, headers)
}

// Delete deletes a store
func (s *StoreService) Delete(ctx context.Context, req *models.DeleteRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/delete", req, headers)
}

// Restore restores a deleted store
func (s *StoreService) Restore(ctx context.Context, req *models.RestoreRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/restore", req, headers)
}

// MoveToArchive moves a store to archive
func (s *StoreService) MoveToArchive(ctx context.Context, req *models.MoveToArchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/move_to_archive", req, headers)
}

// Unarchive restores a store from archive
func (s *StoreService) Unarchive(ctx context.Context, req *models.UnarchiveRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/unarchive", req, headers)
}

// Batch executes multiple commands in a batch
func (s *StoreService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/stores/batch", req, headers)
}

// GetByID retrieves a store by ID
func (s *StoreService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.StoreDto, error) {
	var response models.StoreDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/stores/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of stores
func (s *StoreService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.StoreListDto, error) {
	var response models.StoreListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/stores/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
