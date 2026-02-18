package tests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rsl6/loyalty-client/models"
)

func TestCustomers_GetByID(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	customer, err := ctx.Client.Customers.GetByID(context.Background(), customerID)
	require.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, customerID, customer.ID)
	assert.NotNil(t, customer.FirstName)
	assert.Equal(t, "Иван", *customer.FirstName)
}

func TestCustomers_GetByID_NotFound(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	nonExistentID := uuid.New()

	customer, err := ctx.Client.Customers.GetByID(context.Background(), nonExistentID)
	assert.Error(t, err)
	assert.Nil(t, customer)

	apiErr, ok := err.(*models.APIError)
	assert.True(t, ok)
	assert.Equal(t, 404, apiErr.StatusCode)
}

func TestCustomers_GetByCommunicationValue(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	req := &models.GetByCommunicationValueRequest{
		CommunicationValueType: models.CommunicationEmail,
		Value:                  "ivan@example.com",
	}

	customer, err := ctx.Client.Customers.GetByCommunicationValue(context.Background(), req)
	require.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Иван", *customer.FirstName)
}

func TestCustomers_GetList(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	take := 10
	req := &models.GetListRequest{
		Take: &take,
	}

	list, err := ctx.Client.Customers.GetList(context.Background(), req)
	require.NoError(t, err)
	assert.NotNil(t, list)
	assert.GreaterOrEqual(t, list.Total, int64(1))
	assert.NotEmpty(t, list.Values)
}

func TestCustomers_Rename(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	newFirstName := "Петр"
	newLastName := "Иванов"

	req := &models.RenameCustomerRequest{
		BaseCommand: models.BaseCommand{ID: customerID},
		FirstName:   &newFirstName,
		LastName:    &newLastName,
	}

	err := ctx.Client.Customers.Rename(context.Background(), req, nil)
	require.NoError(t, err)

	// Verify the change
	customer, err := ctx.Client.Customers.GetByID(context.Background(), customerID)
	require.NoError(t, err)
	assert.Equal(t, newFirstName, *customer.FirstName)
	assert.Equal(t, newLastName, *customer.LastName)
}

func TestCustomers_SetCommunicationValue(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	req := &models.SetCustomerCommunicationValueRequest{
		BaseCommand: models.BaseCommand{ID: customerID},
		Value:       "new@example.com",
		ValueType:   models.CommunicationEmail,
	}

	err := ctx.Client.Customers.SetCommunicationValue(context.Background(), req, nil)
	require.NoError(t, err)

	// Verify the change
	customer, err := ctx.Client.Customers.GetByID(context.Background(), customerID)
	require.NoError(t, err)

	found := false
	for _, cv := range customer.CommunicationValues {
		if cv.ValueType == models.CommunicationEmail && cv.Value == "new@example.com" {
			found = true
			break
		}
	}
	assert.True(t, found, "New email should be set")
}

func TestCustomers_GetBalanceByID(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	balance, err := ctx.Client.Customers.GetBalanceByID(context.Background(), customerID)
	require.NoError(t, err)
	assert.NotNil(t, balance)
	assert.NotNil(t, balance.Balances)
	assert.NotEmpty(t, balance.Balances.Balances)
}

func TestCustomers_AllowDisallowSubscription(t *testing.T) {
	ctx := SetupTestContext(t)
	defer ctx.Cleanup()

	customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	// Allow SMS subscription
	allowReq := &models.AllowSubscriptionRequest{
		BaseCommand:      models.BaseCommand{ID: customerID},
		SubscriptionType: models.SubscriptionSmsMailing,
	}
	err := ctx.Client.Customers.AllowSubscription(context.Background(), allowReq, nil)
	require.NoError(t, err)

	// Verify
	customer, err := ctx.Client.Customers.GetByID(context.Background(), customerID)
	require.NoError(t, err)
	assert.Contains(t, customer.Subscriptions, models.SubscriptionSmsMailing)

	// Disallow SMS subscription
	disallowReq := &models.DisallowSubscriptionRequest{
		BaseCommand:      models.BaseCommand{ID: customerID},
		SubscriptionType: models.SubscriptionSmsMailing,
	}
	err = ctx.Client.Customers.DisallowSubscription(context.Background(), disallowReq, nil)
	require.NoError(t, err)

	// Verify
	customer, err = ctx.Client.Customers.GetByID(context.Background(), customerID)
	require.NoError(t, err)
	assert.NotContains(t, customer.Subscriptions, models.SubscriptionSmsMailing)
}
