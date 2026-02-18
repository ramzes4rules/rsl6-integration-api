// Package models contains data structures for RS Loyalty API v2
package models

import (
	"time"

	"github.com/google/uuid"
)

// GenderTypes represents gender enumeration
type GenderTypes string

const (
	GenderMale    GenderTypes = "Male"
	GenderFemale  GenderTypes = "Female"
	GenderUnknown GenderTypes = "Unknown"
)

// CommunicationValueTypes represents communication value types
type CommunicationValueTypes string

const (
	CommunicationPhone CommunicationValueTypes = "Phone"
	CommunicationEmail CommunicationValueTypes = "Email"
)

// SubscriptionTypes represents subscription types
type SubscriptionTypes string

const (
	SubscriptionClientPortalMailing SubscriptionTypes = "ClientPortalMailing"
	SubscriptionEmailMailing        SubscriptionTypes = "EmailMailing"
	SubscriptionPushMailing         SubscriptionTypes = "PushMailing"
	SubscriptionSmsMailing          SubscriptionTypes = "SmsMailing"
	SubscriptionPhoneCallMailing    SubscriptionTypes = "PhoneCallMailing"
	SubscriptionViberMailing        SubscriptionTypes = "ViberMailing"
)

// LoyaltyCardStatuses represents loyalty card status
type LoyaltyCardStatuses string

const (
	LoyaltyCardStatusActive                               LoyaltyCardStatuses = "Active"
	LoyaltyCardStatusBlocked                              LoyaltyCardStatuses = "Blocked"
	LoyaltyCardStatusCreated                              LoyaltyCardStatuses = "Created"
	LoyaltyCardStatusIssued                               LoyaltyCardStatuses = "Issued"
	LoyaltyCardStatusReadyToIssued                        LoyaltyCardStatuses = "ReadyToIssued"
	LoyaltyCardStatusExpired                              LoyaltyCardStatuses = "Expired"
	LoyaltyCardStatusBlockedByCustomerWasBroken           LoyaltyCardStatuses = "BlockedByCustomerWasBroken"
	LoyaltyCardStatusBlockedByCustomerWasLost             LoyaltyCardStatuses = "BlockedByCustomerWasLost"
	LoyaltyCardStatusBlockedByCustomerDeletedPersonalData LoyaltyCardStatuses = "BlockedByCustomerDeletedPersonalData"
)

// TransactionTypes represents transaction types
type TransactionTypes string

const (
	TransactionAccrual  TransactionTypes = "Accrual"
	TransactionSubtract TransactionTypes = "Subtract"
	TransactionBurning  TransactionTypes = "Burning"
)

// PortalTypes represents portal types
type PortalTypes string

const (
	PortalClientPortal      PortalTypes = "ClientPortal"
	PortalMobileApplication PortalTypes = "MobileApplication"
)

// CustomerPropertyTypes represents customer property types
type CustomerPropertyTypes string

const (
	CustomerPropertyInteger      CustomerPropertyTypes = "Integer"
	CustomerPropertyString       CustomerPropertyTypes = "String"
	CustomerPropertyDate         CustomerPropertyTypes = "Date"
	CustomerPropertyEnum         CustomerPropertyTypes = "Enum"
	CustomerPropertyBoolean      CustomerPropertyTypes = "Boolean"
	CustomerPropertyEnumMultiple CustomerPropertyTypes = "EnumMultiple"
)

// BaseCommand represents base command structure with common fields
type BaseCommand struct {
	CommandID     *uuid.UUID `json:"commandId,omitempty"`
	OperationDate *time.Time `json:"operationDate,omitempty"`
	ID            uuid.UUID  `json:"id"`
}

// CustomerDto represents customer data
type CustomerDto struct {
	ID                               uuid.UUID                       `json:"id"`
	FirstName                        *string                         `json:"firstName,omitempty"`
	SecondName                       *string                         `json:"secondName,omitempty"`
	LastName                         *string                         `json:"lastName,omitempty"`
	Address                          *string                         `json:"address,omitempty"`
	Birthday                         *string                         `json:"birthday,omitempty"`
	Gender                           GenderTypes                     `json:"gender"`
	Properties                       []CustomerPropertyDto           `json:"properties"`
	Subscriptions                    []SubscriptionTypes             `json:"subscriptions"`
	CommunicationValues              []CustomerCommunicationValueDto `json:"communicationValues"`
	TerritorialDivisionID            *uuid.UUID                      `json:"territorialDivisionId,omitempty"`
	AllowBySendingVirtualCopyReceipt bool                            `json:"allowBySendingVirtualCopyReceipt"`
	ClientPortalActivationDate       *time.Time                      `json:"clientPortalActivationDate,omitempty"`
	MobileApplicationActivationDate  *time.Time                      `json:"mobileApplicationActivationDate,omitempty"`
	Children                         []CustomerChildDto              `json:"children"`
	StaticSegments                   []uuid.UUID                     `json:"staticSegments"`
}

// CustomerChildDto represents customer child data
type CustomerChildDto struct {
	ID       uuid.UUID   `json:"id"`
	Name     *string     `json:"name,omitempty"`
	Birthday *string     `json:"birthday,omitempty"`
	Gender   GenderTypes `json:"gender"`
}

// CustomerCommunicationValueDto represents customer communication value
type CustomerCommunicationValueDto struct {
	ValueType CommunicationValueTypes `json:"valueType"`
	Value     string                  `json:"value"`
	Confirmed bool                    `json:"confirmed"`
}

// CustomerPropertyDto represents customer property
type CustomerPropertyDto struct {
	PropertyID uuid.UUID                `json:"propertyId"`
	Value      CustomerPropertyValueDto `json:"value"`
}

// CustomerPropertyValueDto represents customer property value
type CustomerPropertyValueDto struct {
	BooleanValue *bool       `json:"booleanValue,omitempty"`
	DateValue    *string     `json:"dateValue,omitempty"`
	StringValue  *string     `json:"stringValue,omitempty"`
	IntValue     *int        `json:"intValue,omitempty"`
	EnumValue    *uuid.UUID  `json:"enumValue,omitempty"`
	EnumValues   []uuid.UUID `json:"enumValues,omitempty"`
}

// CustomerListDto represents list of customers
type CustomerListDto struct {
	Total  int64         `json:"total"`
	Values []CustomerDto `json:"values"`
}

// GetByIdRequest represents request for getting entity by ID
type GetByIdRequest struct {
	ID uuid.UUID `json:"id"`
}

// GetListRequest represents request for getting list of entities
type GetListRequest struct {
	Filter         []uuid.UUID `json:"filter,omitempty"`
	Take           *int        `json:"take,omitempty"`
	Skip           *int        `json:"skip,omitempty"`
	CreateDateFrom *string     `json:"createDateFrom,omitempty"`
	CreateDateTo   *string     `json:"createDateTo,omitempty"`
}

// GetByCommunicationValueRequest represents request for getting customer by communication value
type GetByCommunicationValueRequest struct {
	CommunicationValueType CommunicationValueTypes `json:"communicationValueType"`
	Value                  string                  `json:"value"`
}

// BalanceDto represents balance data
type BalanceDto struct {
	CurrencyID uuid.UUID `json:"currencyId"`
	Value      float64   `json:"value"`
}

// BalanceListDto represents list of balances
type BalanceListDto struct {
	ID       uuid.UUID    `json:"id"`
	Balances []BalanceDto `json:"balances"`
}

// CustomerBalanceDto represents customer balance
type CustomerBalanceDto struct {
	Balances            *BalanceListDto  `json:"balances,omitempty"`
	LoyaltyCardBalances []BalanceListDto `json:"loyaltyCardBalances,omitempty"`
}

// TransactionDto represents transaction data
type TransactionDto struct {
	TransactionID   uuid.UUID        `json:"transactionId"`
	OperationDate   time.Time        `json:"operationDate"`
	TransactionType TransactionTypes `json:"transactionType"`
	Amount          float64          `json:"amount"`
	CurrencyID      uuid.UUID        `json:"currencyId"`
	ExpirationDate  *time.Time       `json:"expirationDate,omitempty"`
	ActivationDate  *time.Time       `json:"activationDate,omitempty"`
}

// TransactionDetailDto represents detailed transaction
type TransactionDetailDto struct {
	TransactionDto
	AccountID uuid.UUID `json:"accountId"`
}

// TransactionListDto represents list of transactions
type TransactionListDto struct {
	Transactions []TransactionDetailDto `json:"transactions"`
}

// GetTransactionsRequest represents request for getting transactions
type GetTransactionsRequest struct {
	FromDate *time.Time `json:"fromDate,omitempty"`
	ToDate   *time.Time `json:"toDate,omitempty"`
	Take     *int       `json:"take,omitempty"`
	Skip     *int       `json:"skip,omitempty"`
}

// CountryDto represents country data
type CountryDto struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	IsDeleted bool      `json:"isDeleted"`
}

// CountryListDto represents list of countries
type CountryListDto struct {
	Total  int64        `json:"total"`
	Values []CountryDto `json:"values"`
}

// CurrencyDto represents currency data
type CurrencyDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	PublicName  *string   `json:"publicName,omitempty"`
	Rate        float64   `json:"rate"`
	IsActive    bool      `json:"isActive"`
	IsDeleted   bool      `json:"isDeleted"`
}

// CurrencyListDto represents list of currencies
type CurrencyListDto struct {
	Total  int64         `json:"total"`
	Values []CurrencyDto `json:"values"`
}

// LoyaltyCardDto represents loyalty card data
type LoyaltyCardDto struct {
	ID           uuid.UUID           `json:"id"`
	Number       string              `json:"number"`
	SeriesID     uuid.UUID           `json:"seriesId"`
	Status       LoyaltyCardStatuses `json:"status"`
	CustomerID   *uuid.UUID          `json:"customerId,omitempty"`
	IssueDate    *time.Time          `json:"issueDate,omitempty"`
	ExpireDate   *time.Time          `json:"expireDate,omitempty"`
	ActivateDate *time.Time          `json:"activateDate,omitempty"`
	IsDeleted    bool                `json:"isDeleted"`
}

// LoyaltyCardListDto represents list of loyalty cards
type LoyaltyCardListDto struct {
	Total  int64            `json:"total"`
	Values []LoyaltyCardDto `json:"values"`
}

// AccrualToCustomerRequest represents request for accrual to customer
type AccrualToCustomerRequest struct {
	BaseCommand
	CustomerID     uuid.UUID  `json:"customerId"`
	CurrencyID     uuid.UUID  `json:"currencyId"`
	Amount         float64    `json:"amount"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	ActivationDate *time.Time `json:"activationDate,omitempty"`
	Comment        *string    `json:"comment,omitempty"`
}

// AccrualToLoyaltyCardRequest represents request for accrual to loyalty card
type AccrualToLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID  uuid.UUID  `json:"loyaltyCardId"`
	CurrencyID     uuid.UUID  `json:"currencyId"`
	Amount         float64    `json:"amount"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	ActivationDate *time.Time `json:"activationDate,omitempty"`
	Comment        *string    `json:"comment,omitempty"`
}

// SubtractFromCustomerRequest represents request for subtract from customer
type SubtractFromCustomerRequest struct {
	BaseCommand
	CustomerID uuid.UUID `json:"customerId"`
	CurrencyID uuid.UUID `json:"currencyId"`
	Amount     float64   `json:"amount"`
	Comment    *string   `json:"comment,omitempty"`
}

// SubtractFromLoyaltyCardRequest represents request for subtract from loyalty card
type SubtractFromLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID uuid.UUID `json:"loyaltyCardId"`
	CurrencyID    uuid.UUID `json:"currencyId"`
	Amount        float64   `json:"amount"`
	Comment       *string   `json:"comment,omitempty"`
}

// CreateCountryRequest represents request for creating country
type CreateCountryRequest struct {
	BaseCommand
	Code string `json:"code"`
}

// SetCountryCodeRequest represents request for setting country code
type SetCountryCodeRequest struct {
	BaseCommand
	Code string `json:"code"`
}

// DeleteCountryRequest represents request for deleting country
type DeleteCountryRequest struct {
	BaseCommand
}

// RestoreCountryRequest represents request for restoring country
type RestoreCountryRequest struct {
	BaseCommand
}

// CreateCurrencyRequest represents request for creating currency
type CreateCurrencyRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// RenameCurrencyRequest represents request for renaming currency
type RenameCurrencyRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetCurrencyDescriptionRequest represents request for setting currency description
type SetCurrencyDescriptionRequest struct {
	BaseCommand
	Description string `json:"description"`
}

// SetCurrencyRateRequest represents request for setting currency rate
type SetCurrencyRateRequest struct {
	BaseCommand
	Rate float64 `json:"rate"`
}

// ActivateCurrencyRequest represents request for activating currency
type ActivateCurrencyRequest struct {
	BaseCommand
}

// DeactivateCurrencyRequest represents request for deactivating currency
type DeactivateCurrencyRequest struct {
	BaseCommand
}

// RenameCustomerRequest represents request for renaming customer
type RenameCustomerRequest struct {
	BaseCommand
	FirstName  *string `json:"firstName,omitempty"`
	SecondName *string `json:"secondName,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
}

// SetCustomerCommunicationValueRequest represents request for setting customer communication value
type SetCustomerCommunicationValueRequest struct {
	BaseCommand
	Value     string                  `json:"value"`
	ValueType CommunicationValueTypes `json:"valueType"`
}

// SetCustomerBirthdayRequest represents request for setting customer birthday
type SetCustomerBirthdayRequest struct {
	BaseCommand
	Birthday string `json:"birthday"`
}

// SetCustomerAddressRequest represents request for setting customer address
type SetCustomerAddressRequest struct {
	BaseCommand
	Address string `json:"address"`
}

// AllowSubscriptionRequest represents request for allowing subscription
type AllowSubscriptionRequest struct {
	BaseCommand
	SubscriptionType SubscriptionTypes `json:"subscriptionType"`
}

// DisallowSubscriptionRequest represents request for disallowing subscription
type DisallowSubscriptionRequest struct {
	BaseCommand
	SubscriptionType SubscriptionTypes `json:"subscriptionType"`
}

// RequestHeaders represents common request headers
type RequestHeaders struct {
	CommandID          *string `json:"command-id,omitempty"`
	OperationDate      *string `json:"operation-date,omitempty"`
	UserID             *string `json:"user-id,omitempty"`
	InteractionChannel *string `json:"interaction-channel,omitempty"`
}

// APIError represents API error response
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}
