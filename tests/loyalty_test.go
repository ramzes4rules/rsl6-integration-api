package tests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ramzes4rules/rsl6-integration-api/models"
)

func TestLoyaltyCards_GetByID(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	cardID := uuid.MustParse("44444444-4444-4444-4444-444444444444")

	card, err := ctx.Client.LoyaltyCards.GetByID(context.Background(), cardID)
	require.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, cardID, card.ID)
	assert.Equal(t, "1234567890", card.Number)
	assert.Equal(t, models.LoyaltyCardStatusActive, card.Status)
}

func TestLoyaltyCards_GetByNumber(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	card, err := ctx.Client.LoyaltyCards.GetByNumber(context.Background(), "1234567890")
	require.NoError(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, "1234567890", card.Number)
}

func TestLoyaltyCards_GetList(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	list, err := ctx.Client.LoyaltyCards.GetList(context.Background(), &models.GetListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, list)
	assert.GreaterOrEqual(t, list.Total, int64(1))
}

func TestLoyaltyCards_GetBalanceByID(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	cardID := uuid.MustParse("44444444-4444-4444-4444-444444444444")

	balance, err := ctx.Client.LoyaltyCards.GetBalanceByID(context.Background(), cardID)
	require.NoError(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, cardID, balance.ID)
	assert.NotEmpty(t, balance.Balances)
}

func TestLoyaltyCards_BlockAndUnblock(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	cardID := uuid.MustParse("44444444-4444-4444-4444-444444444444")

	// Block card
	err := ctx.Client.LoyaltyCards.Block(context.Background(), cardID, nil)
	require.NoError(t, err)

	card, err := ctx.Client.LoyaltyCards.GetByID(context.Background(), cardID)
	require.NoError(t, err)
	assert.Equal(t, models.LoyaltyCardStatusBlocked, card.Status)

	// Unblock card
	err = ctx.Client.LoyaltyCards.Unblock(context.Background(), cardID, nil)
	require.NoError(t, err)

	card, err = ctx.Client.LoyaltyCards.GetByID(context.Background(), cardID)
	require.NoError(t, err)
	assert.Equal(t, models.LoyaltyCardStatusActive, card.Status)
}

func TestAccounts_AccrualToCustomer(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	currencyID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	// Get initial balance
	initialBalance, err := ctx.Client.Customers.GetBalanceByID(context.Background(), customerID)
	require.NoError(t, err)

	var initialValue float64
	for _, b := range initialBalance.Balances.Balances {
		if b.CurrencyID == currencyID {
			initialValue = b.Value
			break
		}
	}

	// Accrue points
	accrualAmount := 100.0
	req := &models.AccrualToCustomerRequest{
		BaseCommand: models.BaseCommand{ID: uuid.New()},
		CustomerID:  customerID,
		CurrencyID:  currencyID,
		Amount:      accrualAmount,
	}
	err = ctx.Client.Accounts.AccrualToCustomer(context.Background(), req, nil)
	require.NoError(t, err)

	// Verify balance increased
	newBalance, err := ctx.Client.Customers.GetBalanceByID(context.Background(), customerID)
	require.NoError(t, err)

	var newValue float64
	for _, b := range newBalance.Balances.Balances {
		if b.CurrencyID == currencyID {
			newValue = b.Value
			break
		}
	}

	assert.Equal(t, initialValue+accrualAmount, newValue)
}

func TestAccounts_SubtractFromCustomer(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	currencyID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	// Get initial balance
	initialBalance, err := ctx.Client.Customers.GetBalanceByID(context.Background(), customerID)
	require.NoError(t, err)

	var initialValue float64
	for _, b := range initialBalance.Balances.Balances {
		if b.CurrencyID == currencyID {
			initialValue = b.Value
			break
		}
	}

	// Subtract points (only if there's enough balance)
	subtractAmount := 50.0
	if initialValue >= subtractAmount {
		req := &models.SubtractFromCustomerRequest{
			BaseCommand: models.BaseCommand{ID: uuid.New()},
			CustomerID:  customerID,
			CurrencyID:  currencyID,
			Amount:      subtractAmount,
		}
		err = ctx.Client.Accounts.SubtractFromCustomer(context.Background(), req, nil)
		require.NoError(t, err)

		// Verify balance decreased
		newBalance, err := ctx.Client.Customers.GetBalanceByID(context.Background(), customerID)
		require.NoError(t, err)

		var newValue float64
		for _, b := range newBalance.Balances.Balances {
			if b.CurrencyID == currencyID {
				newValue = b.Value
				break
			}
		}

		assert.Equal(t, initialValue-subtractAmount, newValue)
	}
}

func TestAccounts_AccrualToLoyaltyCard(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	cardID := uuid.MustParse("44444444-4444-4444-4444-444444444444")
	currencyID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	// Get initial balance
	initialBalance, err := ctx.Client.LoyaltyCards.GetBalanceByID(context.Background(), cardID)
	require.NoError(t, err)

	var initialValue float64
	for _, b := range initialBalance.Balances {
		if b.CurrencyID == currencyID {
			initialValue = b.Value
			break
		}
	}

	// Accrue points
	accrualAmount := 75.0
	req := &models.AccrualToLoyaltyCardRequest{
		BaseCommand:   models.BaseCommand{ID: uuid.New()},
		LoyaltyCardID: cardID,
		CurrencyID:    currencyID,
		Amount:        accrualAmount,
	}
	err = ctx.Client.Accounts.AccrualToLoyaltyCard(context.Background(), req, nil)
	require.NoError(t, err)

	// Verify balance increased
	newBalance, err := ctx.Client.LoyaltyCards.GetBalanceByID(context.Background(), cardID)
	require.NoError(t, err)

	var newValue float64
	for _, b := range newBalance.Balances {
		if b.CurrencyID == currencyID {
			newValue = b.Value
			break
		}
	}

	assert.Equal(t, initialValue+accrualAmount, newValue)
}

func TestAccounts_GetTransactions(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	req := &models.GetTransactionsRequest{}
	transactions, err := ctx.Client.Accounts.GetTransactions(context.Background(), req)
	require.NoError(t, err)
	assert.NotNil(t, transactions)
}
