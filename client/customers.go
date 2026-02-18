package client

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rsl6/loyalty-client/models"
)

// CustomersService handles customer-related operations
type CustomersService struct {
	client *Client
}

// Rename renames a customer
func (s *CustomersService) Rename(ctx context.Context, req *models.RenameCustomerRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customers/rename", req, headers)
}

// SetCommunicationValue sets customer communication value (phone or email)
func (s *CustomersService) SetCommunicationValue(ctx context.Context, req *models.SetCustomerCommunicationValueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customers/set_communication_value", req, headers)
}

// SetBirthday sets customer birthday
func (s *CustomersService) SetBirthday(ctx context.Context, req *models.SetCustomerBirthdayRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customers/set_birthday", req, headers)
}

// SetAddress sets customer address
func (s *CustomersService) SetAddress(ctx context.Context, req *models.SetCustomerAddressRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customers/set_address", req, headers)
}

// AllowSubscription allows subscription for customer
func (s *CustomersService) AllowSubscription(ctx context.Context, req *models.AllowSubscriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customers/allow_subscription", req, headers)
}

// DisallowSubscription disallows subscription for customer
func (s *CustomersService) DisallowSubscription(ctx context.Context, req *models.DisallowSubscriptionRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/customers/disallow_subscription", req, headers)
}

// RemovePersonalData removes personal data from customer
func (s *CustomersService) RemovePersonalData(ctx context.Context, id uuid.UUID, headers *models.RequestHeaders) error {
	req := &models.BaseCommand{ID: id}
	return s.client.doCommand(ctx, "/api/v2/customers/remove_personal_data", req, headers)
}

// GetByID retrieves a customer by ID
func (s *CustomersService) GetByID(ctx context.Context, id uuid.UUID) (*models.CustomerDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.CustomerDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customers/get_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetByCommunicationValue retrieves a customer by communication value (phone or email)
func (s *CustomersService) GetByCommunicationValue(ctx context.Context, req *models.GetByCommunicationValueRequest) (*models.CustomerDto, error) {
	var response models.CustomerDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customers/get_by_communication_value", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves list of customers
func (s *CustomersService) GetList(ctx context.Context, req *models.GetListRequest) (*models.CustomerListDto, error) {
	var response models.CustomerListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customers/get_list", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBalanceByID retrieves customer balance by ID
func (s *CustomersService) GetBalanceByID(ctx context.Context, id uuid.UUID) (*models.CustomerBalanceDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.CustomerBalanceDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customers/get_balance_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTransactionsByID retrieves customer transactions by ID
func (s *CustomersService) GetTransactionsByID(ctx context.Context, id uuid.UUID) (*models.TransactionListDto, error) {
	req := &models.GetByIdRequest{ID: id}
	var response models.TransactionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/customers/get_transactions_by_id", req, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
