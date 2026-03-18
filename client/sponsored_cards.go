// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"context"
	"net/http"

	"github.com/rsl6/rsloyalty/models"
)

// SponsoredCardService handles sponsored card operations
type SponsoredCardService struct {
	client *Client
}

// Create creates a new sponsored card
func (s *SponsoredCardService) Create(ctx context.Context, req *models.CreateSponsoredCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/create", req, headers)
}

// Block blocks a sponsored card
func (s *SponsoredCardService) Block(ctx context.Context, req *models.BlockSponsoredCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/block", req, headers)
}

// Unblock unblocks a sponsored card
func (s *SponsoredCardService) Unblock(ctx context.Context, req *models.UnblockCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/unblock", req, headers)
}

// ConfirmReadinessForIssuance confirms readiness for issuance
func (s *SponsoredCardService) ConfirmReadinessForIssuance(ctx context.Context, req *models.ConfirmReadinessForIssuanceRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/confirm_readiness_for_issuance", req, headers)
}

// Issue issues a sponsored card
func (s *SponsoredCardService) Issue(ctx context.Context, req *models.IssueSponsoredCardRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/issue", req, headers)
}

// CancelIssue cancels issued sponsored card
func (s *SponsoredCardService) CancelIssue(ctx context.Context, req *models.CancelIssueRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/cancel_issue", req, headers)
}

// AddSponsoredCardGroup adds a sponsored card to a group
func (s *SponsoredCardService) AddSponsoredCardGroup(ctx context.Context, req *models.AddCardGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/add_sponsored_card_group", req, headers)
}

// RemoveSponsoredCardGroup removes a sponsored card from a group
func (s *SponsoredCardService) RemoveSponsoredCardGroup(ctx context.Context, req *models.RemoveCardGroupRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/remove_sponsored_card_group", req, headers)
}

// Accrual adds funds to sponsored card
func (s *SponsoredCardService) Accrual(ctx context.Context, req *models.SponsoredCardAccrualRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/accrual", req, headers)
}

// Subtract subtracts funds from sponsored card
func (s *SponsoredCardService) Subtract(ctx context.Context, req *models.SponsoredCardSubtractRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/subtract", req, headers)
}

// SetExpirationDate sets the expiration date of a sponsored card
func (s *SponsoredCardService) SetExpirationDate(ctx context.Context, req *models.SetCardExpirationDateRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/set_expiration_date", req, headers)
}

// SetNumber sets the number of a sponsored card
func (s *SponsoredCardService) SetNumber(ctx context.Context, req *models.SetCardNumberRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/set_number", req, headers)
}

// SetBarcode sets the barcode of a sponsored card
func (s *SponsoredCardService) SetBarcode(ctx context.Context, req *models.SetCardBarcodeRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/set_barcode", req, headers)
}

// Batch executes multiple commands in a batch
func (s *SponsoredCardService) Batch(ctx context.Context, req *models.BatchRequest, headers *models.RequestHeaders) error {
	return s.client.doCommand(ctx, "/api/v2/sponsored_cards/batch", req, headers)
}

// GetByID retrieves a sponsored card by ID
func (s *SponsoredCardService) GetByID(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.SponsoredCardDto, error) {
	var response models.SponsoredCardDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_cards/get_by_id", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetList retrieves a list of sponsored cards
func (s *SponsoredCardService) GetList(ctx context.Context, req *models.GetListRequest, headers *models.RequestHeaders) (*models.SponsoredCardListDto, error) {
	var response models.SponsoredCardListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_cards/get_list", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCheques retrieves cheques for a sponsored card
func (s *SponsoredCardService) GetCheques(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.ChequeListDto, error) {
	var response models.ChequeListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_cards/get_cheques", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTransactions retrieves transactions for a sponsored card
func (s *SponsoredCardService) GetTransactions(ctx context.Context, req *models.GetByIdRequest, headers *models.RequestHeaders) (*models.SponsoredCardTransactionListDto, error) {
	var response models.SponsoredCardTransactionListDto
	err := s.client.doRequestWithResponse(ctx, http.MethodPost, "/api/v2/sponsored_cards/get_transactions", req, headers, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
