// Package config provides the business logic of this config service.
package config

import (
	"errors"
	"time"
)

var (
	// ErrRequiredField will be return if one or more required field is missing or empty
	ErrRequiredField = errors.New("required field can not be missing or empty")
	// ErrNotFound will be return if no record about config item or group found
	ErrNotFound = errors.New("no record found")
	// ErrInvalidGroup will be return if there is one or more invalid group while adding new item config
	ErrInvalidGroup = errors.New("invalid group id")
	// ErrNoRecordRemoved will be return if no record config item or group removed
	ErrNoRecordRemoved = errors.New("no record removed")
)

// GroupedItems is the struct for grouped config items
type GroupedItems struct {
	GroupID   int
	GroupName string
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []Item
}

// Item is the struct for config item with key value pair
type Item struct {
	ID        int
	Key       string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ItemToCreate is the struct for new config item to be created
type ItemToCreate struct {
	GroupID int
	Key     string
	Value   string
}

// ItemToUpdate is the struct for updating config item
type ItemToUpdate struct {
	ID      int
	GroupID int
	Key     string
	Value   string
}

// Group is the struct for config group
type Group struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
