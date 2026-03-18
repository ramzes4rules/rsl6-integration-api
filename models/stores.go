package models

import (
	"github.com/google/uuid"
)

// ========== Store DTOs ==========

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

// StoreListDto represents list of stores
type StoreListDto struct {
	Total  int64      `json:"total"`
	Values []StoreDto `json:"values"`
}

// ========== Store Requests ==========

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
