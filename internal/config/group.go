package config

import (
	"context"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

// GroupServer is the struct for group service layer
type GroupServer interface {
	Groups(ctx context.Context, id []int) ([]Group, error)
	Create(ctx context.Context, group Group) (Group, error)
	Change(ctx context.Context, group Group) (Group, error)
	Remove(ctx context.Context, itemID int) error
}

// NewGroup creates new group service layer
func NewGroup(gs persistent.GroupStore) GroupServer {
	return &groupService{groupStore: gs}
}

type groupService struct {
	groupStore persistent.GroupStore
}

func (gs *groupService) Groups(ctx context.Context, ids []int) ([]Group, error) {
	g, err := gs.groupStore.Groups(ctx, ids)
	if err != nil {
		return nil, err
	}

	if len(g) <= 0 {
		return nil, ErrNotFound
	}

	if len(g) != len(ids) && len(ids) > 0 {
		return nil, ErrInvalidGroup
	}

	groups := make([]Group, len(g))
	for i := 0; i < len(g); i++ {
		groups[i].ID = g[i].ID
		groups[i].Name = g[i].Name
		groups[i].UpdatedAt = g[i].UpdatedAt
		groups[i].CreatedAt = g[i].CreatedAt
	}

	return groups, nil
}

func (gs *groupService) Create(ctx context.Context, group Group) (Group, error) {
	if group.Name == "" {
		return group, ErrRequiredField
	}

	g, err := gs.groupStore.Insert(ctx, persistent.Group{Name: group.Name})
	if err != nil {
		return group, err
	}

	newGroups, err := gs.Groups(ctx, []int{g.ID})
	if err != nil {
		return group, err
	}

	for _, newGroup := range newGroups {
		group.ID = newGroup.ID
		group.Name = newGroup.Name
		group.UpdatedAt = newGroup.UpdatedAt
		group.CreatedAt = newGroup.CreatedAt
	}

	return group, nil
}

func (gs *groupService) Change(ctx context.Context, group Group) (Group, error) {
	if group.ID <= 0 || group.Name == "" {
		return group, ErrRequiredField
	}

	g, err := gs.groupStore.Update(ctx, persistent.Group{ID: group.ID, Name: group.Name})
	if err != nil {
		return group, err
	}

	newGroups, err := gs.Groups(ctx, []int{g.ID})
	if err != nil {
		return group, err
	}

	for _, newGroup := range newGroups {
		group.ID = newGroup.ID
		group.Name = newGroup.Name
		group.UpdatedAt = newGroup.UpdatedAt
		group.CreatedAt = newGroup.CreatedAt
	}

	return group, nil
}

func (gs *groupService) Remove(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrRequiredField
	}

	if err := gs.groupStore.Delete(ctx, id); err != nil {
		if err == persistent.ErrNoRowsAffected {
			return ErrNoRecordRemoved
		}

		return err
	}

	return nil
}
