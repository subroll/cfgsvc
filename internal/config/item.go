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

func (is *itemService) validateGroup(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrInvalidGroup
	}

	groups, err := is.groupStore.Groups(ctx, []int{id})
	if err != nil {
		return err
	}

	if len(groups) <= 0 {
		return ErrInvalidGroup
	}

	return nil
}

func (is *itemService) makeGroupedItems(items []persistent.Item) []GroupedItems {
	groups := make(map[int][]persistent.Item)
	for _, item := range items {
		groups[item.Group.ID] = append(groups[item.Group.ID], item)
	}

	var i int
	groupedItems := make([]GroupedItems, len(groups))
	for _, v := range groups {
		groupedItems[i].GroupID = v[i].Group.ID
		groupedItems[i].GroupName = v[i].Group.Name
		groupedItems[i].CreatedAt = v[i].Group.CreatedAt
		groupedItems[i].UpdatedAt = v[i].Group.UpdatedAt

		for _, item := range v {
			groupedItems[i].Items = append(groupedItems[i].Items, Item{
				ID:        item.ID,
				Key:       item.Key,
				Value:     item.Value,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}

		i++
	}

	return groupedItems
}

func (is *itemService) GroupedItems(ctx context.Context, groupID, itemID int) ([]GroupedItems, error) {
	items, err := is.itemStore.Items(ctx, groupID, itemID)
	if err != nil {
		return nil, err
	}

	if len(items) <= 0 {
		return nil, ErrNotFound
	}

	return is.makeGroupedItems(items), nil
}

func (is *itemService) Create(ctx context.Context, itemToCreate ItemToCreate) ([]GroupedItems, error) {
	if itemToCreate.Key == "" || itemToCreate.Value == "" {
		return nil, ErrRequiredField
	}

	if err := is.validateGroup(ctx, itemToCreate.GroupID); err != nil {
		return nil, err
	}

	item, err := is.itemStore.Insert(ctx, persistent.Item{
		Key:   itemToCreate.Key,
		Value: itemToCreate.Value,
		Group: persistent.Group{ID: itemToCreate.GroupID},
	})
	if err != nil {
		return nil, err
	}

	items, err := is.itemStore.Items(ctx, item.Group.ID, item.ID)
	if err != nil {
		return nil, err
	}

	return is.makeGroupedItems(items), nil
}

func (is *itemService) Change(ctx context.Context, itemToUpdate ItemToUpdate) ([]GroupedItems, error) {
	if itemToUpdate.ID <= 0 || itemToUpdate.Key == "" || itemToUpdate.Value == "" {
		return nil, ErrRequiredField
	}
	if err := is.validateGroup(ctx, itemToUpdate.GroupID); err != nil {
		return nil, err
	}

	item, err := is.itemStore.Update(ctx, persistent.Item{
		ID:    itemToUpdate.ID,
		Key:   itemToUpdate.Key,
		Value: itemToUpdate.Value,
		Group: persistent.Group{ID: itemToUpdate.GroupID},
	})
	if err != nil {
		return nil, err
	}

	items, err := is.itemStore.Items(ctx, item.Group.ID, item.ID)
	if err != nil {
		return nil, err
	}

	return is.makeGroupedItems(items), nil
}

func (is *itemService) Remove(ctx context.Context, itemID int) error {
	if itemID <= 0 {
		return ErrRequiredField
	}

	if err := is.itemStore.Delete(ctx, itemID); err != nil {
		if err == persistent.ErrNoRowsAffected {
			return ErrNoRecordRemoved
		}

		return err
	}

	return nil
}
