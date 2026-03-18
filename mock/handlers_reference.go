package mock

import (
	"net/http"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

// Country handlers

func (s *Server) handleGetCountryByID(w http.ResponseWriter, r *http.Request) {
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
	country, ok := s.countries[req.ID]
	s.mu.RUnlock()

	if !ok {
		s.writeError(w, http.StatusNotFound, "Country not found")
		return
	}

	s.writeJSON(w, country)
}

func (s *Server) handleGetCountryList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetListRequest
	s.readJSON(r, &req)

	s.mu.RLock()
	defer s.mu.RUnlock()

	countries := make([]models.CountryDto, 0, len(s.countries))
	for _, c := range s.countries {
		countries = append(countries, *c)
	}

	response := models.CountryListDto{
		Total:  int64(len(countries)),
		Values: countries,
	}

	s.writeJSON(w, response)
}

func (s *Server) handleCreateCountry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.CreateCountryRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.countries[req.ID] = &models.CountryDto{
		ID:        req.ID,
		Code:      req.Code,
		IsDeleted: false,
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCountryCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCountryCodeRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	country, ok := s.countries[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Country not found")
		return
	}

	country.Code = req.Code
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDeleteCountry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.DeleteCountryRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	country, ok := s.countries[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Country not found")
		return
	}

	country.IsDeleted = true
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleRestoreCountry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.RestoreCountryRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	country, ok := s.countries[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Country not found")
		return
	}

	country.IsDeleted = false
	w.WriteHeader(http.StatusOK)
}

// Currency handlers

func (s *Server) handleGetCurrencyByID(w http.ResponseWriter, r *http.Request) {
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
	currency, ok := s.currencies[req.ID]
	s.mu.RUnlock()

	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	s.writeJSON(w, currency)
}

func (s *Server) handleGetCurrencyList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.GetListRequest
	s.readJSON(r, &req)

	s.mu.RLock()
	defer s.mu.RUnlock()

	currencies := make([]models.CurrencyDto, 0, len(s.currencies))
	for _, c := range s.currencies {
		currencies = append(currencies, *c)
	}

	response := models.CurrencyListDto{
		Total:  int64(len(currencies)),
		Values: currencies,
	}

	s.writeJSON(w, response)
}

func (s *Server) handleCreateCurrency(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.CreateCurrencyRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.currencies[req.ID] = &models.CurrencyDto{
		ID:        req.ID,
		Name:      req.Name,
		Rate:      1.0,
		IsActive:  false,
		IsDeleted: false,
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleRenameCurrency(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.RenameCurrencyRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.Name = req.Name
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCurrencyDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCurrencyDescriptionRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.Description = &req.Description
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCurrencyRate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCurrencyRateRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.Rate = req.Rate
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleActivateCurrency(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.ActivateCurrencyRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.IsActive = true
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDeactivateCurrency(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.DeactivateCurrencyRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.IsActive = false
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDeleteCurrency(w http.ResponseWriter, r *http.Request) {
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

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.IsDeleted = true
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleRestoreCurrency(w http.ResponseWriter, r *http.Request) {
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

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.IsDeleted = false
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCurrencyPublicName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCurrencyPublicNameRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	currency, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	currency.PublicName = &req.PublicName
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCurrencyCalculateRoundRule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCurrencyCalculateRoundRuleRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	// Note: CalculateRoundRule is not stored in CurrencyDto in this mock
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleSetCurrencyCaption(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SetCurrencyCaptionRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.currencies[req.ID]
	if !ok {
		s.writeError(w, http.StatusNotFound, "Currency not found")
		return
	}

	// Note: Caption is not stored in CurrencyDto in this mock
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleCurrenciesBatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BatchRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// In mock, we just accept the batch request without processing individual commands
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleCountriesBatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.BatchRequest
	if err := s.readJSON(r, &req); err != nil {
		s.writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// In mock, we just accept the batch request without processing individual commands
	w.WriteHeader(http.StatusOK)
}
