package tests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rsl6/loyalty-client/models"
)

func TestCountries_GetByID(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	countryID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	country, err := ctx.Client.Countries.GetByID(context.Background(), countryID)
	require.NoError(t, err)
	assert.NotNil(t, country)
	assert.Equal(t, countryID, country.ID)
	assert.Equal(t, "RU", country.Code)
	assert.False(t, country.IsDeleted)
}

func TestCountries_GetList(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	list, err := ctx.Client.Countries.GetList(context.Background(), &models.GetListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, list)
	assert.GreaterOrEqual(t, list.Total, int64(1))
}

func TestCountries_CreateAndSetCode(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	newCountryID := uuid.New()

	// Create country
	createReq := &models.CreateCountryRequest{
		BaseCommand: models.BaseCommand{ID: newCountryID},
		Code:        "US",
	}
	err := ctx.Client.Countries.Create(context.Background(), createReq, nil)
	require.NoError(t, err)

	// Verify creation
	country, err := ctx.Client.Countries.GetByID(context.Background(), newCountryID)
	require.NoError(t, err)
	assert.Equal(t, "US", country.Code)

	// Set new code
	setCodeReq := &models.SetCountryCodeRequest{
		BaseCommand: models.BaseCommand{ID: newCountryID},
		Code:        "USA",
	}
	err = ctx.Client.Countries.SetCode(context.Background(), setCodeReq, nil)
	require.NoError(t, err)

	// Verify code change
	country, err = ctx.Client.Countries.GetByID(context.Background(), newCountryID)
	require.NoError(t, err)
	assert.Equal(t, "USA", country.Code)
}

func TestCountries_DeleteAndRestore(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	newCountryID := uuid.New()

	// Create country first
	createReq := &models.CreateCountryRequest{
		BaseCommand: models.BaseCommand{ID: newCountryID},
		Code:        "DE",
	}
	err := ctx.Client.Countries.Create(context.Background(), createReq, nil)
	require.NoError(t, err)

	// Delete country
	deleteReq := &models.DeleteCountryRequest{
		BaseCommand: models.BaseCommand{ID: newCountryID},
	}
	err = ctx.Client.Countries.Delete(context.Background(), deleteReq, nil)
	require.NoError(t, err)

	// Verify deletion
	country, err := ctx.Client.Countries.GetByID(context.Background(), newCountryID)
	require.NoError(t, err)
	assert.True(t, country.IsDeleted)

	// Restore country
	restoreReq := &models.RestoreCountryRequest{
		BaseCommand: models.BaseCommand{ID: newCountryID},
	}
	err = ctx.Client.Countries.Restore(context.Background(), restoreReq, nil)
	require.NoError(t, err)

	// Verify restoration
	country, err = ctx.Client.Countries.GetByID(context.Background(), newCountryID)
	require.NoError(t, err)
	assert.False(t, country.IsDeleted)
}

func TestCurrencies_GetByID(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	currencyID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	currency, err := ctx.Client.Currencies.GetByID(context.Background(), currencyID)
	require.NoError(t, err)
	assert.NotNil(t, currency)
	assert.Equal(t, currencyID, currency.ID)
	assert.Equal(t, "BONUS", currency.Name)
	assert.True(t, currency.IsActive)
}

func TestCurrencies_GetList(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	list, err := ctx.Client.Currencies.GetList(context.Background(), &models.GetListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, list)
	assert.GreaterOrEqual(t, list.Total, int64(1))
}

func TestCurrencies_CreateAndModify(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	newCurrencyID := uuid.New()

	// Create currency
	createReq := &models.CreateCurrencyRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
		Name:        "POINTS",
	}
	err := ctx.Client.Currencies.Create(context.Background(), createReq, nil)
	require.NoError(t, err)

	// Verify creation
	currency, err := ctx.Client.Currencies.GetByID(context.Background(), newCurrencyID)
	require.NoError(t, err)
	assert.Equal(t, "POINTS", currency.Name)
	assert.False(t, currency.IsActive)

	// Rename currency
	renameReq := &models.RenameCurrencyRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
		Name:        "BONUS_POINTS",
	}
	err = ctx.Client.Currencies.Rename(context.Background(), renameReq, nil)
	require.NoError(t, err)

	// Set description
	descReq := &models.SetCurrencyDescriptionRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
		Description: "Bonus points for loyalty program",
	}
	err = ctx.Client.Currencies.SetDescription(context.Background(), descReq, nil)
	require.NoError(t, err)

	// Set rate
	rateReq := &models.SetCurrencyRateRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
		Rate:        0.5,
	}
	err = ctx.Client.Currencies.SetRate(context.Background(), rateReq, nil)
	require.NoError(t, err)

	// Activate currency
	activateReq := &models.ActivateCurrencyRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
	}
	err = ctx.Client.Currencies.Activate(context.Background(), activateReq, nil)
	require.NoError(t, err)

	// Verify all changes
	currency, err = ctx.Client.Currencies.GetByID(context.Background(), newCurrencyID)
	require.NoError(t, err)
	assert.Equal(t, "BONUS_POINTS", currency.Name)
	assert.NotNil(t, currency.Description)
	assert.Equal(t, "Bonus points for loyalty program", *currency.Description)
	assert.Equal(t, 0.5, currency.Rate)
	assert.True(t, currency.IsActive)
}

func TestCurrencies_ActivateDeactivate(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	newCurrencyID := uuid.New()

	// Create currency
	createReq := &models.CreateCurrencyRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
		Name:        "TEST_CURRENCY",
	}
	err := ctx.Client.Currencies.Create(context.Background(), createReq, nil)
	require.NoError(t, err)

	// Activate
	activateReq := &models.ActivateCurrencyRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
	}
	err = ctx.Client.Currencies.Activate(context.Background(), activateReq, nil)
	require.NoError(t, err)

	currency, err := ctx.Client.Currencies.GetByID(context.Background(), newCurrencyID)
	require.NoError(t, err)
	assert.True(t, currency.IsActive)

	// Deactivate
	deactivateReq := &models.DeactivateCurrencyRequest{
		BaseCommand: models.BaseCommand{ID: newCurrencyID},
	}
	err = ctx.Client.Currencies.Deactivate(context.Background(), deactivateReq, nil)
	require.NoError(t, err)

	currency, err = ctx.Client.Currencies.GetByID(context.Background(), newCurrencyID)
	require.NoError(t, err)
	assert.False(t, currency.IsActive)
}
