package mock

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rsl6/loyalty-client/models"
)

// Loyalty card handlers

func (s *Server) handleGetLoyaltyCardByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetByIdRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.RLock()
	card, ok := s.loyaltyCards[req.ID]
	s.mu.RUnlock()

	if !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	s.writeJSON(w, card)
}

func (s *Server) handleGetLoyaltyCardByNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		Number string `json:"number"`
	}
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, card := range s.loyaltyCards {
		if card.Number == req.Number {
			s.writeJSON(w, card)
			return
		}
	}

	s.writeError(w, http.StatusNotFound, "Loyalty card not found")
}

func (s *Server) handleGetLoyaltyCardList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetListRequest
	s.readJSON(r, &req)

	s.mu.RLock()
	defer s.mu.RUnlock()

	cards := make([]models.LoyaltyCardDto, 0, len(s.loyaltyCards))
	for _, c := range s.loyaltyCards {
		cards = append(cards, *c)
	}

	response := models.LoyaltyCardListDto{
		Total:  int64(len(cards)),
		Values: cards,
	}

	s.writeJSON(w, response)
}

func (s *Server) handleGetLoyaltyCardBalanceByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetByIdRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.loyaltyCards[req.ID]; !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	balances := s.balances[req.ID]
	response := models.BalanceListDto{
		ID:       req.ID,
		Balances: balances,
	}

	s.writeJSON(w, response)
}

func (s *Server) handleGetLoyaltyCardTransactionsByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetByIdRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	transactions := make([]models.TransactionDetailDto, 0)
	for _, tx := range s.transactions {
		if tx.AccountID == req.ID {
			transactions = append(transactions, tx)
		}
	}

	response := models.TransactionListDto{
		Transactions: transactions,
	}

	s.writeJSON(w, response)
}

func (s *Server) handleActivateLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BaseCommand
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.loyaltyCards[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	card.Status = models.LoyaltyCardStatusActive
	now := time.Now()
	card.ActivateDate = &now
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleBlockLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BaseCommand
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.loyaltyCards[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	card.Status = models.LoyaltyCardStatusBlocked
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleUnblockLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BaseCommand
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.loyaltyCards[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	card.Status = models.LoyaltyCardStatusActive
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDeleteLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BaseCommand
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.loyaltyCards[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	card.IsDeleted = true
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleRestoreLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BaseCommand
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.loyaltyCards[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	card.IsDeleted = false
	w.WriteHeader(http.StatusOK)
}

// Account handlers

func (s *Server) handleAccrualToCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.AccrualToCustomerRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.customers[req.CustomerID]; !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Update balance
	balances := s.balances[req.CustomerID]
	found := false
	for i, b := range balances {
		if b.CurrencyID == req.CurrencyID {
			balances[i].Value += req.Amount
			found = true
			break
		}
	}
	if !found {
		balances = append(balances, models.BalanceDto{
			CurrencyID: req.CurrencyID,
			Value:      req.Amount,
		})
	}
	s.balances[req.CustomerID] = balances

	// Add transaction
	s.transactions = append(s.transactions, models.TransactionDetailDto{
		TransactionDto: models.TransactionDto{
			TransactionID:   uuid.New(),
			OperationDate:   time.Now(),
			TransactionType: models.TransactionAccrual,
			Amount:          req.Amount,
			CurrencyID:      req.CurrencyID,
			ExpirationDate:  req.ExpirationDate,
			ActivationDate:  req.ActivationDate,
		},
		AccountID: req.CustomerID,
	})

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleAccrualToLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.AccrualToLoyaltyCardRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.loyaltyCards[req.LoyaltyCardID]; !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	// Update balance
	balances := s.balances[req.LoyaltyCardID]
	found := false
	for i, b := range balances {
		if b.CurrencyID == req.CurrencyID {
			balances[i].Value += req.Amount
			found = true
			break
		}
	}
	if !found {
		balances = append(balances, models.BalanceDto{
			CurrencyID: req.CurrencyID,
			Value:      req.Amount,
		})
	}
	s.balances[req.LoyaltyCardID] = balances

	// Add transaction
	s.transactions = append(s.transactions, models.TransactionDetailDto{
		TransactionDto: models.TransactionDto{
			TransactionID:   uuid.New(),
			OperationDate:   time.Now(),
			TransactionType: models.TransactionAccrual,
			Amount:          req.Amount,
			CurrencyID:      req.CurrencyID,
			ExpirationDate:  req.ExpirationDate,
			ActivationDate:  req.ActivationDate,
		},
		AccountID: req.LoyaltyCardID,
	})

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSubtractFromCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SubtractFromCustomerRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.customers[req.CustomerID]; !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Update balance
	balances := s.balances[req.CustomerID]
	for i, b := range balances {
		if b.CurrencyID == req.CurrencyID {
			if balances[i].Value < req.Amount {
				s.writeError(w, http.StatusBadRequest, "Insufficient balance")
				return
			}
			balances[i].Value -= req.Amount
			break
		}
	}
	s.balances[req.CustomerID] = balances

	// Add transaction
	s.transactions = append(s.transactions, models.TransactionDetailDto{
		TransactionDto: models.TransactionDto{
			TransactionID:   uuid.New(),
			OperationDate:   time.Now(),
			TransactionType: models.TransactionSubtract,
			Amount:          req.Amount,
			CurrencyID:      req.CurrencyID,
		},
		AccountID: req.CustomerID,
	})

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSubtractFromLoyaltyCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SubtractFromLoyaltyCardRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.loyaltyCards[req.LoyaltyCardID]; !ok {
		s.writeError(w, http.StatusNotFound, "Loyalty card not found")
		return
	}

	// Update balance
	balances := s.balances[req.LoyaltyCardID]
	for i, b := range balances {
		if b.CurrencyID == req.CurrencyID {
			if balances[i].Value < req.Amount {
				s.writeError(w, http.StatusBadRequest, "Insufficient balance")
				return
			}
			balances[i].Value -= req.Amount
			break
		}
	}
	s.balances[req.LoyaltyCardID] = balances

	// Add transaction
	s.transactions = append(s.transactions, models.TransactionDetailDto{
		TransactionDto: models.TransactionDto{
			TransactionID:   uuid.New(),
			OperationDate:   time.Now(),
			TransactionType: models.TransactionSubtract,
			Amount:          req.Amount,
			CurrencyID:      req.CurrencyID,
		},
		AccountID: req.LoyaltyCardID,
	})

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleGetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetTransactionsRequest
	s.readJSON(r, &req)

	s.mu.RLock()
	defer s.mu.RUnlock()

	transactions := make([]models.TransactionDetailDto, 0)
	for _, tx := range s.transactions {
		// Filter by date range if specified
		if req.FromDate != nil && tx.OperationDate.Before(*req.FromDate) {
			continue
		}
		if req.ToDate != nil && tx.OperationDate.After(*req.ToDate) {
			continue
		}
		transactions = append(transactions, tx)
	}

	// Apply pagination
	skip := 0
	if req.Skip != nil {
		skip = *req.Skip
	}
	if skip >= len(transactions) {
		transactions = []models.TransactionDetailDto{}
	} else {
		take := len(transactions) - skip
		if req.Take != nil && *req.Take < take {
			take = *req.Take
		}
		transactions = transactions[skip : skip+take]
	}

	response := models.TransactionListDto{
		Transactions: transactions,
	}

	s.writeJSON(w, response)
}
