// Package client provides HTTP client for RS Loyalty API v2
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rsl6/rsloyalty/models"
)

// Config represents client configuration
type Config struct {
	BaseURL    string
	Timeout    time.Duration
	HTTPClient *http.Client
	Headers    map[string]string
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "http://localhost:8080",
		Timeout: 30 * time.Second,
	}
}

// Client represents RS Loyalty API client
type Client struct {
	config     *Config
	httpClient *http.Client

	// Services
	Accounts                    *AccountsService
	Countries                   *CountriesService
	Currencies                  *CurrenciesService
	Customers                   *CustomersService
	LoyaltyCards                *LoyaltyCardsService
	CustomerCards               *CustomerCardsService
	CustomerProperty            *CustomerPropertyService
	CustomerSegment             *CustomerSegmentService
	ExternalGiftCardSeries      *ExternalGiftCardSeriesService
	ExternalLoyaltyCardSeries   *ExternalLoyaltyCardSeriesService
	ExternalSponsoredCardSeries *ExternalSponsoredCardSeriesService
	GiftCards                   *GiftCardService
	GiftCardBlockReason         *GiftCardBlockReasonService
	GiftCardGroup               *GiftCardGroupService
	GiftCardIssueReason         *GiftCardIssueReasonService
	Items                       *ItemService
	ItemCategories              *ItemCategoryService
	ItemGroups                  *ItemGroupService
	ItemProperties              *ItemPropertyService
	LoyaltyCardBlockReason      *LoyaltyCardBlockReasonService
	LoyaltyCardGroup            *LoyaltyCardGroupService
	LoyaltyCardIssueReason      *LoyaltyCardIssueReasonService
	ManualAccrualReason         *ManualAccrualReasonService
	OpeningHours                *OpeningHoursService
	Pos                         *PosService
	PosType                     *PosTypeService
	SegmentGroup                *SegmentGroupService
	SponsoredCards              *SponsoredCardService
	SponsoredCardBlockReason    *SponsoredCardBlockReasonService
	SponsoredCardGroup          *SponsoredCardGroupService
	SponsoredCardIssueReason    *SponsoredCardIssueReasonService
	SponsoredCardOwner          *SponsoredCardOwnerService
	StaticSegment               *StaticSegmentService
	Store                       *StoreService
	StoreCluster                *StoreClusterService
	StoreFormat                 *StoreFormatService
	StoreProperty               *StorePropertyService
	TerritorialDivision         *TerritorialDivisionService
}

// NewClient creates a new RS Loyalty API client
func NewClient(config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	c := &Client{
		config:     config,
		httpClient: httpClient,
	}

	// Initialize services
	c.Accounts = &AccountsService{client: c}
	c.Countries = &CountriesService{client: c}
	c.Currencies = &CurrenciesService{client: c}
	c.Customers = &CustomersService{client: c}
	c.LoyaltyCards = &LoyaltyCardsService{client: c}
	c.CustomerCards = &CustomerCardsService{client: c}
	c.CustomerProperty = &CustomerPropertyService{client: c}
	c.CustomerSegment = &CustomerSegmentService{client: c}
	c.ExternalGiftCardSeries = &ExternalGiftCardSeriesService{client: c}
	c.ExternalLoyaltyCardSeries = &ExternalLoyaltyCardSeriesService{client: c}
	c.ExternalSponsoredCardSeries = &ExternalSponsoredCardSeriesService{client: c}
	c.GiftCards = &GiftCardService{client: c}
	c.GiftCardBlockReason = &GiftCardBlockReasonService{client: c}
	c.GiftCardGroup = &GiftCardGroupService{client: c}
	c.GiftCardIssueReason = &GiftCardIssueReasonService{client: c}
	c.Items = &ItemService{client: c}
	c.ItemCategories = &ItemCategoryService{client: c}
	c.ItemGroups = &ItemGroupService{client: c}
	c.ItemProperties = &ItemPropertyService{client: c}
	c.LoyaltyCardBlockReason = &LoyaltyCardBlockReasonService{client: c}
	c.LoyaltyCardGroup = &LoyaltyCardGroupService{client: c}
	c.LoyaltyCardIssueReason = &LoyaltyCardIssueReasonService{client: c}
	c.ManualAccrualReason = &ManualAccrualReasonService{client: c}
	c.OpeningHours = &OpeningHoursService{client: c}
	c.Pos = &PosService{client: c}
	c.PosType = &PosTypeService{client: c}
	c.SegmentGroup = &SegmentGroupService{client: c}
	c.SponsoredCards = &SponsoredCardService{client: c}
	c.SponsoredCardBlockReason = &SponsoredCardBlockReasonService{client: c}
	c.SponsoredCardGroup = &SponsoredCardGroupService{client: c}
	c.SponsoredCardIssueReason = &SponsoredCardIssueReasonService{client: c}
	c.SponsoredCardOwner = &SponsoredCardOwnerService{client: c}
	c.StaticSegment = &StaticSegmentService{client: c}
	c.Store = &StoreService{client: c}
	c.StoreCluster = &StoreClusterService{client: c}
	c.StoreFormat = &StoreFormatService{client: c}
	c.StoreProperty = &StorePropertyService{client: c}
	c.TerritorialDivision = &TerritorialDivisionService{client: c}

	return c
}

// doRequest performs HTTP request
func (c *Client) doRequest(ctx context.Context, method, endpoint string, body interface{}, headers *models.RequestHeaders) (*http.Response, error) {
	url := c.config.BaseURL + endpoint

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom headers from config
	for key, value := range c.config.Headers {
		req.Header.Set(key, value)
	}

	// Set request-specific headers
	if headers != nil {
		if headers.CommandID != nil {
			req.Header.Set("command-id", *headers.CommandID)
		}
		if headers.OperationDate != nil {
			req.Header.Set("operation-date", *headers.OperationDate)
		}
		if headers.UserID != nil {
			req.Header.Set("user-id", *headers.UserID)
		}
		if headers.InteractionChannel != nil {
			req.Header.Set("interaction-channel", *headers.InteractionChannel)
		}
	}

	return c.httpClient.Do(req)
}

// doRequestWithResponse performs HTTP request and decodes response
func (c *Client) doRequestWithResponse(ctx context.Context, method, endpoint string, body interface{}, headers *models.RequestHeaders, response interface{}) error {
	resp, err := c.doRequest(ctx, method, endpoint, body, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &models.APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("API request failed with status %d", resp.StatusCode),
			Details:    string(bodyBytes),
		}
	}

	if response != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// doCommand performs command request (POST without response body)
func (c *Client) doCommand(ctx context.Context, endpoint string, body interface{}, headers *models.RequestHeaders) error {
	resp, err := c.doRequest(ctx, http.MethodPost, endpoint, body, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &models.APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("API command failed with status %d", resp.StatusCode),
			Details:    string(bodyBytes),
		}
	}

	return nil
}

// SetHeader sets a header that will be sent with all requests
func (c *Client) SetHeader(key, value string) {
	if c.config.Headers == nil {
		c.config.Headers = make(map[string]string)
	}
	c.config.Headers[key] = value
}

// GetBaseURL returns the base URL of the API
func (c *Client) GetBaseURL() string {
	return c.config.BaseURL
}
