package models

import (
	"time"

	"github.com/google/uuid"
)

// ========== Customer Enums ==========

// ========== Customer DTOs ==========

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

// CustomerBalanceDto represents customer balance
type CustomerBalanceDto struct {
	Balances            *BalanceListDto  `json:"balances,omitempty"`
	LoyaltyCardBalances []BalanceListDto `json:"loyaltyCardBalances,omitempty"`
}

// ========== Customer Requests ==========

// GetByCommunicationValueRequest represents request for getting customer by communication value
type GetByCommunicationValueRequest struct {
	CommunicationValueType CommunicationValueTypes `json:"communicationValueType"`
	Value                  string                  `json:"value"`
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
