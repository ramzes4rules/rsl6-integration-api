package client

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rsl6/loyalty-client/models"
)

// LoyaltyCardsService handles loyalty card-related operations
type LoyaltyCardsService struct {
	client *Client
}

// GetByID retrieves a loyalty card by ID
func (s *LoyaltyCardsService) GetByID(ctx context.Context, id uuid.UUID) (*models.LoyaltyCardDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.LoyaltyCardDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_cards/get_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetByNumber retrieves a loyalty card by number
func (s *LoyaltyCardsService) GetByNumber(ctx context.Context, number string) (*models.LoyaltyCardDto, error) {
	req := map[string]string{"number": number}
	var response models.LoyaltyCardDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_cards/get_by_number", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves list of loyalty cards
func (s *LoyaltyCardsService) GetList(ctx context.Context, req *models.GetListRequest) (*models.LoyaltyCardListDto, error) {
	var response models.LoyaltyCardListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_cards/get_list", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBalanceByID retrieves loyalty card balance by ID
func (s *LoyaltyCardsService) GetBalanceByID(ctx context.Context, id uuid.UUID) (*models.BalanceListDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.BalanceListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_cards/get_balance_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTransactionsByID retrieves loyalty card transactions by ID
func (s *LoyaltyCardsService) GetTransactionsByID(ctx context.Context, id uuid.UUID) (*models.TransactionListDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.TransactionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/loyalty_cards/get_transactions_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Activate activates a loyalty card
func (s *LoyaltyCardsService) Activate(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/loyalty_cards/activate", req, headers)
}

// Block blocks a loyalty card
func (s *LoyaltyCardsService) Block(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/loyalty_cards/block", req, headers)
}

// Unblock unblocks a loyalty card
func (s *LoyaltyCardsService) Unblock(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/loyalty_cards/unblock", req, headers)
}

// Delete deletes a loyalty card
func (s *LoyaltyCardsService) Delete(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/loyalty_cards/delete", req, headers)
}

// Restore restores a deleted loyalty card
func (s *LoyaltyCardsService) Restore(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/loyalty_cards/restore", req, headers)
}
