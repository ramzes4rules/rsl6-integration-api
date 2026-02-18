package client

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rsl6/loyalty-client/models"
)

// CountriesService handles country-related operations
type CountriesService struct {
	client *Client
}

// Create creates a new country
func (s *CountriesService) Create(ctx context.Context, req *models.CreateCountryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/countries/create", req, headers)
}

// SetCode sets the country code
func (s *CountriesService) SetCode(ctx context.Context, req *models.SetCountryCodeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/countries/set_code", req, headers)
}

// Delete deletes a country
func (s *CountriesService) Delete(ctx context.Context, req *models.DeleteCountryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/countries/delete", req, headers)
}

// Restore restores a deleted country
func (s *CountriesService) Restore(ctx context.Context, req *models.RestoreCountryRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/countries/restore", req, headers)
}

// GetByID retrieves a country by ID
func (s *CountriesService) GetByID(ctx context.Context, id uuid.UUID) (*models.CountryDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.CountryDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/countries/get_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves list of countries
func (s *CountriesService) GetList(ctx context.Context, req *models.GetListRequest) (*models.CountryListDto, error) {
	var response models.CountryListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/countries/get_list", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
