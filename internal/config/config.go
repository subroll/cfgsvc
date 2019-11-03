// Package config provides the business logic of this config service.
package config

import (
	"time"
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
