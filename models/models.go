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

// SetCurrencyPublicNameRequest represents request for setting currency public name
type SetCurrencyPublicNameRequest struct {
	BaseCommand
	PublicName string `json:"publicName"`
}

// SetCurrencyCalculateRoundRuleRequest represents request for setting currency calculate round rule
type SetCurrencyCalculateRoundRuleRequest struct {
	BaseCommand
	CalculateRoundRule string `json:"calculateRoundRule"`
}

// SetCurrencyCaptionRequest represents request for setting currency caption
type SetCurrencyCaptionRequest struct {
	BaseCommand
	Caption string `json:"caption"`
}

// BatchRequest represents batch request for executing multiple commands
type BatchRequest struct {
	Commands []interface{} `json:"commands"`
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

// ========== CustomerCards Commands ==========

// BindLoyaltyCardRequest represents request for binding loyalty card to customer
type BindLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID      uuid.UUID `json:"loyaltyCardId"`
	PinCode            *string   `json:"pinCode,omitempty"`
	InteractionChannel *string   `json:"interactionChannel,omitempty"`
}

// UnbindLoyaltyCardRequest represents request for unbinding loyalty card from customer
type UnbindLoyaltyCardRequest struct {
	BaseCommand
	LoyaltyCardID uuid.UUID `json:"loyaltyCardId"`
}

// CustomerCardsBatchRequest represents batch request for customer cards commands
type CustomerCardsBatchRequest struct {
	Commands []interface{} `json:"commands"`
}

// ========== CustomerProperty Commands ==========

// CreateCustomerPropertyRequest represents request for creating customer property
type CreateCustomerPropertyRequest struct {
	BaseCommand
	Name       *string               `json:"name,omitempty"`
	Type       CustomerPropertyTypes `json:"type"`
	ExternalID *string               `json:"externalId,omitempty"`
}

// RenameCustomerPropertyRequest represents request for renaming customer property
type RenameCustomerPropertyRequest struct {
	BaseCommand
	Name *string `json:"name,omitempty"`
}

// DeleteCustomerPropertyRequest represents request for deleting customer property
type DeleteCustomerPropertyRequest struct {
	BaseCommand
}

// RestoreCustomerPropertyRequest represents request for restoring customer property
type RestoreCustomerPropertyRequest struct {
	BaseCommand
}

// AddEnumCustomerPropertyRequest represents request for adding enum value to customer property
type AddEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
	Name        *string   `json:"name,omitempty"`
}

// RenameEnumCustomerPropertyRequest represents request for renaming enum value of customer property
type RenameEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
	Name        *string   `json:"name,omitempty"`
}

// DeleteEnumCustomerPropertyRequest represents request for deleting enum value from customer property
type DeleteEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
}

// RestoreEnumCustomerPropertyRequest represents request for restoring enum value of customer property
type RestoreEnumCustomerPropertyRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
}

// CustomerPropertyBatchRequest represents batch request for customer property commands
type CustomerPropertyBatchRequest struct {
	Commands []interface{} `json:"commands"`
}

// PropertyDefinitionDto represents customer property definition
type PropertyDefinitionDto struct {
	ID        uuid.UUID              `json:"id"`
	Type      CustomerPropertyTypes  `json:"type"`
	Name      *string                `json:"name,omitempty"`
	IsDeleted bool                   `json:"isDeleted"`
	Enums     []EnumPropertyValueDto `json:"enums,omitempty"`
}

// EnumPropertyValueDto represents enum property value
type EnumPropertyValueDto struct {
	ID        uuid.UUID `json:"id"`
	Name      *string   `json:"name,omitempty"`
	IsDeleted bool      `json:"isDeleted"`
}

// PropertyDefinitionListDto represents list of customer property definitions
type PropertyDefinitionListDto struct {
	Total  int64                   `json:"total"`
	Values []PropertyDefinitionDto `json:"values"`
}

// ========== CustomerSegment Commands ==========

// AddToStaticSegmentRequest represents request for adding customer to static segment
type AddToStaticSegmentRequest struct {
	BaseCommand
	StaticSegmentID uuid.UUID `json:"staticSegmentId"`
}

// RemoveFromStaticSegmentRequest represents request for removing customer from static segment
type RemoveFromStaticSegmentRequest struct {
	BaseCommand
	StaticSegmentID uuid.UUID `json:"staticSegmentId"`
}

// CustomerSegmentBatchRequest represents batch request for customer segment commands
type CustomerSegmentBatchRequest struct {
	Commands []interface{} `json:"commands"`
}

// ========== Common Request Types ==========

// DeleteRequest represents a generic delete request
type DeleteRequest struct {
	BaseCommand
}

// RestoreRequest represents a generic restore request
type RestoreRequest struct {
	BaseCommand
}

// MoveToArchiveRequest represents a generic move to archive request
type MoveToArchiveRequest struct {
	BaseCommand
}

// UnarchiveRequest represents a generic unarchive request
type UnarchiveRequest struct {
	BaseCommand
}

// RenameRequest represents a generic rename request
type RenameRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetDescriptionRequest represents a generic set description request
type SetDescriptionRequest struct {
	BaseCommand
	Description string `json:"description"`
}

// SetColorRequest represents a generic set color request
type SetColorRequest struct {
	BaseCommand
	Color string `json:"color"`
}

// SetParentRequest represents a generic set parent request
type SetParentRequest struct {
	BaseCommand
	ParentID *uuid.UUID `json:"parentId,omitempty"`
}

// ========== External Card Series ==========

// CreateExternalGiftCardSeriesRequest represents request for creating external gift card series
type CreateExternalGiftCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// CreateExternalLoyaltyCardSeriesRequest represents request for creating external loyalty card series
type CreateExternalLoyaltyCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// CreateExternalSponsoredCardSeriesRequest represents request for creating external sponsored card series
type CreateExternalSponsoredCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// RenameExternalCardSeriesRequest represents request for renaming external card series
type RenameExternalCardSeriesRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// AddCirculationRequest represents request for adding circulation to external card series
type AddCirculationRequest struct {
	BaseCommand
	CirculationID uuid.UUID `json:"circulationId"`
	Name          string    `json:"name"`
}

// RenameCirculationRequest represents request for renaming circulation
type RenameCirculationRequest struct {
	BaseCommand
	CirculationID uuid.UUID `json:"circulationId"`
	Name          string    `json:"name"`
}

// ExternalCardSeriesDto represents external card series data
type ExternalCardSeriesDto struct {
	ID           uuid.UUID                    `json:"id"`
	Name         string                       `json:"name"`
	IsDeleted    bool                         `json:"isDeleted"`
	Circulations []ExternalCardCirculationDto `json:"circulations,omitempty"`
}

// ExternalCardCirculationDto represents external card circulation data
type ExternalCardCirculationDto struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// ExternalCardSeriesListDto represents list of external card series
type ExternalCardSeriesListDto struct {
	Total  int64                   `json:"total"`
	Values []ExternalCardSeriesDto `json:"values"`
}

// ========== Gift Cards ==========

// GiftCardStatuses represents gift card status
type GiftCardStatuses string

const (
	GiftCardStatusCreated       GiftCardStatuses = "Created"
	GiftCardStatusReadyToIssued GiftCardStatuses = "ReadyToIssued"
	GiftCardStatusIssued        GiftCardStatuses = "Issued"
	GiftCardStatusActive        GiftCardStatuses = "Active"
	GiftCardStatusBlocked       GiftCardStatuses = "Blocked"
	GiftCardStatusExpired       GiftCardStatuses = "Expired"
	GiftCardStatusFullyRedeemed GiftCardStatuses = "FullyRedeemed"
)

// CreateGiftCardRequest represents request for creating a gift card
type CreateGiftCardRequest struct {
	BaseCommand
	SeriesID uuid.UUID `json:"seriesId"`
}

// SetCardNumberRequest represents request for setting card number
type SetCardNumberRequest struct {
	BaseCommand
	Number string `json:"number"`
}

// SetCardBarcodeRequest represents request for setting card barcode
type SetCardBarcodeRequest struct {
	BaseCommand
	Barcode string `json:"barcode"`
}

// SetCardPinCodeRequest represents request for setting card pin code
type SetCardPinCodeRequest struct {
	BaseCommand
	PinCode string `json:"pinCode"`
}

// SetCardExpirationDateRequest represents request for setting card expiration date
type SetCardExpirationDateRequest struct {
	BaseCommand
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// SetGiftCardValidityPeriodRequest represents request for setting gift card validity period
type SetGiftCardValidityPeriodRequest struct {
	BaseCommand
	ValidityPeriod int `json:"validityPeriod"`
}

// ConfirmReadinessForIssuanceRequest represents request for confirming readiness for issuance
type ConfirmReadinessForIssuanceRequest struct {
	BaseCommand
}

// IssueGiftCardRequest represents request for issuing a gift card
type IssueGiftCardRequest struct {
	BaseCommand
	Amount        float64    `json:"amount"`
	IssueReasonID *uuid.UUID `json:"issueReasonId,omitempty"`
}

// CancelIssueRequest represents request for canceling card issue
type CancelIssueRequest struct {
	BaseCommand
}

// BlockGiftCardRequest represents request for blocking a gift card
type BlockGiftCardRequest struct {
	BaseCommand
	BlockReasonID *uuid.UUID `json:"blockReasonId,omitempty"`
}

// UnblockCardRequest represents request for unblocking a card
type UnblockCardRequest struct {
	BaseCommand
}

// SubtractGiftCardRequest represents request for subtracting from gift card
type SubtractGiftCardRequest struct {
	BaseCommand
	Amount float64 `json:"amount"`
}

// RefundGiftCardRequest represents request for refunding to gift card
type RefundGiftCardRequest struct {
	BaseCommand
	Amount float64 `json:"amount"`
}

// AddCardGroupRequest represents request for adding card to group
type AddCardGroupRequest struct {
	BaseCommand
	GroupID uuid.UUID `json:"groupId"`
}

// RemoveCardGroupRequest represents request for removing card from group
type RemoveCardGroupRequest struct {
	BaseCommand
	GroupID uuid.UUID `json:"groupId"`
}

// GiftCardDto represents gift card data
type GiftCardDto struct {
	ID             uuid.UUID        `json:"id"`
	Number         string           `json:"number"`
	Barcode        *string          `json:"barcode,omitempty"`
	SeriesID       uuid.UUID        `json:"seriesId"`
	Status         GiftCardStatuses `json:"status"`
	Balance        float64          `json:"balance"`
	InitialAmount  float64          `json:"initialAmount"`
	IssueDate      *time.Time       `json:"issueDate,omitempty"`
	ExpirationDate *time.Time       `json:"expirationDate,omitempty"`
	IsDeleted      bool             `json:"isDeleted"`
	Groups         []uuid.UUID      `json:"groups,omitempty"`
}

// GiftCardListDto represents list of gift cards
type GiftCardListDto struct {
	Total  int64         `json:"total"`
	Values []GiftCardDto `json:"values"`
}

// ========== Reasons (Block, Issue, Manual Accrual) ==========

// CreateReasonRequest represents request for creating a reason
type CreateReasonRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// ReasonDto represents reason data
type ReasonDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IsDeleted  bool      `json:"isDeleted"`
	IsArchived bool      `json:"isArchived"`
}

// ReasonListDto represents list of reasons
type ReasonListDto struct {
	Total  int64       `json:"total"`
	Values []ReasonDto `json:"values"`
}

// ========== Groups (Card Groups) ==========

// CreateGroupRequest represents request for creating a group
type CreateGroupRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// GroupDto represents group data
type GroupDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Color      *string   `json:"color,omitempty"`
	IsDeleted  bool      `json:"isDeleted"`
	IsArchived bool      `json:"isArchived"`
}

// GroupListDto represents list of groups
type GroupListDto struct {
	Total  int64      `json:"total"`
	Values []GroupDto `json:"values"`
}

// ========== Items ==========

// CreateItemRequest represents request for creating an item
type CreateItemRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetPublicNameRequest represents request for setting public name
type SetPublicNameRequest struct {
	BaseCommand
	PublicName string `json:"publicName"`
}

// SetArticleRequest represents request for setting article
type SetArticleRequest struct {
	BaseCommand
	Article string `json:"article"`
}

// SetExternalIDRequest represents request for setting external ID
type SetExternalIDRequest struct {
	BaseCommand
	ExternalID string `json:"externalId"`
}

// SetItemGroupIDRequest represents request for setting item group ID
type SetItemGroupIDRequest struct {
	BaseCommand
	ItemGroupID *uuid.UUID `json:"itemGroupId,omitempty"`
}

// SetRestrictionRequest represents request for setting restriction
type SetRestrictionRequest struct {
	BaseCommand
	Restriction string `json:"restriction"`
}

// SetBooleanPropertyValueRequest represents request for setting boolean property value
type SetBooleanPropertyValueRequest struct {
	BaseCommand
	PropertyID uuid.UUID `json:"propertyId"`
	Value      bool      `json:"value"`
}

// SetIntPropertyValueRequest represents request for setting int property value
type SetIntPropertyValueRequest struct {
	BaseCommand
	PropertyID uuid.UUID `json:"propertyId"`
	Value      int       `json:"value"`
}

// SetStringPropertyValueRequest represents request for setting string property value
type SetStringPropertyValueRequest struct {
	BaseCommand
	PropertyID uuid.UUID `json:"propertyId"`
	Value      string    `json:"value"`
}

// SetDatePropertyValueRequest represents request for setting date property value
type SetDatePropertyValueRequest struct {
	BaseCommand
	PropertyID uuid.UUID `json:"propertyId"`
	Value      string    `json:"value"`
}

// SetEnumPropertyValueRequest represents request for setting enum property value
type SetEnumPropertyValueRequest struct {
	BaseCommand
	PropertyID uuid.UUID `json:"propertyId"`
	Value      uuid.UUID `json:"value"`
}

// SetEnumsPropertyValueRequest represents request for setting multiple enum property values
type SetEnumsPropertyValueRequest struct {
	BaseCommand
	PropertyID uuid.UUID   `json:"propertyId"`
	Values     []uuid.UUID `json:"values"`
}

// AddMediaContentRequest represents request for adding media content
type AddMediaContentRequest struct {
	BaseCommand
	MediaContentID uuid.UUID `json:"mediaContentId"`
	Name           string    `json:"name"`
	Source         string    `json:"source"`
	Type           string    `json:"type"`
}

// RemoveMediaContentRequest represents request for removing media content
type RemoveMediaContentRequest struct {
	BaseCommand
	MediaContentID uuid.UUID `json:"mediaContentId"`
}

// RenameMediaContentRequest represents request for renaming media content
type RenameMediaContentRequest struct {
	BaseCommand
	MediaContentID uuid.UUID `json:"mediaContentId"`
	Name           string    `json:"name"`
}

// SetMediaContentSourceRequest represents request for setting media content source
type SetMediaContentSourceRequest struct {
	BaseCommand
	MediaContentID uuid.UUID `json:"mediaContentId"`
	Source         string    `json:"source"`
}

// SetMediaContentTypeRequest represents request for setting media content type
type SetMediaContentTypeRequest struct {
	BaseCommand
	MediaContentID uuid.UUID `json:"mediaContentId"`
	Type           string    `json:"type"`
}

// SetDefaultImageRequest represents request for setting default image
type SetDefaultImageRequest struct {
	BaseCommand
	MediaContentID uuid.UUID `json:"mediaContentId"`
}

// ClearDefaultImageRequest represents request for clearing default image
type ClearDefaultImageRequest struct {
	BaseCommand
}

// AddSaleItemRequest represents request for adding sale item
type AddSaleItemRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
	ExternalID string    `json:"externalId"`
}

// DeleteSaleItemRequest represents request for deleting sale item
type DeleteSaleItemRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
}

// RestoreSaleItemRequest represents request for restoring sale item
type RestoreSaleItemRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
}

// SetSaleItemExternalIDRequest represents request for setting sale item external ID
type SetSaleItemExternalIDRequest struct {
	BaseCommand
	SaleItemID uuid.UUID `json:"saleItemId"`
	ExternalID string    `json:"externalId"`
}

// SetSaleItemPropertyValueRequest represents request for setting sale item property value
type SetSaleItemPropertyValueRequest struct {
	BaseCommand
	SaleItemID uuid.UUID   `json:"saleItemId"`
	PropertyID uuid.UUID   `json:"propertyId"`
	Value      interface{} `json:"value"`
}

// ItemDto represents item data
type ItemDto struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	PublicName  *string    `json:"publicName,omitempty"`
	Description *string    `json:"description,omitempty"`
	Article     *string    `json:"article,omitempty"`
	ExternalID  *string    `json:"externalId,omitempty"`
	ItemGroupID *uuid.UUID `json:"itemGroupId,omitempty"`
	IsDeleted   bool       `json:"isDeleted"`
}

// ItemListDto represents list of items
type ItemListDto struct {
	Total  int64     `json:"total"`
	Values []ItemDto `json:"values"`
}

// ========== Item Categories ==========

// CreateItemCategoryRequest represents request for creating an item category
type CreateItemCategoryRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// AddItemToCategoryRequest represents request for adding item to category
type AddItemToCategoryRequest struct {
	BaseCommand
	ItemID uuid.UUID `json:"itemId"`
}

// RemoveItemFromCategoryRequest represents request for removing item from category
type RemoveItemFromCategoryRequest struct {
	BaseCommand
	ItemID uuid.UUID `json:"itemId"`
}

// AddItemGroupToCategoryRequest represents request for adding item group to category
type AddItemGroupToCategoryRequest struct {
	BaseCommand
	ItemGroupID uuid.UUID `json:"itemGroupId"`
}

// RemoveItemGroupFromCategoryRequest represents request for removing item group from category
type RemoveItemGroupFromCategoryRequest struct {
	BaseCommand
	ItemGroupID uuid.UUID `json:"itemGroupId"`
}

// ItemCategoryDto represents item category data
type ItemCategoryDto struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Color       *string     `json:"color,omitempty"`
	IsDeleted   bool        `json:"isDeleted"`
	IsArchived  bool        `json:"isArchived"`
	Items       []uuid.UUID `json:"items,omitempty"`
	ItemGroups  []uuid.UUID `json:"itemGroups,omitempty"`
}

// ItemCategoryListDto represents list of item categories
type ItemCategoryListDto struct {
	Total  int64             `json:"total"`
	Values []ItemCategoryDto `json:"values"`
}

// ========== Item Groups ==========

// CreateItemGroupRequest represents request for creating an item group
type CreateItemGroupRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// ItemGroupDto represents item group data
type ItemGroupDto struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	ParentID  *uuid.UUID `json:"parentId,omitempty"`
	IsDeleted bool       `json:"isDeleted"`
}

// ItemGroupListDto represents list of item groups
type ItemGroupListDto struct {
	Total  int64          `json:"total"`
	Values []ItemGroupDto `json:"values"`
}

// ========== Item Properties ==========

// CreateItemPropertyRequest represents request for creating an item property
type CreateItemPropertyRequest struct {
	BaseCommand
	Name string                `json:"name"`
	Type CustomerPropertyTypes `json:"type"`
}

// AddEnumRequest represents request for adding enum value
type AddEnumRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
	Name        string    `json:"name"`
}

// RenameEnumRequest represents request for renaming enum value
type RenameEnumRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
	Name        string    `json:"name"`
}

// DeleteEnumRequest represents request for deleting enum value
type DeleteEnumRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
}

// RestoreEnumRequest represents request for restoring enum value
type RestoreEnumRequest struct {
	BaseCommand
	EnumValueID uuid.UUID `json:"enumValueId"`
}

// ========== Opening Hours ==========

// CreateOpeningHoursRequest represents request for creating opening hours
type CreateOpeningHoursRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetWorkPeriodRequest represents request for setting work period
type SetWorkPeriodRequest struct {
	BaseCommand
	DayOfWeek int    `json:"dayOfWeek"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// OpeningHoursDto represents opening hours data
type OpeningHoursDto struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	IsDeleted   bool            `json:"isDeleted"`
	IsArchived  bool            `json:"isArchived"`
	WorkPeriods []WorkPeriodDto `json:"workPeriods,omitempty"`
}

// WorkPeriodDto represents work period data
type WorkPeriodDto struct {
	DayOfWeek int    `json:"dayOfWeek"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// OpeningHoursListDto represents list of opening hours
type OpeningHoursListDto struct {
	Total  int64             `json:"total"`
	Values []OpeningHoursDto `json:"values"`
}

// ========== POS ==========

// CreatePosRequest represents request for creating a POS
type CreatePosRequest struct {
	BaseCommand
	Name    string    `json:"name"`
	StoreID uuid.UUID `json:"storeId"`
}

// SetPosTypeRequest represents request for setting POS type
type SetPosTypeRequest struct {
	BaseCommand
	PosTypeID *uuid.UUID `json:"posTypeId,omitempty"`
}

// SetPosAuthorizationRequest represents request for setting POS authorization
type SetPosAuthorizationRequest struct {
	BaseCommand
	Login    string `json:"login"`
	Password string `json:"password"`
}

// MoveToStoreRequest represents request for moving POS to store
type MoveToStoreRequest struct {
	BaseCommand
	StoreID uuid.UUID `json:"storeId"`
}

// BlockPosRequest represents request for blocking POS
type BlockPosRequest struct {
	BaseCommand
}

// UnblockPosRequest represents request for unblocking POS
type UnblockPosRequest struct {
	BaseCommand
}

// PosDto represents POS data
type PosDto struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	StoreID    uuid.UUID  `json:"storeId"`
	PosTypeID  *uuid.UUID `json:"posTypeId,omitempty"`
	IsBlocked  bool       `json:"isBlocked"`
	IsDeleted  bool       `json:"isDeleted"`
	IsArchived bool       `json:"isArchived"`
}

// PosListDto represents list of POSes
type PosListDto struct {
	Total  int64    `json:"total"`
	Values []PosDto `json:"values"`
}

// ========== POS Type ==========

// CreatePosTypeRequest represents request for creating a POS type
type CreatePosTypeRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// PosTypeDto represents POS type data
type PosTypeDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IsDeleted  bool      `json:"isDeleted"`
	IsArchived bool      `json:"isArchived"`
}

// PosTypeListDto represents list of POS types
type PosTypeListDto struct {
	Total  int64        `json:"total"`
	Values []PosTypeDto `json:"values"`
}

// ========== Segment Groups ==========

// SegmentGroupDto represents segment group data
type SegmentGroupDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDeleted   bool      `json:"isDeleted"`
	IsArchived  bool      `json:"isArchived"`
}

// SegmentGroupListDto represents list of segment groups
type SegmentGroupListDto struct {
	Total  int64             `json:"total"`
	Values []SegmentGroupDto `json:"values"`
}

// ========== Sponsored Cards ==========

// SponsoredCardStatuses represents sponsored card status
type SponsoredCardStatuses string

const (
	SponsoredCardStatusCreated       SponsoredCardStatuses = "Created"
	SponsoredCardStatusReadyToIssued SponsoredCardStatuses = "ReadyToIssued"
	SponsoredCardStatusIssued        SponsoredCardStatuses = "Issued"
	SponsoredCardStatusActive        SponsoredCardStatuses = "Active"
	SponsoredCardStatusBlocked       SponsoredCardStatuses = "Blocked"
	SponsoredCardStatusExpired       SponsoredCardStatuses = "Expired"
)

// CreateSponsoredCardRequest represents request for creating a sponsored card
type CreateSponsoredCardRequest struct {
	BaseCommand
	SeriesID uuid.UUID `json:"seriesId"`
}

// BlockSponsoredCardRequest represents request for blocking a sponsored card
type BlockSponsoredCardRequest struct {
	BaseCommand
	BlockReasonID *uuid.UUID `json:"blockReasonId,omitempty"`
}

// IssueSponsoredCardRequest represents request for issuing a sponsored card
type IssueSponsoredCardRequest struct {
	BaseCommand
	OwnerID       uuid.UUID  `json:"ownerId"`
	IssueReasonID *uuid.UUID `json:"issueReasonId,omitempty"`
}

// SponsoredCardAccrualRequest represents request for accrual to sponsored card
type SponsoredCardAccrualRequest struct {
	BaseCommand
	CurrencyID     uuid.UUID  `json:"currencyId"`
	Amount         float64    `json:"amount"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// SponsoredCardSubtractRequest represents request for subtract from sponsored card
type SponsoredCardSubtractRequest struct {
	BaseCommand
	CurrencyID uuid.UUID `json:"currencyId"`
	Amount     float64   `json:"amount"`
}

// SponsoredCardDto represents sponsored card data
type SponsoredCardDto struct {
	ID             uuid.UUID             `json:"id"`
	Number         string                `json:"number"`
	Barcode        *string               `json:"barcode,omitempty"`
	SeriesID       uuid.UUID             `json:"seriesId"`
	Status         SponsoredCardStatuses `json:"status"`
	OwnerID        *uuid.UUID            `json:"ownerId,omitempty"`
	IssueDate      *time.Time            `json:"issueDate,omitempty"`
	ExpirationDate *time.Time            `json:"expirationDate,omitempty"`
	IsDeleted      bool                  `json:"isDeleted"`
	Groups         []uuid.UUID           `json:"groups,omitempty"`
	Balances       []BalanceDto          `json:"balances,omitempty"`
}

// SponsoredCardListDto represents list of sponsored cards
type SponsoredCardListDto struct {
	Total  int64              `json:"total"`
	Values []SponsoredCardDto `json:"values"`
}

// SponsoredCardTransactionDto represents sponsored card transaction data
type SponsoredCardTransactionDto struct {
	TransactionID   uuid.UUID        `json:"transactionId"`
	OperationDate   time.Time        `json:"operationDate"`
	TransactionType TransactionTypes `json:"transactionType"`
	Amount          float64          `json:"amount"`
	CurrencyID      uuid.UUID        `json:"currencyId"`
	ExpirationDate  *time.Time       `json:"expirationDate,omitempty"`
}

// SponsoredCardTransactionListDto represents list of sponsored card transactions
type SponsoredCardTransactionListDto struct {
	Transactions []SponsoredCardTransactionDto `json:"transactions"`
}

// ChequeDto represents cheque data
type ChequeDto struct {
	ChequeID      uuid.UUID  `json:"chequeId"`
	OperationDate time.Time  `json:"operationDate"`
	StoreID       *uuid.UUID `json:"storeId,omitempty"`
	PosID         *uuid.UUID `json:"posId,omitempty"`
	TotalAmount   float64    `json:"totalAmount"`
}

// ChequeListDto represents list of cheques
type ChequeListDto struct {
	Cheques []ChequeDto `json:"cheques"`
}

// ========== Sponsored Card Owners ==========

// CreateSponsoredCardOwnerRequest represents request for creating a sponsored card owner
type CreateSponsoredCardOwnerRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SponsoredCardOwnerDto represents sponsored card owner data
type SponsoredCardOwnerDto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IsDeleted  bool      `json:"isDeleted"`
	IsArchived bool      `json:"isArchived"`
}

// SponsoredCardOwnerListDto represents list of sponsored card owners
type SponsoredCardOwnerListDto struct {
	Total  int64                   `json:"total"`
	Values []SponsoredCardOwnerDto `json:"values"`
}

// ========== Static Segments ==========

// CreateStaticSegmentRequest represents request for creating a static segment
type CreateStaticSegmentRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetStaticSegmentGroupRequest represents request for setting static segment group
type SetStaticSegmentGroupRequest struct {
	BaseCommand
	GroupID *uuid.UUID `json:"groupId,omitempty"`
}

// StaticSegmentDto represents static segment data
type StaticSegmentDto struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Color       *string    `json:"color,omitempty"`
	GroupID     *uuid.UUID `json:"groupId,omitempty"`
	IsDeleted   bool       `json:"isDeleted"`
	IsArchived  bool       `json:"isArchived"`
}

// StaticSegmentListDto represents list of static segments
type StaticSegmentListDto struct {
	Total  int64              `json:"total"`
	Values []StaticSegmentDto `json:"values"`
}

// ========== Stores ==========

// CreateStoreRequest represents request for creating a store
type CreateStoreRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// SetStoreInternalNumberRequest represents request for setting store internal number
type SetStoreInternalNumberRequest struct {
	BaseCommand
	InternalNumber string `json:"internalNumber"`
}

// SetStoreTimeZoneRequest represents request for setting store time zone
type SetStoreTimeZoneRequest struct {
	BaseCommand
	TimeZone string `json:"timeZone"`
}

// SetStoreAddressRequest represents request for setting store address
type SetStoreAddressRequest struct {
	BaseCommand
	Address string `json:"address"`
}

// CloseStoreRequest represents request for closing a store
type CloseStoreRequest struct {
	BaseCommand
}

// OpenStoreRequest represents request for opening a store
type OpenStoreRequest struct {
	BaseCommand
}

// SetStoreRetailSpaceRequest represents request for setting store retail space
type SetStoreRetailSpaceRequest struct {
	BaseCommand
	RetailSpace float64 `json:"retailSpace"`
}

// SetStoreFormatRequest represents request for setting store format
type SetStoreFormatRequest struct {
	BaseCommand
	FormatID *uuid.UUID `json:"formatId,omitempty"`
}

// SetStoreTerritorialDivisionRequest represents request for setting store territorial division
type SetStoreTerritorialDivisionRequest struct {
	BaseCommand
	TerritorialDivisionID *uuid.UUID `json:"territorialDivisionId,omitempty"`
}

// SetStoreOpeningHoursRequest represents request for setting store opening hours
type SetStoreOpeningHoursRequest struct {
	BaseCommand
	OpeningHoursID *uuid.UUID `json:"openingHoursId,omitempty"`
}

// SetStoreLocationCoordinatesRequest represents request for setting store location coordinates
type SetStoreLocationCoordinatesRequest struct {
	BaseCommand
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

// StoreDto represents store data
type StoreDto struct {
	ID                    uuid.UUID         `json:"id"`
	Name                  string            `json:"name"`
	PublicName            *string           `json:"publicName,omitempty"`
	InternalNumber        *string           `json:"internalNumber,omitempty"`
	TimeZone              *string           `json:"timeZone,omitempty"`
	Address               *string           `json:"address,omitempty"`
	RetailSpace           *float64          `json:"retailSpace,omitempty"`
	FormatID              *uuid.UUID        `json:"formatId,omitempty"`
	TerritorialDivisionID *uuid.UUID        `json:"territorialDivisionId,omitempty"`
	OpeningHoursID        *uuid.UUID        `json:"openingHoursId,omitempty"`
	Latitude              *float64          `json:"latitude,omitempty"`
	Longitude             *float64          `json:"longitude,omitempty"`
	IsClosed              bool              `json:"isClosed"`
	IsDeleted             bool              `json:"isDeleted"`
	IsArchived            bool              `json:"isArchived"`
	MediaContents         []MediaContentDto `json:"mediaContents,omitempty"`
}

// MediaContentDto represents media content data
type MediaContentDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Source    string    `json:"source"`
	Type      string    `json:"type"`
	IsDefault bool      `json:"isDefault"`
}

// StoreListDto represents list of stores
type StoreListDto struct {
	Total  int64      `json:"total"`
	Values []StoreDto `json:"values"`
}

// ========== Store Clusters ==========

// AddStoreToClusterRequest represents request for adding store to cluster
type AddStoreToClusterRequest struct {
	BaseCommand
	StoreID uuid.UUID `json:"storeId"`
}

// RemoveStoreFromClusterRequest represents request for removing store from cluster
type RemoveStoreFromClusterRequest struct {
	BaseCommand
	StoreID uuid.UUID `json:"storeId"`
}

// StoreClusterDto represents store cluster data
type StoreClusterDto struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	Color      *string     `json:"color,omitempty"`
	IsDeleted  bool        `json:"isDeleted"`
	IsArchived bool        `json:"isArchived"`
	Stores     []uuid.UUID `json:"stores,omitempty"`
}

// StoreClusterListDto represents list of store clusters
type StoreClusterListDto struct {
	Total  int64             `json:"total"`
	Values []StoreClusterDto `json:"values"`
}

// ========== Store Formats ==========

// CreateStoreFormatRequest represents request for creating a store format
type CreateStoreFormatRequest struct {
	BaseCommand
	Name string `json:"name"`
}

// StoreFormatDto represents store format data
type StoreFormatDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"isDeleted"`
}

// StoreFormatListDto represents list of store formats
type StoreFormatListDto struct {
	Total  int64            `json:"total"`
	Values []StoreFormatDto `json:"values"`
}

// ========== Store Properties ==========

// CreateStorePropertyRequest represents request for creating a store property
type CreateStorePropertyRequest struct {
	BaseCommand
	Name string                `json:"name"`
	Type CustomerPropertyTypes `json:"type"`
}

// ========== Territorial Divisions ==========

// CreateTerritorialDivisionRequest represents request for creating a territorial division
type CreateTerritorialDivisionRequest struct {
	BaseCommand
	Name     string     `json:"name"`
	ParentID *uuid.UUID `json:"parentId,omitempty"`
}

// TerritorialDivisionDto represents territorial division data
type TerritorialDivisionDto struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	ParentID   *uuid.UUID `json:"parentId,omitempty"`
	IsDeleted  bool       `json:"isDeleted"`
	IsArchived bool       `json:"isArchived"`
}

// TerritorialDivisionListDto represents list of territorial divisions
type TerritorialDivisionListDto struct {
	Total  int64                    `json:"total"`
	Values []TerritorialDivisionDto `json:"values"`
}
