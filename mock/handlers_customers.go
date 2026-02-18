package mock

import (
	"net/http"

	"github.com/rsl6/loyalty-client/models"
)

// Customer handlers

func (s *Server) handleGetCustomerByID(w http.ResponseWriter, r *http.Request) {
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
	customer, ok := s.customers[req.ID]
	s.mu.RUnlock()

	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	s.writeJSON(w, customer)
}

func (s *Server) handleGetCustomerByCommunicationValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetByCommunicationValueRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, customer := range s.customers {
		for _, cv := range customer.CommunicationValues {
			if cv.ValueType == req.CommunicationValueType && cv.Value == req.Value {
				s.writeJSON(w, customer)
				return
			}
		}
	}

	s.writeError(w, http.StatusNotFound, "Customer not found")
}

func (s *Server) handleGetCustomerList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetListRequest
	s.readJSON(r, &req)

	s.mu.RLock()
	defer s.mu.RUnlock()

	customers := make([]models.CustomerDto, 0, len(s.customers))
	for _, c := range s.customers {
		customers = append(customers, *c)
	}

	// Apply pagination
	skip := 0
	if req.Skip != nil {
		skip = *req.Skip
	}
	take := len(customers)
	if req.Take != nil && *req.Take < take {
		take = *req.Take
	}

	if skip >= len(customers) {
		customers = []models.CustomerDto{}
	} else {
		end := skip + take
		if end > len(customers) {
			end = len(customers)
		}
		customers = customers[skip:end]
	}

	response := models.CustomerListDto{
		Total:  int64(len(s.customers)),
		Values: customers,
	}

	s.writeJSON(w, response)
}

func (s *Server) handleGetCustomerBalanceByID(w http.ResponseWriter, r *http.Request) {
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

	if _, ok := s.customers[req.ID]; !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	balances := s.balances[req.ID]
	response := models.CustomerBalanceDto{
		Balances: &models.BalanceListDto{
			ID:       req.ID,
			Balances: balances,
		},
	}

	s.writeJSON(w, response)
}

func (s *Server) handleGetCustomerTransactionsByID(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) handleRenameCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.RenameCustomerRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	if req.FirstName != nil {
		customer.FirstName = req.FirstName
	}
	if req.SecondName != nil {
		customer.SecondName = req.SecondName
	}
	if req.LastName != nil {
		customer.LastName = req.LastName
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCommunicationValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCustomerCommunicationValueRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Update or add communication value
	found := false
	for i, cv := range customer.CommunicationValues {
		if cv.ValueType == req.ValueType {
			customer.CommunicationValues[i].Value = req.Value
			customer.CommunicationValues[i].Confirmed = false
			found = true
			break
		}
	}

	if !found {
		customer.CommunicationValues = append(customer.CommunicationValues, models.CustomerCommunicationValueDto{
			ValueType: req.ValueType,
			Value:     req.Value,
			Confirmed: false,
		})
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetBirthday(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCustomerBirthdayRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	customer.Birthday = &req.Birthday
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCustomerAddressRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	customer.Address = &req.Address
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleAllowSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.AllowSubscriptionRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Check if subscription already exists
	for _, sub := range customer.Subscriptions {
		if sub == req.SubscriptionType {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	customer.Subscriptions = append(customer.Subscriptions, req.SubscriptionType)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDisallowSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.DisallowSubscriptionRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Remove subscription
	subs := make([]models.SubscriptionTypes, 0)
	for _, sub := range customer.Subscriptions {
		if sub != req.SubscriptionType {
			subs = append(subs, sub)
		}
	}
	customer.Subscriptions = subs

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleRemovePersonalData(w http.ResponseWriter, r *http.Request) {
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

	customer, ok := s.customers[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Customer not found")
		return
	}

	// Clear personal data
	customer.FirstName = nil
	customer.SecondName = nil
	customer.LastName = nil
	customer.Address = nil
	customer.Birthday = nil
	customer.CommunicationValues = []models.CustomerCommunicationValueDto{}

	w.WriteHeader(http.StatusOK)
}
