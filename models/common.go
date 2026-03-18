// Package models contains data structures for RS Loyalty API v2
package models

import (
	"time"

	"github.com/google/uuid"
)

// ========== Common Enums ==========

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

// PropertyTypes represents property types (used for customer, item, store properties)
type PropertyTypes string

const (
	PropertyTypeInteger      PropertyTypes = "Integer"
	PropertyTypeString       PropertyTypes = "String"
	PropertyTypeDate         PropertyTypes = "Date"
	PropertyTypeEnum         PropertyTypes = "Enum"
	PropertyTypeBoolean      PropertyTypes = "Boolean"
	PropertyTypeEnumMultiple PropertyTypes = "EnumMultiple"
)

// CustomerPropertyTypes is an alias for backward compatibility
type CustomerPropertyTypes = PropertyTypes

// ========== Base Command ==========

// BaseCommand represents base command structure with common fields
type BaseCommand struct {
	CommandID     *uuid.UUID `json:"commandId,omitempty"`
	OperationDate *time.Time `json:"operationDate,omitempty"`
	ID            uuid.UUID  `json:"id"`
}

// ========== Common Request Types ==========

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

// BatchRequest represents batch request for executing multiple commands
type BatchRequest struct {
	Commands []interface{} `json:"commands"`
}

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

// ========== Common Property Requests ==========

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

// ========== Common Enum Requests ==========

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

// ========== Common Media Content Requests ==========

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

// ========== Common DTOs ==========

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

// MediaContentDto represents media content data
type MediaContentDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Source    string    `json:"source"`
	Type      string    `json:"type"`
	IsDefault bool      `json:"isDefault"`
}

// PropertyDefinitionDto represents property definition
type PropertyDefinitionDto struct {
	ID        uuid.UUID              `json:"id"`
	Type      PropertyTypes          `json:"type"`
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

// PropertyDefinitionListDto represents list of property definitions
type PropertyDefinitionListDto struct {
	Total  int64                   `json:"total"`
	Values []PropertyDefinitionDto `json:"values"`
}

// ========== Common Headers ==========

// RequestHeaders represents common request headers
type RequestHeaders struct {
	CommandID          *string `json:"command-id,omitempty"`
	OperationDate      *string `json:"operation-date,omitempty"`
	UserID             *string `json:"user-id,omitempty"`
	InteractionChannel *string `json:"interaction-channel,omitempty"`
}

// ========== API Error ==========

// APIError represents API error response
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

// ========== Common Reason DTOs ==========

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

// ========== Common Group DTOs ==========

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

// ========== Common Card Requests ==========

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

// ConfirmReadinessForIssuanceRequest represents request for confirming readiness for issuance
type ConfirmReadinessForIssuanceRequest struct {
	BaseCommand
}

// CancelIssueRequest represents request for canceling card issue
type CancelIssueRequest struct {
	BaseCommand
}

// UnblockCardRequest represents request for unblocking a card
type UnblockCardRequest struct {
	BaseCommand
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

// SetPublicNameRequest represents request for setting public name
type SetPublicNameRequest struct {
	BaseCommand
	PublicName string `json:"publicName"`
}
