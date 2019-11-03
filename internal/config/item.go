package config

import (
	"context"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

// ItemServer is the struct for item service layer
type ItemServer interface {
	GroupedItems(ctx context.Context, groupID, itemID int) ([]GroupedItems, error)
	Create(ctx context.Context, item ItemToCreate) ([]GroupedItems, error)
	Change(ctx context.Context, item ItemToUpdate) ([]GroupedItems, error)
	Remove(ctx context.Context, itemID int) error
}

// NewItem creates new config item service layer
func NewItem(is persistent.ItemStore, gs persistent.GroupStore) ItemServer {
	return &itemService{itemStore: is, groupStore: gs}
}

type itemService struct {
	itemStore  persistent.ItemStore
	groupStore persistent.GroupStore
}

func (is *itemService) GroupedItems(ctx context.Context, groupID, itemID int) ([]GroupedItems, error) {
	return nil, nil
}

func (is *itemService) Create(ctx context.Context, itemToCreate ItemToCreate) ([]GroupedItems, error) {
	return nil, nil
}

func (is *itemService) Change(ctx context.Context, itemToUpdate ItemToUpdate) ([]GroupedItems, error) {
	return nil, nil
}

func (is *itemService) Remove(ctx context.Context, itemID int) error {
	return nil
}
