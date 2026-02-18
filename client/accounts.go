package client

import (
	"context"
	"net/http"

	"github.com/rsl6/loyalty-client/models"
)

// AccountsService handles account-related operations
type AccountsService struct {
	client *Client
}

// AccrualToCustomer performs accrual to customer account
func (s *AccountsService) AccrualToCustomer(ctx context.Context, req *models.AccrualToCustomerRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/accounts/accrual_to_customer", req, headers)
}

// AccrualToLoyaltyCard performs accrual to loyalty card account
func (s *AccountsService) AccrualToLoyaltyCard(ctx context.Context, req *models.AccrualToLoyaltyCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/accounts/accrual_to_loyalty_card", req, headers)
}

// SubtractFromCustomer performs subtraction from customer account
func (s *AccountsService) SubtractFromCustomer(ctx context.Context, req *models.SubtractFromCustomerRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/accounts/subtract_from_customer", req, headers)
}

// SubtractFromLoyaltyCard performs subtraction from loyalty card account
func (s *AccountsService) SubtractFromLoyaltyCard(ctx context.Context, req *models.SubtractFromLoyaltyCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/accounts/subtract_from_loyalty_card", req, headers)
}

// GetTransactions retrieves list of transactions
func (s *AccountsService) GetTransactions(ctx context.Context, req *models.GetTransactionsRequest) (*models.TransactionListDto, error) {
	var response models.TransactionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/accounts/get_transactions", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
