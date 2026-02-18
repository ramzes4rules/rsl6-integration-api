package client

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rsl6/loyalty-client/models"
)

// CurrenciesService handles currency-related operations
type CurrenciesService struct {
	client *Client
}

// Create creates a new currency
func (s *CurrenciesService) Create(ctx context.Context, req *models.CreateCurrencyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/currencies/create", req, headers)
}

// Rename renames a currency
func (s *CurrenciesService) Rename(ctx context.Context, req *models.RenameCurrencyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/currencies/rename", req, headers)
}

// SetDescription sets the currency description
func (s *CurrenciesService) SetDescription(ctx context.Context, req *models.SetCurrencyDescriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/currencies/set_description", req, headers)
}

// SetRate sets the currency rate
func (s *CurrenciesService) SetRate(ctx context.Context, req *models.SetCurrencyRateRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/currencies/set_rate", req, headers)
}

// Activate activates a currency
func (s *CurrenciesService) Activate(ctx context.Context, req *models.ActivateCurrencyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/currencies/activate", req, headers)
}

// Deactivate deactivates a currency
func (s *CurrenciesService) Deactivate(ctx context.Context, req *models.DeactivateCurrencyRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/currencies/deactivate", req, headers)
}

// Delete deletes a currency
func (s *CurrenciesService) Delete(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/currencies/delete", req, headers)
}

// Restore restores a deleted currency
func (s *CurrenciesService) Restore(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/currencies/restore", req, headers)
}

// GetByID retrieves a currency by ID
func (s *CurrenciesService) GetByID(ctx context.Context, id uuid.UUID) (*models.CurrencyDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.CurrencyDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/currencies/get_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves list of currencies
func (s *CurrenciesService) GetList(ctx context.Context, req *models.GetListRequest) (*models.CurrencyListDto, error) {
	var response models.CurrencyListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/currencies/get_list", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
