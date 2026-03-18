package models

import (
	"github.com/google/uuid"
)

// ========== Opening Hours DTOs ==========

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

// ========== Opening Hours Requests ==========

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
