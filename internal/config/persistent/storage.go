package persistent

import (
	"context"
	"time"
)

// Item is the data structure of config item
type Item struct {
	ID        int
	Key       string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Group     Group
}

// Group is the data structure of config group
type Group struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ItemStore is the interface for item persistent storage layer that should be implemented
type ItemStore interface {
	Items(ctx context.Context, groupID, itemID int) ([]Item, error)
	Insert(ctx context.Context, item Item) (Item, error)
	Update(ctx context.Context, item Item) (Item, error)
	Delete(ctx context.Context, id int) error
}

// GroupStore is the interface for item group persistent storage layer that should be implemented
type GroupStore interface {
	Groups(ctx context.Context, ids []int) ([]Group, error)
	Insert(ctx context.Context, group Group) (Group, error)
	Update(ctx context.Context, group Group) (Group, error)
	Delete(ctx context.Context, id int) error
}
