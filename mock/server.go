// Package mock provides mock server for RS Loyalty API v2 testing
package mock

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// Server represents mock API server
type Server struct {
	mux *http.ServeMux

	mu           sync.RWMutex
	customers    map[uuid.UUID]*models.CustomerDto
	countries    map[uuid.UUID]*models.CountryDto
	currencies   map[uuid.UUID]*models.CurrencyDto
	loyaltyCards map[uuid.UUID]*models.LoyaltyCardDto
	balances     map[uuid.UUID][]models.BalanceDto
	transactions []models.TransactionDetailDto
}

// NewServer creates a new mock server
func NewServer() *Server {
	s := &Server{
		mux:          http.NewServeMux(),
		customers:    make(map[uuid.UUID]*models.CustomerDto),
		countries:    make(map[uuid.UUID]*models.CountryDto),
		currencies:   make(map[uuid.UUID]*models.CurrencyDto),
		loyaltyCards: make(map[uuid.UUID]*models.LoyaltyCardDto),
		balances:     make(map[uuid.UUID][]models.BalanceDto),
		transactions: make([]models.TransactionDetailDto, 0),
	}

	s.registerRoutes()
	s.seedTestData()

	return s
}

// ServeHTTP implements http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// GetMux returns the HTTP multiplexer
func (s *Server) GetMux() *http.ServeMux {
	return s.mux
}

func (s *Server) registerRoutes() {
	// Customers endpoints
	s.mux.HandleFunc("/api/v2/customers/get_by_id", s.handleGetCustomerByID)
	s.mux.HandleFunc("/api/v2/customers/get_by_communication_value", s.handleGetCustomerByCommunicationValue)
	s.mux.HandleFunc("/api/v2/customers/get_list", s.handleGetCustomerList)
	s.mux.HandleFunc("/api/v2/customers/get_balance_by_id", s.handleGetCustomerBalanceByID)
	s.mux.HandleFunc("/api/v2/customers/get_transactions_by_id", s.handleGetCustomerTransactionsByID)
	s.mux.HandleFunc("/api/v2/customers/rename", s.handleRenameCustomer)
	s.mux.HandleFunc("/api/v2/customers/set_communication_value", s.handleSetCommunicationValue)
	s.mux.HandleFunc("/api/v2/customers/set_birthday", s.handleSetBirthday)
	s.mux.HandleFunc("/api/v2/customers/set_address", s.handleSetAddress)
	s.mux.HandleFunc("/api/v2/customers/allow_subscription", s.handleAllowSubscription)
	s.mux.HandleFunc("/api/v2/customers/disallow_subscription", s.handleDisallowSubscription)
	s.mux.HandleFunc("/api/v2/customers/remove_personal_data", s.handleRemovePersonalData)

	// Countries endpoints
	s.mux.HandleFunc("/api/v2/countries/get_by_id", s.handleGetCountryByID)
	s.mux.HandleFunc("/api/v2/countries/get_list", s.handleGetCountryList)
	s.mux.HandleFunc("/api/v2/countries/create", s.handleCreateCountry)
	s.mux.HandleFunc("/api/v2/countries/set_code", s.handleSetCountryCode)
	s.mux.HandleFunc("/api/v2/countries/delete", s.handleDeleteCountry)
	s.mux.HandleFunc("/api/v2/countries/restore", s.handleRestoreCountry)
	s.mux.HandleFunc("/api/v2/countries/batch", s.handleCountriesBatch)

	// Currencies endpoints
	s.mux.HandleFunc("/api/v2/currencies/get_by_id", s.handleGetCurrencyByID)
	s.mux.HandleFunc("/api/v2/currencies/get_list", s.handleGetCurrencyList)
	s.mux.HandleFunc("/api/v2/currencies/create", s.handleCreateCurrency)
	s.mux.HandleFunc("/api/v2/currencies/rename", s.handleRenameCurrency)
	s.mux.HandleFunc("/api/v2/currencies/set_description", s.handleSetCurrencyDescription)
	s.mux.HandleFunc("/api/v2/currencies/set_public_name", s.handleSetCurrencyPublicName)
	s.mux.HandleFunc("/api/v2/currencies/set_rate", s.handleSetCurrencyRate)
	s.mux.HandleFunc("/api/v2/currencies/set_calculate_round_rule", s.handleSetCurrencyCalculateRoundRule)
	s.mux.HandleFunc("/api/v2/currencies/set_caption", s.handleSetCurrencyCaption)
	s.mux.HandleFunc("/api/v2/currencies/activate", s.handleActivateCurrency)
	s.mux.HandleFunc("/api/v2/currencies/deactivate", s.handleDeactivateCurrency)
	s.mux.HandleFunc("/api/v2/currencies/delete", s.handleDeleteCurrency)
	s.mux.HandleFunc("/api/v2/currencies/restore", s.handleRestoreCurrency)
	s.mux.HandleFunc("/api/v2/currencies/batch", s.handleCurrenciesBatch)

	// Loyalty cards endpoints
	s.mux.HandleFunc("/api/v2/loyalty_cards/get_by_id", s.handleGetLoyaltyCardByID)
	s.mux.HandleFunc("/api/v2/loyalty_cards/get_by_number", s.handleGetLoyaltyCardByNumber)
	s.mux.HandleFunc("/api/v2/loyalty_cards/get_list", s.handleGetLoyaltyCardList)
	s.mux.HandleFunc("/api/v2/loyalty_cards/get_balance_by_id", s.handleGetLoyaltyCardBalanceByID)
	s.mux.HandleFunc("/api/v2/loyalty_cards/get_transactions_by_id", s.handleGetLoyaltyCardTransactionsByID)
	s.mux.HandleFunc("/api/v2/loyalty_cards/activate", s.handleActivateLoyaltyCard)
	s.mux.HandleFunc("/api/v2/loyalty_cards/block", s.handleBlockLoyaltyCard)
	s.mux.HandleFunc("/api/v2/loyalty_cards/unblock", s.handleUnblockLoyaltyCard)
	s.mux.HandleFunc("/api/v2/loyalty_cards/delete", s.handleDeleteLoyaltyCard)
	s.mux.HandleFunc("/api/v2/loyalty_cards/restore", s.handleRestoreLoyaltyCard)

	// Accounts endpoints
	s.mux.HandleFunc("/api/v2/accounts/accrual_to_customer", s.handleAccrualToCustomer)
	s.mux.HandleFunc("/api/v2/accounts/accrual_to_loyalty_card", s.handleAccrualToLoyaltyCard)
	s.mux.HandleFunc("/api/v2/accounts/subtract_from_customer", s.handleSubtractFromCustomer)
	s.mux.HandleFunc("/api/v2/accounts/subtract_from_loyalty_card", s.handleSubtractFromLoyaltyCard)
	s.mux.HandleFunc("/api/v2/accounts/get_transactions", s.handleGetTransactions)
	s.mux.HandleFunc("/api/v2/accounts/batch", s.handleAccountsBatch)
}

func (s *Server) seedTestData() {
	// Seed test customer
	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	firstName := "Иван"
	lastName := "Петров"
	email := "ivan@example.com"
	phone := "+79001234567"
	birthday := "1990-05-15"

	s.customers[customerID] = &models.CustomerDto{
		ID:        customerID,
		FirstName: &firstName,
		LastName:  &lastName,
		Birthday:  &birthday,
		Gender:    models.GenderMale,
		CommunicationValues: []models.CustomerCommunicationValueDto{
			{ValueType: models.CommunicationEmail, Value: email, Confirmed: true},
			{ValueType: models.CommunicationPhone, Value: phone, Confirmed: true},
		},
		Subscriptions:                    []models.SubscriptionTypes{models.SubscriptionEmailMailing},
		AllowBySendingVirtualCopyReceipt: true,
		Properties:                       []models.CustomerPropertyDto{},
		Children:                         []models.CustomerChildDto{},
		StaticSegments:                   []uuid.UUID{},
	}

	// Seed test country
	countryID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	s.countries[countryID] = &models.CountryDto{
		ID:        countryID,
		Code:      "RU",
		IsDeleted: false,
	}

	// Seed test currency
	currencyID := uuid.MustParse("33333333-3333-3333-3333-333333333333")
	description := "Бонусные баллы"
	publicName := "Баллы"
	s.currencies[currencyID] = &models.CurrencyDto{
		ID:          currencyID,
		Name:        "BONUS",
		Description: &description,
		PublicName:  &publicName,
		Rate:        1.0,
		IsActive:    true,
		IsDeleted:   false,
	}

	// Seed test loyalty card
	loyaltyCardID := uuid.MustParse("44444444-4444-4444-4444-444444444444")
	seriesID := uuid.MustParse("55555555-5555-5555-5555-555555555555")
	issueDate := time.Now().AddDate(0, -6, 0)
	s.loyaltyCards[loyaltyCardID] = &models.LoyaltyCardDto{
		ID:         loyaltyCardID,
		Number:     "1234567890",
		SeriesID:   seriesID,
		Status:     models.LoyaltyCardStatusActive,
		CustomerID: &customerID,
		IssueDate:  &issueDate,
		IsDeleted:  false,
	}

	// Seed test balances
	s.balances[customerID] = []models.BalanceDto{
		{CurrencyID: currencyID, Value: 1000.0},
	}
	s.balances[loyaltyCardID] = []models.BalanceDto{
		{CurrencyID: currencyID, Value: 500.0},
	}

	// Seed test transactions
	s.transactions = []models.TransactionDetailDto{
		{
			TransactionDto: models.TransactionDto{
				TransactionID:   uuid.New(),
				OperationDate:   time.Now().AddDate(0, 0, -7),
				TransactionType: models.TransactionAccrual,
				Amount:          100.0,
				CurrencyID:      currencyID,
			},
			AccountID: customerID,
		},
		{
			TransactionDto: models.TransactionDto{
				TransactionID:   uuid.New(),
				OperationDate:   time.Now().AddDate(0, 0, -3),
				TransactionType: models.TransactionSubtract,
				Amount:          50.0,
				CurrencyID:      currencyID,
			},
			AccountID: customerID,
		},
	}
}

// Helper functions
func (s *Server) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (s *Server) writeError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (s *Server) readJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// AddCustomer adds a customer to the mock data
func (s *Server) AddCustomer(customer *models.CustomerDto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.customers[customer.ID] = customer
}

// AddCountry adds a country to the mock data
func (s *Server) AddCountry(country *models.CountryDto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.countries[country.ID] = country
}

// AddCurrency adds a currency to the mock data
func (s *Server) AddCurrency(currency *models.CurrencyDto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.currencies[currency.ID] = currency
}

// AddLoyaltyCard adds a loyalty card to the mock data
func (s *Server) AddLoyaltyCard(card *models.LoyaltyCardDto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.loyaltyCards[card.ID] = card
}

// SetBalance sets balance for an entity
func (s *Server) SetBalance(entityID uuid.UUID, balances []models.BalanceDto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.balances[entityID] = balances
}

// AddTransaction adds a transaction
func (s *Server) AddTransaction(tx models.TransactionDetailDto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions = append(s.transactions, tx)
}

// Reset clears all mock data and re-seeds test data
func (s *Server) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.customers = make(map[uuid.UUID]*models.CustomerDto)
	s.countries = make(map[uuid.UUID]*models.CountryDto)
	s.currencies = make(map[uuid.UUID]*models.CurrencyDto)
	s.loyaltyCards = make(map[uuid.UUID]*models.LoyaltyCardDto)
	s.balances = make(map[uuid.UUID][]models.BalanceDto)
	s.transactions = make([]models.TransactionDetailDto, 0)
	s.seedTestData()
}

// GetCustomer returns customer by ID (for testing)
func (s *Server) GetCustomer(id uuid.UUID) *models.CustomerDto {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.customers[id]
}

// GetCountry returns country by ID (for testing)
func (s *Server) GetCountry(id uuid.UUID) *models.CountryDto {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.countries[id]
}

// GetCurrency returns currency by ID (for testing)
func (s *Server) GetCurrency(id uuid.UUID) *models.CurrencyDto {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.currencies[id]
}

// GetLoyaltyCard returns loyalty card by ID (for testing)
func (s *Server) GetLoyaltyCard(id uuid.UUID) *models.LoyaltyCardDto {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.loyaltyCards[id]
}
